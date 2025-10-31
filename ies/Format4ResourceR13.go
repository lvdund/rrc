package ies

import "rrc/utils"

// Format4-resource-r13 ::= SEQUENCE
type Format4ResourceR13 struct {
	StartingprbFormat4R13 utils.INTEGER `lb:0,ub:109`
	NumberofprbFormat4R13 utils.INTEGER `lb:0,ub:7`
}
