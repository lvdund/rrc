package ies

import "rrc/utils"

// MBSFN-AreaInfo-r9-mcch-Config-r9 ::= SEQUENCE
type MbsfnAreainfoR9McchConfigR9 struct {
	McchRepetitionperiodR9   MbsfnAreainfoR9McchConfigR9McchRepetitionperiodR9
	McchOffsetR9             utils.INTEGER `lb:0,ub:10`
	McchModificationperiodR9 MbsfnAreainfoR9McchConfigR9McchModificationperiodR9
	SfAllocinfoR9            utils.BITSTRING `lb:6,ub:6`
	SignallingmcsR9          MbsfnAreainfoR9McchConfigR9SignallingmcsR9
}
