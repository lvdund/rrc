package ies

// DL-CarrierConfigDedicated-NB-r13 ::= SEQUENCE
// Extensible
type DlCarrierconfigdedicatedNbR13 struct {
	DlCarrierfreqR13           CarrierfreqNbR13
	DownlinkbitmapnonanchorR13 *DlCarrierconfigdedicatedNbR13DownlinkbitmapnonanchorR13
	DlGapnonanchorR13          *DlCarrierconfigdedicatedNbR13DlGapnonanchorR13
	InbandcarrierinfoR13       *DlCarrierconfigdedicatedNbR13InbandcarrierinfoR13
}
