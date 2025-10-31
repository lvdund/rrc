package ies

import "rrc/utils"

// RelayParameters-r17-remoteUE-Operation-L2-r17 ::= ENUMERATED
type RelayparametersR17RemoteueOperationL2R17 struct {
	Value utils.ENUMERATED
}

const (
	RelayparametersR17RemoteueOperationL2R17EnumeratedNothing = iota
	RelayparametersR17RemoteueOperationL2R17EnumeratedSupported
)
