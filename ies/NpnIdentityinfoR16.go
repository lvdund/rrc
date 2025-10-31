package ies

import "rrc/utils"

// NPN-IdentityInfo-r16 ::= SEQUENCE
// Extensible
type NpnIdentityinfoR16 struct {
	NpnIdentitylistR16            []NpnIdentityR16 `lb:1,ub:maxNPNR16`
	TrackingareacodeR16           Trackingareacode
	RanacR16                      *RanAreacode
	CellidentityR16               Cellidentity
	CellreservedforoperatoruseR16 NpnIdentityinfoR16CellreservedforoperatoruseR16
	IabSupportR16                 *utils.BOOLEAN
	GnbIdLengthR17                *utils.INTEGER `lb:0,ub:32,ext`
}
