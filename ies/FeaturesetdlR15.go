package ies

import "rrc/utils"

// FeatureSetDL-r15 ::= SEQUENCE
type FeaturesetdlR15 struct {
	MimoCaParametersperbobcR15 *MimoCaParametersperbobcR15
	FeaturesetperccListdlR15   interface{}
}
