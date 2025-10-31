package ies

import "rrc/utils"

// PUCCH-Config ::= SEQUENCE
// Extensible
type PucchConfig struct {
	Resourcesettoaddmodlist                       *[]PucchResourceset   `lb:1,ub:maxNrofPUCCHResourcesets`
	Resourcesettoreleaselist                      *[]PucchResourcesetid `lb:1,ub:maxNrofPUCCHResourcesets`
	Resourcetoaddmodlist                          *[]PucchResource      `lb:1,ub:maxNrofPUCCHResources`
	Resourcetoreleaselist                         *[]PucchResourceid    `lb:1,ub:maxNrofPUCCHResources`
	Format1                                       *utils.Setuprelease[PucchFormatconfig]
	Format2                                       *utils.Setuprelease[PucchFormatconfig]
	Format3                                       *utils.Setuprelease[PucchFormatconfig]
	Format4                                       *utils.Setuprelease[PucchFormatconfig]
	Schedulingrequestresourcetoaddmodlist         *[]Schedulingrequestresourceconfig `lb:1,ub:maxNrofSRResources`
	Schedulingrequestresourcetoreleaselist        *[]Schedulingrequestresourceid     `lb:1,ub:maxNrofSRResources`
	MultiCsiPucchResourcelist                     *[]PucchResourceid                 `lb:1,ub:2`
	DlDatatoulAck                                 *[]utils.INTEGER                   `lb:1,ub:8`
	Spatialrelationinfotoaddmodlist               *[]PucchSpatialrelationinfo        `lb:1,ub:maxNrofSpatialRelationInfos`
	Spatialrelationinfotoreleaselist              *[]PucchSpatialrelationinfoid      `lb:1,ub:maxNrofSpatialRelationInfos`
	PucchPowercontrol                             *PucchPowercontrol
	ResourcetoaddmodlistextV1610                  *[]PucchResourceextV1610                                  `lb:1,ub:maxNrofPUCCHResources,ext`
	DlDatatoulAckR16                              *utils.Setuprelease[DlDatatoulAckR16]                     `ext`
	UlAccessconfiglistdci11R16                    *utils.Setuprelease[UlAccessconfiglistdci11R16]           `ext`
	SubslotlengthforpucchR16                      *PucchConfigSubslotlengthforpucchR16                      `ext`
	DlDatatoulAckDci12R16                         *utils.Setuprelease[DlDatatoulAckDci12R16]                `ext`
	NumberofbitsforpucchResourceindicatordci12R16 *utils.INTEGER                                            `lb:0,ub:3,ext`
	DmrsUplinktransformprecodingpucchR16          *PucchConfigDmrsUplinktransformprecodingpucchR16          `ext`
	SpatialrelationinfotoaddmodlistsizeextV1610   *[]PucchSpatialrelationinfo                               `lb:1,ub:maxNrofSpatialRelationInfosDiffR16,ext`
	SpatialrelationinfotoreleaselistsizeextV1610  *[]PucchSpatialrelationinfoid                             `lb:1,ub:maxNrofSpatialRelationInfosDiffR16,ext`
	SpatialrelationinfotoaddmodlistextV1610       *[]PucchSpatialrelationinfoextR16                         `lb:1,ub:maxNrofSpatialRelationInfosR16,ext`
	SpatialrelationinfotoreleaselistextV1610      *[]PucchSpatialrelationinfoidR16                          `lb:1,ub:maxNrofSpatialRelationInfosR16,ext`
	ResourcegrouptoaddmodlistR16                  *[]PucchResourcegroupR16                                  `lb:1,ub:maxNrofPUCCHResourcegroupsR16,ext`
	ResourcegrouptoreleaselistR16                 *[]PucchResourcegroupidR16                                `lb:1,ub:maxNrofPUCCHResourcegroupsR16,ext`
	SpsPucchAnListR16                             *utils.Setuprelease[SpsPucchAnListR16]                    `ext`
	SchedulingrequestresourcetoaddmodlistextV1610 *[]SchedulingrequestresourceconfigextV1610                `lb:1,ub:maxNrofSRResources,ext`
	Format0R17                                    *utils.Setuprelease[PucchFormatconfig]                    `ext`
	Format2extR17                                 *utils.Setuprelease[PucchFormatconfigextR17]              `ext`
	Format3extR17                                 *utils.Setuprelease[PucchFormatconfigextR17]              `ext`
	Format4extR17                                 *utils.Setuprelease[PucchFormatconfigextR17]              `ext`
	UlAccessconfiglistdci12R17                    *utils.Setuprelease[UlAccessconfiglistdci12R17]           `ext`
	MappingpatternR17                             *PucchConfigMappingpatternR17                             `ext`
	PowercontrolsetinfotoaddmodlistR17            *[]PucchPowercontrolsetinfoR17                            `lb:1,ub:maxNrofPowerControlSetInfosR17,ext`
	PowercontrolsetinfotoreleaselistR17           *[]PucchPowercontrolsetinfoidR17                          `lb:1,ub:maxNrofPowerControlSetInfosR17,ext`
	Secondtpcfielddci11R17                        *PucchConfigSecondtpcfielddci11R17                        `ext`
	Secondtpcfielddci12R17                        *PucchConfigSecondtpcfielddci12R17                        `ext`
	DlDatatoulAckR17                              *utils.Setuprelease[DlDatatoulAckR17]                     `ext`
	DlDatatoulAckDci12R17                         *utils.Setuprelease[DlDatatoulAckDci12R17]                `ext`
	UlAccessconfiglistdci11R17                    *utils.Setuprelease[UlAccessconfiglistdci11R17]           `ext`
	SchedulingrequestresourcetoaddmodlistextV1700 *[]SchedulingrequestresourceconfigextV1700                `lb:1,ub:maxNrofSRResources,ext`
	DmrsBundlingpucchConfigR17                    *utils.Setuprelease[DmrsBundlingpucchConfigR17]           `ext`
	DlDatatoulAckV1700                            *utils.Setuprelease[DlDatatoulAckV1700]                   `ext`
	DlDatatoulAckMulticastdciFormat41R17          *utils.Setuprelease[DlDatatoulAckMulticastdciFormat41R17] `ext`
	SpsPucchAnListmulticastR17                    *utils.Setuprelease[SpsPucchAnListR16]                    `ext`
}
