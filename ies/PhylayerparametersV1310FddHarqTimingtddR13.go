package ies

import "rrc/utils"

// PhyLayerParameters-v1310-fdd-HARQ-TimingTDD-r13 ::= ENUMERATED
type PhylayerparametersV1310FddHarqTimingtddR13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1310FddHarqTimingtddR13EnumeratedNothing = iota
	PhylayerparametersV1310FddHarqTimingtddR13EnumeratedSupported
)
