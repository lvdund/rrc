package ies

// SIB1-uac-BarringInfo-uac-AccessCategory1-SelectionAssistanceInfo ::= CHOICE
const (
	Sib1UacBarringinfoUacAccesscategory1SelectionassistanceinfoChoiceNothing = iota
	Sib1UacBarringinfoUacAccesscategory1SelectionassistanceinfoChoicePlmncommon
	Sib1UacBarringinfoUacAccesscategory1SelectionassistanceinfoChoiceIndividualplmnlist
)

type Sib1UacBarringinfoUacAccesscategory1Selectionassistanceinfo struct {
	Choice             uint64
	Plmncommon         *UacAccesscategory1Selectionassistanceinfo
	Individualplmnlist *[]UacAccesscategory1Selectionassistanceinfo `lb:2,ub:maxPLMN`
}
