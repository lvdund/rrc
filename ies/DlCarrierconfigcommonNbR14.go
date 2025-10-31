package ies

// DL-CarrierConfigCommon-NB-r14 ::= SEQUENCE
// Extensible
type DlCarrierconfigcommonNbR14 struct {
	DlCarrierfreqR14           CarrierfreqNbR13
	DownlinkbitmapnonanchorR14 DlCarrierconfigcommonNbR14DownlinkbitmapnonanchorR14
	DlGapnonanchorR14          DlCarrierconfigcommonNbR14DlGapnonanchorR14
	InbandcarrierinfoR14       *DlCarrierconfigcommonNbR14InbandcarrierinfoR14
	NrsPoweroffsetnonanchorR14 DlCarrierconfigcommonNbR14NrsPoweroffsetnonanchorR14
}
