package ies

// SystemInformationBlockType3-intraFreqCellReselectionInfo ::= SEQUENCE
type Systeminformationblocktype3Intrafreqcellreselectioninfo struct {
	QRxlevmin            QRxlevmin
	PMax                 *PMax
	SIntrasearch         *Reselectionthreshold
	Allowedmeasbandwidth *Allowedmeasbandwidth
	Presenceantennaport1 Presenceantennaport1
	Neighcellconfig      Neighcellconfig
	TReselectioneutra    TReselection
	TReselectioneutraSf  *Speedstatescalefactors
}
