package ies

import "rrc/utils"

// SL-ConfiguredGrantConfig-r16 ::= SEQUENCE
// Extensible
type SlConfiguredgrantconfigR16 struct {
	SlConfigindexcgR16            SlConfigindexcgR16
	SlPeriodcgR16                 *SlPeriodcgR16
	SlNrofharqProcessesR16        *utils.INTEGER `lb:0,ub:16`
	SlHarqProcidOffsetR16         *utils.INTEGER `lb:0,ub:15`
	SlCgMaxtransnumlistR16        *SlCgMaxtransnumlistR16
	RrcConfiguredsidelinkgrantR16 *SlConfiguredgrantconfigR16RrcConfiguredsidelinkgrantR16
	SlN1pucchAnType2R16           *PucchResourceid `ext`
}
