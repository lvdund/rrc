package ies

import "rrc/utils"

// MBSFN-AreaInfo-r9-mcch-Config-r9-mcch-ModificationPeriod-r9 ::= ENUMERATED
type MbsfnAreainfoR9McchConfigR9McchModificationperiodR9 struct {
	Value utils.ENUMERATED
}

const (
	MbsfnAreainfoR9McchConfigR9McchModificationperiodR9EnumeratedNothing = iota
	MbsfnAreainfoR9McchConfigR9McchModificationperiodR9EnumeratedRf512
	MbsfnAreainfoR9McchConfigR9McchModificationperiodR9EnumeratedRf1024
)
