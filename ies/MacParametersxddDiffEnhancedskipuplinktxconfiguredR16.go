package ies

import "rrc/utils"

// MAC-ParametersXDD-Diff-enhancedSkipUplinkTxConfigured-r16 ::= ENUMERATED
type MacParametersxddDiffEnhancedskipuplinktxconfiguredR16 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersxddDiffEnhancedskipuplinktxconfiguredR16EnumeratedNothing = iota
	MacParametersxddDiffEnhancedskipuplinktxconfiguredR16EnumeratedSupported
)
