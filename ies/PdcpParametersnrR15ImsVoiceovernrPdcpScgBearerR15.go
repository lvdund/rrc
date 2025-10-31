package ies

import "rrc/utils"

// PDCP-ParametersNR-r15-ims-VoiceOverNR-PDCP-SCG-Bearer-r15 ::= ENUMERATED
type PdcpParametersnrR15ImsVoiceovernrPdcpScgBearerR15 struct {
	Value utils.ENUMERATED
}

const (
	PdcpParametersnrR15ImsVoiceovernrPdcpScgBearerR15EnumeratedNothing = iota
	PdcpParametersnrR15ImsVoiceovernrPdcpScgBearerR15EnumeratedSupported
)
