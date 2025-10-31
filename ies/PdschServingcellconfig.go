package ies

import "rrc/utils"

// PDSCH-ServingCellConfig ::= SEQUENCE
// Extensible
type PdschServingcellconfig struct {
	Codeblockgrouptransmission             *utils.Setuprelease[PdschCodeblockgrouptransmission]
	Xoverhead                              *PdschServingcellconfigXoverhead
	NrofharqProcessesforpdsch              *PdschServingcellconfigNrofharqProcessesforpdsch
	PucchCell                              *Servcellindex
	MaxmimoLayers                          *utils.INTEGER                                              `lb:0,ub:8,ext`
	Processingtype2enabled                 *utils.BOOLEAN                                              `ext`
	PdschCodeblockgrouptransmissionlistR16 *utils.Setuprelease[PdschCodeblockgrouptransmissionlistR16] `ext`
	DownlinkharqFeedbackdisabledR17        *utils.Setuprelease[DownlinkharqFeedbackdisabledR17]        `ext`
	NrofharqProcessesforpdschV1700         *PdschServingcellconfigNrofharqProcessesforpdschV1700       `ext`
}
