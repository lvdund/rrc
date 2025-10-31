package ies

import "rrc/utils"

// PLMN-IdentityInfo ::= SEQUENCE
// Extensible
type PlmnIdentityinfo struct {
	PlmnIdentitylist           []PlmnIdentity `lb:1,ub:maxPLMN`
	Trackingareacode           *Trackingareacode
	Ranac                      *RanAreacode
	Cellidentity               Cellidentity
	Cellreservedforoperatoruse PlmnIdentityinfoCellreservedforoperatoruse
	IabSupportR16              *utils.BOOLEAN      `ext`
	TrackingarealistR17        *[]Trackingareacode `lb:1,ub:maxTACR17,ext`
	GnbIdLengthR17             *utils.INTEGER      `lb:0,ub:32,ext`
}
