package ies

import "rrc/utils"

// SDAP-Config ::= SEQUENCE
// Extensible
type SdapConfig struct {
	PduSession              PduSessionid
	SdapHeaderdl            SdapConfigSdapHeaderdl
	SdapHeaderul            SdapConfigSdapHeaderul
	Defaultdrb              utils.BOOLEAN
	MappedqosFlowstoadd     *[]Qfi `lb:1,ub:maxNrofQFIs`
	MappedqosFlowstorelease *[]Qfi `lb:1,ub:maxNrofQFIs`
}
