package ies

import "rrc/utils"

// DL-CarrierConfigDedicated-NB-r13 ::= SEQUENCE
// Extensible
type DlCarrierconfigdedicatedNbR13 struct {
	DlCarrierfreqR13           CarrierfreqNbR13
	DownlinkbitmapnonanchorR13 *interface{}
	DlGapnonanchorR13          *interface{}
	InbandcarrierinfoR13       *interface{}
}
