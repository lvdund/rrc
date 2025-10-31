package ies

import "rrc/utils"

// CellAccessRelatedInfo ::= SEQUENCE
// Extensible
type Cellaccessrelatedinfo struct {
	PlmnIdentityinfolist        PlmnIdentityinfolist
	Cellreservedforotheruse     *utils.BOOLEAN
	CellreservedforfutureuseR16 *utils.BOOLEAN          `ext`
	NpnIdentityinfolistR16      *NpnIdentityinfolistR16 `ext`
	SnpnAccessinfolistR17       *[]SnpnAccessinfoR17    `lb:1,ub:maxNPNR16,ext`
}
