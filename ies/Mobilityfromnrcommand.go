package ies

import "rrc/utils"

// MobilityFromNRCommand-IEs ::= SEQUENCE
// Extensible
type Mobilityfromnrcommand struct {
	TargetratType             MobilityfromnrcommandIesTargetratType
	TargetratMessagecontainer utils.OCTETSTRING
	NasSecurityparamfromnr    *utils.OCTETSTRING
	Latenoncriticalextension  *utils.OCTETSTRING
	Noncriticalextension      *MobilityfromnrcommandV1610
}
