package ies

import "rrc/utils"

// CarrierFreqGERAN ::= SEQUENCE
type Carrierfreqgeran struct {
	Arfcn         ArfcnValuegeran
	Bandindicator Bandindicatorgeran
}
