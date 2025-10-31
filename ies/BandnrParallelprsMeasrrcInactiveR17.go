package ies

import "rrc/utils"

// BandNR-parallelPRS-MeasRRC-Inactive-r17 ::= ENUMERATED
type BandnrParallelprsMeasrrcInactiveR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrParallelprsMeasrrcInactiveR17EnumeratedNothing = iota
	BandnrParallelprsMeasrrcInactiveR17EnumeratedSupported
)
