package ies

import "rrc/utils"

// MeasResultServFreq-r13 ::= SEQUENCE
// Extensible
type MeasresultservfreqR13 struct {
	ServfreqidR13              ServcellindexR13
	MeasresultscellR13         *interface{}
	MeasresultbestneighcellR13 *interface{}
}
