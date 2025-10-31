package ies

import "rrc/utils"

// IMS-ParametersCommon-voiceOverSCG-BearerEUTRA-5GC ::= ENUMERATED
type ImsParameterscommonVoiceoverscgBearereutra5gc struct {
	Value utils.ENUMERATED
}

const (
	ImsParameterscommonVoiceoverscgBearereutra5gcEnumeratedNothing = iota
	ImsParameterscommonVoiceoverscgBearereutra5gcEnumeratedSupported
)
