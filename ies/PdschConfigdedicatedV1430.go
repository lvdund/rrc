package ies

import "rrc/utils"

// PDSCH-ConfigDedicated-v1430 ::= SEQUENCE
type PdschConfigdedicatedV1430 struct {
	CePdschMaxbandwidthR14     *utils.ENUMERATED
	CePdschTenprocessesR14     *utils.ENUMERATED
	CeHarqAckbundlingR14       *utils.ENUMERATED
	CeSchedulingenhancementR14 *utils.ENUMERATED
	Tbsindexalt2R14            *utils.ENUMERATED
}
