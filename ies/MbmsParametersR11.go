package ies

// MBMS-Parameters-r11 ::= SEQUENCE
type MbmsParametersR11 struct {
	MbmsScellR11          *MbmsParametersR11MbmsScellR11
	MbmsNonservingcellR11 *MbmsParametersR11MbmsNonservingcellR11
}
