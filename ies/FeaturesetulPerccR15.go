package ies

import "rrc/utils"

// FeatureSetUL-PerCC-r15 ::= SEQUENCE
type FeaturesetulPerccR15 struct {
	SupportedmimoCapabilityulR15 *MimoCapabilityulR10
	Ul256qamR15                  *utils.ENUMERATED
}
