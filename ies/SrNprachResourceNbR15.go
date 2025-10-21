package ies

import "rrc/utils"

// SR-NPRACH-Resource-NB-r15 ::= SEQUENCE
type SrNprachResourceNbR15 struct {
	NprachCarrierindexR15    utils.INTEGER
	NprachResourceindexR15   utils.INTEGER
	NprachSubcarrierindexR15 interface{}
	P0SrR15                  utils.INTEGER
	AlphaR15                 utils.ENUMERATED
}
