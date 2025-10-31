package ies

import "rrc/utils"

// BandNR-oneShotHARQ-feedbackPhy-Priority-r17 ::= ENUMERATED
type BandnrOneshotharqFeedbackphyPriorityR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrOneshotharqFeedbackphyPriorityR17EnumeratedNothing = iota
	BandnrOneshotharqFeedbackphyPriorityR17EnumeratedSupported
)
