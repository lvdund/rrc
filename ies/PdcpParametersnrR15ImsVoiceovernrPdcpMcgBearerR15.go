package ies

import "rrc/utils"

// PDCP-ParametersNR-r15-ims-VoiceOverNR-PDCP-MCG-Bearer-r15 ::= ENUMERATED
type PdcpParametersnrR15ImsVoiceovernrPdcpMcgBearerR15 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersnrR15ImsVoiceovernrPdcpMcgBearerR15EnumeratedNothing = iota
	PdcpParametersnrR15ImsVoiceovernrPdcpMcgBearerR15EnumeratedSupported
)
