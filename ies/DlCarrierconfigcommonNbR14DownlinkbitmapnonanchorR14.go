package ies

// DL-CarrierConfigCommon-NB-r14-downlinkBitmapNonAnchor-r14 ::= CHOICE
const (
	DlCarrierconfigcommonNbR14DownlinkbitmapnonanchorR14ChoiceNothing = iota
	DlCarrierconfigcommonNbR14DownlinkbitmapnonanchorR14ChoiceUsenobitmapR14
	DlCarrierconfigcommonNbR14DownlinkbitmapnonanchorR14ChoiceUseanchorbitmapR14
	DlCarrierconfigcommonNbR14DownlinkbitmapnonanchorR14ChoiceExplicitbitmapconfigurationR14
)

type DlCarrierconfigcommonNbR14DownlinkbitmapnonanchorR14 struct {
	Choice                         uint64
	UsenobitmapR14                 *struct{}
	UseanchorbitmapR14             *struct{}
	ExplicitbitmapconfigurationR14 *DlBitmapNbR13
}
