package ies

// LAA-Parameters-r13 ::= SEQUENCE
type LaaParametersR13 struct {
	CrosscarrierschedulinglaaDlR13 *LaaParametersR13CrosscarrierschedulinglaaDlR13
	CsiRsDrsRrmMeasurementslaaR13  *LaaParametersR13CsiRsDrsRrmMeasurementslaaR13
	DownlinklaaR13                 *LaaParametersR13DownlinklaaR13
	EndingdwptsR13                 *LaaParametersR13EndingdwptsR13
	SecondslotstartingpositionR13  *LaaParametersR13SecondslotstartingpositionR13
	Tm9LaaR13                      *LaaParametersR13Tm9LaaR13
	Tm10LaaR13                     *LaaParametersR13Tm10LaaR13
}
