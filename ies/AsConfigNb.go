package ies

import "rrc/utils"

// AS-Config-NB ::= SEQUENCE
// Extensible
type AsConfigNb struct {
	SourceradioresourceconfigR13     RadioresourceconfigdedicatedNbR13
	SourcesecurityalgorithmconfigR13 Securityalgorithmconfig
	SourceueIdentityR13              CRnti
	SourcedlCarrierfreqR13           CarrierfreqNbR13
}
