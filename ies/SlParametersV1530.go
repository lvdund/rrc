package ies

// SL-Parameters-v1530 ::= SEQUENCE
type SlParametersV1530 struct {
	SlssSupportedtxfreqR15               *SlParametersV1530SlssSupportedtxfreqR15
	Sl64qamTxR15                         *SlParametersV1530Sl64qamTxR15
	SlTxdiversityR15                     *SlParametersV1530SlTxdiversityR15
	UeCategoryslR15                      *UeCategoryslR15
	V2xSupportedbandcombinationlistV1530 *V2xSupportedbandcombinationV1530
}
