package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-r13 ::= SEQUENCE
type PucchConfigdedicatedR13 struct {
	AcknackrepetitionR13                      interface{}
	TddAcknackfeedbackmodeR13                 *utils.ENUMERATED
	PucchFormatR13                            *interface{}
	TwoantennaportactivatedpucchFormat1a1bR13 *utils.ENUMERATED
	SimultaneouspucchPuschR13                 *utils.ENUMERATED
	N1pucchAnRepp1R13                         *utils.INTEGER
	NpucchParamR13                            *interface{}
	NkapucchParamR13                          *interface{}
	SpatialbundlingpucchR13                   bool
	SpatialbundlingpuschR13                   bool
	HarqTimingtddR13                          bool
	CodebooksizedeterminationR13              *utils.ENUMERATED
	MaximumpayloadcoderateR13                 *utils.INTEGER
	PucchNumrepetitionceR13                   *interface{}
}
