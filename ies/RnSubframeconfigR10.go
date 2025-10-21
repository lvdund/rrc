package ies

import "rrc/utils"

// RN-SubframeConfig-r10 ::= SEQUENCE
// Extensible
type RnSubframeconfigR10 struct {
	SubframeconfigpatternR10 *interface{}
	RpdcchConfigR10          *interface{}
}
