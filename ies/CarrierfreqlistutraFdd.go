package ies

// CarrierFreqListUTRA-FDD ::= SEQUENCE OF CarrierFreqUTRA-FDD
// SIZE (1..maxUTRA-FDD-Carrier)
type CarrierfreqlistutraFdd struct {
	Value []CarrierfrequtraFdd `lb:1,ub:maxUTRAFddCarrier`
}
