package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff ::= SEQUENCE
// Extensible
type PhyParametersfrxDiff struct {
	Dynamicsfi                                    *PhyParametersfrxDiffDynamicsfi
	Dummy1                                        *utils.BITSTRING `lb:2,ub:2`
	TwoflDmrs                                     *utils.BITSTRING `lb:2,ub:2`
	Dummy2                                        *utils.BITSTRING `lb:2,ub:2`
	Dummy3                                        *utils.BITSTRING `lb:2,ub:2`
	SupporteddmrsTypedl                           *PhyParametersfrxDiffSupporteddmrsTypedl
	SupporteddmrsTypeul                           *PhyParametersfrxDiffSupporteddmrsTypeul
	Semiopenloopcsi                               *PhyParametersfrxDiffSemiopenloopcsi
	CsiReportwithoutpmi                           *PhyParametersfrxDiffCsiReportwithoutpmi
	CsiReportwithoutcqi                           *PhyParametersfrxDiffCsiReportwithoutcqi
	Oneportsptrs                                  *utils.BITSTRING `lb:2,ub:2`
	TwopucchF02Consecsymbols                      *PhyParametersfrxDiffTwopucchF02Consecsymbols
	PucchF2Withfh                                 *PhyParametersfrxDiffPucchF2Withfh
	PucchF3Withfh                                 *PhyParametersfrxDiffPucchF3Withfh
	PucchF4Withfh                                 *PhyParametersfrxDiffPucchF4Withfh
	PucchF02withoutfh                             *PhyParametersfrxDiffPucchF02withoutfh
	PucchF134withoutfh                            *PhyParametersfrxDiffPucchF134withoutfh
	MuxSrHarqAckCsiPucchMultiperslot              *PhyParametersfrxDiffMuxSrHarqAckCsiPucchMultiperslot
	UciCodeblocksegmentation                      *PhyParametersfrxDiffUciCodeblocksegmentation
	OnepucchLongandshortformat                    *PhyParametersfrxDiffOnepucchLongandshortformat
	TwopucchAnyothersinslot                       *PhyParametersfrxDiffTwopucchAnyothersinslot
	IntraslotfreqhoppingPusch                     *PhyParametersfrxDiffIntraslotfreqhoppingPusch
	PuschLbrm                                     *PhyParametersfrxDiffPuschLbrm
	PdcchBlinddetectionca                         *utils.INTEGER `lb:0,ub:16`
	TpcPuschRnti                                  *PhyParametersfrxDiffTpcPuschRnti
	TpcPucchRnti                                  *PhyParametersfrxDiffTpcPucchRnti
	TpcSrsRnti                                    *PhyParametersfrxDiffTpcSrsRnti
	AbsolutetpcCommand                            *PhyParametersfrxDiffAbsolutetpcCommand
	TwodifferenttpcLoopPusch                      *PhyParametersfrxDiffTwodifferenttpcLoopPusch
	TwodifferenttpcLoopPucch                      *PhyParametersfrxDiffTwodifferenttpcLoopPucch
	PuschHalfpiBpsk                               *PhyParametersfrxDiffPuschHalfpiBpsk
	PucchF34HalfpiBpsk                            *PhyParametersfrxDiffPucchF34HalfpiBpsk
	AlmostcontiguouscpOfdmUl                      *PhyParametersfrxDiffAlmostcontiguouscpOfdmUl
	SpCsiRs                                       *PhyParametersfrxDiffSpCsiRs
	SpCsiIm                                       *PhyParametersfrxDiffSpCsiIm
	TddMultidlUlSwitchperslot                     *PhyParametersfrxDiffTddMultidlUlSwitchperslot
	Multiplecoreset                               *PhyParametersfrxDiffMultiplecoreset
	CsiRsImReceptionforfeedback                   *CsiRsImReceptionforfeedback                                       `ext`
	CsiRsProcframeworkforsrs                      *CsiRsProcframeworkforsrs                                          `ext`
	CsiReportframework                            *CsiReportframework                                                `ext`
	MuxSrHarqAckCsiPucchOnceperslot               *PhyParametersfrxDiffMuxSrHarqAckCsiPucchOnceperslot               `ext`
	MuxSrHarqAckPucch                             *PhyParametersfrxDiffMuxSrHarqAckPucch                             `ext`
	MuxMultiplegroupctrlchOverlap                 *PhyParametersfrxDiffMuxMultiplegroupctrlchOverlap                 `ext`
	DlSchedulingoffsetPdschTypea                  *PhyParametersfrxDiffDlSchedulingoffsetPdschTypea                  `ext`
	DlSchedulingoffsetPdschTypeb                  *PhyParametersfrxDiffDlSchedulingoffsetPdschTypeb                  `ext`
	UlSchedulingoffset                            *PhyParametersfrxDiffUlSchedulingoffset                            `ext`
	Dl64qamMcsTablealt                            *PhyParametersfrxDiffDl64qamMcsTablealt                            `ext`
	Ul64qamMcsTablealt                            *PhyParametersfrxDiffUl64qamMcsTablealt                            `ext`
	CqiTablealt                                   *PhyParametersfrxDiffCqiTablealt                                   `ext`
	OneflDmrsTwoadditionaldmrsUl                  *PhyParametersfrxDiffOneflDmrsTwoadditionaldmrsUl                  `ext`
	TwoflDmrsTwoadditionaldmrsUl                  *PhyParametersfrxDiffTwoflDmrsTwoadditionaldmrsUl                  `ext`
	OneflDmrsThreeadditionaldmrsUl                *PhyParametersfrxDiffOneflDmrsThreeadditionaldmrsUl                `ext`
	PdcchBlinddetectionnrdc                       *PhyParametersfrxDiffPdcchBlinddetectionnrdc                       `ext`
	MuxHarqAckPuschDiffsymbol                     *PhyParametersfrxDiffMuxHarqAckPuschDiffsymbol                     `ext`
	Type1HarqAckCodebookR16                       *PhyParametersfrxDiffType1HarqAckCodebookR16                       `ext`
	EnhancedpowercontrolR16                       *PhyParametersfrxDiffEnhancedpowercontrolR16                       `ext`
	SimultaneoustciActmultipleccR16               *PhyParametersfrxDiffSimultaneoustciActmultipleccR16               `ext`
	SimultaneousspatialrelationmultipleccR16      *PhyParametersfrxDiffSimultaneousspatialrelationmultipleccR16      `ext`
	CliRssiFdmDlR16                               *PhyParametersfrxDiffCliRssiFdmDlR16                               `ext`
	CliSrsRsrpFdmDlR16                            *PhyParametersfrxDiffCliSrsRsrpFdmDlR16                            `ext`
	MaxlayersmimoAdaptationR16                    *PhyParametersfrxDiffMaxlayersmimoAdaptationR16                    `ext`
	AggregationfactorspsDlR16                     *PhyParametersfrxDiffAggregationfactorspsDlR16                     `ext`
	MaxtotalresourcesforonefreqrangeR16           *PhyParametersfrxDiffMaxtotalresourcesforonefreqrangeR16           `ext`
	CsiReportframeworkextR16                      *CsiReportframeworkextR16                                          `ext`
	TwotciActServingcellinccListR16               *PhyParametersfrxDiffTwotciActServingcellinccListR16               `ext`
	CriRiCqiWithoutnonPmiPortindR16               *PhyParametersfrxDiffCriRiCqiWithoutnonPmiPortindR16               `ext`
	Cqi4BitssubbandtnNonsharedspectrumchaccessR17 *PhyParametersfrxDiffCqi4BitssubbandtnNonsharedspectrumchaccessR17 `ext`
}
