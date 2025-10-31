package ies

// DL-PPW-PreConfigToAddModList-r17 ::= SEQUENCE OF DL-PPW-PreConfig-r17
// SIZE (1..maxNrofPPW-Config-r17)
type DlPpwPreconfigtoaddmodlistR17 struct {
	Value []DlPpwPreconfigR17 `lb:1,ub:maxNrofPPWConfigR17`
}
