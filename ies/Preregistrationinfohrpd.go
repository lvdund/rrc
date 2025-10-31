package ies

import "rrc/utils"

// PreRegistrationInfoHRPD ::= SEQUENCE
type Preregistrationinfohrpd struct {
	Preregistrationallowed             utils.BOOLEAN
	Preregistrationzoneid              *Preregistrationzoneidhrpd
	Secondarypreregistrationzoneidlist *Secondarypreregistrationzoneidlisthrpd
}
