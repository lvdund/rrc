package ies

import "rrc/utils"

// ChannelRasterOffset-NB-r13 ::= ENUMERATED
type ChannelrasteroffsetNbR13 struct {
	Value utils.ENUMERATED
}

const (
	ChannelrasteroffsetNbR13EnumeratedNothing = iota
	ChannelrasteroffsetNbR13EnumeratedKhz_7dot5
	ChannelrasteroffsetNbR13EnumeratedKhz_2dot5
	ChannelrasteroffsetNbR13EnumeratedKhz2dot5
	ChannelrasteroffsetNbR13EnumeratedKhz7dot5
)
