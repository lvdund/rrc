package ies

// CarrierFreqListEUTRA ::= SEQUENCE OF CarrierFreqEUTRA
// SIZE (1..maxEUTRA-Carrier)
type Carrierfreqlisteutra struct {
	Value []Carrierfreqeutra `lb:1,ub:maxEUTRACarrier`
}
