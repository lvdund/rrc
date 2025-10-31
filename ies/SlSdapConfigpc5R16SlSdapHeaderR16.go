package ies

import "rrc/utils"

// SL-SDAP-ConfigPC5-r16-sl-SDAP-Header-r16 ::= ENUMERATED
type SlSdapConfigpc5R16SlSdapHeaderR16 struct {
	Value utils.ENUMERATED
}

const (
	SlSdapConfigpc5R16SlSdapHeaderR16EnumeratedNothing = iota
	SlSdapConfigpc5R16SlSdapHeaderR16EnumeratedPresent
	SlSdapConfigpc5R16SlSdapHeaderR16EnumeratedAbsent
)
