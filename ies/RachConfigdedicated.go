package ies

// RACH-ConfigDedicated ::= SEQUENCE
// Extensible
type RachConfigdedicated struct {
	Cfra                       *Cfra
	RaPrioritization           *RaPrioritization
	RaPrioritizationtwostepR16 *RaPrioritization `ext`
	CfraTwostepR16             *CfraTwostepR16   `ext`
}
