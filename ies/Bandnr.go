package ies

import "rrc/utils"

// BandNR ::= SEQUENCE
// Extensible
type Bandnr struct {
	Bandnr                                                       Freqbandindicatornr
	ModifiedmprBehaviour                                         *utils.BITSTRING `lb:8,ub:8`
	MimoParametersperband                                        *MimoParametersperband
	Extendedcp                                                   *BandnrExtendedcp
	Multipletci                                                  *BandnrMultipletci
	BwpWithoutrestriction                                        *BandnrBwpWithoutrestriction
	BwpSamenumerology                                            *BandnrBwpSamenumerology
	BwpDiffnumerology                                            *BandnrBwpDiffnumerology
	CrosscarrierschedulingSamescs                                *BandnrCrosscarrierschedulingSamescs
	Pdsch256qamFr2                                               *BandnrPdsch256qamFr2
	Pusch256qam                                                  *BandnrPusch256qam
	UePowerclass                                                 *BandnrUePowerclass
	RatematchinglteCrs                                           *BandnrRatematchinglteCrs
	ChannelbwsDl                                                 *BandnrChannelbwsDl
	ChannelbwsUl                                                 *BandnrChannelbwsUl
	MaxuplinkdutycyclePc2Fr1                                     *BandnrMaxuplinkdutycyclePc2Fr1                                     `ext`
	PucchSpatialrelinfomacCe                                     *BandnrPucchSpatialrelinfomacCe                                     `ext`
	PowerboostingPi2bpsk                                         *BandnrPowerboostingPi2bpsk                                         `ext`
	MaxuplinkdutycycleFr2                                        *BandnrMaxuplinkdutycycleFr2                                        `ext`
	ChannelbwsDlV1590                                            *BandnrChannelbwsDlV1590                                            `ext`
	ChannelbwsUlV1590                                            *BandnrChannelbwsUlV1590                                            `ext`
	Asymmetricbandwidthcombinationset                            *utils.BITSTRING                                                    `lb:1,ub:32,ext`
	SharedspectrumchaccessparamsperbandR16                       *SharedspectrumchaccessparamsperbandR16                             `ext`
	CanceloverlappingpuschR16                                    *BandnrCanceloverlappingpuschR16                                    `ext`
	MultipleratematchingeutraCrsR16                              *BandnrMultipleratematchingeutraCrsR16                              `ext`
	OverlapratematchingeutraCrsR16                               *BandnrOverlapratematchingeutraCrsR16                               `ext`
	PdschMappingtypebAltR16                                      *BandnrPdschMappingtypebAltR16                                      `ext`
	OneslotperiodictrsR16                                        *BandnrOneslotperiodictrsR16                                        `ext`
	OlpcSrsPosR16                                                *OlpcSrsPosR16                                                      `ext`
	SpatialrelationssrsPosR16                                    *SpatialrelationssrsPosR16                                          `ext`
	SimulsrsMimoTranswithinbandR16                               *BandnrSimulsrsMimoTranswithinbandR16                               `ext`
	ChannelbwDlIabR16                                            *BandnrChannelbwDlIabR16                                            `ext`
	ChannelbwUlIabR16                                            *BandnrChannelbwUlIabR16                                            `ext`
	Rastershift7dot5IabR16                                       *BandnrRastershift7dot5IabR16                                       `ext`
	UePowerclassV1610                                            *BandnrUePowerclassV1610                                            `ext`
	CondhandoverR16                                              *BandnrCondhandoverR16                                              `ext`
	CondhandoverfailureR16                                       *BandnrCondhandoverfailureR16                                       `ext`
	CondhandovertwotriggereventsR16                              *BandnrCondhandovertwotriggereventsR16                              `ext`
	CondpscellchangeR16                                          *BandnrCondpscellchangeR16                                          `ext`
	CondpscellchangetwotriggereventsR16                          *BandnrCondpscellchangetwotriggereventsR16                          `ext`
	MprPowerboostFr2R16                                          *BandnrMprPowerboostFr2R16                                          `ext`
	ActiveconfiguredgrantR16                                     *BandnrActiveconfiguredgrantR16                                     `ext`
	Jointreleaseconfiguredgranttype2R16                          *BandnrJointreleaseconfiguredgranttype2R16                          `ext`
	SpsR16                                                       *BandnrSpsR16                                                       `ext`
	JointreleasespsR16                                           *BandnrJointreleasespsR16                                           `ext`
	SimulsrsTranswithinbandR16                                   *BandnrSimulsrsTranswithinbandR16                                   `ext`
	TrsAdditionalbandwidthR16                                    *BandnrTrsAdditionalbandwidthR16                                    `ext`
	HandoverintrafIabR16                                         *BandnrHandoverintrafIabR16                                         `ext`
	SimultxSrsAntswitchingintrabandulCaR16                       *SimulsrsForantennaswitchingR16                                     `ext`
	SharedspectrumchaccessparamsperbandV1630                     *SharedspectrumchaccessparamsperbandV1630                           `ext`
	HandoverutraFddR16                                           *BandnrHandoverutraFddR16                                           `ext`
	EnhancedulTransientperiodR16                                 *BandnrEnhancedulTransientperiodR16                                 `ext`
	SharedspectrumchaccessparamsperbandV1640                     *SharedspectrumchaccessparamsperbandV1640                           `ext`
	Type1PuschRepetitionmultislotsV1650                          *BandnrType1PuschRepetitionmultislotsV1650                          `ext`
	Type2PuschRepetitionmultislotsV1650                          *BandnrType2PuschRepetitionmultislotsV1650                          `ext`
	PuschRepetitionmultislotsV1650                               *BandnrPuschRepetitionmultislotsV1650                               `ext`
	ConfiguredulGranttype1V1650                                  *BandnrConfiguredulGranttype1V1650                                  `ext`
	ConfiguredulGranttype2V1650                                  *BandnrConfiguredulGranttype2V1650                                  `ext`
	SharedspectrumchaccessparamsperbandV1650                     *SharedspectrumchaccessparamsperbandV1650                           `ext`
	EnhancedskipuplinktxconfiguredV1660                          *BandnrEnhancedskipuplinktxconfiguredV1660                          `ext`
	EnhancedskipuplinktxdynamicV1660                             *BandnrEnhancedskipuplinktxdynamicV1660                             `ext`
	MaxuplinkdutycyclePc1dot5MpeFr1R16                           *BandnrMaxuplinkdutycyclePc1dot5MpeFr1R16                           `ext`
	TxdiversityR16                                               *BandnrTxdiversityR16                                               `ext`
	Pdsch1024qamFr1R17                                           *BandnrPdsch1024qamFr1R17                                           `ext`
	UePowerclassV1700                                            *BandnrUePowerclassV1700                                            `ext`
	Fr22AccessparamsperbandR17                                   *Fr22AccessparamsperbandR17                                         `ext`
	RlmRelaxationR17                                             *BandnrRlmRelaxationR17                                             `ext`
	BfdRelaxationR17                                             *BandnrBfdRelaxationR17                                             `ext`
	CgSdtR17                                                     *BandnrCgSdtR17                                                     `ext`
	LocationbasedcondhandoverR17                                 *BandnrLocationbasedcondhandoverR17                                 `ext`
	TimebasedcondhandoverR17                                     *BandnrTimebasedcondhandoverR17                                     `ext`
	Eventa4basedcondhandoverR17                                  *BandnrEventa4basedcondhandoverR17                                  `ext`
	MnInitiatedcondpscellchangenrdcR17                           *BandnrMnInitiatedcondpscellchangenrdcR17                           `ext`
	SnInitiatedcondpscellchangenrdcR17                           *BandnrSnInitiatedcondpscellchangenrdcR17                           `ext`
	PdcchSkippingwithoutsssgR17                                  *BandnrPdcchSkippingwithoutsssgR17                                  `ext`
	SssgSwitching1bitindR17                                      *BandnrSssgSwitching1bitindR17                                      `ext`
	SssgSwitching2bitindR17                                      *BandnrSssgSwitching2bitindR17                                      `ext`
	PdcchSkippingwithsssgR17                                     *BandnrPdcchSkippingwithsssgR17                                     `ext`
	SearchspacesetgrpSwitchcap2R17                               *BandnrSearchspacesetgrpSwitchcap2R17                               `ext`
	UplinkprecompensationR17                                     *BandnrUplinkprecompensationR17                                     `ext`
	UplinkTaReportingR17                                         *BandnrUplinkTaReportingR17                                         `ext`
	MaxHarqProcessnumberR17                                      *BandnrMaxHarqProcessnumberR17                                      `ext`
	Type2HarqCodebookR17                                         *BandnrType2HarqCodebookR17                                         `ext`
	Type1HarqCodebookR17                                         *BandnrType1HarqCodebookR17                                         `ext`
	Type3HarqCodebookR17                                         *BandnrType3HarqCodebookR17                                         `ext`
	UeSpecificKOffsetR17                                         *BandnrUeSpecificKOffsetR17                                         `ext`
	MultipdschSingledciFr21Scs120khzR17                          *BandnrMultipdschSingledciFr21Scs120khzR17                          `ext`
	MultipuschSingledciFr21Scs120khzR17                          *BandnrMultipuschSingledciFr21Scs120khzR17                          `ext`
	ParallelprsMeasrrcInactiveR17                                *BandnrParallelprsMeasrrcInactiveR17                                `ext`
	NrUeTxtegIdMaxsupportR17                                     *BandnrNrUeTxtegIdMaxsupportR17                                     `ext`
	PrsProcessingrrcInactiveR17                                  *BandnrPrsProcessingrrcInactiveR17                                  `ext`
	PrsProcessingwindowtype1aR17                                 *BandnrPrsProcessingwindowtype1aR17                                 `ext`
	PrsProcessingwindowtype1bR17                                 *BandnrPrsProcessingwindowtype1bR17                                 `ext`
	PrsProcessingwindowtype2R17                                  *BandnrPrsProcessingwindowtype2R17                                  `ext`
	SrsAllposresourcesrrcInactiveR17                             *SrsAllposresourcesrrcInactiveR17                                   `ext`
	OlpcSrsPosrrcInactiveR17                                     *OlpcSrsPosR16                                                      `ext`
	SpatialrelationssrsPosrrcInactiveR17                         *SpatialrelationssrsPosR16                                          `ext`
	MaxnumberpuschTypeaRepetitionR17                             *BandnrMaxnumberpuschTypeaRepetitionR17                             `ext`
	PuschtypeaRepetitionsavailslotR17                            *BandnrPuschtypeaRepetitionsavailslotR17                            `ext`
	TbProcessingmultislotpuschR17                                *BandnrTbProcessingmultislotpuschR17                                `ext`
	TbProcessingrepmultislotpuschR17                             *BandnrTbProcessingrepmultislotpuschR17                             `ext`
	MaxdurationdmrsBundlingR17                                   *BandnrMaxdurationdmrsBundlingR17                                   `ext`
	PuschRepetitionmsg3R17                                       *BandnrPuschRepetitionmsg3R17                                       `ext`
	SharedspectrumchaccessparamsperbandV1710                     *SharedspectrumchaccessparamsperbandV1710                           `ext`
	ParallelmeasurementwithoutrestrictionR17                     *BandnrParallelmeasurementwithoutrestrictionR17                     `ext`
	MaxnumberNgsoSatelliteswithinonesmtcR17                      *BandnrMaxnumberNgsoSatelliteswithinonesmtcR17                      `ext`
	K1RangeextensionR17                                          *BandnrK1RangeextensionR17                                          `ext`
	AperiodiccsiRsFastscellactivationR17                         *BandnrAperiodiccsiRsFastscellactivationR17                         `ext`
	AperiodiccsiRsAdditionalbandwidthR17                         *BandnrAperiodiccsiRsAdditionalbandwidthR17                         `ext`
	BwpWithoutcdSsbOrncdSsbRedcapR17                             *BandnrBwpWithoutcdSsbOrncdSsbRedcapR17                             `ext`
	HalfduplexfddTypeaRedcapR17                                  *BandnrHalfduplexfddTypeaRedcapR17                                  `ext`
	PossrsRrcInactiveOutsideinitialulBwpR17                      *PossrsRrcInactiveOutsideinitialulBwpR17                            `ext`
	ChannelbwsDlScs480khzFr22R17                                 *utils.BITSTRING                                                    `lb:8,ub:8,ext`
	ChannelbwsUlScs480khzFr22R17                                 *utils.BITSTRING                                                    `lb:8,ub:8,ext`
	ChannelbwsDlScs960khzFr22R17                                 *utils.BITSTRING                                                    `lb:8,ub:8,ext`
	ChannelbwsUlScs960khzFr22R17                                 *utils.BITSTRING                                                    `lb:8,ub:8,ext`
	UlGapfr2R17                                                  *BandnrUlGapfr2R17                                                  `ext`
	OneshotharqFeedbacktriggeredbydci12R17                       *BandnrOneshotharqFeedbacktriggeredbydci12R17                       `ext`
	OneshotharqFeedbackphyPriorityR17                            *BandnrOneshotharqFeedbackphyPriorityR17                            `ext`
	Enhancedtype3HarqCodebookfeedbackR17                         *BandnrEnhancedtype3HarqCodebookfeedbackR17                         `ext`
	TriggeredharqCodebookretxR17                                 *BandnrTriggeredharqCodebookretxR17                                 `ext`
	UeOneshotulTimingadjR17                                      *BandnrUeOneshotulTimingadjR17                                      `ext`
	PucchRepetitionF02R17                                        *BandnrPucchRepetitionF02R17                                        `ext`
	Cqi4BitssubbandntnSharedspectrumchaccessR17                  *BandnrCqi4BitssubbandntnSharedspectrumchaccessR17                  `ext`
	MuxHarqAckDiffprioritiesR17                                  *BandnrMuxHarqAckDiffprioritiesR17                                  `ext`
	TaBasedpdcNtnSharedspectrumchaccessR17                       *BandnrTaBasedpdcNtnSharedspectrumchaccessR17                       `ext`
	AckNackFeedbackformulticastwithdciEnablerR17                 *BandnrAckNackFeedbackformulticastwithdciEnablerR17                 `ext`
	MaxnumbergRntiR17                                            *utils.INTEGER                                                      `lb:0,ub:8,ext`
	DynamicmulticastdciFormat42R17                               *BandnrDynamicmulticastdciFormat42R17                               `ext`
	MaxmodulationorderformulticastR17                            *BandnrMaxmodulationorderformulticastR17                            `ext`
	DynamicslotrepetitionmulticasttnNonsharedspectrumchaccessR17 *BandnrDynamicslotrepetitionmulticasttnNonsharedspectrumchaccessR17 `ext`
	DynamicslotrepetitionmulticastntnSharedspectrumchaccessR17   *BandnrDynamicslotrepetitionmulticastntnSharedspectrumchaccessR17   `ext`
	NackOnlyfeedbackformulticastwithdciEnablerR17                *BandnrNackOnlyfeedbackformulticastwithdciEnablerR17                `ext`
	AckNackFeedbackforspsMulticastwithdciEnablerR17              *BandnrAckNackFeedbackforspsMulticastwithdciEnablerR17              `ext`
	MaxnumbergCsRntiR17                                          *utils.INTEGER                                                      `lb:0,ub:8,ext`
	ReLevelratematchingformulticastR17                           *BandnrReLevelratematchingformulticastR17                           `ext`
	Pdsch1024qam2mimoFr1R17                                      *BandnrPdsch1024qam2mimoFr1R17                                      `ext`
	PrsMeasurementwithoutmgR17                                   *BandnrPrsMeasurementwithoutmgR17                                   `ext`
	MaxnumberLeoSatellitespercarrierR17                          *utils.INTEGER                                                      `lb:0,ub:4,ext`
	PrsProcessingcapabilityoutsidemginppwR17                     *[]PrsProcessingcapabilityoutsidemginppwpertypeR17                  `lb:1,ub:3,ext`
	SrsSemipersistentPosresourcesrrcInactiveR17                  *BandnrSrsSemipersistentPosresourcesrrcInactiveR17                  `ext`
	ChannelbwsDlScs120khzFr22R17                                 *utils.BITSTRING                                                    `lb:8,ub:8,ext`
	ChannelbwsUlScs120khzFr22R17                                 *utils.BITSTRING                                                    `lb:8,ub:8,ext`
	DmrsBundlingpuschReptypeaR17                                 *BandnrDmrsBundlingpuschReptypeaR17                                 `ext`
	DmrsBundlingpuschReptypebR17                                 *BandnrDmrsBundlingpuschReptypebR17                                 `ext`
	DmrsBundlingpuschMultislotR17                                *BandnrDmrsBundlingpuschMultislotR17                                `ext`
	DmrsBundlingpucchRepR17                                      *BandnrDmrsBundlingpucchRepR17                                      `ext`
	InterslotfreqhopinterslotbundlingpuschR17                    *BandnrInterslotfreqhopinterslotbundlingpuschR17                    `ext`
	InterslotfreqhoppucchR17                                     *BandnrInterslotfreqhoppucchR17                                     `ext`
	DmrsBundlingrestartR17                                       *BandnrDmrsBundlingrestartR17                                       `ext`
	DmrsBundlingnonbacktobacktxR17                               *BandnrDmrsBundlingnonbacktobacktxR17                               `ext`
}
