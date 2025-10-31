package ies

// UAC-Param-NB-r16 ::= CHOICE
const (
	UacParamNbR16ChoiceNothing = iota
	UacParamNbR16ChoiceUacBarringcommon
	UacParamNbR16ChoiceUacBarringperplmnList
)

type UacParamNbR16 struct {
	Choice                uint64
	UacBarringcommon      *UacBarringNbR16
	UacBarringperplmnList *[]UacBarringNbR16 `lb:1,ub:maxPLMNR11`
}
