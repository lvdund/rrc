package ies

import "rrc/utils"

// RelayParameters-r17-remoteUE-PathSwitchToIdleInactiveRelay-r17 ::= ENUMERATED
type RelayparametersR17RemoteuePathswitchtoidleinactiverelayR17 struct {
	Value utils.ENUMERATED
}

const (
	RelayparametersR17RemoteuePathswitchtoidleinactiverelayR17EnumeratedNothing = iota
	RelayparametersR17RemoteuePathswitchtoidleinactiverelayR17EnumeratedSupported
)
