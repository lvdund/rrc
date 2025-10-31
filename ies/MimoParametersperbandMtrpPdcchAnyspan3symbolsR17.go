package ies

import "rrc/utils"

// MIMO-ParametersPerBand-mTRP-PDCCH-anySpan-3Symbols-r17 ::= ENUMERATED
type MimoParametersperbandMtrpPdcchAnyspan3symbolsR17 struct {
	Value utils.ENUMERATED
}

const (
	MimoParametersperbandMtrpPdcchAnyspan3symbolsR17EnumeratedNothing = iota
	MimoParametersperbandMtrpPdcchAnyspan3symbolsR17EnumeratedSupported
)
