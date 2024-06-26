package v0

// MfaVote represents the vote for a MFA request
type MfaVote string

const (
	Approve MfaVote = "approve"
	Reject  MfaVote = "reject"
)

// String returns the string representation of MfaVote
func (v MfaVote) String() string {
	return string(v)
}

// IsValid checks if the MfaVote value is valid
func (v MfaVote) IsValid() bool {
	switch v {
	case Approve, Reject:
		return true
	}
	return false
}
