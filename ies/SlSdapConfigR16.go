package ies

import "rrc/utils"

// SL-SDAP-Config-r16 ::= SEQUENCE
// Extensible
type SlSdapConfigR16 struct {
	SlSdapHeaderR16     SlSdapConfigR16SlSdapHeaderR16
	SlDefaultrbR16      utils.BOOLEAN
	SlMappedqosFlowsR16 *SlSdapConfigR16SlMappedqosFlowsR16
	SlCasttypeR16       *SlSdapConfigR16SlCasttypeR16
}
