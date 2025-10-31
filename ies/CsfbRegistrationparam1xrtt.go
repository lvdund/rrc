package ies

import "rrc/utils"

// CSFB-RegistrationParam1XRTT ::= SEQUENCE
type CsfbRegistrationparam1xrtt struct {
	Sid                utils.BITSTRING `lb:15,ub:15`
	Nid                utils.BITSTRING `lb:16,ub:16`
	Multiplesid        utils.BOOLEAN
	Multiplenid        utils.BOOLEAN
	Homereg            utils.BOOLEAN
	Foreignsidreg      utils.BOOLEAN
	Foreignnidreg      utils.BOOLEAN
	Parameterreg       utils.BOOLEAN
	Powerupreg         utils.BOOLEAN
	Registrationperiod utils.BITSTRING `lb:7,ub:7`
	Registrationzone   utils.BITSTRING `lb:12,ub:12`
	Totalzone          utils.BITSTRING `lb:3,ub:3`
	Zonetimer          utils.BITSTRING `lb:3,ub:3`
}
