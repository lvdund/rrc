package ies

// FeatureSetUplinkPerCC ::= SEQUENCE
type Featuresetuplinkpercc struct {
	Supportedsubcarrierspacingul  Subcarrierspacing
	Supportedbandwidthul          Supportedbandwidth
	Channelbw90mhz                *FeaturesetuplinkperccChannelbw90mhz
	MimoCbPusch                   *FeaturesetuplinkperccMimoCbPusch
	MaxnumbermimoLayersnoncbPusch *MimoLayersul
	Supportedmodulationorderul    *Modulationorder
}
