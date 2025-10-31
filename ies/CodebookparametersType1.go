package ies

// CodebookParameters-type1 ::= SEQUENCE
type CodebookparametersType1 struct {
	Singlepanel CodebookparametersType1Singlepanel
	Multipanel  *CodebookparametersType1Multipanel
}
