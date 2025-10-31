package ies

import "rrc/utils"

// RRCResume-v1610-IEs-mrdc-SecondaryCellGroup-r16 ::= CHOICE
const (
	RrcresumeV1610IesMrdcSecondarycellgroupR16ChoiceNothing = iota
	RrcresumeV1610IesMrdcSecondarycellgroupR16ChoiceNrScgR16
	RrcresumeV1610IesMrdcSecondarycellgroupR16ChoiceEutraScgR16
)

type RrcresumeV1610IesMrdcSecondarycellgroupR16 struct {
	Choice      uint64
	NrScgR16    *utils.OCTETSTRING
	EutraScgR16 *utils.OCTETSTRING
}
