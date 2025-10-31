package ies

import "rrc/utils"

// BandNR-prs-ProcessingWindowType1A-r17 ::= ENUMERATED
type BandnrPrsProcessingwindowtype1aR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPrsProcessingwindowtype1aR17EnumeratedNothing = iota
	BandnrPrsProcessingwindowtype1aR17EnumeratedOption1
	BandnrPrsProcessingwindowtype1aR17EnumeratedOption2
	BandnrPrsProcessingwindowtype1aR17EnumeratedOption3
)
