package ies

import "rrc/utils"

// PUSCH-ServingCellConfig-nrofHARQ-ProcessesForPUSCH-r17 ::= ENUMERATED
type PuschServingcellconfigNrofharqProcessesforpuschR17 struct {
	Value utils.ENUMERATED
}

const (
	PuschServingcellconfigNrofharqProcessesforpuschR17EnumeratedNothing = iota
	PuschServingcellconfigNrofharqProcessesforpuschR17EnumeratedN32
)
