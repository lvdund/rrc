package ies

import "rrc/utils"

// BandParametersSidelinkDiscovery-r17-tx-IUC-Scheme1-Mode2Sidelink-r17 ::= ENUMERATED
type BandparameterssidelinkdiscoveryR17TxIucScheme1Mode2sidelinkR17 struct {
	Value utils.ENUMERATED
}

const (
	BandparameterssidelinkdiscoveryR17TxIucScheme1Mode2sidelinkR17EnumeratedNothing = iota
	BandparameterssidelinkdiscoveryR17TxIucScheme1Mode2sidelinkR17EnumeratedSupported
)
