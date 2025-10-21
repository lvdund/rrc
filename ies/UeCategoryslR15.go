package ies

import "rrc/utils"

// UE-CategorySL-r15 ::= SEQUENCE
type UeCategoryslR15 struct {
	UeCategoryslCTxR15 utils.INTEGER
	UeCategoryslCRxR15 utils.INTEGER
}
