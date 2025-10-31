package ies

import "rrc/utils"

// MAC-ParametersXDD-Diff-skipUplinkTxDynamic ::= ENUMERATED
type MacParametersxddDiffSkipuplinktxdynamic struct {
	Value utils.ENUMERATED
}

const (
	MacParametersxddDiffSkipuplinktxdynamicEnumeratedNothing = iota
	MacParametersxddDiffSkipuplinktxdynamicEnumeratedSupported
)
