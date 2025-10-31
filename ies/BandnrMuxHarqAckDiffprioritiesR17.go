package ies

import "rrc/utils"

// BandNR-mux-HARQ-ACK-DiffPriorities-r17 ::= ENUMERATED
type BandnrMuxHarqAckDiffprioritiesR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrMuxHarqAckDiffprioritiesR17EnumeratedNothing = iota
	BandnrMuxHarqAckDiffprioritiesR17EnumeratedSupported
)
