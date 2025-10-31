package ies

import "rrc/utils"

// MAC-ParametersCommon-harq-FeedbackDisabled-r17 ::= ENUMERATED
type MacParameterscommonHarqFeedbackdisabledR17 struct {
	Value utils.ENUMERATED
}

const (
	MacParameterscommonHarqFeedbackdisabledR17EnumeratedNothing = iota
	MacParameterscommonHarqFeedbackdisabledR17EnumeratedSupported
)
