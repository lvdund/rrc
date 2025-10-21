package ies

import "rrc/utils"

// MeasResultForRSSI-r13 ::= SEQUENCE
// Extensible
type MeasresultforrssiR13 struct {
	RssiResultR13       RssiRangeR13
	ChanneloccupancyR13 utils.INTEGER
}
