package ies

import "rrc/utils"

// BandNR-searchSpaceSetGrp-switchCap2-r17 ::= ENUMERATED
type BandnrSearchspacesetgrpSwitchcap2R17 struct {
	Value utils.ENUMERATED
}

const (
	BandnrSearchspacesetgrpSwitchcap2R17EnumeratedNothing = iota
	BandnrSearchspacesetgrpSwitchcap2R17EnumeratedSupported
)
