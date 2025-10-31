package ies

import "rrc/utils"

// MeasAndMobParametersCommon-concurrentMeasGap-r17 ::= CHOICE
const (
	MeasandmobparameterscommonConcurrentmeasgapR17ChoiceNothing = iota
	MeasandmobparameterscommonConcurrentmeasgapR17ChoiceConcurrentperueOnlymeasgapR17
	MeasandmobparameterscommonConcurrentmeasgapR17ChoiceConcurrentperuePerfrcombmeasgapR17
)

type MeasandmobparameterscommonConcurrentmeasgapR17 struct {
	Choice                             uint64
	ConcurrentperueOnlymeasgapR17      *utils.ENUMERATED
	ConcurrentperuePerfrcombmeasgapR17 *utils.ENUMERATED
}
