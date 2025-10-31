package ies

import "rrc/utils"

// MAC-ParametersXDD-Diff-enhancedSkipUplinkTxDynamic-r16 ::= ENUMERATED
type MacParametersxddDiffEnhancedskipuplinktxdynamicR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersxddDiffEnhancedskipuplinktxdynamicR16EnumeratedNothing = iota
	MacParametersxddDiffEnhancedskipuplinktxdynamicR16EnumeratedSupported
)
