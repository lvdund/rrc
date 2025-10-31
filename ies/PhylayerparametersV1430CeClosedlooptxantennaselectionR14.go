package ies

import "rrc/utils"

// PhyLayerParameters-v1430-ce-ClosedLoopTxAntennaSelection-r14 ::= ENUMERATED
type PhylayerparametersV1430CeClosedlooptxantennaselectionR14 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1430CeClosedlooptxantennaselectionR14EnumeratedNothing = iota
	PhylayerparametersV1430CeClosedlooptxantennaselectionR14EnumeratedSupported
)
