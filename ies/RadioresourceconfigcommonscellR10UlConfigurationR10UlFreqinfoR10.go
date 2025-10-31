package ies

// RadioResourceConfigCommonSCell-r10-ul-Configuration-r10-ul-FreqInfo-r10 ::= SEQUENCE
type RadioresourceconfigcommonscellR10UlConfigurationR10UlFreqinfoR10 struct {
	UlCarrierfreqR10                   *ArfcnValueeutra
	UlBandwidthR10                     *RadioresourceconfigcommonscellR10UlConfigurationR10UlFreqinfoR10UlBandwidthR10
	AdditionalspectrumemissionscellR10 Additionalspectrumemission
}
