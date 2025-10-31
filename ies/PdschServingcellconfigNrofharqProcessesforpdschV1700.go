package ies

import "rrc/utils"

// PDSCH-ServingCellConfig-nrofHARQ-ProcessesForPDSCH-v1700 ::= ENUMERATED
type PdschServingcellconfigNrofharqProcessesforpdschV1700 struct {
	Value utils.ENUMERATED
}

const (
	PdschServingcellconfigNrofharqProcessesforpdschV1700EnumeratedNothing = iota
	PdschServingcellconfigNrofharqProcessesforpdschV1700EnumeratedN32
)
