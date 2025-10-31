package ies

// BandParametersTxSL-r14 ::= SEQUENCE
type BandparameterstxslR14 struct {
	V2xBandwidthclasstxslR14 V2xBandwidthclassslR14
	V2xEnbScheduledR14       *BandparameterstxslR14V2xEnbScheduledR14
	V2xHighpowerR14          *BandparameterstxslR14V2xHighpowerR14
}
