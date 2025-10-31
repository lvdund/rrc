package ies

import "rrc/utils"

// PDCCH-ConfigLAA-r14-maxNumberOfSchedSubframes-Format0B-r14 ::= ENUMERATED
type PdcchConfiglaaR14MaxnumberofschedsubframesFormat0bR14 struct {
	Value utils.ENUMERATED
}

const (
	PdcchConfiglaaR14MaxnumberofschedsubframesFormat0bR14EnumeratedNothing = iota
	PdcchConfiglaaR14MaxnumberofschedsubframesFormat0bR14EnumeratedSf2
	PdcchConfiglaaR14MaxnumberofschedsubframesFormat0bR14EnumeratedSf3
	PdcchConfiglaaR14MaxnumberofschedsubframesFormat0bR14EnumeratedSf4
)
