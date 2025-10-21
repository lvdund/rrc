package ies

import "rrc/utils"

// MeasResultList3EUTRA-r15 ::= SEQUENCE OF MeasResult3EUTRA-r15
// SIZE (1..maxFreq)
type Measresultlist3eutraR15 struct {
	Value []Measresult3eutraR15
}
