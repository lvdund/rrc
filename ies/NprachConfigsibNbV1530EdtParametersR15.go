package ies

// NPRACH-ConfigSIB-NB-v1530-edt-Parameters-r15 ::= SEQUENCE
type NprachConfigsibNbV1530EdtParametersR15 struct {
	EdtSmalltbsSubsetR15       *bool
	EdtTbsInfolistR15          EdtTbsInfolistNbR15
	NprachParameterslistedtR15 *NprachParameterslistNbR14
}
