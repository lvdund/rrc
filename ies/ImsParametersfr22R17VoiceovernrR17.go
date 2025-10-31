package ies

import "rrc/utils"

// IMS-ParametersFR2-2-r17-voiceOverNR-r17 ::= ENUMERATED
type ImsParametersfr22R17VoiceovernrR17 struct {
	Value utils.ENUMERATED
}

const (
	ImsParametersfr22R17VoiceovernrR17EnumeratedNothing = iota
	ImsParametersfr22R17VoiceovernrR17EnumeratedSupported
)
