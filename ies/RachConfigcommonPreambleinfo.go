package ies

// RACH-ConfigCommon-preambleInfo ::= SEQUENCE
// Extensible
type RachConfigcommonPreambleinfo struct {
	NumberofraPreambles   RachConfigcommonPreambleinfoNumberofraPreambles
	Preamblesgroupaconfig *RachConfigcommonPreambleinfoPreamblesgroupaconfig
}
