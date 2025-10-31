package ies

import "rrc/utils"

// BandNR-ta-BasedPDC-NTN-SharedSpectrumChAccess-r17 ::= ENUMERATED
type BandnrTaBasedpdcNtnSharedspectrumchaccessR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrTaBasedpdcNtnSharedspectrumchaccessR17EnumeratedNothing = iota
	BandnrTaBasedpdcNtnSharedspectrumchaccessR17EnumeratedSupported
)
