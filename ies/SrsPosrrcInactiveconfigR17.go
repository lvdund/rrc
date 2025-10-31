package ies

// SRS-PosRRC-InactiveConfig-r17 ::= SEQUENCE
type SrsPosrrcInactiveconfigR17 struct {
	SrsPosconfignulR17                   *SrsPosconfigR17
	SrsPosconfigsulR17                   *SrsPosconfigR17
	BwpNulR17                            *Bwp
	BwpSulR17                            *Bwp
	InactivepossrsTimealignmenttimerR17  *Timealignmenttimer
	InactivepossrsRsrpChangethresholdR17 *RsrpChangethresholdR17
}
