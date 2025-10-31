package ies

import "rrc/utils"

// MRDC-SecondaryCellGroupConfig-mrdc-SecondaryCellGroup ::= CHOICE
const (
	MrdcSecondarycellgroupconfigMrdcSecondarycellgroupChoiceNothing = iota
	MrdcSecondarycellgroupconfigMrdcSecondarycellgroupChoiceNrScg
	MrdcSecondarycellgroupconfigMrdcSecondarycellgroupChoiceEutraScg
)

type MrdcSecondarycellgroupconfigMrdcSecondarycellgroup struct {
	Choice   uint64
	NrScg    *utils.OCTETSTRING
	EutraScg *utils.OCTETSTRING
}
