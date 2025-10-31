package ies

// DL-CarrierConfigCommon-NB-r14-dl-GapNonAnchor-r14 ::= CHOICE
const (
	DlCarrierconfigcommonNbR14DlGapnonanchorR14ChoiceNothing = iota
	DlCarrierconfigcommonNbR14DlGapnonanchorR14ChoiceUsenogapR14
	DlCarrierconfigcommonNbR14DlGapnonanchorR14ChoiceUseanchorgapconfigR14
	DlCarrierconfigcommonNbR14DlGapnonanchorR14ChoiceExplicitgapconfigurationR14
)

type DlCarrierconfigcommonNbR14DlGapnonanchorR14 struct {
	Choice                      uint64
	UsenogapR14                 *struct{}
	UseanchorgapconfigR14       *struct{}
	ExplicitgapconfigurationR14 *DlGapconfigNbR13
}
