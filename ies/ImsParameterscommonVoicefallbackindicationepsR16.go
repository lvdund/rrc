package ies

import "rrc/utils"

// IMS-ParametersCommon-voiceFallbackIndicationEPS-r16 ::= ENUMERATED
type ImsParameterscommonVoicefallbackindicationepsR16 struct {
	Value utils.ENUMERATED
}

const (
	ImsParameterscommonVoicefallbackindicationepsR16EnumeratedNothing = iota
	ImsParameterscommonVoicefallbackindicationepsR16EnumeratedSupported
)
