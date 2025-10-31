package ies

import "rrc/utils"

// SDAP-Config-sdap-HeaderUL ::= ENUMERATED
type SdapConfigSdapHeaderul struct {
	Value utils.ENUMERATED
}

const (
	SdapConfigSdapHeaderulEnumeratedNothing = iota
	SdapConfigSdapHeaderulEnumeratedPresent
	SdapConfigSdapHeaderulEnumeratedAbsent
)
