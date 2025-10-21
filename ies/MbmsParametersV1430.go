package ies

import "rrc/utils"

// MBMS-Parameters-v1430 ::= SEQUENCE
type MbmsParametersV1430 struct {
	FembmsdedicatedcellR14            *utils.ENUMERATED
	FembmsmixedcellR14                *utils.ENUMERATED
	SubcarrierspacingmbmsKhz7dot5R14  *utils.ENUMERATED
	SubcarrierspacingmbmsKhz1dot25R14 *utils.ENUMERATED
}
