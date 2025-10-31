package ies

// CarrierBandwidthEUTRA ::= SEQUENCE
type Carrierbandwidtheutra struct {
	DlBandwidth CarrierbandwidtheutraDlBandwidth
	UlBandwidth *CarrierbandwidtheutraUlBandwidth
}
