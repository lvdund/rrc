package ies

import "rrc/utils"

// CA-ParametersNR-v1730-nack-OnlyFeedbackSpecificResourceForSPS-Multicast-r17 ::= ENUMERATED
type CaParametersnrV1730NackOnlyfeedbackspecificresourceforspsMulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1730NackOnlyfeedbackspecificresourceforspsMulticastR17EnumeratedNothing = iota
	CaParametersnrV1730NackOnlyfeedbackspecificresourceforspsMulticastR17EnumeratedSupported
)
