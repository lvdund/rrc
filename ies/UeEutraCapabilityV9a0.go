package ies

import "rrc/utils"

// UE-EUTRA-Capability-v9a0-IEs ::= SEQUENCE
type UeEutraCapabilityV9a0 struct {
	Featuregroupindrel9addR9    *utils.BITSTRING `lb:32,ub:32`
	FddAddUeEutraCapabilitiesR9 *UeEutraCapabilityaddxddModeR9
	TddAddUeEutraCapabilitiesR9 *UeEutraCapabilityaddxddModeR9
	Noncriticalextension        *UeEutraCapabilityV9c0
}
