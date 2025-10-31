package ies

import "rrc/utils"

// MeasResults-measResultNeighCells ::= CHOICE
// Extensible
const (
	MeasresultsMeasresultneighcellsChoiceNothing = iota
	MeasresultsMeasresultneighcellsChoiceMeasresultlistnr
	MeasresultsMeasresultneighcellsChoiceMeasresultlisteutra
	MeasresultsMeasresultneighcellsChoiceMeasresultlistutraFddR16
	MeasresultsMeasresultneighcellsChoiceSlMeasresultscandrelayR17
)

type MeasresultsMeasresultneighcells struct {
	Choice                    uint64
	Measresultlistnr          *Measresultlistnr
	Measresultlisteutra       *Measresultlisteutra
	MeasresultlistutraFddR16  *MeasresultlistutraFddR16
	SlMeasresultscandrelayR17 *utils.OCTETSTRING
}
