package ies

import "rrc/utils"

// SPUCCH-Config-r15-setup-dummy ::= SEQUENCE
type SpucchConfigR15SetupDummy struct {
	N3spucchAnListR15 []utils.INTEGER `lb:1,ub:4`
}
