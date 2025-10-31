package ies

import "rrc/utils"

// SI-RequestResources ::= SEQUENCE
type SiRequestresources struct {
	RaPreamblestartindex     utils.INTEGER  `lb:0,ub:63`
	RaAssociationperiodindex *utils.INTEGER `lb:0,ub:15`
	RaSsbOccasionmaskindex   *utils.INTEGER `lb:0,ub:15`
}
