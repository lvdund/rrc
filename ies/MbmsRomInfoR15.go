package ies

import "rrc/utils"

// MBMS-ROM-Info-r15 ::= SEQUENCE
type MbmsRomInfoR15 struct {
	MbmsRomFreqR15              ArfcnValueeutraR9
	MbmsRomSubcarrierspacingR15 utils.ENUMERATED
	MbmsBandwidthR15            utils.ENUMERATED
}
