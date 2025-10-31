package ies

// BandSidelink-r16 ::= SEQUENCE
// Extensible
type BandsidelinkR16 struct {
	FreqbandsidelinkR16                           Freqbandindicatornr
	SlReceptionR16                                *BandsidelinkR16SlReceptionR16
	SlTransmissionmode1R16                        *BandsidelinkR16SlTransmissionmode1R16
	SyncSidelinkR16                               *BandsidelinkR16SyncSidelinkR16
	SlTx256qamR16                                 *BandsidelinkR16SlTx256qamR16
	PsfchFormatzerosidelinkR16                    *BandsidelinkR16PsfchFormatzerosidelinkR16
	Lowse64qamMcsTablesidelinkR16                 *BandsidelinkR16Lowse64qamMcsTablesidelinkR16
	EnbSyncSidelinkR16                            *BandsidelinkR16EnbSyncSidelinkR16
	SlTransmissionmode2R16                        *BandsidelinkR16SlTransmissionmode2R16                        `ext`
	CongestioncontrolsidelinkR16                  *BandsidelinkR16CongestioncontrolsidelinkR16                  `ext`
	FewersymbolslotsidelinkR16                    *BandsidelinkR16FewersymbolslotsidelinkR16                    `ext`
	SlOpenlooppcRsrpReportsidelinkR16             *BandsidelinkR16SlOpenlooppcRsrpReportsidelinkR16             `ext`
	SlRx256qamR16                                 *BandsidelinkR16SlRx256qamR16                                 `ext`
	UePowerclasssidelinkR16                       *BandsidelinkR16UePowerclasssidelinkR16                       `ext`
	SlTransmissionmode2RandomresourceselectionR17 *BandsidelinkR16SlTransmissionmode2RandomresourceselectionR17 `ext`
	SyncSidelinkV1710                             *BandsidelinkR16SyncSidelinkV1710                             `ext`
	EnbSyncSidelinkV1710                          *BandsidelinkR16EnbSyncSidelinkV1710                          `ext`
	RxIucScheme1Preferredmode2sidelinkR17         *BandsidelinkR16RxIucScheme1Preferredmode2sidelinkR17         `ext`
	RxIucScheme1Nonpreferredmode2sidelinkR17      *BandsidelinkR16RxIucScheme1Nonpreferredmode2sidelinkR17      `ext`
	RxIucScheme2Mode2sidelinkR17                  *BandsidelinkR16RxIucScheme2Mode2sidelinkR17                  `ext`
	RxIucScheme1SciR17                            *BandsidelinkR16RxIucScheme1SciR17                            `ext`
	RxIucScheme1SciExplicitreqR17                 *BandsidelinkR16RxIucScheme1SciExplicitreqR17                 `ext`
}
