package ies

import "rrc/utils"

// CSI-RS-ConfigNonPrecoded-r13 ::= SEQUENCE
type CsiRsConfignonprecodedR13 struct {
	PCAndcbsrlistR13                    *PCAndcbsrPairR13
	Codebookconfign1R13                 utils.ENUMERATED
	Codebookconfign2R13                 utils.ENUMERATED
	CodebookoversamplingrateconfigO1R13 *utils.ENUMERATED
	CodebookoversamplingrateconfigO2R13 *utils.ENUMERATED
	CodebookconfigR13                   utils.INTEGER
	CsiImConfigidlistR13                *interface{}
	CsiRsConfignzpEmimoR13              *CsiRsConfignzpEmimoR13
}
