package ies

// DownlinkConfigCommon ::= SEQUENCE
// Extensible
type Downlinkconfigcommon struct {
	Frequencyinfodl             *Frequencyinfodl
	Initialdownlinkbwp          *BwpDownlinkcommon
	InitialdownlinkbwpRedcapR17 *BwpDownlinkcommon `ext`
}
