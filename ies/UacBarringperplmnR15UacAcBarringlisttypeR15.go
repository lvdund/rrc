package ies

// UAC-BarringPerPLMN-r15-uac-AC-BarringListType-r15 ::= CHOICE
const (
	UacBarringperplmnR15UacAcBarringlisttypeR15ChoiceNothing = iota
	UacBarringperplmnR15UacAcBarringlisttypeR15ChoiceUacImplicitacBarringlistR15
	UacBarringperplmnR15UacAcBarringlisttypeR15ChoiceUacExplicitacBarringlistR15
)

type UacBarringperplmnR15UacAcBarringlisttypeR15 struct {
	Choice                      uint64
	UacImplicitacBarringlistR15 *[]UacBarringinfosetindexR15 `lb:maxAccessCat1R15,ub:maxAccessCat1R15`
	UacExplicitacBarringlistR15 *UacBarringpercatlistR15
}
