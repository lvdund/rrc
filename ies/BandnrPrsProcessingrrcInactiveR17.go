package ies

import "rrc/utils"

// BandNR-prs-ProcessingRRC-Inactive-r17 ::= ENUMERATED
type BandnrPrsProcessingrrcInactiveR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrPrsProcessingrrcInactiveR17EnumeratedNothing = iota
	BandnrPrsProcessingrrcInactiveR17EnumeratedSupported
)
