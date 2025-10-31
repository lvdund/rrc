package ies

// Tx-PoolMeasList-r16 ::= SEQUENCE OF SL-ResourcePoolID-r16
// SIZE (1..maxNrofSL-PoolToMeasureNR-r16)
type TxPoolmeaslistR16 struct {
	Value []SlResourcepoolidR16 `lb:1,ub:maxNrofSLPooltomeasurenrR16`
}
