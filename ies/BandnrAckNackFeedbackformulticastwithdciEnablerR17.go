package ies

import "rrc/utils"

// BandNR-ack-NACK-FeedbackForMulticastWithDCI-Enabler-r17 ::= ENUMERATED
type BandnrAckNackFeedbackformulticastwithdciEnablerR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrAckNackFeedbackformulticastwithdciEnablerR17EnumeratedNothing = iota
	BandnrAckNackFeedbackformulticastwithdciEnablerR17EnumeratedSupported
)
