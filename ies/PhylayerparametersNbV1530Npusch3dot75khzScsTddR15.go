package ies

import "rrc/utils"

// PhyLayerParameters-NB-v1530-npusch-3dot75kHz-SCS-TDD-r15 ::= ENUMERATED
type PhylayerparametersNbV1530Npusch3dot75khzScsTddR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersNbV1530Npusch3dot75khzScsTddR15EnumeratedNothing = iota
	PhylayerparametersNbV1530Npusch3dot75khzScsTddR15EnumeratedSupported
)
