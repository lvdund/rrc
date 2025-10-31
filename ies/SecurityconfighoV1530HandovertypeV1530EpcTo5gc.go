package ies

import "rrc/utils"

// SecurityConfigHO-v1530-handoverType-v1530-epc-To5GC ::= SEQUENCE
type SecurityconfighoV1530HandovertypeV1530EpcTo5gc struct {
	SecurityalgorithmconfigR15 Securityalgorithmconfig
	NasContainerR15            utils.OCTETSTRING
}
