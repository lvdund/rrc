package ies

// MBMS-ROM-Info-r15 ::= SEQUENCE
type MbmsRomInfoR15 struct {
	MbmsRomFreqR15              ArfcnValueeutraR9
	MbmsRomSubcarrierspacingR15 MbmsRomInfoR15MbmsRomSubcarrierspacingR15
	MbmsBandwidthR15            MbmsRomInfoR15MbmsBandwidthR15
}
