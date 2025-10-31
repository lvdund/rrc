package ies

import "rrc/utils"

// Phy-ParametersCommon-srs-AdditionalRepetition-r17 ::= ENUMERATED
type PhyParameterscommonSrsAdditionalrepetitionR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSrsAdditionalrepetitionR17EnumeratedNothing = iota
	PhyParameterscommonSrsAdditionalrepetitionR17EnumeratedSupported
)
