package ies

// LAA-Parameters-v1530 ::= SEQUENCE
type LaaParametersV1530 struct {
	AulR15           *LaaParametersV1530AulR15
	LaaPuschMode1R15 *LaaParametersV1530LaaPuschMode1R15
	LaaPuschMode2R15 *LaaParametersV1530LaaPuschMode2R15
	LaaPuschMode3R15 *LaaParametersV1530LaaPuschMode3R15
}
