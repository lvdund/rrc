package ies

// BSR-Config ::= SEQUENCE
// Extensible
type BsrConfig struct {
	PeriodicbsrTimer           BsrConfigPeriodicbsrTimer
	RetxbsrTimer               BsrConfigRetxbsrTimer
	LogicalchannelsrDelaytimer *BsrConfigLogicalchannelsrDelaytimer
}
