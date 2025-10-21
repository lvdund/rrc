package ies

import "rrc/utils"

// SecondaryPreRegistrationZoneIdListHRPD ::= SEQUENCE OF PreRegistrationZoneIdHRPD
// SIZE (1..2)
type Secondarypreregistrationzoneidlisthrpd struct {
	Value utils.Sequence[Preregistrationzoneidhrpd]
}
