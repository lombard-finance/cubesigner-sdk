package client

import (
	"fmt"
	"os"

	"github.com/lombard-finance/cubesigner-sdk/api"
	v0 "github.com/lombard-finance/cubesigner-sdk/api/v0"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	DomainChainId          = int64(0x42415244)
	LombardWhaleDepositStr = "Lombard Approval Deposit"
	VersionStr             = "1"
)

type ClientArgs struct {
	address string
	orgID   string
	token   string
	logger  *logrus.Entry
	timeout time.Duration
}

// / The message that we will sign using EIP712-style signing
type WhaleDepositMessage struct {
	ToAddress           string
	ChainId             uint64
	LbtcContractAddress string
	ReferralId          string
	FinalityProviderPk  string
	Utxo                string
}

func TestClient_SignEip712(t *testing.T) {
	token := os.Getenv("CUBESIGNER_TOKEN")
	if len(token) == 0 {
		panic("must set CUBESIGNER_TOKEN envvar, e.v.,\nexport CUBESIGNER_TOKEN=<base64_token>")
	}
	pubkey := os.Getenv("CUBESIGNER_EVM_ADDR")
	if len(pubkey) == 0 {
		panic("must set CUBESIGNER_EVM_ADDR envvar, e.g.,\nexport CUBESIGNER_EVM_ADDR=0xbd9ee03f626377e3a9d86f933aa04398ecd05e5c")
	}
	org_id := os.Getenv("CUBESIGNER_ORG_ID")
	if len(org_id) == 0 {
		panic("must set CUBESIGNER_ORG_ID envvar, e.g.,\nexport CUBESIGNER_ORG_ID=Org#9aa34a9c-3ff3-47fa-a98b-8c288812731c")
	}
	cubesigner_env := os.Getenv("CUBESIGNER_ENV")
	if len(cubesigner_env) == 0 {
		cubesigner_env = "https://gamma.signer.cubist.dev"
	}

	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	domainChainIdString := fmt.Sprintf("%v", DomainChainId)

	type args struct {
		pubkey          string
		request         *v0.Eip712SignRequest
		mfaId           *string
		mfaConfirmation *string
	}
	tests := []struct {
		name       string
		clientArgs ClientArgs
		args       args
		want       *v0.EvmSignResponse
		want1      string
		wantErr    bool
	}{
		{
			name: "",
			clientArgs: ClientArgs{
				address: cubesigner_env,
				orgID:   org_id,
				token:   token,
				logger:  logrus.NewEntry(logger),
				timeout: time.Second * 5,
			},
			args: args{
				pubkey: pubkey,
				request: &v0.Eip712SignRequest{
					ChainId: DomainChainId,
					TypedData: v0.TypedData{
						Domain: v0.TypedDataDomain{
							ChainId: *api.NewNullableString(&domainChainIdString),
							Name:    *api.NewNullableString(&LombardWhaleDepositStr),
							Version: *api.NewNullableString(&VersionStr),
						},
						Message: map[string]interface{}{
							"ToAddress":           "0x17a604740a25d703f9848857beb9e88d0ab0d3f8\n",
							"ChainId":             1,
							"LbtcContractAddress": "0x8236a87084f8B84306f72007F36F2618A5634494",
							"ReferralId":          "lombard",
							"FinalityProviderPk":  "03d5a0bb72d71993e435d6c5a70e2aa4db500a62cfaae33c56050deefee64ec0",
							"Utxo":                "c862eae3f63015d0b9bab1c877e44c6ef89a7e56c310a83122e85403f628ef2a:0",
						},
						PrimaryType: "ApprovalMessage",
						Types: map[string][]v0.TypedDataTypesValueInner{
							"ApprovalMessage": {
								{
									Name: "ToAddress",
									Type: "string",
								},
								{
									Name: "ChainId",
									Type: "uint256",
								},
								{
									Name: "LbtcContractAddress",
									Type: "string",
								},
								{
									Name: "ReferralId",
									Type: "string",
								},
								{
									Name: "FinalityProviderPk",
									Type: "string",
								},
								{
									Name: "Utxo",
									Type: "string",
								},
							},
						},
					},
				},
			},
			want:    nil,
			want1:   "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, err := New(tt.clientArgs.address, tt.clientArgs.orgID, tt.clientArgs.token, tt.clientArgs.logger, tt.clientArgs.timeout)
			require.NoError(t, err)
			got, got1, err := cli.SignEip712(tt.args.pubkey, tt.args.request, tt.args.mfaId, tt.args.mfaConfirmation)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignEip712() error = %v, wantErr %v", err, tt.wantErr)
				t.Errorf("Did you set the '\"AllowEip712Signing\"' policy on Key#%v ?\n", pubkey)
				return
			}
			fmt.Printf("Response: %v\nStatus: %v\n", got.Signature, got1)
		})
	}
}
