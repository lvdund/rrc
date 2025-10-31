package ies

// LogMeasResultWLAN-r16 ::= SEQUENCE
// Extensible
type LogmeasresultwlanR16 struct {
	WlanIdentifiersR16 WlanIdentifiersR16
	RssiwlanR16        *WlanRssiRangeR16
	RttWlanR16         *WlanRttR16
}
