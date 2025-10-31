package ies

import "rrc/utils"

// PhyLayerParameters-v1430-ce-PDSCH-PUSCH-MaxBandwidth-r14 ::= ENUMERATED
type PhylayerparametersV1430CePdschPuschMaxbandwidthR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430CePdschPuschMaxbandwidthR14EnumeratedNothing = iota
	PhylayerparametersV1430CePdschPuschMaxbandwidthR14EnumeratedBw5
	PhylayerparametersV1430CePdschPuschMaxbandwidthR14EnumeratedBw20
)
