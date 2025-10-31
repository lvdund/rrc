package ies

import "rrc/utils"

// SRS-Resource-partialFreqSounding-r17-startRBIndexFScaling-r17 ::= CHOICE
const (
	SrsResourcePartialfreqsoundingR17StartrbindexfscalingR17ChoiceNothing = iota
	SrsResourcePartialfreqsoundingR17StartrbindexfscalingR17ChoiceStartrbindexandfreqscalingfactor2R17
	SrsResourcePartialfreqsoundingR17StartrbindexfscalingR17ChoiceStartrbindexandfreqscalingfactor4R17
)

type SrsResourcePartialfreqsoundingR17StartrbindexfscalingR17 struct {
	Choice                               uint64
	Startrbindexandfreqscalingfactor2R17 *utils.INTEGER `lb:0,ub:1`
	Startrbindexandfreqscalingfactor4R17 *utils.INTEGER `lb:0,ub:3`
}
