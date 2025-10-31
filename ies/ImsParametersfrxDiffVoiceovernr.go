package ies

import "rrc/utils"

// IMS-ParametersFRX-Diff-voiceOverNR ::= ENUMERATED
type ImsParametersfrxDiffVoiceovernr struct {
	Value utils.ENUMERATED
}

const (
	ImsParametersfrxDiffVoiceovernrEnumeratedNothing = iota
	ImsParametersfrxDiffVoiceovernrEnumeratedSupported
)
