package ies

import "rrc/utils"

// CellAccessRelatedInfo-r14 ::= SEQUENCE
type CellaccessrelatedinfoR14 struct {
	PlmnIdentitylistR14 PlmnIdentitylist
	TrackingareacodeR14 Trackingareacode
	CellidentityR14     Cellidentity
}
