package ies

import "rrc/utils"

// CellInfoUTRA-TDD-r9 ::= SEQUENCE
type CellinfoutraTddR9 struct {
	PhyscellidR9        PhyscellidutraTdd
	UtraBcchContainerR9 utils.OCTETSTRING
}
