package ies

import "rrc/utils"

// DL-UL-CCs-r15 ::= SEQUENCE
type DlUlCcsR15 struct {
	MaxnumberdlCcsR15 *utils.INTEGER
	MaxnumberulCcsR15 *utils.INTEGER
}
