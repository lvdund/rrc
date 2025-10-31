package ies

// MeasResultServFreq-r10 ::= SEQUENCE
// Extensible
type MeasresultservfreqR10 struct {
	ServfreqidR10              ServcellindexR10
	MeasresultscellR10         *MeasresultservfreqR10MeasresultscellR10
	MeasresultbestneighcellR10 *MeasresultservfreqR10MeasresultbestneighcellR10
}
