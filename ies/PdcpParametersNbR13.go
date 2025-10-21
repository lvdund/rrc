package ies

import "rrc/utils"

// PDCP-Parameters-NB-r13 ::= SEQUENCE
// Extensible
type PdcpParametersNbR13 struct {
	SupportedrohcProfilesR13        interface{}
	MaxnumberrohcContextsessionsR13 utils.ENUMERATED
}
