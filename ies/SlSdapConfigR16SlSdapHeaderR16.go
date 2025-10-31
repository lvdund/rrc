package ies

import "rrc/utils"

// SL-SDAP-Config-r16-sl-SDAP-Header-r16 ::= ENUMERATED
type SlSdapConfigR16SlSdapHeaderR16 struct {
	Value utils.ENUMERATED
}

const (
	SlSdapConfigR16SlSdapHeaderR16EnumeratedNothing = iota
	SlSdapConfigR16SlSdapHeaderR16EnumeratedPresent
	SlSdapConfigR16SlSdapHeaderR16EnumeratedAbsent
)
