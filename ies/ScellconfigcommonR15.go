package ies

// SCellConfigCommon-r15 ::= SEQUENCE
type ScellconfigcommonR15 struct {
	RadioresourceconfigcommonscellR15    *RadioresourceconfigcommonscellR10
	RadioresourceconfigdedicatedscellR15 *RadioresourceconfigdedicatedscellR10
	AntennainfodedicatedscellR15         *AntennainfodedicatedV10i0
}
