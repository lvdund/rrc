package ies

import "rrc/utils"

// IRAT-ParametersNR-v1540-ims-VoiceOverNR-FR1-r15 ::= ENUMERATED
type IratParametersnrV1540ImsVoiceovernrFr1R15 struct {
	Value utils.ENUMERATED
}

const (
	IratParametersnrV1540ImsVoiceovernrFr1R15EnumeratedNothing = iota
	IratParametersnrV1540ImsVoiceovernrFr1R15EnumeratedSupported
)
