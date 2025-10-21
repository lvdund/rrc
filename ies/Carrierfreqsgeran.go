package ies

import "rrc/utils"

// CarrierFreqsGERAN ::= SEQUENCE
type Carrierfreqsgeran struct {
	Startingarfcn   ArfcnValuegeran
	Bandindicator   Bandindicatorgeran
	Followingarfcns interface{}
}
