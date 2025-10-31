package ies

import "rrc/utils"

// SL-LogicalChannelConfig-r16-sl-HARQ-FeedbackEnabled-r16 ::= ENUMERATED
type SlLogicalchannelconfigR16SlHarqFeedbackenabledR16 struct {
	Value utils.ENUMERATED
}

const (
	SlLogicalchannelconfigR16SlHarqFeedbackenabledR16EnumeratedNothing = iota
	SlLogicalchannelconfigR16SlHarqFeedbackenabledR16EnumeratedEnabled
	SlLogicalchannelconfigR16SlHarqFeedbackenabledR16EnumeratedDisabled
)
