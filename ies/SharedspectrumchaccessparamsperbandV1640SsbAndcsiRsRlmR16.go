package ies

import "rrc/utils"

// SharedSpectrumChAccessParamsPerBand-v1640-ssb-AndCSI-RS-RLM-r16 ::= ENUMERATED
type SharedspectrumchaccessparamsperbandV1640SsbAndcsiRsRlmR16 struct {
	Value utils.ENUMERATED
}

const (
	SharedspectrumchaccessparamsperbandV1640SsbAndcsiRsRlmR16EnumeratedNothing = iota
	SharedspectrumchaccessparamsperbandV1640SsbAndcsiRsRlmR16EnumeratedSupported
)
