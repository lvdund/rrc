package ies

import "rrc/utils"

// SDAP-Config-sdap-HeaderDL ::= ENUMERATED
type SdapConfigSdapHeaderdl struct {
	Value utils.ENUMERATED
}

const (
	SdapConfigSdapHeaderdlEnumeratedNothing = iota
	SdapConfigSdapHeaderdlEnumeratedPresent
	SdapConfigSdapHeaderdlEnumeratedAbsent
)
