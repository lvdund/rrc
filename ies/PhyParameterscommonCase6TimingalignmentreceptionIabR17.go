package ies

import "rrc/utils"

// Phy-ParametersCommon-case6-TimingAlignmentReception-IAB-r17 ::= ENUMERATED
type PhyParameterscommonCase6TimingalignmentreceptionIabR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonCase6TimingalignmentreceptionIabR17EnumeratedNothing = iota
	PhyParameterscommonCase6TimingalignmentreceptionIabR17EnumeratedSupported
)
