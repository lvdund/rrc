package ies

import "rrc/utils"

// SL-CarrierFreqInfo-v1310 ::= SEQUENCE
// Extensible
type SlCarrierfreqinfoV1310 struct {
	DiscresourcesnonpsR13 *SlResourcesinterfreqR13
	DiscresourcespsR13    *SlResourcesinterfreqR13
	DiscconfigotherR13    *SlDiscconfigotherinterfreqR13
}
