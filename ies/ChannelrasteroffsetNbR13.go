package ies

import "rrc/utils"

// ChannelRasterOffset-NB-r13 ::= ENUMERATED
type ChannelrasteroffsetNbR13 struct {
	Value utils.ENUMERATED
}

const (
	ChannelrasteroffsetNbR13Khz7dot5 = 0
	ChannelrasteroffsetNbR13Khz2dot5 = 1
	ChannelrasteroffsetNbR13Khz2dot5 = 2
	ChannelrasteroffsetNbR13Khz7dot5 = 3
)
