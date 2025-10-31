package ies

import "rrc/utils"

// SRS-ConfigAdd-r16 ::= SEQUENCE
type SrsConfigaddR16 struct {
	SrsRepnumaddR16              SrsConfigaddR16SrsRepnumaddR16
	SrsBandwidthaddR16           SrsConfigaddR16SrsBandwidthaddR16
	SrsHoppingbandwidthaddR16    SrsConfigaddR16SrsHoppingbandwidthaddR16
	SrsFreqdomainposaddR16       utils.INTEGER `lb:0,ub:23`
	SrsAntennaportaddR16         SrsAntennaport
	SrsCyclicshiftaddR16         SrsConfigaddR16SrsCyclicshiftaddR16
	SrsTransmissioncombnumaddR16 SrsConfigaddR16SrsTransmissioncombnumaddR16
	SrsTransmissioncombaddR16    utils.INTEGER `lb:0,ub:3`
	SrsStartposaddR16            utils.INTEGER `lb:0,ub:13`
	SrsDurationaddR16            utils.INTEGER `lb:0,ub:13`
	SrsGuardsymbolasAddR16       *SrsConfigaddR16SrsGuardsymbolasAddR16
	SrsGuardsymbolfhAddR16       *SrsConfigaddR16SrsGuardsymbolfhAddR16
}
