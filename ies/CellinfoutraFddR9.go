package ies

import "rrc/utils"

// CellInfoUTRA-FDD-r9 ::= SEQUENCE
type CellinfoutraFddR9 struct {
	PhyscellidR9        PhyscellidutraFdd
	UtraBcchContainerR9 utils.OCTETSTRING
}
