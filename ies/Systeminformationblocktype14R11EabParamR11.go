package ies

// SystemInformationBlockType14-r11-eab-Param-r11 ::= CHOICE
const (
	Systeminformationblocktype14R11EabParamR11ChoiceNothing = iota
	Systeminformationblocktype14R11EabParamR11ChoiceEabCommonR11
	Systeminformationblocktype14R11EabParamR11ChoiceEabPerplmnListR11
)

type Systeminformationblocktype14R11EabParamR11 struct {
	Choice            uint64
	EabCommonR11      *EabConfigR11
	EabPerplmnListR11 *[]EabConfigplmnR11 `lb:1,ub:maxPLMNR11`
}
