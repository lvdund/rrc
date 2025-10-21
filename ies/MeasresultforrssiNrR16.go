package ies

import "rrc/utils"

// MeasResultForRSSI-NR-r16 ::= SEQUENCE
// Extensible
type MeasresultforrssiNrR16 struct {
	RssiResultnrR16       RssiRangeR13
	ChanneloccupancynrR16 utils.INTEGER
}
