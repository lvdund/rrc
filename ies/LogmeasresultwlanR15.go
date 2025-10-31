package ies

// LogMeasResultWLAN-r15 ::= SEQUENCE
// Extensible
type LogmeasresultwlanR15 struct {
	WlanIdentifiersR15 WlanIdentifiersR12
	RssiwlanR15        *WlanRssiRangeR13
	RttWlanR15         *WlanRttR15
}
