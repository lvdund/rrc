package ies

import "rrc/utils"

// CellAccessRelatedInfo-EUTRA-EPC ::= SEQUENCE
type CellaccessrelatedinfoEutraEpc struct {
	PlmnIdentitylistEutraEpc PlmnIdentitylistEutraEpc
	TrackingareacodeEutraEpc utils.BITSTRING `lb:16,ub:16`
	CellidentityEutraEpc     utils.BITSTRING `lb:28,ub:28`
}
