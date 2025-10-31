package ies

// MBMS-Parameters-v1430 ::= SEQUENCE
type MbmsParametersV1430 struct {
	FembmsdedicatedcellR14            *MbmsParametersV1430FembmsdedicatedcellR14
	FembmsmixedcellR14                *MbmsParametersV1430FembmsmixedcellR14
	SubcarrierspacingmbmsKhz7dot5R14  *MbmsParametersV1430SubcarrierspacingmbmsKhz7dot5R14
	SubcarrierspacingmbmsKhz1dot25R14 *MbmsParametersV1430SubcarrierspacingmbmsKhz1dot25R14
}
