package ies

// CarrierFreqsGERAN ::= SEQUENCE
type Carrierfreqsgeran struct {
	Startingarfcn   ArfcnValuegeran
	Bandindicator   Bandindicatorgeran
	Followingarfcns CarrierfreqsgeranFollowingarfcns
}
