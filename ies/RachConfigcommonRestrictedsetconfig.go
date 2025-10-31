package ies

import "rrc/utils"

// RACH-ConfigCommon-restrictedSetConfig ::= ENUMERATED
type RachConfigcommonRestrictedsetconfig struct {
	Value utils.ENUMERATED
}

const (
	RachConfigcommonRestrictedsetconfigEnumeratedNothing = iota
	RachConfigcommonRestrictedsetconfigEnumeratedUnrestrictedset
	RachConfigcommonRestrictedsetconfigEnumeratedRestrictedsettypea
	RachConfigcommonRestrictedsetconfigEnumeratedRestrictedsettypeb
)
