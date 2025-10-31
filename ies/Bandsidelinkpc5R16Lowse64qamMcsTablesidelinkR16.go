package ies

import "rrc/utils"

// BandSidelinkPC5-r16-lowSE-64QAM-MCS-TableSidelink-r16 ::= ENUMERATED
type Bandsidelinkpc5R16Lowse64qamMcsTablesidelinkR16 struct {
	Value utils.ENUMERATED
}

const (
	Bandsidelinkpc5R16Lowse64qamMcsTablesidelinkR16EnumeratedNothing = iota
	Bandsidelinkpc5R16Lowse64qamMcsTablesidelinkR16EnumeratedSupported
)
