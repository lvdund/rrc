package ies

import "rrc/utils"

// MobilityStateParameters ::= SEQUENCE
type Mobilitystateparameters struct {
	TEvaluation       utils.ENUMERATED
	THystnormal       utils.ENUMERATED
	NCellchangemedium utils.INTEGER
	NCellchangehigh   utils.INTEGER
}
