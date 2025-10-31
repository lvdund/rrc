package ies

import "rrc/utils"

// CSI-RS-ConfigNonPrecoded-r13 ::= SEQUENCE
type CsiRsConfignonprecodedR13 struct {
	PCAndcbsrlistR13                    *PCAndcbsrPairR13
	Codebookconfign1R13                 CsiRsConfignonprecodedR13Codebookconfign1R13
	Codebookconfign2R13                 CsiRsConfignonprecodedR13Codebookconfign2R13
	CodebookoversamplingrateconfigO1R13 *CsiRsConfignonprecodedR13CodebookoversamplingrateconfigO1R13
	CodebookoversamplingrateconfigO2R13 *CsiRsConfignonprecodedR13CodebookoversamplingrateconfigO2R13
	CodebookconfigR13                   utils.INTEGER       `lb:0,ub:4`
	CsiImConfigidlistR13                *[]CsiImConfigidR13 `lb:1,ub:2`
	CsiRsConfignzpEmimoR13              *CsiRsConfignzpEmimoR13
}
