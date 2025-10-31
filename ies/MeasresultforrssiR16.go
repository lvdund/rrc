package ies

import "rrc/utils"

// MeasResultForRSSI-r16 ::= SEQUENCE
type MeasresultforrssiR16 struct {
	RssiResultR16       RssiRangeR16
	ChanneloccupancyR16 utils.INTEGER `lb:0,ub:100`
}
