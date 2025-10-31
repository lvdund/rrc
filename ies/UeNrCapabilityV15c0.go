package ies

import "rrc/utils"

// UE-NR-Capability-v15c0 ::= SEQUENCE
type UeNrCapabilityV15c0 struct {
	NrdcParametersV15c0     *NrdcParametersV15c0
	Partialfr2FallbackrxReq *utils.BOOLEAN
	Noncriticalextension    *UeNrCapabilityV15g0
}
