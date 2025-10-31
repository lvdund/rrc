package ies

// SL-DiscSysInfoReport-r13-freqInfo-r13 ::= SEQUENCE
type SlDiscsysinforeportR13FreqinfoR13 struct {
	UlCarrierfreqR13              *ArfcnValueeutra
	UlBandwidthR13                *SlDiscsysinforeportR13FreqinfoR13UlBandwidthR13
	AdditionalspectrumemissionR13 *Additionalspectrumemission
}
