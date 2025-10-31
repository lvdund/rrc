package ies

import "rrc/utils"

// CA-ParametersNR-v1720-nack-OnlyFeedbackSpecificResourceForMulticast-r17 ::= ENUMERATED
type CaParametersnrV1720NackOnlyfeedbackspecificresourceformulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1720NackOnlyfeedbackspecificresourceformulticastR17EnumeratedNothing = iota
	CaParametersnrV1720NackOnlyfeedbackspecificresourceformulticastR17EnumeratedSupported
)
