package ies

import "rrc/utils"

// CSFB-RegistrationParam1XRTT ::= SEQUENCE
type CsfbRegistrationparam1xrtt struct {
	Sid                utils.BITSTRING
	Nid                utils.BITSTRING
	Multiplesid        bool
	Multiplenid        bool
	Homereg            bool
	Foreignsidreg      bool
	Foreignnidreg      bool
	Parameterreg       bool
	Powerupreg         bool
	Registrationperiod utils.BITSTRING
	Registrationzone   utils.BITSTRING
	Totalzone          utils.BITSTRING
	Zonetimer          utils.BITSTRING
}
