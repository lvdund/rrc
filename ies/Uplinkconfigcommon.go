package ies

// UplinkConfigCommon ::= SEQUENCE
type Uplinkconfigcommon struct {
	Frequencyinfoul  *Frequencyinfoul
	Initialuplinkbwp *BwpUplinkcommon
	Dummy            Timealignmenttimer
}
