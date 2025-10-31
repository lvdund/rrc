package ies

import "rrc/utils"

// IMS-ParametersCommon-voiceOverEUTRA-5GC ::= ENUMERATED
type ImsParameterscommonVoiceovereutra5gc struct {
	Value utils.ENUMERATED
}

const (
	ImsParameterscommonVoiceovereutra5gcEnumeratedNothing = iota
	ImsParameterscommonVoiceovereutra5gcEnumeratedSupported
)
