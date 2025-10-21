package ies

import "rrc/utils"

// MeasResultList2UTRA-r9 ::= SEQUENCE OF MeasResult2UTRA-r9
// SIZE (1..maxFreq)
type Measresultlist2utraR9 struct {
	Value []Measresult2utraR9
}
