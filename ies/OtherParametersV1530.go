package ies

import "rrc/utils"

// Other-Parameters-v1530 ::= SEQUENCE
type OtherParametersV1530 struct {
	AssistinfobitforlcR15     *utils.ENUMERATED
	TimereferenceprovisionR15 *utils.ENUMERATED
	FlightpathplanR15         *utils.ENUMERATED
}
