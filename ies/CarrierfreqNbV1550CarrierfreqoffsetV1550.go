package ies

import "rrc/utils"

// CarrierFreq-NB-v1550-carrierFreqOffset-v1550 ::= ENUMERATED
type CarrierfreqNbV1550CarrierfreqoffsetV1550 struct {
	Value utils.ENUMERATED
}

const (
	CarrierfreqNbV1550CarrierfreqoffsetV1550EnumeratedNothing = iota
	CarrierfreqNbV1550CarrierfreqoffsetV1550EnumeratedV_8dot5
	CarrierfreqNbV1550CarrierfreqoffsetV1550EnumeratedV_4dot5
	CarrierfreqNbV1550CarrierfreqoffsetV1550EnumeratedV3dot5
	CarrierfreqNbV1550CarrierfreqoffsetV1550EnumeratedV7dot5
)
