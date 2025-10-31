package ies

import "rrc/utils"

// CA-ParametersNR-v1720-ack-NACK-FeedbackForMulticast-r17 ::= ENUMERATED
type CaParametersnrV1720AckNackFeedbackformulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1720AckNackFeedbackformulticastR17EnumeratedNothing = iota
	CaParametersnrV1720AckNackFeedbackformulticastR17EnumeratedSupported
)
