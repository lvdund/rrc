package ies

import "rrc/utils"

// MeasParameters-v16c0-nr-CellIndividualOffset-r16 ::= ENUMERATED
type MeasparametersV16c0NrCellindividualoffsetR16 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV16c0NrCellindividualoffsetR16EnumeratedNothing = iota
	MeasparametersV16c0NrCellindividualoffsetR16EnumeratedSupported
)
