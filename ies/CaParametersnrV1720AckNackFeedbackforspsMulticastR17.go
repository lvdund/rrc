package ies

import "rrc/utils"

// CA-ParametersNR-v1720-ack-NACK-FeedbackForSPS-Multicast-r17 ::= ENUMERATED
type CaParametersnrV1720AckNackFeedbackforspsMulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1720AckNackFeedbackforspsMulticastR17EnumeratedNothing = iota
	CaParametersnrV1720AckNackFeedbackforspsMulticastR17EnumeratedSupported
)
