package ies

import "rrc/utils"

// MBSFN-AreaInfo-r9-mcch-Config-r9-mcch-RepetitionPeriod-r9 ::= ENUMERATED
type MbsfnAreainfoR9McchConfigR9McchRepetitionperiodR9 struct {
	Value utils.ENUMERATED
}

const (
	MbsfnAreainfoR9McchConfigR9McchRepetitionperiodR9EnumeratedNothing = iota
	MbsfnAreainfoR9McchConfigR9McchRepetitionperiodR9EnumeratedRf32
	MbsfnAreainfoR9McchConfigR9McchRepetitionperiodR9EnumeratedRf64
	MbsfnAreainfoR9McchConfigR9McchRepetitionperiodR9EnumeratedRf128
	MbsfnAreainfoR9McchConfigR9McchRepetitionperiodR9EnumeratedRf256
)
