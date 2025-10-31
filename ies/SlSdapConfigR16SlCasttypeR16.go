package ies

import "rrc/utils"

// SL-SDAP-Config-r16-sl-CastType-r16 ::= ENUMERATED
type SlSdapConfigR16SlCasttypeR16 struct {
	Value utils.ENUMERATED
}

const (
	SlSdapConfigR16SlCasttypeR16EnumeratedNothing = iota
	SlSdapConfigR16SlCasttypeR16EnumeratedBroadcast
	SlSdapConfigR16SlCasttypeR16EnumeratedGroupcast
	SlSdapConfigR16SlCasttypeR16EnumeratedUnicast
	SlSdapConfigR16SlCasttypeR16EnumeratedSpare1
)
