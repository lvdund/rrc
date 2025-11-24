package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectUTRA_FDD_r16 struct {
	CarrierFreq_r16            ARFCN_ValueUTRA_FDD_r16        `madatory`
	Utra_FDD_Q_OffsetRange_r16 *UTRA_FDD_Q_OffsetRange_r16    `optional`
	CellsToRemoveList_r16      *UTRA_FDD_CellIndexList_r16    `optional`
	CellsToAddModList_r16      *CellsToAddModListUTRA_FDD_r16 `optional`
}

func (ie *MeasObjectUTRA_FDD_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Utra_FDD_Q_OffsetRange_r16 != nil, ie.CellsToRemoveList_r16 != nil, ie.CellsToAddModList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CarrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq_r16", err)
	}
	if ie.Utra_FDD_Q_OffsetRange_r16 != nil {
		if err = ie.Utra_FDD_Q_OffsetRange_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Utra_FDD_Q_OffsetRange_r16", err)
		}
	}
	if ie.CellsToRemoveList_r16 != nil {
		if err = ie.CellsToRemoveList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CellsToRemoveList_r16", err)
		}
	}
	if ie.CellsToAddModList_r16 != nil {
		if err = ie.CellsToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CellsToAddModList_r16", err)
		}
	}
	return nil
}

func (ie *MeasObjectUTRA_FDD_r16) Decode(r *uper.UperReader) error {
	var err error
	var Utra_FDD_Q_OffsetRange_r16Present bool
	if Utra_FDD_Q_OffsetRange_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CellsToRemoveList_r16Present bool
	if CellsToRemoveList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CellsToAddModList_r16Present bool
	if CellsToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CarrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq_r16", err)
	}
	if Utra_FDD_Q_OffsetRange_r16Present {
		ie.Utra_FDD_Q_OffsetRange_r16 = new(UTRA_FDD_Q_OffsetRange_r16)
		if err = ie.Utra_FDD_Q_OffsetRange_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Utra_FDD_Q_OffsetRange_r16", err)
		}
	}
	if CellsToRemoveList_r16Present {
		ie.CellsToRemoveList_r16 = new(UTRA_FDD_CellIndexList_r16)
		if err = ie.CellsToRemoveList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CellsToRemoveList_r16", err)
		}
	}
	if CellsToAddModList_r16Present {
		ie.CellsToAddModList_r16 = new(CellsToAddModListUTRA_FDD_r16)
		if err = ie.CellsToAddModList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CellsToAddModList_r16", err)
		}
	}
	return nil
}
