package ies

import "rrc/utils"

// BandNR-halfDuplexFDD-TypeA-RedCap-r17 ::= ENUMERATED
type BandnrHalfduplexfddTypeaRedcapR17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrHalfduplexfddTypeaRedcapR17EnumeratedNothing = iota
	BandnrHalfduplexfddTypeaRedcapR17EnumeratedSupported
)
