package ies

import "rrc/utils"

// IRAT-ParametersUTRA-v9c0-voiceOverPS-HS-UTRA-TDD128-r9 ::= ENUMERATED
type IratParametersutraV9c0VoiceoverpsHsUtraTdd128R9 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersutraV9c0VoiceoverpsHsUtraTdd128R9EnumeratedNothing = iota
	IratParametersutraV9c0VoiceoverpsHsUtraTdd128R9EnumeratedSupported
)
