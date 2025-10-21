package ies

import "rrc/utils"

// SRS-ConfigAdd-r16 ::= SEQUENCE
type SrsConfigaddR16 struct {
	SrsRepnumaddR16              utils.ENUMERATED
	SrsBandwidthaddR16           utils.ENUMERATED
	SrsHoppingbandwidthaddR16    utils.ENUMERATED
	SrsFreqdomainposaddR16       utils.INTEGER
	SrsAntennaportaddR16         SrsAntennaport
	SrsCyclicshiftaddR16         utils.ENUMERATED
	SrsTransmissioncombnumaddR16 utils.ENUMERATED
	SrsTransmissioncombaddR16    utils.INTEGER
	SrsStartposaddR16            utils.INTEGER
	SrsDurationaddR16            utils.INTEGER
	SrsGuardsymbolasAddR16       *utils.ENUMERATED
	SrsGuardsymbolfhAddR16       *utils.ENUMERATED
}
