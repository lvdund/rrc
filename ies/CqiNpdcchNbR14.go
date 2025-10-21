package ies

import "rrc/utils"

// CQI-NPDCCH-NB-r14 ::= ENUMERATED
type CqiNpdcchNbR14 struct {
	Value utils.ENUMERATED
}

const (
	CqiNpdcchNbR14Nomeasurements = 0
	CqiNpdcchNbR14CandidaterepA  = 1
	CqiNpdcchNbR14CandidaterepB  = 2
	CqiNpdcchNbR14CandidaterepC  = 3
	CqiNpdcchNbR14CandidaterepD  = 4
	CqiNpdcchNbR14CandidaterepE  = 5
	CqiNpdcchNbR14CandidaterepF  = 6
	CqiNpdcchNbR14CandidaterepG  = 7
	CqiNpdcchNbR14CandidaterepH  = 8
	CqiNpdcchNbR14CandidaterepI  = 9
	CqiNpdcchNbR14CandidaterepJ  = 10
	CqiNpdcchNbR14CandidaterepK  = 11
	CqiNpdcchNbR14CandidaterepL  = 12
)
