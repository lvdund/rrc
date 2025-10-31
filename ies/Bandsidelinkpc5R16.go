package ies

// BandSidelinkPC5-r16 ::= SEQUENCE
// Extensible
type Bandsidelinkpc5R16 struct {
	FreqbandsidelinkR16                      Freqbandindicatornr
	SlReceptionR16                           *Bandsidelinkpc5R16SlReceptionR16
	SlTx256qamR16                            *Bandsidelinkpc5R16SlTx256qamR16
	Lowse64qamMcsTablesidelinkR16            *Bandsidelinkpc5R16Lowse64qamMcsTablesidelinkR16
	CsiReportsidelinkR16                     *Bandsidelinkpc5R16CsiReportsidelinkR16                     `ext`
	RanktworeceptionR16                      *Bandsidelinkpc5R16RanktworeceptionR16                      `ext`
	SlOpenlooppcRsrpReportsidelinkR16        *Bandsidelinkpc5R16SlOpenlooppcRsrpReportsidelinkR16        `ext`
	SlRx256qamR16                            *Bandsidelinkpc5R16SlRx256qamR16                            `ext`
	RxIucScheme1Preferredmode2sidelinkR17    *Bandsidelinkpc5R16RxIucScheme1Preferredmode2sidelinkR17    `ext`
	RxIucScheme1Nonpreferredmode2sidelinkR17 *Bandsidelinkpc5R16RxIucScheme1Nonpreferredmode2sidelinkR17 `ext`
	RxIucScheme2Mode2sidelinkR17             *Bandsidelinkpc5R16RxIucScheme2Mode2sidelinkR17             `ext`
	RxIucScheme1SciR17                       *Bandsidelinkpc5R16RxIucScheme1SciR17                       `ext`
	RxIucScheme1SciExplicitreqR17            *Bandsidelinkpc5R16RxIucScheme1SciExplicitreqR17            `ext`
	Scheme2ConflictdeterminationrsrpR17      *Bandsidelinkpc5R16Scheme2ConflictdeterminationrsrpR17      `ext`
}
