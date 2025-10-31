package ies

// RN-SubframeConfig-r10 ::= SEQUENCE
// Extensible
type RnSubframeconfigR10 struct {
	SubframeconfigpatternR10 *RnSubframeconfigR10SubframeconfigpatternR10
	RpdcchConfigR10          *RnSubframeconfigR10RpdcchConfigR10
}
