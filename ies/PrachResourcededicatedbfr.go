package ies

// PRACH-ResourceDedicatedBFR ::= CHOICE
const (
	PrachResourcededicatedbfrChoiceNothing = iota
	PrachResourcededicatedbfrChoiceSsb
	PrachResourcededicatedbfrChoiceCsiRs
)

type PrachResourcededicatedbfr struct {
	Choice uint64
	Ssb    *BfrSsbResource
	CsiRs  *BfrCsirsResource
}
