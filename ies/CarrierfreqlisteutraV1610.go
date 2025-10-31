package ies

// CarrierFreqListEUTRA-v1610 ::= SEQUENCE OF CarrierFreqEUTRA-v1610
// SIZE (1..maxEUTRA-Carrier)
type CarrierfreqlisteutraV1610 struct {
	Value []CarrierfreqeutraV1610 `lb:1,ub:maxEUTRACarrier`
}
