package ies

// MeasResultServFreq-r13 ::= SEQUENCE
// Extensible
type MeasresultservfreqR13 struct {
	ServfreqidR13              ServcellindexR13
	MeasresultscellR13         *MeasresultservfreqR13MeasresultscellR13
	MeasresultbestneighcellR13 *MeasresultservfreqR13MeasresultbestneighcellR13
}
