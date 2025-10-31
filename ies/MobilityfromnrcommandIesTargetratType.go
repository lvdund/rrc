package ies

import "rrc/utils"

// MobilityFromNRCommand-IEs-targetRAT-Type ::= utils.ENUMERATED // Extensible
type MobilityfromnrcommandIesTargetratType struct {
	Value utils.ENUMERATED
}

const (
	MobilityfromnrcommandIesTargetratTypeEnumeratedNothing = iota
	MobilityfromnrcommandIesTargetratTypeEnumeratedEutra
	MobilityfromnrcommandIesTargetratTypeEnumeratedUtra_Fdd_V1610
	MobilityfromnrcommandIesTargetratTypeEnumeratedSpare2
	MobilityfromnrcommandIesTargetratTypeEnumeratedSpare1
)
