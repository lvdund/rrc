package ies

import "rrc/utils"

// SPS-ConfigUL-STTI-r15-setup ::= SEQUENCE
// Extensible
type SpsConfigulSttiR15Setup struct {
	SemipersistschedintervalulSttiR15          SpsConfigulSttiR15SetupSemipersistschedintervalulSttiR15
	Implicitreleaseafter                       SpsConfigulSttiR15SetupImplicitreleaseafter
	P0PersistentR15                            *SpsConfigulSttiR15SetupP0PersistentR15
	TwointervalsconfigR15                      *bool
	P0Persistentsubframeset2R15                *SpsConfigulSttiR15SetupP0Persistentsubframeset2R15
	NumberofconfulSpsProcessesSttiR15          *utils.INTEGER `lb:0,ub:12`
	SttiStarttimeulR15                         utils.INTEGER  `lb:0,ub:5`
	TpcPdcchConfigpuschSpsR15                  *TpcPdcchConfig
	CyclicshiftspsSttiR15                      *SpsConfigulSttiR15SetupCyclicshiftspsSttiR15
	IfdmaConfigSpsR15                          *utils.BOOLEAN
	HarqProcidOffsetR15                        *utils.INTEGER `lb:0,ub:15`
	RvSpsSttiUlRepetitionsR15                  *SpsConfigulSttiR15SetupRvSpsSttiUlRepetitionsR15
	SpsConfigindexR15                          *SpsConfigindexR15
	TbsScalingfactorsubslotspsUlRepetitionsR15 *SpsConfigulSttiR15SetupTbsScalingfactorsubslotspsUlRepetitionsR15
	TotalnumberpuschSpsSttiUlRepetitionsR15    *SpsConfigulSttiR15SetupTotalnumberpuschSpsSttiUlRepetitionsR15
}
