package ies

import "rrc/utils"

// PhyLayerParameters-v1310-supportedBlindDecoding-r13-pdcch-CandidateReductions-r13 ::= ENUMERATED
type PhylayerparametersV1310SupportedblinddecodingR13PdcchCandidatereductionsR13 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1310SupportedblinddecodingR13PdcchCandidatereductionsR13EnumeratedNothing = iota
	PhylayerparametersV1310SupportedblinddecodingR13PdcchCandidatereductionsR13EnumeratedSupported
)
