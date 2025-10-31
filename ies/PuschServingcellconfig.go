package ies

import "rrc/utils"

// PUSCH-ServingCellConfig ::= SEQUENCE
// Extensible
type PuschServingcellconfig struct {
	Codeblockgrouptransmission   *utils.Setuprelease[PuschCodeblockgrouptransmission]
	Ratematching                 *PuschServingcellconfigRatematching
	Xoverhead                    *PuschServingcellconfigXoverhead
	MaxmimoLayers                *utils.INTEGER                                      `lb:0,ub:4,ext`
	Processingtype2enabled       *utils.BOOLEAN                                      `ext`
	MaxmimoLayersdci02R16        *utils.Setuprelease[MaxmimoLayersdci02R16]          `ext`
	NrofharqProcessesforpuschR17 *PuschServingcellconfigNrofharqProcessesforpuschR17 `ext`
	UplinkharqModeR17            *utils.Setuprelease[UplinkharqModeR17]              `ext`
}
