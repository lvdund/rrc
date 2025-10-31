package ies

import "rrc/utils"

// BandNR-tb-ProcessingRepMultiSlotPUSCH-r17 ::= ENUMERATED
type BandnrTbProcessingrepmultislotpuschR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrTbProcessingrepmultislotpuschR17EnumeratedNothing = iota
	BandnrTbProcessingrepmultislotpuschR17EnumeratedSupported
)
