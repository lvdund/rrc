package ies

// RadioResourceConfigCommonSCell-r10-nonUL-Configuration-r10 ::= SEQUENCE
type RadioresourceconfigcommonscellR10NonulConfigurationR10 struct {
	DlBandwidthR10             RadioresourceconfigcommonscellR10NonulConfigurationR10DlBandwidthR10
	AntennainfocommonR10       Antennainfocommon
	MbsfnSubframeconfiglistR10 *MbsfnSubframeconfiglist
	PhichConfigR10             PhichConfig
	PdschConfigcommonR10       PdschConfigcommon
	TddConfigR10               *TddConfig
}
