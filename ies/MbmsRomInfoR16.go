package ies

// MBMS-ROM-Info-r16 ::= SEQUENCE
type MbmsRomInfoR16 struct {
	MbmsRomFreqR16              ArfcnValueeutraR9
	MbmsRomSubcarrierspacingR16 MbmsRomInfoR16MbmsRomSubcarrierspacingR16
	MbmsBandwidthR16            MbmsRomInfoR16MbmsBandwidthR16
}
