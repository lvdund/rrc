package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CarrierFreqEUTRA_v1700 struct {
	Eutra_FreqNeighHSDN_CellList_r17 *EUTRA_FreqNeighHSDN_CellList_r17 `optional`
}

func (ie *CarrierFreqEUTRA_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Eutra_FreqNeighHSDN_CellList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Eutra_FreqNeighHSDN_CellList_r17 != nil {
		if err = ie.Eutra_FreqNeighHSDN_CellList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Eutra_FreqNeighHSDN_CellList_r17", err)
		}
	}
	return nil
}

func (ie *CarrierFreqEUTRA_v1700) Decode(r *aper.AperReader) error {
	var err error
	var Eutra_FreqNeighHSDN_CellList_r17Present bool
	if Eutra_FreqNeighHSDN_CellList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Eutra_FreqNeighHSDN_CellList_r17Present {
		ie.Eutra_FreqNeighHSDN_CellList_r17 = new(EUTRA_FreqNeighHSDN_CellList_r17)
		if err = ie.Eutra_FreqNeighHSDN_CellList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra_FreqNeighHSDN_CellList_r17", err)
		}
	}
	return nil
}
