package ies

import "rrc/utils"

// ConfiguredGrantConfig ::= SEQUENCE
// Extensible
type Configuredgrantconfig struct {
	Frequencyhopping                 *ConfiguredgrantconfigFrequencyhopping
	CgDmrsConfiguration              DmrsUplinkconfig
	McsTable                         *ConfiguredgrantconfigMcsTable
	McsTabletransformprecoder        *ConfiguredgrantconfigMcsTabletransformprecoder
	UciOnpusch                       *utils.Setuprelease[CgUciOnpusch]
	Resourceallocation               ConfiguredgrantconfigResourceallocation
	RbgSize                          *ConfiguredgrantconfigRbgSize
	Powercontrollooptouse            ConfiguredgrantconfigPowercontrollooptouse
	P0PuschAlpha                     P0PuschAlphasetid
	Transformprecoder                *ConfiguredgrantconfigTransformprecoder
	NrofharqProcesses                utils.INTEGER `lb:0,ub:16`
	Repk                             ConfiguredgrantconfigRepk
	RepkRv                           *ConfiguredgrantconfigRepkRv
	Periodicity                      ConfiguredgrantconfigPeriodicity
	Configuredgranttimer             *utils.INTEGER                                   `lb:0,ub:64`
	RrcConfigureduplinkgrant         *ConfiguredgrantconfigRrcConfigureduplinkgrant   `ext`
	CgRetransmissiontimerR16         *utils.INTEGER                                   `lb:0,ub:64,ext`
	CgMindfiDelayR16                 *ConfiguredgrantconfigCgMindfiDelayR16           `ext`
	CgNrofpuschInslotR16             *utils.INTEGER                                   `lb:0,ub:7,ext`
	CgNrofslotsR16                   *utils.INTEGER                                   `lb:0,ub:40,ext`
	CgStartingoffsetsR16             *CgStartingoffsetsR16                            `ext`
	CgUciMultiplexingR16             *ConfiguredgrantconfigCgUciMultiplexingR16       `ext`
	CgCotSharingoffsetR16            *utils.INTEGER                                   `lb:0,ub:39,ext`
	BetaoffsetcgUciR16               *utils.INTEGER                                   `lb:0,ub:31,ext`
	CgCotSharinglistR16              *[]CgCotSharingR16                               `lb:1,ub:1709,ext`
	HarqProcidOffsetR16              *utils.INTEGER                                   `lb:0,ub:15,ext`
	HarqProcidOffset2R16             *utils.INTEGER                                   `lb:0,ub:15,ext`
	ConfiguredgrantconfigindexR16    *ConfiguredgrantconfigindexR16                   `ext`
	ConfiguredgrantconfigindexmacR16 *ConfiguredgrantconfigindexmacR16                `ext`
	PeriodicityextR16                *utils.INTEGER                                   `lb:0,ub:5120,ext`
	Startingfromrv0R16               *ConfiguredgrantconfigStartingfromrv0R16         `ext`
	PhyPriorityindexR16              *ConfiguredgrantconfigPhyPriorityindexR16        `ext`
	AutonomoustxR16                  *ConfiguredgrantconfigAutonomoustxR16            `ext`
	CgBetaoffsetscrosspri0R17        *utils.Setuprelease[BetaoffsetscrosspriselcgR17] `ext`
	CgBetaoffsetscrosspri1R17        *utils.Setuprelease[BetaoffsetscrosspriselcgR17] `ext`
	MappingpatternR17                *ConfiguredgrantconfigMappingpatternR17          `ext`
	SequenceoffsetforrvR17           *utils.INTEGER                                   `lb:0,ub:3,ext`
	P0PuschAlpha2R17                 *P0PuschAlphasetid                               `ext`
	Powercontrollooptouse2R17        *ConfiguredgrantconfigPowercontrollooptouse2R17  `ext`
	CgCotSharinglistR17              *[]CgCotSharingR17                               `lb:1,ub:50722,ext`
	PeriodicityextR17                *utils.INTEGER                                   `lb:0,ub:40960,ext`
	RepkV1710                        *ConfiguredgrantconfigRepkV1710                  `ext`
	NrofharqProcessesV1700           *utils.INTEGER                                   `lb:0,ub:32,ext`
	HarqProcidOffset2V1700           *utils.INTEGER                                   `lb:0,ub:31,ext`
	ConfiguredgranttimerV1700        *utils.INTEGER                                   `lb:0,ub:288,ext`
	CgMindfiDelayV1710               *utils.INTEGER                                   `lb:0,ub:3584,ext`
	HarqProcidOffsetV1730            *utils.INTEGER                                   `lb:0,ub:31,ext`
	CgNrofslotsR17                   *utils.INTEGER                                   `lb:0,ub:320,ext`
}
