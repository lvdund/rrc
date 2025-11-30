package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UL_GapFR2_Config_r17 struct {
	GapOffset_r17              int64                         `lb:0,ub:159,madatory`
	Ugl_r17                    UL_GapFR2_Config_r17_ugl_r17  `madatory`
	Ugrp_r17                   UL_GapFR2_Config_r17_ugrp_r17 `madatory`
	RefFR2_ServCellAsyncCA_r17 *ServCellIndex                `optional`
}

func (ie *UL_GapFR2_Config_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.RefFR2_ServCellAsyncCA_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.GapOffset_r17, &aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("WriteInteger GapOffset_r17", err)
	}
	if err = ie.Ugl_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Ugl_r17", err)
	}
	if err = ie.Ugrp_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Ugrp_r17", err)
	}
	if ie.RefFR2_ServCellAsyncCA_r17 != nil {
		if err = ie.RefFR2_ServCellAsyncCA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RefFR2_ServCellAsyncCA_r17", err)
		}
	}
	return nil
}

func (ie *UL_GapFR2_Config_r17) Decode(r *aper.AperReader) error {
	var err error
	var RefFR2_ServCellAsyncCA_r17Present bool
	if RefFR2_ServCellAsyncCA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_GapOffset_r17 int64
	if tmp_int_GapOffset_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
		return utils.WrapError("ReadInteger GapOffset_r17", err)
	}
	ie.GapOffset_r17 = tmp_int_GapOffset_r17
	if err = ie.Ugl_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Ugl_r17", err)
	}
	if err = ie.Ugrp_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Ugrp_r17", err)
	}
	if RefFR2_ServCellAsyncCA_r17Present {
		ie.RefFR2_ServCellAsyncCA_r17 = new(ServCellIndex)
		if err = ie.RefFR2_ServCellAsyncCA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RefFR2_ServCellAsyncCA_r17", err)
		}
	}
	return nil
}
