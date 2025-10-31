package ies

// SecurityConfigHO-v1530-handoverType-v1530 ::= CHOICE
const (
	SecurityconfighoV1530HandovertypeV1530ChoiceNothing = iota
	SecurityconfighoV1530HandovertypeV1530ChoiceIntra5gc
	SecurityconfighoV1530HandovertypeV1530ChoiceFivegcToepc
	SecurityconfighoV1530HandovertypeV1530ChoiceEpcTo5gc
)

type SecurityconfighoV1530HandovertypeV1530 struct {
	Choice      uint64
	Intra5gc    *SecurityconfighoV1530HandovertypeV1530Intra5gc
	FivegcToepc *SecurityconfighoV1530HandovertypeV1530FivegcToepc
	EpcTo5gc    *SecurityconfighoV1530HandovertypeV1530EpcTo5gc
}
