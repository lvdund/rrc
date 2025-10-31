package ies

import "rrc/utils"

// SecurityConfigHO-handoverType-interRAT ::= SEQUENCE
type SecurityconfighoHandovertypeInterrat struct {
	Securityalgorithmconfig Securityalgorithmconfig
	NasSecurityparamtoeutra utils.OCTETSTRING `lb:6,ub:6`
}
