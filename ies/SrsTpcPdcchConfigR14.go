package ies

// SRS-TPC-PDCCH-Config-r14 ::= CHOICE
const (
	SrsTpcPdcchConfigR14ChoiceNothing = iota
	SrsTpcPdcchConfigR14ChoiceRelease
	SrsTpcPdcchConfigR14ChoiceSetup
)

type SrsTpcPdcchConfigR14 struct {
	Choice  uint64
	Release *struct{}
	Setup   *SrsTpcPdcchConfigR14Setup
}
