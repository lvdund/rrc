package ies

// RLC-Config-NB-r13 ::= CHOICE
// Extensible
const (
	RlcConfigNbR13ChoiceNothing = iota
	RlcConfigNbR13ChoiceAm
	RlcConfigNbR13ChoiceUmBiDirectionalR15
	RlcConfigNbR13ChoiceUmUniDirectionalUlR15
	RlcConfigNbR13ChoiceUmUniDirectionalDlR15
)

type RlcConfigNbR13 struct {
	Choice                uint64
	Am                    *RlcConfigNbR13Am
	UmBiDirectionalR15    *struct{}
	UmUniDirectionalUlR15 *struct{}
	UmUniDirectionalDlR15 *struct{}
}
