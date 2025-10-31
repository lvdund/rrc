package ies

import "rrc/utils"

// BandNR-dmrs-BundlingRestart-r17 ::= ENUMERATED
type BandnrDmrsBundlingrestartR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrDmrsBundlingrestartR17EnumeratedNothing = iota
	BandnrDmrsBundlingrestartR17EnumeratedSupported
)
