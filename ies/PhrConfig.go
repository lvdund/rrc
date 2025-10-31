package ies

import "rrc/utils"

// PHR-Config ::= SEQUENCE
// Extensible
type PhrConfig struct {
	PhrPeriodictimer       PhrConfigPhrPeriodictimer
	PhrProhibittimer       PhrConfigPhrProhibittimer
	PhrTxPowerfactorchange PhrConfigPhrTxPowerfactorchange
	Multiplephr            utils.BOOLEAN
	Dummy                  utils.BOOLEAN
	PhrType2othercell      utils.BOOLEAN
	PhrModeothercg         PhrConfigPhrModeothercg
	MpeReportingFr2R16     *utils.Setuprelease[MpeConfigFr2R16] `ext`
	MpeReportingFr2R17     *utils.Setuprelease[MpeConfigFr2R17] `ext`
	TwophrmodeR17          *PhrConfigTwophrmodeR17              `ext`
}
