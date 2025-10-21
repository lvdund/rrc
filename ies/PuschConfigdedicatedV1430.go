package ies

import "rrc/utils"

// PUSCH-ConfigDedicated-v1430 ::= SEQUENCE
type PuschConfigdedicatedV1430 struct {
	CePuschNbMaxtbsR14     *utils.ENUMERATED
	CePuschMaxbandwidthR14 *utils.ENUMERATED
	TddPuschUpptsR14       *TddPuschUpptsR14
	UlDmrsIfdmaR14         bool
	Enable256qamR14        *Enable256qamR14
}
