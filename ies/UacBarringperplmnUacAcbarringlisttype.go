package ies

// UAC-BarringPerPLMN-uac-ACBarringListType ::= CHOICE
const (
	UacBarringperplmnUacAcbarringlisttypeChoiceNothing = iota
	UacBarringperplmnUacAcbarringlisttypeChoiceUacImplicitacbarringlist
	UacBarringperplmnUacAcbarringlisttypeChoiceUacExplicitacbarringlist
)

type UacBarringperplmnUacAcbarringlisttype struct {
	Choice                   uint64
	UacImplicitacbarringlist *[]UacBarringinfosetindex `lb:maxAccessCat1,ub:maxAccessCat1`
	UacExplicitacbarringlist *UacBarringpercatlist
}
