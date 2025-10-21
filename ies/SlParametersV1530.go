package ies

import "rrc/utils"

// SL-Parameters-v1530 ::= SEQUENCE
type SlParametersV1530 struct {
	SlssSupportedtxfreqR15               *utils.ENUMERATED
	Sl64qamTxR15                         *utils.ENUMERATED
	SlTxdiversityR15                     *utils.ENUMERATED
	UeCategoryslR15                      *UeCategoryslR15
	V2xSupportedbandcombinationlistV1530 *V2xSupportedbandcombinationV1530
}
