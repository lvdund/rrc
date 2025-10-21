package ies

import "rrc/utils"

// LAA-Parameters-r13 ::= SEQUENCE
type LaaParametersR13 struct {
	CrosscarrierschedulinglaaDlR13 *utils.ENUMERATED
	CsiRsDrsRrmMeasurementslaaR13  *utils.ENUMERATED
	DownlinklaaR13                 *utils.ENUMERATED
	EndingdwptsR13                 *utils.ENUMERATED
	SecondslotstartingpositionR13  *utils.ENUMERATED
	Tm9LaaR13                      *utils.ENUMERATED
	Tm10LaaR13                     *utils.ENUMERATED
}
