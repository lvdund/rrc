package ies

import "rrc/utils"

// BandSidelink-r16-lowSE-64QAM-MCS-TableSidelink-r16 ::= ENUMERATED
type BandsidelinkR16Lowse64qamMcsTablesidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	BandsidelinkR16Lowse64qamMcsTablesidelinkR16EnumeratedNothing = iota
	BandsidelinkR16Lowse64qamMcsTablesidelinkR16EnumeratedSupported
)
