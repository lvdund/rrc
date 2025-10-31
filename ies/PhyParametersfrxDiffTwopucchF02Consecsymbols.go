package ies

import "rrc/utils"

// Phy-ParametersFRX-Diff-twoPUCCH-F0-2-ConsecSymbols ::= ENUMERATED
type PhyParametersfrxDiffTwopucchF02Consecsymbols struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersfrxDiffTwopucchF02ConsecsymbolsEnumeratedNothing = iota
	PhyParametersfrxDiffTwopucchF02ConsecsymbolsEnumeratedSupported
)
