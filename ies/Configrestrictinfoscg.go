package ies

import "rrc/utils"

// ConfigRestrictInfoSCG ::= SEQUENCE
// Extensible
type Configrestrictinfoscg struct {
	AllowedbcListmrdc                     *Bandcombinationinfolist
	PowercoordinationFr1                  *ConfigrestrictinfoscgPowercoordinationFr1
	Servcellindexrangescg                 *ConfigrestrictinfoscgServcellindexrangescg
	Maxmeasfreqsscg                       *utils.INTEGER                                `lb:0,ub:maxMeasFreqsMN`
	Dummy                                 *utils.INTEGER                                `lb:0,ub:maxMeasIdentitiesMN`
	Selectedbandentriesmnlist             *[]Selectedbandentriesmn                      `lb:1,ub:maxBandComb,ext`
	PdcchBlinddetectionscg                *utils.INTEGER                                `lb:0,ub:15,ext`
	MaxnumberrohcContextsessionssn        *utils.INTEGER                                `lb:0,ub:16384,ext`
	Maxintrafreqmeasidentitiesscg         *utils.INTEGER                                `lb:0,ub:maxMeasIdentitiesMN,ext`
	Maxinterfreqmeasidentitiesscg         *utils.INTEGER                                `lb:0,ub:maxMeasIdentitiesMN,ext`
	PMaxnrFr1McgR16                       *PMax                                         `ext`
	PowercoordinationFr2R16               *ConfigrestrictinfoscgPowercoordinationFr2R16 `ext`
	NrdcPcModeFr1R16                      *ConfigrestrictinfoscgNrdcPcModeFr1R16        `ext`
	NrdcPcModeFr2R16                      *ConfigrestrictinfoscgNrdcPcModeFr2R16        `ext`
	MaxmeassrsResourcescgR16              *utils.INTEGER                                `lb:0,ub:maxNrofCLISrsResourcesR16,ext`
	MaxmeascliResourcescgR16              *utils.INTEGER                                `lb:0,ub:maxNrofCLIRssiResourcesR16,ext`
	MaxnumberehcContextssnR16             *utils.INTEGER                                `lb:0,ub:65536,ext`
	AllowedreducedconfigforoverheatingR16 *Overheatingassistance                        `ext`
	MaxtoffsetR16                         *TOffsetR16                                   `ext`
	AllowedreducedconfigforoverheatingR17 *OverheatingassistanceR17                     `ext`
	MaxnumberudcDrbR17                    *utils.INTEGER                                `lb:0,ub:2,ext`
	MaxnumbercpccandidatesR17             *utils.INTEGER                                `lb:0,ub:maxNrofCondCells1R17,ext`
}
