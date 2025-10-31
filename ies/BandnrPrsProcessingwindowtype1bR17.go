package ies

import "rrc/utils"

// BandNR-prs-ProcessingWindowType1B-r17 ::= ENUMERATED
type BandnrPrsProcessingwindowtype1bR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPrsProcessingwindowtype1bR17EnumeratedNothing = iota
	BandnrPrsProcessingwindowtype1bR17EnumeratedOption1
	BandnrPrsProcessingwindowtype1bR17EnumeratedOption2
	BandnrPrsProcessingwindowtype1bR17EnumeratedOption3
)
