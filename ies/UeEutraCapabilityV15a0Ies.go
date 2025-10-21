package ies

import "rrc/utils"

// UE-EUTRA-Capability-v15a0-IEs ::= SEQUENCE
type UeEutraCapabilityV15a0Ies struct {
	NeighcellsiAcquisitionparametersV15a0 NeighcellsiAcquisitionparametersV15a0
	Eutra5gcParametersR15                 *Eutra5gcParametersR15
	FddAddUeEutraCapabilitiesV15a0        *UeEutraCapabilityaddxddModeV15a0
	TddAddUeEutraCapabilitiesV15a0        *UeEutraCapabilityaddxddModeV15a0
	Noncriticalextension                  *UeEutraCapabilityV1610Ies
}
