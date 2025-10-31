package ies

import "rrc/utils"

// BandNR-cqi-4-BitsSubbandNTN-SharedSpectrumChAccess-r17 ::= ENUMERATED
type BandnrCqi4BitssubbandntnSharedspectrumchaccessR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrCqi4BitssubbandntnSharedspectrumchaccessR17EnumeratedNothing = iota
	BandnrCqi4BitssubbandntnSharedspectrumchaccessR17EnumeratedSupported
)
