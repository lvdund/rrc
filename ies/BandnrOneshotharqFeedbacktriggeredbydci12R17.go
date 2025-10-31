package ies

import "rrc/utils"

// BandNR-oneShotHARQ-feedbackTriggeredByDCI-1-2-r17 ::= ENUMERATED
type BandnrOneshotharqFeedbacktriggeredbydci12R17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrOneshotharqFeedbacktriggeredbydci12R17EnumeratedNothing = iota
	BandnrOneshotharqFeedbacktriggeredbydci12R17EnumeratedSupported
)
