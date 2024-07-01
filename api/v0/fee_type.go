package v0

// FeeType represents the type of fee.
type FeeType string

const (
	FixedFee FeeType = "fixed"
	RateFee  FeeType = "rate"
)
