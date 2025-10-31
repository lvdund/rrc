package ies

import "rrc/utils"

// BandSidelink-r16-fewerSymbolSlotSidelink-r16 ::= ENUMERATED
type BandsidelinkR16FewersymbolslotsidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16FewersymbolslotsidelinkR16EnumeratedNothing = iota
	BandsidelinkR16FewersymbolslotsidelinkR16EnumeratedSupported
)
