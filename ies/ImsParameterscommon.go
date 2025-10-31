package ies

// IMS-ParametersCommon ::= SEQUENCE
// Extensible
type ImsParameterscommon struct {
	Voiceovereutra5gc             *ImsParameterscommonVoiceovereutra5gc
	VoiceoverscgBearereutra5gc    *ImsParameterscommonVoiceoverscgBearereutra5gc    `ext`
	VoicefallbackindicationepsR16 *ImsParameterscommonVoicefallbackindicationepsR16 `ext`
}
