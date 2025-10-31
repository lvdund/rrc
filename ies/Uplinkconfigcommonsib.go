package ies

// UplinkConfigCommonSIB ::= SEQUENCE
type Uplinkconfigcommonsib struct {
	Frequencyinfoul          FrequencyinfoulSib
	Initialuplinkbwp         BwpUplinkcommon
	Timealignmenttimercommon Timealignmenttimer
}
