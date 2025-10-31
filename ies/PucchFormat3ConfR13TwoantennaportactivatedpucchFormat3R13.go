package ies

// PUCCH-Format3-Conf-r13-twoAntennaPortActivatedPUCCH-Format3-r13 ::= CHOICE
const (
	PucchFormat3ConfR13TwoantennaportactivatedpucchFormat3R13ChoiceNothing = iota
	PucchFormat3ConfR13TwoantennaportactivatedpucchFormat3R13ChoiceRelease
	PucchFormat3ConfR13TwoantennaportactivatedpucchFormat3R13ChoiceSetup
)

type PucchFormat3ConfR13TwoantennaportactivatedpucchFormat3R13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *PucchFormat3ConfR13TwoantennaportactivatedpucchFormat3R13Setup
}
