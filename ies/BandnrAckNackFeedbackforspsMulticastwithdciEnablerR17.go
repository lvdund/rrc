package ies

import "rrc/utils"

// BandNR-ack-NACK-FeedbackForSPS-MulticastWithDCI-Enabler-r17 ::= ENUMERATED
type BandnrAckNackFeedbackforspsMulticastwithdciEnablerR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrAckNackFeedbackforspsMulticastwithdciEnablerR17EnumeratedNothing = iota
	BandnrAckNackFeedbackforspsMulticastwithdciEnablerR17EnumeratedSupported
)
