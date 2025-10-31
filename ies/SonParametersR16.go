package ies

// SON-Parameters-r16 ::= SEQUENCE
// Extensible
type SonParametersR16 struct {
	RachReportR16        *SonParametersR16RachReportR16
	RlfreportchoR17      *SonParametersR16RlfreportchoR17      `ext`
	RlfreportdapsR17     *SonParametersR16RlfreportdapsR17     `ext`
	SuccessHoReportR17   *SonParametersR16SuccessHoReportR17   `ext`
	TwosteprachReportR17 *SonParametersR16TwosteprachReportR17 `ext`
	PscellMhiReportR17   *SonParametersR16PscellMhiReportR17   `ext`
	OndemandsiReportR17  *SonParametersR16OndemandsiReportR17  `ext`
}
