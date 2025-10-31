package ies

// BandParametersRxSL-r14 ::= SEQUENCE
type BandparametersrxslR14 struct {
	V2xBandwidthclassrxslR14 V2xBandwidthclassslR14
	V2xHighreceptionR14      *BandparametersrxslR14V2xHighreceptionR14
}
