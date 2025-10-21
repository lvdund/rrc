package ies

import "rrc/utils"

// Format5-resource-r13 ::= SEQUENCE
type Format5ResourceR13 struct {
	StartingprbFormat5R13 utils.INTEGER
	CdmIndexFormat5R13    utils.INTEGER
}
