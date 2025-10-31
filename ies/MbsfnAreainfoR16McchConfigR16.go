package ies

import "rrc/utils"

// MBSFN-AreaInfo-r16-mcch-Config-r16 ::= SEQUENCE
type MbsfnAreainfoR16McchConfigR16 struct {
	McchRepetitionperiodR16   MbsfnAreainfoR16McchConfigR16McchRepetitionperiodR16
	McchModificationperiodR16 MbsfnAreainfoR16McchConfigR16McchModificationperiodR16
	McchOffsetR16             utils.INTEGER   `lb:0,ub:10`
	SfAllocinfoR16            utils.BITSTRING `lb:10,ub:10`
	SignallingmcsR16          MbsfnAreainfoR16McchConfigR16SignallingmcsR16
}
