package ies

import "rrc/utils"

// Phy-ParametersCommon-case7-TimingAlignmentReception-IAB-r17 ::= ENUMERATED
type PhyParameterscommonCase7TimingalignmentreceptionIabR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonCase7TimingalignmentreceptionIabR17EnumeratedNothing = iota
	PhyParameterscommonCase7TimingalignmentreceptionIabR17EnumeratedSupported
)
