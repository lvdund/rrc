package ies

import "rrc/utils"

// ServingCellConfig-crs-RateMatch-PerCORESETPoolIndex-r16 ::= ENUMERATED
type ServingcellconfigCrsRatematchPercoresetpoolindexR16 struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigCrsRatematchPercoresetpoolindexR16EnumeratedNothing = iota
	ServingcellconfigCrsRatematchPercoresetpoolindexR16EnumeratedEnabled
)
