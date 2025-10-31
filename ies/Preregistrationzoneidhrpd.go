package ies

import "rrc/utils"

// PreRegistrationZoneIdHRPD ::= utils.INTEGER (0..255)
type Preregistrationzoneidhrpd struct {
	Value utils.INTEGER `lb:0,ub:255`
}
