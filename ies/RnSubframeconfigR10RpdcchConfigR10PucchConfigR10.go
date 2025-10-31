package ies

// RN-SubframeConfig-r10-rpdcch-Config-r10-pucch-Config-r10 ::= CHOICE
const (
	RnSubframeconfigR10RpdcchConfigR10PucchConfigR10ChoiceNothing = iota
	RnSubframeconfigR10RpdcchConfigR10PucchConfigR10ChoiceTdd
	RnSubframeconfigR10RpdcchConfigR10PucchConfigR10ChoiceFdd
)

type RnSubframeconfigR10RpdcchConfigR10PucchConfigR10 struct {
	Choice uint64
	Tdd    *RnSubframeconfigR10RpdcchConfigR10PucchConfigR10Tdd
	Fdd    *RnSubframeconfigR10RpdcchConfigR10PucchConfigR10Fdd
}
