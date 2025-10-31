package ies

import "rrc/utils"

// T-Reselection-NB-r13 ::= ENUMERATED
type TReselectionNbR13 struct {
	Value utils.ENUMERATED
}

const (
	TReselectionNbR13EnumeratedNothing = iota
	TReselectionNbR13EnumeratedS0
	TReselectionNbR13EnumeratedS3
	TReselectionNbR13EnumeratedS6
	TReselectionNbR13EnumeratedS9
	TReselectionNbR13EnumeratedS12
	TReselectionNbR13EnumeratedS15
	TReselectionNbR13EnumeratedS18
	TReselectionNbR13EnumeratedS21
)
