package ies

import "rrc/utils"

// RSS-MeasPowerBias-r16 ::= ENUMERATED
type RssMeaspowerbiasR16 struct {
	Value utils.ENUMERATED
}

const (
	RssMeaspowerbiasR16EnumeratedNothing = iota
	RssMeaspowerbiasR16EnumeratedDb_6
	RssMeaspowerbiasR16EnumeratedDb_3
	RssMeaspowerbiasR16EnumeratedDb0
	RssMeaspowerbiasR16EnumeratedDb3
	RssMeaspowerbiasR16EnumeratedDb6
	RssMeaspowerbiasR16EnumeratedDb9
	RssMeaspowerbiasR16EnumeratedDb12
	RssMeaspowerbiasR16EnumeratedRssnotused
)
