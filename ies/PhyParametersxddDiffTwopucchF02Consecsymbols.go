package ies

import "rrc/utils"

// Phy-ParametersXDD-Diff-twoPUCCH-F0-2-ConsecSymbols ::= ENUMERATED
type PhyParametersxddDiffTwopucchF02Consecsymbols struct {
	Value utils.ENUMERATED
}

const (
	PhyParametersxddDiffTwopucchF02ConsecsymbolsEnumeratedNothing = iota
	PhyParametersxddDiffTwopucchF02ConsecsymbolsEnumeratedSupported
)
