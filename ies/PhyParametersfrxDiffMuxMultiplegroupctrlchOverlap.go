package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-mux-MultipleGroupCtrlCH-Overlap ::= ENUMERATED
type PhyParametersfrxDiffMuxMultiplegroupctrlchOverlap struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffMuxMultiplegroupctrlchOverlapEnumeratedNothing = iota
	PhyParametersfrxDiffMuxMultiplegroupctrlchOverlapEnumeratedSupported
)
