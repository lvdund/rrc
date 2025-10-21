package ies

import "rrc/utils"

// RSS-MeasPowerBias-r16 ::= ENUMERATED
type RssMeaspowerbiasR16 struct {
	Value utils.ENUMERATED
}

const (
	RssMeaspowerbiasR16Db6        = 0
	RssMeaspowerbiasR16Db3        = 1
	RssMeaspowerbiasR16Db0        = 2
	RssMeaspowerbiasR16Db3        = 3
	RssMeaspowerbiasR16Db6        = 4
	RssMeaspowerbiasR16Db9        = 5
	RssMeaspowerbiasR16Db12       = 6
	RssMeaspowerbiasR16Rssnotused = 7
)
