package ies

import "rrc/utils"

// UE-EUTRA-Capability-v920-IEs-deviceType-r9 ::= ENUMERATED
type UeEutraCapabilityV920IesDevicetypeR9 struct {
	Value utils.ENUMERATED
}

const (
	UeEutraCapabilityV920IesDevicetypeR9EnumeratedNothing = iota
	UeEutraCapabilityV920IesDevicetypeR9EnumeratedNobenfrombatconsumpopt
)
