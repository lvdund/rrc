package ies

import "rrc/utils"

// BandNR-prs-ProcessingWindowType2-r17 ::= ENUMERATED
type BandnrPrsProcessingwindowtype2R17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPrsProcessingwindowtype2R17EnumeratedNothing = iota
	BandnrPrsProcessingwindowtype2R17EnumeratedOption1
	BandnrPrsProcessingwindowtype2R17EnumeratedOption2
	BandnrPrsProcessingwindowtype2R17EnumeratedOption3
)
