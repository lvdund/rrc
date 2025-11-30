package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResult2UTRA_FDD_r16 struct {
	CarrierFreq_r16             ARFCN_ValueUTRA_FDD_r16    `madatory`
	MeasResultNeighCellList_r16 MeasResultListUTRA_FDD_r16 `madatory`
}

func (ie *MeasResult2UTRA_FDD_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CarrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq_r16", err)
	}
	if err = ie.MeasResultNeighCellList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultNeighCellList_r16", err)
	}
	return nil
}

func (ie *MeasResult2UTRA_FDD_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CarrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq_r16", err)
	}
	if err = ie.MeasResultNeighCellList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MeasResultNeighCellList_r16", err)
	}
	return nil
}
