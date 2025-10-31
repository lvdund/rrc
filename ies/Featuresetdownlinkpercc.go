package ies

// FeatureSetDownlinkPerCC ::= SEQUENCE
type Featuresetdownlinkpercc struct {
	Supportedsubcarrierspacingdl Subcarrierspacing
	Supportedbandwidthdl         Supportedbandwidth
	Channelbw90mhz               *FeaturesetdownlinkperccChannelbw90mhz
	MaxnumbermimoLayerspdsch     *MimoLayersdl
	Supportedmodulationorderdl   *Modulationorder
}
