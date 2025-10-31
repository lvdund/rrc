package ies

import "rrc/utils"

// PhyLayerParameters-v1430-ce-PUSCH-NB-MaxTBS-r14 ::= ENUMERATED
type PhylayerparametersV1430CePuschNbMaxtbsR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430CePuschNbMaxtbsR14EnumeratedNothing = iota
	PhylayerparametersV1430CePuschNbMaxtbsR14EnumeratedSupported
)
