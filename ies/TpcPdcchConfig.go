package ies

// TPC-PDCCH-Config ::= CHOICE
const (
	TpcPdcchConfigChoiceNothing = iota
	TpcPdcchConfigChoiceRelease
	TpcPdcchConfigChoiceSetup
)

type TpcPdcchConfig struct {
	Choice  uint64
	Release *struct{}
	Setup   *TpcPdcchConfigSetup
}
