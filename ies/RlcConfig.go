package ies

// RLC-Config ::= CHOICE
// Extensible
const (
	RlcConfigChoiceNothing = iota
	RlcConfigChoiceAm
	RlcConfigChoiceUmBiDirectional
	RlcConfigChoiceUmUniDirectionalUl
	RlcConfigChoiceUmUniDirectionalDl
)

type RlcConfig struct {
	Choice             uint64
	Am                 *RlcConfigAm
	UmBiDirectional    *RlcConfigUmBiDirectional
	UmUniDirectionalUl *RlcConfigUmUniDirectionalUl
	UmUniDirectionalDl *RlcConfigUmUniDirectionalDl
}
