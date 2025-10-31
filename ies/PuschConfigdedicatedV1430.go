package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-v1430 ::= SEQUENCE
type PuschConfigdedicatedV1430 struct {
	CePuschNbMaxtbsR14     *PuschConfigdedicatedV1430CePuschNbMaxtbsR14
	CePuschMaxbandwidthR14 *PuschConfigdedicatedV1430CePuschMaxbandwidthR14
	TddPuschUpptsR14       *TddPuschUpptsR14
	UlDmrsIfdmaR14         utils.BOOLEAN
	Enable256qamR14        *Enable256qamR14
}
