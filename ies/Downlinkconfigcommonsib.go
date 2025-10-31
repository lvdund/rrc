package ies

// DownlinkConfigCommonSIB ::= SEQUENCE
// Extensible
type Downlinkconfigcommonsib struct {
	Frequencyinfodl             FrequencyinfodlSib
	Initialdownlinkbwp          BwpDownlinkcommon
	BcchConfig                  BcchConfig
	PcchConfig                  PcchConfig
	PeiConfigR17                *PeiConfigR17      `ext`
	InitialdownlinkbwpRedcapR17 *BwpDownlinkcommon `ext`
}
