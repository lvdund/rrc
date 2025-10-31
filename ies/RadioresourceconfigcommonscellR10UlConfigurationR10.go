package ies

// RadioResourceConfigCommonSCell-r10-ul-Configuration-r10 ::= SEQUENCE
type RadioresourceconfigcommonscellR10UlConfigurationR10 struct {
	UlFreqinfoR10                    *RadioresourceconfigcommonscellR10UlConfigurationR10UlFreqinfoR10
	PMaxR10                          *PMax
	UplinkpowercontrolcommonscellR10 UplinkpowercontrolcommonscellR10
	SoundingrsUlConfigcommonR10      SoundingrsUlConfigcommon
	UlCyclicprefixlengthR10          UlCyclicprefixlength
	PrachConfigscellR10              *PrachConfigscellR10
	PuschConfigcommonR10             PuschConfigcommon
}
