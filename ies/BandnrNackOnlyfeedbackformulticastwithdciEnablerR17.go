package ies

import "rrc/utils"

// BandNR-nack-OnlyFeedbackForMulticastWithDCI-Enabler-r17 ::= ENUMERATED
type BandnrNackOnlyfeedbackformulticastwithdciEnablerR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrNackOnlyfeedbackformulticastwithdciEnablerR17EnumeratedNothing = iota
	BandnrNackOnlyfeedbackformulticastwithdciEnablerR17EnumeratedSupported
)
