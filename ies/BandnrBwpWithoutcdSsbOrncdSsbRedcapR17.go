package ies

import "rrc/utils"

// BandNR-bwp-WithoutCD-SSB-OrNCD-SSB-RedCap-r17 ::= ENUMERATED
type BandnrBwpWithoutcdSsbOrncdSsbRedcapR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrBwpWithoutcdSsbOrncdSsbRedcapR17EnumeratedNothing = iota
	BandnrBwpWithoutcdSsbOrncdSsbRedcapR17EnumeratedSupported
)
