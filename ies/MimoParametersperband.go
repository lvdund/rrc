package ies

import "rrc/utils"

// MIMO-ParametersPerBand ::= SEQUENCE
// Extensible
type MimoParametersperband struct {
	TciStatepdsch                              *MimoParametersperbandTciStatepdsch
	AdditionalactivetciStatepdcch              *MimoParametersperbandAdditionalactivetciStatepdcch
	PuschTranscoherence                        *MimoParametersperbandPuschTranscoherence
	BeamcorrespondencewithoutulBeamsweeping    *MimoParametersperbandBeamcorrespondencewithoutulBeamsweeping
	Periodicbeamreport                         *MimoParametersperbandPeriodicbeamreport
	Aperiodicbeamreport                        *MimoParametersperbandAperiodicbeamreport
	SpBeamreportpucch                          *MimoParametersperbandSpBeamreportpucch
	SpBeamreportpusch                          *MimoParametersperbandSpBeamreportpusch
	Dummy1                                     *Dummyg
	Maxnumberrxbeam                            *utils.INTEGER `lb:0,ub:8`
	Maxnumberrxtxbeamswitchdl                  *MimoParametersperbandMaxnumberrxtxbeamswitchdl
	Maxnumbernongroupbeamreporting             *MimoParametersperbandMaxnumbernongroupbeamreporting
	Groupbeamreporting                         *MimoParametersperbandGroupbeamreporting
	Uplinkbeammanagement                       *MimoParametersperbandUplinkbeammanagement
	MaxnumbercsiRsBfd                          *utils.INTEGER `lb:0,ub:64`
	MaxnumberssbBfd                            *utils.INTEGER `lb:0,ub:64`
	MaxnumbercsiRsSsbCbd                       *utils.INTEGER `lb:0,ub:256`
	Dummy2                                     *MimoParametersperbandDummy2
	TwoportsptrsUl                             *MimoParametersperbandTwoportsptrsUl
	Dummy5                                     *SrsResources
	Dummy3                                     *utils.INTEGER `lb:0,ub:4`
	Beamreporttiming                           *MimoParametersperbandBeamreporttiming
	PtrsDensityrecommendationsetdl             *MimoParametersperbandPtrsDensityrecommendationsetdl
	PtrsDensityrecommendationsetul             *MimoParametersperbandPtrsDensityrecommendationsetul
	Dummy4                                     *Dummyh
	Aperiodictrs                               *MimoParametersperbandAperiodictrs
	Dummy6                                     *utils.BOOLEAN                                                   `ext`
	BeammanagementssbCsiRs                     *BeammanagementssbCsiRs                                          `ext`
	Beamswitchtiming                           *MimoParametersperbandBeamswitchtiming                           `ext`
	Codebookparameters                         *Codebookparameters                                              `ext`
	CsiRsImReceptionforfeedback                *CsiRsImReceptionforfeedback                                     `ext`
	CsiRsProcframeworkforsrs                   *CsiRsProcframeworkforsrs                                        `ext`
	CsiReportframework                         *CsiReportframework                                              `ext`
	CsiRsFortracking                           *CsiRsFortracking                                                `ext`
	SrsAssoccsiRs                              *[]SupportedcsiRsResource                                        `lb:1,ub:maxNrofCSIRsResources,ext`
	Spatialrelations                           *Spatialrelations                                                `ext`
	DefaultqclTwotciR16                        *MimoParametersperbandDefaultqclTwotciR16                        `ext`
	CodebookparametersperbandR16               *CodebookparametersV1610                                         `ext`
	SimulSpatialrelationupdatepucchresgroupR16 *MimoParametersperbandSimulSpatialrelationupdatepucchresgroupR16 `ext`
	MaxnumberscellbfrR16                       *MimoParametersperbandMaxnumberscellbfrR16                       `ext`
	SimultaneousreceptiondifftypedR16          *MimoParametersperbandSimultaneousreceptiondifftypedR16          `ext`
	SsbCsirsSinrMeasurementR16                 *MimoParametersperbandSsbCsirsSinrMeasurementR16                 `ext`
	NongroupsinrReportingR16                   *MimoParametersperbandNongroupsinrReportingR16                   `ext`
	GroupsinrReportingR16                      *MimoParametersperbandGroupsinrReportingR16                      `ext`
	MultidciMultitrpParametersR16              *MimoParametersperbandMultidciMultitrpParametersR16              `ext`
	SingledciSdmSchemeParametersR16            *MimoParametersperbandSingledciSdmSchemeParametersR16            `ext`
	SupportfdmSchemeaR16                       *MimoParametersperbandSupportfdmSchemeaR16                       `ext`
	SupportcodewordsoftcombiningR16            *MimoParametersperbandSupportcodewordsoftcombiningR16            `ext`
	SupporttdmSchemeaR16                       *MimoParametersperbandSupporttdmSchemeaR16                       `ext`
	SupportinterSlottdmR16                     *MimoParametersperbandSupportinterSlottdmR16                     `ext`
	LowpaprDmrsPdschR16                        *MimoParametersperbandLowpaprDmrsPdschR16                        `ext`
	LowpaprDmrsPuschwithoutprecodingR16        *MimoParametersperbandLowpaprDmrsPuschwithoutprecodingR16        `ext`
	LowpaprDmrsPucchR16                        *MimoParametersperbandLowpaprDmrsPucchR16                        `ext`
	LowpaprDmrsPuschwithprecodingR16           *MimoParametersperbandLowpaprDmrsPuschwithprecodingR16           `ext`
	CsiReportframeworkextR16                   *CsiReportframeworkextR16                                        `ext`
	CodebookparametersadditionR16              *CodebookparametersadditionR16                                   `ext`
	CodebookcomboparametersadditionR16         *CodebookcomboparametersadditionR16                              `ext`
	BeamcorrespondencessbBasedR16              *MimoParametersperbandBeamcorrespondencessbBasedR16              `ext`
	BeamcorrespondencecsiRsBasedR16            *MimoParametersperbandBeamcorrespondencecsiRsBasedR16            `ext`
	BeamswitchtimingR16                        *MimoParametersperbandBeamswitchtimingR16                        `ext`
	SemiPersistentl1SinrReportPucchR16         *MimoParametersperbandSemiPersistentl1SinrReportPucchR16         `ext`
	SemiPersistentl1SinrReportPuschR16         *MimoParametersperbandSemiPersistentl1SinrReportPuschR16         `ext`
	SpatialrelationsV1640                      *MimoParametersperbandSpatialrelationsV1640                      `ext`
	Support64candidatebeamrsBfrR16             *MimoParametersperbandSupport64candidatebeamrsBfrR16             `ext`
	MaxmimoLayersformultiDciMtrpR16            *MimoParametersperbandMaxmimoLayersformultiDciMtrpR16            `ext`
	SupportedsinrMeasV1670                     *utils.BITSTRING                                                 `lb:4,ub:4,ext`
	SrsIncreasedrepetitionR17                  *MimoParametersperbandSrsIncreasedrepetitionR17                  `ext`
	SrsPartialfrequencysoundingR17             *MimoParametersperbandSrsPartialfrequencysoundingR17             `ext`
	SrsStartrbLocationhoppingpartialR17        *MimoParametersperbandSrsStartrbLocationhoppingpartialR17        `ext`
	SrsCombeightR17                            *MimoParametersperbandSrsCombeightR17                            `ext`
	Codebookparametersfetype2R17               *Codebookparametersfetype2R17                                    `ext`
	MtrpPuschTwocsiRsR17                       *MimoParametersperbandMtrpPuschTwocsiRsR17                       `ext`
	MtrpPucchInterslotR17                      *MimoParametersperbandMtrpPucchInterslotR17                      `ext`
	MtrpPucchCyclicmappingR17                  *MimoParametersperbandMtrpPucchCyclicmappingR17                  `ext`
	MtrpPucchSecondtpcR17                      *MimoParametersperbandMtrpPucchSecondtpcR17                      `ext`
	MtrpBfrTwobfdRsSetR17                      *MimoParametersperbandMtrpBfrTwobfdRsSetR17                      `ext`
	MtrpBfrPucchSrPercgR17                     *MimoParametersperbandMtrpBfrPucchSrPercgR17                     `ext`
	MtrpBfrAssociationPucchSrR17               *MimoParametersperbandMtrpBfrAssociationPucchSrR17               `ext`
	SfnSimultwotciAcrossmulticcR17             *MimoParametersperbandSfnSimultwotciAcrossmulticcR17             `ext`
	SfnDefaultdlBeamsetupR17                   *MimoParametersperbandSfnDefaultdlBeamsetupR17                   `ext`
	SfnDefaultulBeamsetupR17                   *MimoParametersperbandSfnDefaultulBeamsetupR17                   `ext`
	SrsTriggeringoffsetR17                     *MimoParametersperbandSrsTriggeringoffsetR17                     `ext`
	SrsTriggeringdciR17                        *MimoParametersperbandSrsTriggeringdciR17                        `ext`
	CodebookcomboparametermixedtypeR17         *CodebookcomboparametermixedtypeR17                              `ext`
	UnifiedjointtciR17                         *MimoParametersperbandUnifiedjointtciR17                         `ext`
	UnifiedjointtciMultimacCeR17               *MimoParametersperbandUnifiedjointtciMultimacCeR17               `ext`
	UnifiedjointtciPerbwpCaR17                 *MimoParametersperbandUnifiedjointtciPerbwpCaR17                 `ext`
	UnifiedjointtciListsharingcaR17            *MimoParametersperbandUnifiedjointtciListsharingcaR17            `ext`
	UnifiedjointtciCommonmulticcR17            *MimoParametersperbandUnifiedjointtciCommonmulticcR17            `ext`
	UnifiedjointtciBeamaligndlrsR17            *MimoParametersperbandUnifiedjointtciBeamaligndlrsR17            `ext`
	UnifiedjointtciPcAssociationR17            *MimoParametersperbandUnifiedjointtciPcAssociationR17            `ext`
	UnifiedjointtciLegacyR17                   *MimoParametersperbandUnifiedjointtciLegacyR17                   `ext`
	UnifiedjointtciLegacySrsR17                *MimoParametersperbandUnifiedjointtciLegacySrsR17                `ext`
	UnifiedjointtciLegacyCoreset0R17           *MimoParametersperbandUnifiedjointtciLegacyCoreset0R17           `ext`
	UnifiedjointtciScellbfrR17                 *MimoParametersperbandUnifiedjointtciScellbfrR17                 `ext`
	UnifiedjointtciIntercellR17                *MimoParametersperbandUnifiedjointtciIntercellR17                `ext`
	UnifiedseparatetciR17                      *MimoParametersperbandUnifiedseparatetciR17                      `ext`
	UnifiedseparatetciMultimacCeR17            *MimoParametersperbandUnifiedseparatetciMultimacCeR17            `ext`
	UnifiedseparatetciPerbwpCaR17              *MimoParametersperbandUnifiedseparatetciPerbwpCaR17              `ext`
	UnifiedseparatetciListsharingcaR17         *MimoParametersperbandUnifiedseparatetciListsharingcaR17         `ext`
	UnifiedseparatetciCommonmulticcR17         *MimoParametersperbandUnifiedseparatetciCommonmulticcR17         `ext`
	UnifiedseparatetciIntercellR17             *MimoParametersperbandUnifiedseparatetciIntercellR17             `ext`
	UnifiedjointtciMtrpIntercellBmR17          *MimoParametersperbandUnifiedjointtciMtrpIntercellBmR17          `ext`
	MpeMitigationR17                           *MimoParametersperbandMpeMitigationR17                           `ext`
	SrsPortreportR17                           *MimoParametersperbandSrsPortreportR17                           `ext`
	MtrpPdcchIndividualR17                     *MimoParametersperbandMtrpPdcchIndividualR17                     `ext`
	MtrpPdcchAnyspan3symbolsR17                *MimoParametersperbandMtrpPdcchAnyspan3symbolsR17                `ext`
	MtrpPdcchTwoqclTypedR17                    *MimoParametersperbandMtrpPdcchTwoqclTypedR17                    `ext`
	MtrpPuschCsiRsR17                          *MimoParametersperbandMtrpPuschCsiRsR17                          `ext`
	MtrpPuschCyclicmappingR17                  *MimoParametersperbandMtrpPuschCyclicmappingR17                  `ext`
	MtrpPuschSecondtpcR17                      *MimoParametersperbandMtrpPuschSecondtpcR17                      `ext`
	MtrpPuschTwophrReportingR17                *MimoParametersperbandMtrpPuschTwophrReportingR17                `ext`
	MtrpPuschACsiR17                           *MimoParametersperbandMtrpPuschACsiR17                           `ext`
	MtrpPuschSpCsiR17                          *MimoParametersperbandMtrpPuschSpCsiR17                          `ext`
	MtrpPuschCgR17                             *MimoParametersperbandMtrpPuschCgR17                             `ext`
	MtrpPucchMacCeR17                          *MimoParametersperbandMtrpPucchMacCeR17                          `ext`
	MtrpPucchMaxnumPcFr1R17                    *utils.INTEGER                                                   `lb:0,ub:8,ext`
	MtrpInterCellR17                           *MimoParametersperbandMtrpInterCellR17                           `ext`
	MtrpGroupbasedl1RsrpR17                    *MimoParametersperbandMtrpGroupbasedl1RsrpR17                    `ext`
	MtrpBfdRsMacCeR17                          *MimoParametersperbandMtrpBfdRsMacCeR17                          `ext`
	MtrpCsiEnhancementperbandR17               *MimoParametersperbandMtrpCsiEnhancementperbandR17               `ext`
	CodebookcomboparametermultitrpR17          *CodebookcomboparametermultitrpR17                               `ext`
	MtrpCsiAdditionalcsiR17                    *MimoParametersperbandMtrpCsiAdditionalcsiR17                    `ext`
	MtrpCsiNMax2R17                            *MimoParametersperbandMtrpCsiNMax2R17                            `ext`
	MtrpCsiCmrR17                              *MimoParametersperbandMtrpCsiCmrR17                              `ext`
	SrsPartialfreqsoundingR17                  *MimoParametersperbandSrsPartialfreqsoundingR17                  `ext`
	BeamswitchtimingV1710                      *MimoParametersperbandBeamswitchtimingV1710                      `ext`
	BeamswitchtimingR17                        *MimoParametersperbandBeamswitchtimingR17                        `ext`
	BeamreporttimingV1710                      *MimoParametersperbandBeamreporttimingV1710                      `ext`
	MaxnumberrxtxbeamswitchdlV1710             *MimoParametersperbandMaxnumberrxtxbeamswitchdlV1710             `ext`
	SrsPortreportspApR17                       *MimoParametersperbandSrsPortreportspApR17                       `ext`
	MaxnumberrxbeamV1720                       *utils.INTEGER                                                   `lb:0,ub:12,ext`
	SfnImplicitrsTwotciR17                     *MimoParametersperbandSfnImplicitrsTwotciR17                     `ext`
	SfnQclTypedCollisionTwotciR17              *MimoParametersperbandSfnQclTypedCollisionTwotciR17              `ext`
	MtrpCsiNumcpuR17                           *MimoParametersperbandMtrpCsiNumcpuR17                           `ext`
	SupportrepnumpdschTdraDci12R17             *MimoParametersperbandSupportrepnumpdschTdraDci12R17             `ext`
}
