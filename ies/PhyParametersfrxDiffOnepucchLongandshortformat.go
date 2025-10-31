package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-onePUCCH-LongAndShortFormat ::= ENUMERATED
type PhyParametersfrxDiffOnepucchLongandshortformat struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffOnepucchLongandshortformatEnumeratedNothing = iota
	PhyParametersfrxDiffOnepucchLongandshortformatEnumeratedSupported
)
