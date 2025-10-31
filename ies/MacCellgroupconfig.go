package ies

import "rrc/utils"

// MAC-CellGroupConfig ::= SEQUENCE
// Extensible
type MacCellgroupconfig struct {
	DrxConfig                          *utils.Setuprelease[DrxConfig]
	Schedulingrequestconfig            *Schedulingrequestconfig
	BsrConfig                          *BsrConfig
	TagConfig                          *TagConfig
	PhrConfig                          *utils.Setuprelease[PhrConfig]
	Skipuplinktxdynamic                utils.BOOLEAN
	CsiMask                            *utils.BOOLEAN                                  `ext`
	Datainactivitytimer                *utils.Setuprelease[Datainactivitytimer]        `ext`
	UseprebsrR16                       *utils.BOOLEAN                                  `ext`
	SchedulingrequestidLbtScellR16     *Schedulingrequestid                            `ext`
	LchBasedprioritizationR16          *MacCellgroupconfigLchBasedprioritizationR16    `ext`
	SchedulingrequestidBfrScellR16     *Schedulingrequestid                            `ext`
	DrxConfigsecondarygroupR16         *utils.Setuprelease[DrxConfigsecondarygroupR16] `ext`
	EnhancedskipuplinktxdynamicR16     *utils.BOOLEAN                                  `ext`
	EnhancedskipuplinktxconfiguredR16  *utils.BOOLEAN                                  `ext`
	IntracgPrioritizationR17           *MacCellgroupconfigIntracgPrioritizationR17     `ext`
	DrxConfigslR17                     *utils.Setuprelease[DrxConfigslR17]             `ext`
	DrxConfigextV1700                  *utils.Setuprelease[DrxConfigextV1700]          `ext`
	SchedulingrequestidBfrR17          *Schedulingrequestid                            `ext`
	SchedulingrequestidBfr2R17         *Schedulingrequestid                            `ext`
	SchedulingrequestconfigV1700       *SchedulingrequestconfigV1700                   `ext`
	TarConfigR17                       *utils.Setuprelease[TarConfigR17]               `ext`
	GRntiConfigtoaddmodlistR17         *[]MbsRntiSpecificconfigR17                     `lb:1,ub:maxGRntiR17,ext`
	GRntiConfigtoreleaselistR17        *[]MbsRntiSpecificconfigidR17                   `lb:1,ub:maxGRntiR17,ext`
	GCsRntiConfigtoaddmodlistR17       *[]MbsRntiSpecificconfigR17                     `lb:1,ub:maxGCsRntiR17,ext`
	GCsRntiConfigtoreleaselistR17      *[]MbsRntiSpecificconfigidR17                   `lb:1,ub:maxGCsRntiR17,ext`
	AllowcsiSrsTxMulticastdrxActiveR17 *utils.BOOLEAN                                  `ext`
	SchedulingrequestidPosmgRequestR17 *Schedulingrequestid                            `ext`
	DrxLasttransmissionulR17           *MacCellgroupconfigDrxLasttransmissionulR17     `ext`
}
