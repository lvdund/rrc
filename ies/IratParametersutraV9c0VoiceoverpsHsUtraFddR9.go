package ies

import "rrc/utils"

// IRAT-ParametersUTRA-v9c0-voiceOverPS-HS-UTRA-FDD-r9 ::= ENUMERATED
type IratParametersutraV9c0VoiceoverpsHsUtraFddR9 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersutraV9c0VoiceoverpsHsUtraFddR9EnumeratedNothing = iota
	IratParametersutraV9c0VoiceoverpsHsUtraFddR9EnumeratedSupported
)
