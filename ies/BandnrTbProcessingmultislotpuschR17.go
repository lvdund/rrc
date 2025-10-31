package ies

import "rrc/utils"

// BandNR-tb-ProcessingMultiSlotPUSCH-r17 ::= ENUMERATED
type BandnrTbProcessingmultislotpuschR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrTbProcessingmultislotpuschR17EnumeratedNothing = iota
	BandnrTbProcessingmultislotpuschR17EnumeratedSupported
)
