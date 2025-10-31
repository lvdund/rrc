package ies

import "rrc/utils"

// Phy-ParametersCommon ::= SEQUENCE
// Extensible
type PhyParameterscommon struct {
	CsiRsCfraForho                                           *PhyParameterscommonCsiRsCfraForho
	DynamicprbBundlingdl                                     *PhyParameterscommonDynamicprbBundlingdl
	SpCsiReportpucch                                         *PhyParameterscommonSpCsiReportpucch
	SpCsiReportpusch                                         *PhyParameterscommonSpCsiReportpusch
	NzpCsiRsIntefmgmt                                        *PhyParameterscommonNzpCsiRsIntefmgmt
	Type2SpCsiFeedbackLongpucch                              *PhyParameterscommonType2SpCsiFeedbackLongpucch
	Precodergranularitycoreset                               *PhyParameterscommonPrecodergranularitycoreset
	DynamicharqAckCodebook                                   *PhyParameterscommonDynamicharqAckCodebook
	SemistaticharqAckCodebook                                *PhyParameterscommonSemistaticharqAckCodebook
	SpatialbundlingharqAck                                   *PhyParameterscommonSpatialbundlingharqAck
	DynamicbetaoffsetindHarqAckCsi                           *PhyParameterscommonDynamicbetaoffsetindHarqAckCsi
	PucchRepetitionF134                                      *PhyParameterscommonPucchRepetitionF134
	RaType0Pusch                                             *PhyParameterscommonRaType0Pusch
	DynamicswitchraType01Pdsch                               *PhyParameterscommonDynamicswitchraType01Pdsch
	DynamicswitchraType01Pusch                               *PhyParameterscommonDynamicswitchraType01Pusch
	PdschMappingtypea                                        *PhyParameterscommonPdschMappingtypea
	PdschMappingtypeb                                        *PhyParameterscommonPdschMappingtypeb
	InterleavingvrbToprbPdsch                                *PhyParameterscommonInterleavingvrbToprbPdsch
	InterslotfreqhoppingPusch                                *PhyParameterscommonInterslotfreqhoppingPusch
	Type1PuschRepetitionmultislots                           *PhyParameterscommonType1PuschRepetitionmultislots
	Type2PuschRepetitionmultislots                           *PhyParameterscommonType2PuschRepetitionmultislots
	PuschRepetitionmultislots                                *PhyParameterscommonPuschRepetitionmultislots
	PdschRepetitionmultislots                                *PhyParameterscommonPdschRepetitionmultislots
	Downlinksps                                              *PhyParameterscommonDownlinksps
	ConfiguredulGranttype1                                   *PhyParameterscommonConfiguredulGranttype1
	ConfiguredulGranttype2                                   *PhyParameterscommonConfiguredulGranttype2
	PreEmptindicationDl                                      *PhyParameterscommonPreEmptindicationDl
	CbgTransindicationDl                                     *PhyParameterscommonCbgTransindicationDl
	CbgTransindicationUl                                     *PhyParameterscommonCbgTransindicationUl
	CbgFlushindicationDl                                     *PhyParameterscommonCbgFlushindicationDl
	DynamicharqAckCodebCbgRetxDl                             *PhyParameterscommonDynamicharqAckCodebCbgRetxDl
	RatematchingresrcsetsemiStatic                           *PhyParameterscommonRatematchingresrcsetsemiStatic
	Ratematchingresrcsetdynamic                              *PhyParameterscommonRatematchingresrcsetdynamic
	BwpSwitchingdelay                                        *PhyParameterscommonBwpSwitchingdelay
	Dummy                                                    *PhyParameterscommonDummy                                                    `ext`
	Maxnumbersearchspaces                                    *PhyParameterscommonMaxnumbersearchspaces                                    `ext`
	Ratematchingctrlresrcsetdynamic                          *PhyParameterscommonRatematchingctrlresrcsetdynamic                          `ext`
	MaxlayersmimoIndication                                  *PhyParameterscommonMaxlayersmimoIndication                                  `ext`
	Spcellplacement                                          *Carrieraggregationvariant                                                   `ext`
	TwosteprachR16                                           *PhyParameterscommonTwosteprachR16                                           `ext`
	DciFormat12and02R16                                      *PhyParameterscommonDciFormat12and02R16                                      `ext`
	MonitoringdciSamesearchspaceR16                          *PhyParameterscommonMonitoringdciSamesearchspaceR16                          `ext`
	Type2CgReleasedci01R16                                   *PhyParameterscommonType2CgReleasedci01R16                                   `ext`
	Type2CgReleasedci02R16                                   *PhyParameterscommonType2CgReleasedci02R16                                   `ext`
	SpsReleasedci11R16                                       *PhyParameterscommonSpsReleasedci11R16                                       `ext`
	SpsReleasedci12R16                                       *PhyParameterscommonSpsReleasedci12R16                                       `ext`
	CsiTriggerstatenonActivebwpR16                           *PhyParameterscommonCsiTriggerstatenonActivebwpR16                           `ext`
	SeparatesmtcInteriabSupportR16                           *PhyParameterscommonSeparatesmtcInteriabSupportR16                           `ext`
	SeparaterachIabSupportR16                                *PhyParameterscommonSeparaterachIabSupportR16                                `ext`
	UlFlexibledlSlotformatsemistaticIabR16                   *PhyParameterscommonUlFlexibledlSlotformatsemistaticIabR16                   `ext`
	UlFlexibledlSlotformatdynamicsIabR16                     *PhyParameterscommonUlFlexibledlSlotformatdynamicsIabR16                     `ext`
	DftSOfdmWaveformulIabR16                                 *PhyParameterscommonDftSOfdmWaveformulIabR16                                 `ext`
	Dci25AiRntiSupportIabR16                                 *PhyParameterscommonDci25AiRntiSupportIabR16                                 `ext`
	TDeltareceptionsupportIabR16                             *PhyParameterscommonTDeltareceptionsupportIabR16                             `ext`
	GuardsymbolreportreceptionIabR16                         *PhyParameterscommonGuardsymbolreportreceptionIabR16                         `ext`
	HarqackCbSpatialbundlingpucchGroupR16                    *PhyParameterscommonHarqackCbSpatialbundlingpucchGroupR16                    `ext`
	CrossslotschedulingR16                                   *PhyParameterscommonCrossslotschedulingR16                                   `ext`
	MaxnumbersrsPospathlossestimateallservingcellsR16        *PhyParameterscommonMaxnumbersrsPospathlossestimateallservingcellsR16        `ext`
	ExtendedcgPeriodicitiesR16                               *PhyParameterscommonExtendedcgPeriodicitiesR16                               `ext`
	ExtendedspsPeriodicitiesR16                              *PhyParameterscommonExtendedspsPeriodicitiesR16                              `ext`
	CodebookvariantslistR16                                  *CodebookvariantslistR16                                                     `ext`
	PuschRepetitiontypeaR16                                  *PhyParameterscommonPuschRepetitiontypeaR16                                  `ext`
	DciDlPriorityindicatorR16                                *PhyParameterscommonDciDlPriorityindicatorR16                                `ext`
	DciUlPriorityindicatorR16                                *PhyParameterscommonDciUlPriorityindicatorR16                                `ext`
	MaxnumberpathlossrsUpdateR16                             *PhyParameterscommonMaxnumberpathlossrsUpdateR16                             `ext`
	Type2HarqAckCodebookR16                                  *PhyParameterscommonType2HarqAckCodebookR16                                  `ext`
	MaxtotalresourcesforacrossfreqrangesR16                  *PhyParameterscommonMaxtotalresourcesforacrossfreqrangesR16                  `ext`
	HarqackSeparatemultidciMultitrpR16                       *PhyParameterscommonHarqackSeparatemultidciMultitrpR16                       `ext`
	HarqackJointmultidciMultitrpR16                          *PhyParameterscommonHarqackJointmultidciMultitrpR16                          `ext`
	BwpSwitchingmulticcsR16                                  *PhyParameterscommonBwpSwitchingmulticcsR16                                  `ext`
	TargetsmtcScgR16                                         *PhyParameterscommonTargetsmtcScgR16                                         `ext`
	SupportrepetitionzerooffsetrvR16                         *PhyParameterscommonSupportrepetitionzerooffsetrvR16                         `ext`
	CbgTransinorderpuschUlR16                                *PhyParameterscommonCbgTransinorderpuschUlR16                                `ext`
	BwpSwitchingmultidormancyccsR16                          *PhyParameterscommonBwpSwitchingmultidormancyccsR16                          `ext`
	SupportretxDiffCoresetpoolMultiDciTrpR16                 *PhyParameterscommonSupportretxDiffCoresetpoolMultiDciTrpR16                 `ext`
	PdcchMonitoringanyoccasionswithspangapcrosscarrierschR16 *PhyParameterscommonPdcchMonitoringanyoccasionswithspangapcrosscarrierschR16 `ext`
	Newbeamidentifications2portcsiRsR16                      *PhyParameterscommonNewbeamidentifications2portcsiRsR16                      `ext`
	Pathlossestimation2portcsiRsR16                          *PhyParameterscommonPathlossestimation2portcsiRsR16                          `ext`
	MuxHarqAckWithoutpucchOnpuschR16                         *PhyParameterscommonMuxHarqAckWithoutpucchOnpuschR16                         `ext`
	GuardsymbolreportreceptionIabR17                         *PhyParameterscommonGuardsymbolreportreceptionIabR17                         `ext`
	RestrictedIabDuBeamreceptionR17                          *PhyParameterscommonRestrictedIabDuBeamreceptionR17                          `ext`
	RecommendedIabMtBeamtransmissionR17                      *PhyParameterscommonRecommendedIabMtBeamtransmissionR17                      `ext`
	Case6TimingalignmentreceptionIabR17                      *PhyParameterscommonCase6TimingalignmentreceptionIabR17                      `ext`
	Case7TimingalignmentreceptionIabR17                      *PhyParameterscommonCase7TimingalignmentreceptionIabR17                      `ext`
	DlTxPoweradjustmentIabR17                                *PhyParameterscommonDlTxPoweradjustmentIabR17                                `ext`
	DesiredUlTxPoweradjustmentR17                            *PhyParameterscommonDesiredUlTxPoweradjustmentR17                            `ext`
	FdmSoftresourceavailabilityDynamicindicationR17          *PhyParameterscommonFdmSoftresourceavailabilityDynamicindicationR17          `ext`
	UpdatedTDeltarangerecptionR17                            *PhyParameterscommonUpdatedTDeltarangerecptionR17                            `ext`
	SlotbaseddynamicpucchRepR17                              *PhyParameterscommonSlotbaseddynamicpucchRepR17                              `ext`
	SpsHarqAckDeferralR17                                    *PhyParameterscommonSpsHarqAckDeferralR17                                    `ext`
	UnifiedjointtciCommonupdateR17                           *utils.INTEGER                                                               `lb:0,ub:4,ext`
	MtrpPdcchSinglespanR17                                   *PhyParameterscommonMtrpPdcchSinglespanR17                                   `ext`
	SupportedactivatedprsProcessingwindowR17                 *PhyParameterscommonSupportedactivatedprsProcessingwindowR17                 `ext`
	CgTimedomainallocationextensionR17                       *PhyParameterscommonCgTimedomainallocationextensionR17                       `ext`
	TaBasedpdcTnNonsharedspectrumchaccessR17                 *PhyParameterscommonTaBasedpdcTnNonsharedspectrumchaccessR17                 `ext`
	DirectionalcollisiondcIabR17                             *PhyParameterscommonDirectionalcollisiondcIabR17                             `ext`
	PriorityindicatorindciMulticastR17                       *PhyParameterscommonPriorityindicatorindciMulticastR17                       `ext`
	PriorityindicatorindciSpsMulticastR17                    *PhyParameterscommonPriorityindicatorindciSpsMulticastR17                    `ext`
	TwoharqAckCodebookforunicastandmulticastR17              *PhyParameterscommonTwoharqAckCodebookforunicastandmulticastR17              `ext`
	MultipucchHarqAckFormulticastunicastR17                  *PhyParameterscommonMultipucchHarqAckFormulticastunicastR17                  `ext`
	SrsAdditionalrepetitionR17                               *PhyParameterscommonSrsAdditionalrepetitionR17                               `ext`
	PuschRepetitionCgSdtR17                                  *PhyParameterscommonPuschRepetitionCgSdtR17                                  `ext`
}
