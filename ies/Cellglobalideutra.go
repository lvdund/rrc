package ies

import "rrc/utils"

// CellGlobalIdEUTRA ::= SEQUENCE
type Cellglobalideutra struct {
	PlmnIdentity PlmnIdentity
	Cellidentity Cellidentity
}
