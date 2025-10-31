package ies

import "rrc/utils"

// SPS-Config ::= SEQUENCE
// Extensible
type SpsConfig struct {
	Periodicity               SpsConfigPeriodicity
	NrofharqProcesses         utils.INTEGER `lb:0,ub:8`
	N1pucchAn                 *PucchResourceid
	McsTable                  *SpsConfigMcsTable
	SpsConfigindexR16         *SpsConfigindexR16                  `ext`
	HarqProcidOffsetR16       *utils.INTEGER                      `lb:0,ub:15,ext`
	PeriodicityextR16         *utils.INTEGER                      `lb:0,ub:5120,ext`
	HarqCodebookidR16         *utils.INTEGER                      `lb:0,ub:2,ext`
	PdschAggregationfactorR16 *SpsConfigPdschAggregationfactorR16 `ext`
	SpsHarqDeferralR17        *utils.INTEGER                      `lb:0,ub:32,ext`
	N1pucchAnPucchsscellR17   *PucchResourceid                    `ext`
	PeriodicityextR17         *utils.INTEGER                      `lb:0,ub:40960,ext`
	NrofharqProcessesV1710    *utils.INTEGER                      `lb:0,ub:32,ext`
	HarqProcidOffsetV1700     *utils.INTEGER                      `lb:0,ub:31,ext`
}
