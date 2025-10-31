package ies

// CarrierFreqListEUTRA-v1700 ::= SEQUENCE OF CarrierFreqEUTRA-v1700
// SIZE (1..maxEUTRA-Carrier)
type CarrierfreqlisteutraV1700 struct {
	Value []CarrierfreqeutraV1700 `lb:1,ub:maxEUTRACarrier`
}
