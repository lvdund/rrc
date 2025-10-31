package ies

// SystemInformationBlockType14-NB-r13-ab-Param-r13 ::= CHOICE
const (
	Systeminformationblocktype14NbR13AbParamR13ChoiceNothing = iota
	Systeminformationblocktype14NbR13AbParamR13ChoiceAbCommonR13
	Systeminformationblocktype14NbR13AbParamR13ChoiceAbPerplmnListR13
)

type Systeminformationblocktype14NbR13AbParamR13 struct {
	Choice           uint64
	AbCommonR13      *AbConfigNbR13
	AbPerplmnListR13 *[]AbConfigplmnNbR13 `lb:1,ub:maxPLMNR11`
}
