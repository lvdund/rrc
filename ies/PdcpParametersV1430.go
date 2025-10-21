package ies

import "rrc/utils"

// PDCP-Parameters-v1430 ::= SEQUENCE
type PdcpParametersV1430 struct {
	SupporteduplinkonlyrohcProfilesR14 interface{}
	MaxnumberrohcContextsessionsR14    utils.ENUMERATED
}
