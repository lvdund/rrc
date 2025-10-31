package ies

import "rrc/utils"

// RelayParameters-r17-relayUE-Operation-L2-r17 ::= ENUMERATED
type RelayparametersR17RelayueOperationL2R17 struct {
	Value utils.ENUMERATED
}

const (
	RelayparametersR17RelayueOperationL2R17EnumeratedNothing = iota
	RelayparametersR17RelayueOperationL2R17EnumeratedSupported
)
