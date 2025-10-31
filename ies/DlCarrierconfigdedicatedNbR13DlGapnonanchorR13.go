package ies

// DL-CarrierConfigDedicated-NB-r13-dl-GapNonAnchor-r13 ::= CHOICE
const (
	DlCarrierconfigdedicatedNbR13DlGapnonanchorR13ChoiceNothing = iota
	DlCarrierconfigdedicatedNbR13DlGapnonanchorR13ChoiceUsenogapR13
	DlCarrierconfigdedicatedNbR13DlGapnonanchorR13ChoiceUseanchorgapconfigR13
	DlCarrierconfigdedicatedNbR13DlGapnonanchorR13ChoiceExplicitgapconfigurationR13
	DlCarrierconfigdedicatedNbR13DlGapnonanchorR13ChoiceSpare
)

type DlCarrierconfigdedicatedNbR13DlGapnonanchorR13 struct {
	Choice                      uint64
	UsenogapR13                 *struct{}
	UseanchorgapconfigR13       *struct{}
	ExplicitgapconfigurationR13 *DlGapconfigNbR13
	Spare                       *struct{}
}
