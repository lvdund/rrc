package ies

import "rrc/utils"

// MBMS-ROM-Info-r16 ::= SEQUENCE
type MbmsRomInfoR16 struct {
	MbmsRomFreqR16              ArfcnValueeutraR9
	MbmsRomSubcarrierspacingR16 utils.ENUMERATED
	MbmsBandwidthR16            utils.ENUMERATED
}
