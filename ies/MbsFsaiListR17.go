package ies

// MBS-FSAI-List-r17 ::= SEQUENCE OF MBS-FSAI-r17
// SIZE (1..maxFSAI-MBS-r17)
type MbsFsaiListR17 struct {
	Value []MbsFsaiR17 `lb:1,ub:maxFSAIMbsR17`
}
