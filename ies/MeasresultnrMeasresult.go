package ies

// MeasResultNR-measResult ::= SEQUENCE
type MeasresultnrMeasresult struct {
	Cellresults    *MeasresultnrMeasresultCellresults
	Rsindexresults *MeasresultnrMeasresultRsindexresults
}
