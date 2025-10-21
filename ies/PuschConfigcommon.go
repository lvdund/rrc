package ies

import "rrc/utils"

// PUSCH-ConfigCommon ::= SEQUENCE
type PuschConfigcommon struct {
	PuschConfigbasic        interface{}
	UlReferencesignalspusch UlReferencesignalspusch
}
