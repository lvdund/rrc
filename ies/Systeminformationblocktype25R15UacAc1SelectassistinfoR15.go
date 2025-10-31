package ies

// SystemInformationBlockType25-r15-uac-AC1-SelectAssistInfo-r15 ::= CHOICE
const (
	Systeminformationblocktype25R15UacAc1SelectassistinfoR15ChoiceNothing = iota
	Systeminformationblocktype25R15UacAc1SelectassistinfoR15ChoicePlmncommonR15
	Systeminformationblocktype25R15UacAc1SelectassistinfoR15ChoiceIndividualplmnlistR15
)

type Systeminformationblocktype25R15UacAc1SelectassistinfoR15 struct {
	Choice                uint64
	PlmncommonR15         *UacAc1SelectassistinfoR15
	IndividualplmnlistR15 *[]UacAc1SelectassistinfoR15 `lb:2,ub:maxPLMNR11`
}
