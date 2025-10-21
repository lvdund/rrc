package ies

import "rrc/utils"

// DL-CarrierConfigCommon-NB-r14 ::= SEQUENCE
// Extensible
type DlCarrierconfigcommonNbR14 struct {
	DlCarrierfreqR14           CarrierfreqNbR13
	DownlinkbitmapnonanchorR14 interface{}
	DlGapnonanchorR14          interface{}
	InbandcarrierinfoR14       *interface{}
	NrsPoweroffsetnonanchorR14 utils.ENUMERATED
}
