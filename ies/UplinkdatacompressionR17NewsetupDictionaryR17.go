package ies

import "rrc/utils"

// UplinkDataCompression-r17-newSetup-dictionary-r17 ::= ENUMERATED
type UplinkdatacompressionR17NewsetupDictionaryR17 struct {
	Value utils.ENUMERATED
}

const (
	UplinkdatacompressionR17NewsetupDictionaryR17EnumeratedNothing = iota
	UplinkdatacompressionR17NewsetupDictionaryR17EnumeratedSip_Sdp
	UplinkdatacompressionR17NewsetupDictionaryR17EnumeratedOperator
)
