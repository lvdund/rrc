package ies

// FeatureSetDL-r15 ::= SEQUENCE
type FeaturesetdlR15 struct {
	MimoCaParametersperbobcR15 *MimoCaParametersperbobcR15
	FeaturesetperccListdlR15   []FeaturesetdlPerccIdR15 `lb:1,ub:maxServCellR13`
}
