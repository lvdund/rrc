package ies

import "rrc/utils"

// PhyLayerParameters-v1310-supportedBlindDecoding-r13-skipMonitoringDCI-Format0-1A-r13 ::= ENUMERATED
type PhylayerparametersV1310SupportedblinddecodingR13SkipmonitoringdciFormat01aR13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1310SupportedblinddecodingR13SkipmonitoringdciFormat01aR13EnumeratedNothing = iota
	PhylayerparametersV1310SupportedblinddecodingR13SkipmonitoringdciFormat01aR13EnumeratedSupported
)
