package ies

import "rrc/utils"

// PhyLayerParameters-v1530-urllc-Capabilities-r15-pusch-SPS-SlotRepPSCell-r15 ::= ENUMERATED
type PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSlotreppscellR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSlotreppscellR15EnumeratedNothing = iota
	PhylayerparametersV1530UrllcCapabilitiesR15PuschSpsSlotreppscellR15EnumeratedSupported
)
