package ies

import "rrc/utils"

// PDCCH-ConfigLAA-r14-maxNumberOfSchedSubframes-Format4B-r14 ::= ENUMERATED
type PdcchConfiglaaR14MaxnumberofschedsubframesFormat4bR14 struct {
	Value utils.ENUMERATED
}

const (
	PdcchConfiglaaR14MaxnumberofschedsubframesFormat4bR14EnumeratedNothing = iota
	PdcchConfiglaaR14MaxnumberofschedsubframesFormat4bR14EnumeratedSf2
	PdcchConfiglaaR14MaxnumberofschedsubframesFormat4bR14EnumeratedSf3
	PdcchConfiglaaR14MaxnumberofschedsubframesFormat4bR14EnumeratedSf4
)
