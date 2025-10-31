package ies

import "rrc/utils"

// SecurityConfigHO-v1530-handoverType-v1530-intra5GC ::= SEQUENCE
type SecurityconfighoV1530HandovertypeV1530Intra5gc struct {
	SecurityalgorithmconfigR15 *Securityalgorithmconfig
	KeychangeindicatorR15      utils.BOOLEAN
	NexthopchainingcountR15    Nexthopchainingcount
	NasContainerR15            *utils.OCTETSTRING
}
