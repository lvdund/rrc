package ies

import "rrc/utils"

// TRS-ResourceSet-r17 ::= SEQUENCE
// Extensible
type TrsResourcesetR17 struct {
	PowercontroloffsetssR17        TrsResourcesetR17PowercontroloffsetssR17
	ScramblingidInfoR17            TrsResourcesetR17ScramblingidInfoR17
	FirstofdmsymbolintimedomainR17 utils.INTEGER `lb:0,ub:9`
	StartingrbR17                  utils.INTEGER `lb:0,ub:maxNrofPhysicalResourceBlocks1`
	NrofrbsR17                     utils.INTEGER `lb:0,ub:maxNrofPhysicalResourceBlocksPlus1`
	SsbIndexR17                    SsbIndex
	PeriodicityandoffsetR17        TrsResourcesetR17PeriodicityandoffsetR17
	FrequencydomainallocationR17   utils.BITSTRING `lb:4,ub:4`
	IndbitidR17                    utils.INTEGER   `lb:0,ub:5`
	NrofresourcesR17               TrsResourcesetR17NrofresourcesR17
}
