package ies

// RLC-Config-r15-mode-r15 ::= CHOICE
const (
	RlcConfigR15ModeR15ChoiceNothing = iota
	RlcConfigR15ModeR15ChoiceAmR15
	RlcConfigR15ModeR15ChoiceUmBiDirectionalR15
	RlcConfigR15ModeR15ChoiceUmUniDirectionalUlR15
	RlcConfigR15ModeR15ChoiceUmUniDirectionalDlR15
)

type RlcConfigR15ModeR15 struct {
	Choice                uint64
	AmR15                 *RlcConfigR15ModeR15AmR15
	UmBiDirectionalR15    *RlcConfigR15ModeR15UmBiDirectionalR15
	UmUniDirectionalUlR15 *RlcConfigR15ModeR15UmUniDirectionalUlR15
	UmUniDirectionalDlR15 *RlcConfigR15ModeR15UmUniDirectionalDlR15
}
