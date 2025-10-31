package ies

import "rrc/utils"

// CellInfoUTRA-TDD-r10 ::= SEQUENCE
type CellinfoutraTddR10 struct {
	PhyscellidR10        PhyscellidutraTdd
	CarrierfreqR10       ArfcnValueutra
	UtraBcchContainerR10 utils.OCTETSTRING
}
