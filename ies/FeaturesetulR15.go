package ies

// FeatureSetUL-r15 ::= SEQUENCE
type FeaturesetulR15 struct {
	FeaturesetperccListulR15 []FeaturesetulPerccIdR15 `lb:1,ub:maxServCellR13`
}
