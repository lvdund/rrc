package ies

// SharedSpectrumChAccessParamsPerBand-v1640 ::= SEQUENCE
type SharedspectrumchaccessparamsperbandV1640 struct {
	CsiRsrpAndrsrqMeaswithssbR16    *SharedspectrumchaccessparamsperbandV1640CsiRsrpAndrsrqMeaswithssbR16
	CsiRsrpAndrsrqMeaswithoutssbR16 *SharedspectrumchaccessparamsperbandV1640CsiRsrpAndrsrqMeaswithoutssbR16
	CsiSinrMeasR16                  *SharedspectrumchaccessparamsperbandV1640CsiSinrMeasR16
	SsbAndcsiRsRlmR16               *SharedspectrumchaccessparamsperbandV1640SsbAndcsiRsRlmR16
	CsiRsCfraForhoR16               *SharedspectrumchaccessparamsperbandV1640CsiRsCfraForhoR16
}
