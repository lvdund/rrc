package ies

import "rrc/utils"

// MobilityStateParameters ::= SEQUENCE
type Mobilitystateparameters struct {
	TEvaluation       MobilitystateparametersTEvaluation
	THystnormal       MobilitystateparametersTHystnormal
	NCellchangemedium utils.INTEGER `lb:0,ub:16`
	NCellchangehigh   utils.INTEGER `lb:0,ub:16`
}
