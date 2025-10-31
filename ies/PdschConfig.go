package ies

import "rrc/utils"

// PDSCH-Config ::= SEQUENCE
// Extensible
type PdschConfig struct {
	Datascramblingidentitypdsch                       *utils.INTEGER `lb:0,ub:1023`
	DmrsDownlinkforpdschMappingtypea                  *utils.Setuprelease[DmrsDownlinkconfig]
	DmrsDownlinkforpdschMappingtypeb                  *utils.Setuprelease[DmrsDownlinkconfig]
	TciStatestoaddmodlist                             *[]TciState   `lb:1,ub:maxNrofTCIStates`
	TciStatestoreleaselist                            *[]TciStateid `lb:1,ub:maxNrofTCIStates`
	VrbToprbInterleaver                               *PdschConfigVrbToprbInterleaver
	Resourceallocation                                PdschConfigResourceallocation
	PdschTimedomainallocationlist                     *utils.Setuprelease[PdschTimedomainresourceallocationlist]
	PdschAggregationfactor                            *PdschConfigPdschAggregationfactor
	Ratematchpatterntoaddmodlist                      *[]Ratematchpattern   `lb:1,ub:maxNrofRateMatchPatterns`
	Ratematchpatterntoreleaselist                     *[]Ratematchpatternid `lb:1,ub:maxNrofRateMatchPatterns`
	Ratematchpatterngroup1                            *Ratematchpatterngroup
	Ratematchpatterngroup2                            *Ratematchpatterngroup
	RbgSize                                           PdschConfigRbgSize
	McsTable                                          *PdschConfigMcsTable
	Maxnrofcodewordsscheduledbydci                    *PdschConfigMaxnrofcodewordsscheduledbydci
	PrbBundlingtype                                   *PdschConfigPrbBundlingtype
	ZpCsiRsResourcetoaddmodlist                       *[]ZpCsiRsResource      `lb:1,ub:maxNrofZPCsiRsResources`
	ZpCsiRsResourcetoreleaselist                      *[]ZpCsiRsResourceid    `lb:1,ub:maxNrofZPCsiRsResources`
	AperiodicZpCsiRsResourcesetstoaddmodlist          *[]ZpCsiRsResourceset   `lb:1,ub:maxNrofZPCsiRsResourcesets`
	AperiodicZpCsiRsResourcesetstoreleaselist         *[]ZpCsiRsResourcesetid `lb:1,ub:maxNrofZPCsiRsResourcesets`
	SpZpCsiRsResourcesetstoaddmodlist                 *[]ZpCsiRsResourceset   `lb:1,ub:maxNrofZPCsiRsResourcesets`
	SpZpCsiRsResourcesetstoreleaselist                *[]ZpCsiRsResourcesetid `lb:1,ub:maxNrofZPCsiRsResourcesets`
	PZpCsiRsResourceset                               *utils.Setuprelease[ZpCsiRsResourceset]
	MaxmimoLayersR16                                  *utils.Setuprelease[MaxmimoLayersdlR16]                       `ext`
	Minimumschedulingoffsetk0R16                      *utils.Setuprelease[Minschedulingoffsetk0ValuesR16]           `ext`
	Antennaportsfieldpresencedci12R16                 *PdschConfigAntennaportsfieldpresencedci12R16                 `ext`
	AperiodiczpCsiRsResourcesetstoaddmodlistdci12R16  *[]ZpCsiRsResourceset                                         `lb:1,ub:maxNrofZPCsiRsResourcesets,ext`
	AperiodiczpCsiRsResourcesetstoreleaselistdci12R16 *[]ZpCsiRsResourcesetid                                       `lb:1,ub:maxNrofZPCsiRsResourcesets,ext`
	DmrsDownlinkforpdschMappingtypeaDci12R16          *utils.Setuprelease[DmrsDownlinkconfig]                       `ext`
	DmrsDownlinkforpdschMappingtypebDci12R16          *utils.Setuprelease[DmrsDownlinkconfig]                       `ext`
	DmrsSequenceinitializationdci12R16                *PdschConfigDmrsSequenceinitializationdci12R16                `ext`
	HarqProcessnumbersizedci12R16                     *utils.INTEGER                                                `lb:0,ub:4,ext`
	McsTabledci12R16                                  *PdschConfigMcsTabledci12R16                                  `ext`
	NumberofbitsforrvDci12R16                         *utils.INTEGER                                                `lb:0,ub:2,ext`
	PdschTimedomainallocationlistdci12R16             *utils.Setuprelease[PdschTimedomainresourceallocationlistR16] `ext`
	PrbBundlingtypedci12R16                           *PdschConfigPrbBundlingtypedci12R16                           `ext`
	Priorityindicatordci12R16                         *PdschConfigPriorityindicatordci12R16                         `ext`
	Ratematchpatterngroup1dci12R16                    *Ratematchpatterngroup                                        `ext`
	Ratematchpatterngroup2dci12R16                    *Ratematchpatterngroup                                        `ext`
	Resourceallocationtype1granularitydci12R16        *PdschConfigResourceallocationtype1granularitydci12R16        `ext`
	VrbToprbInterleaverdci12R16                       *PdschConfigVrbToprbInterleaverdci12R16                       `ext`
	Referenceofslivdci12R16                           *PdschConfigReferenceofslivdci12R16                           `ext`
	Resourceallocationdci12R16                        *PdschConfigResourceallocationdci12R16                        `ext`
	Priorityindicatordci11R16                         *PdschConfigPriorityindicatordci11R16                         `ext`
	Datascramblingidentitypdsch2R16                   *utils.INTEGER                                                `lb:0,ub:1023,ext`
	PdschTimedomainallocationlistR16                  *utils.Setuprelease[PdschTimedomainresourceallocationlistR16] `ext`
	RepetitionschemeconfigR16                         *utils.Setuprelease[RepetitionschemeconfigR16]                `ext`
	RepetitionschemeconfigV1630                       *utils.Setuprelease[RepetitionschemeconfigV1630]              `ext`
	PdschHarqAckOneshotfeedbackdci12R17               *PdschConfigPdschHarqAckOneshotfeedbackdci12R17               `ext`
	PdschHarqAckEnhtype3dci12R17                      *PdschConfigPdschHarqAckEnhtype3dci12R17                      `ext`
	PdschHarqAckEnhtype3dciField12R17                 *PdschConfigPdschHarqAckEnhtype3dciField12R17                 `ext`
	PdschHarqAckRetxdci12R17                          *PdschConfigPdschHarqAckRetxdci12R17                          `ext`
	PucchSscelldyndci12R17                            *PdschConfigPucchSscelldyndci12R17                            `ext`
	DlOrjointtciStatelistR17                          *PdschConfigDlOrjointtciStatelistR17                          `ext`
	BeamapptimeR17                                    *PdschConfigBeamapptimeR17                                    `ext`
	Dummy                                             *utils.Setuprelease[DummyTdraList]                            `ext`
	DmrsFdOccDisabledforrank1PdschR17                 *utils.BOOLEAN                                                `ext`
	Minimumschedulingoffsetk0R17                      *utils.Setuprelease[Minschedulingoffsetk0ValuesR17]           `ext`
	HarqProcessnumbersizedci12V1700                   *utils.INTEGER                                                `lb:0,ub:5,ext`
	HarqProcessnumbersizedci11R17                     *utils.INTEGER                                                `lb:0,ub:5,ext`
	McsTableR17                                       *PdschConfigMcsTableR17                                       `ext`
	McsTabledci12R17                                  *PdschConfigMcsTabledci12R17                                  `ext`
	XoverheadmulticastR17                             *PdschConfigXoverheadmulticastR17                             `ext`
	Priorityindicatordci42R17                         *PdschConfigPriorityindicatordci42R17                         `ext`
	Sizedci42R17                                      *utils.INTEGER                                                `lb:0,ub:maxDCI42SizeR17,ext`
	PdschTimedomainallocationlistformultipdschR17     *utils.Setuprelease[MultipdschTdraListR17]                    `ext`
}
