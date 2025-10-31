package ies

// DL-CarrierConfigDedicated-NB-r13-downlinkBitmapNonAnchor-r13 ::= CHOICE
const (
	DlCarrierconfigdedicatedNbR13DownlinkbitmapnonanchorR13ChoiceNothing = iota
	DlCarrierconfigdedicatedNbR13DownlinkbitmapnonanchorR13ChoiceUsenobitmapR13
	DlCarrierconfigdedicatedNbR13DownlinkbitmapnonanchorR13ChoiceUseanchorbitmapR13
	DlCarrierconfigdedicatedNbR13DownlinkbitmapnonanchorR13ChoiceExplicitbitmapconfigurationR13
	DlCarrierconfigdedicatedNbR13DownlinkbitmapnonanchorR13ChoiceSpare
)

type DlCarrierconfigdedicatedNbR13DownlinkbitmapnonanchorR13 struct {
	Choice                         uint64
	UsenobitmapR13                 *struct{}
	UseanchorbitmapR13             *struct{}
	ExplicitbitmapconfigurationR13 *DlBitmapNbR13
	Spare                          *struct{}
}
