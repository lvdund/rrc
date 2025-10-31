package ies

// MAC-ParametersCommon ::= SEQUENCE
// Extensible
type MacParameterscommon struct {
	LcpRestriction                        *MacParameterscommonLcpRestriction
	Dummy                                 *MacParameterscommonDummy
	LchToscellrestriction                 *MacParameterscommonLchToscellrestriction
	Recommendedbitrate                    *MacParameterscommonRecommendedbitrate                    `ext`
	Recommendedbitratequery               *MacParameterscommonRecommendedbitratequery               `ext`
	RecommendedbitratemultiplierR16       *MacParameterscommonRecommendedbitratemultiplierR16       `ext`
	PreemptivebsrR16                      *MacParameterscommonPreemptivebsrR16                      `ext`
	AutonomoustransmissionR16             *MacParameterscommonAutonomoustransmissionR16             `ext`
	LchPrioritybasedprioritizationR16     *MacParameterscommonLchPrioritybasedprioritizationR16     `ext`
	LchToconfiguredgrantmappingR16        *MacParameterscommonLchToconfiguredgrantmappingR16        `ext`
	LchTograntpriorityrestrictionR16      *MacParameterscommonLchTograntpriorityrestrictionR16      `ext`
	SinglephrPR16                         *MacParameterscommonSinglephrPR16                         `ext`
	UlLbtFailuredetectionrecoveryR16      *MacParameterscommonUlLbtFailuredetectionrecoveryR16      `ext`
	TddMpePMprReportingR16                *MacParameterscommonTddMpePMprReportingR16                `ext`
	LcidExtensioniabR16                   *MacParameterscommonLcidExtensioniabR16                   `ext`
	SpcellBfrCbraR16                      *MacParameterscommonSpcellBfrCbraR16                      `ext`
	SrsResourceidExtR16                   *MacParameterscommonSrsResourceidExtR16                   `ext`
	EnhanceduudrxForsidelinkR17           *MacParameterscommonEnhanceduudrxForsidelinkR17           `ext`
	MgActivationrequestprsMeasR17         *MacParameterscommonMgActivationrequestprsMeasR17         `ext`
	MgActivationcommprsMeasR17            *MacParameterscommonMgActivationcommprsMeasR17            `ext`
	IntracgPrioritizationR17              *MacParameterscommonIntracgPrioritizationR17              `ext`
	JointprioritizationcgRetxTimerR17     *MacParameterscommonJointprioritizationcgRetxTimerR17     `ext`
	SurvivaltimeR17                       *MacParameterscommonSurvivaltimeR17                       `ext`
	LcgExtensioniabR17                    *MacParameterscommonLcgExtensioniabR17                    `ext`
	HarqFeedbackdisabledR17               *MacParameterscommonHarqFeedbackdisabledR17               `ext`
	UplinkHarqModebR17                    *MacParameterscommonUplinkHarqModebR17                    `ext`
	SrTriggeredbyTaReportR17              *MacParameterscommonSrTriggeredbyTaReportR17              `ext`
	ExtendeddrxCycleinactiveR17           *MacParameterscommonExtendeddrxCycleinactiveR17           `ext`
	SimultaneoussrPuschDiffpucchGroupsR17 *MacParameterscommonSimultaneoussrPuschDiffpucchGroupsR17 `ext`
	LasttransmissionulR17                 *MacParameterscommonLasttransmissionulR17                 `ext`
}
