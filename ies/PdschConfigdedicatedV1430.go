package ies

// PDSCH-ConfigDedicated-v1430 ::= SEQUENCE
type PdschConfigdedicatedV1430 struct {
	CePdschMaxbandwidthR14     *PdschConfigdedicatedV1430CePdschMaxbandwidthR14
	CePdschTenprocessesR14     *PdschConfigdedicatedV1430CePdschTenprocessesR14
	CeHarqAckbundlingR14       *PdschConfigdedicatedV1430CeHarqAckbundlingR14
	CeSchedulingenhancementR14 *PdschConfigdedicatedV1430CeSchedulingenhancementR14
	Tbsindexalt2R14            *PdschConfigdedicatedV1430Tbsindexalt2R14
}
