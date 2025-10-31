package ies

import "rrc/utils"

// CA-ParametersNR-v1720-nack-OnlyFeedbackForMulticast-r17 ::= ENUMERATED
type CaParametersnrV1720NackOnlyfeedbackformulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1720NackOnlyfeedbackformulticastR17EnumeratedNothing = iota
	CaParametersnrV1720NackOnlyfeedbackformulticastR17EnumeratedSupported
)
