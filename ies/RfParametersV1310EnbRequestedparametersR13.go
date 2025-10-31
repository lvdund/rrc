package ies

import "rrc/utils"

// RF-Parameters-v1310-eNB-RequestedParameters-r13 ::= SEQUENCE
type RfParametersV1310EnbRequestedparametersR13 struct {
	ReducedintnoncontcombrequestedR13 *bool
	RequestedccsdlR13                 *utils.INTEGER `lb:0,ub:32`
	RequestedccsulR13                 *utils.INTEGER `lb:0,ub:32`
	SkipfallbackcombrequestedR13      *bool
}
