package ies

// TPC-PDCCH-ConfigSCell-r13 ::= CHOICE
const (
	TpcPdcchConfigscellR13ChoiceNothing = iota
	TpcPdcchConfigscellR13ChoiceRelease
	TpcPdcchConfigscellR13ChoiceSetup
)

type TpcPdcchConfigscellR13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *TpcPdcchConfigscellR13Setup
}
