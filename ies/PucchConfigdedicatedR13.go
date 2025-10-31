package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-r13 ::= SEQUENCE
type PucchConfigdedicatedR13 struct {
	AcknackrepetitionR13                      PucchConfigdedicatedR13AcknackrepetitionR13
	TddAcknackfeedbackmodeR13                 *PucchConfigdedicatedR13TddAcknackfeedbackmodeR13
	PucchFormatR13                            *PucchConfigdedicatedR13PucchFormatR13
	TwoantennaportactivatedpucchFormat1a1bR13 *bool
	SimultaneouspucchPuschR13                 *bool
	N1pucchAnRepp1R13                         *utils.INTEGER `lb:0,ub:2047`
	NpucchParamR13                            *PucchConfigdedicatedR13NpucchParamR13
	NkapucchParamR13                          *PucchConfigdedicatedR13NkapucchParamR13
	SpatialbundlingpucchR13                   utils.BOOLEAN
	SpatialbundlingpuschR13                   utils.BOOLEAN
	HarqTimingtddR13                          utils.BOOLEAN
	CodebooksizedeterminationR13              *PucchConfigdedicatedR13CodebooksizedeterminationR13
	MaximumpayloadcoderateR13                 *utils.INTEGER `lb:0,ub:7`
	PucchNumrepetitionceR13                   *PucchConfigdedicatedR13PucchNumrepetitionceR13
}
