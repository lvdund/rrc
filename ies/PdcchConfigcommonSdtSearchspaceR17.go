package ies

// PDCCH-ConfigCommon-sdt-SearchSpace-r17 ::= CHOICE
const (
	PdcchConfigcommonSdtSearchspaceR17ChoiceNothing = iota
	PdcchConfigcommonSdtSearchspaceR17ChoiceNewsearchspace
	PdcchConfigcommonSdtSearchspaceR17ChoiceExistingsearchspace
)

type PdcchConfigcommonSdtSearchspaceR17 struct {
	Choice              uint64
	Newsearchspace      *Searchspace
	Existingsearchspace *Searchspaceid
}
