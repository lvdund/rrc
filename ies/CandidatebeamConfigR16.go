package ies

// CandidateBeamConfig-r16 ::= CHOICE
const (
	CandidatebeamConfigR16ChoiceNothing = iota
	CandidatebeamConfigR16ChoiceSSBIndex
	CandidatebeamConfigR16ChoiceCsiRsResourceId
)

type CandidatebeamConfigR16 struct {
	Choice   *uint64
	SsbR16   *SsbIndex
	CsiRsR16 *NzpCsiRsResourceid
}