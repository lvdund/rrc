package ies

// Phy-ParametersMRDC ::= SEQUENCE
// Extensible
type PhyParametersmrdc struct {
	NaicsCapabilityList          *[]NaicsCapabilityEntry                        `lb:1,ub:maxNrofNAICSEntries`
	Spcellplacement              *Carrieraggregationvariant                     `ext`
	TddPcellulTxAllulSubframeR16 *PhyParametersmrdcTddPcellulTxAllulSubframeR16 `ext`
	FddPcellulTxAllulSubframeR16 *PhyParametersmrdcFddPcellulTxAllulSubframeR16 `ext`
}
