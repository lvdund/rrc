package ies

import "rrc/utils"

// CFRA-resources-ssb ::= SEQUENCE
type CfraResourcesSsb struct {
	SsbResourcelist        []CfraSsbResource `lb:1,ub:maxRASsbResources`
	RaSsbOccasionmaskindex utils.INTEGER     `lb:0,ub:15`
}
