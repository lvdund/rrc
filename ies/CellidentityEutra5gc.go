package ies

import "rrc/utils"

// CellIdentity-EUTRA-5GC ::= CHOICE
const (
	CellidentityEutra5gcChoiceNothing = iota
	CellidentityEutra5gcChoiceCellidentityEutra
	CellidentityEutra5gcChoiceCellidIndex
)

type CellidentityEutra5gc struct {
	Choice            uint64
	CellidentityEutra *utils.BITSTRING `lb:28,ub:28`
	CellidIndex       *utils.INTEGER   `lb:1,ub:maxPLMN`
}
