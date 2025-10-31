package ies

// MeasObjectRxTxDiff-r17-dl-Ref-r17 ::= CHOICE
// Extensible
const (
	MeasobjectrxtxdiffR17DlRefR17ChoiceNothing = iota
	MeasobjectrxtxdiffR17DlRefR17ChoicePrsRefR17
	MeasobjectrxtxdiffR17DlRefR17ChoiceCsiRsRefR17
)

type MeasobjectrxtxdiffR17DlRefR17 struct {
	Choice      uint64
	PrsRefR17   *struct{}
	CsiRsRefR17 *struct{}
}
