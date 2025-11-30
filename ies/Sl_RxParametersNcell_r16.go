package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Sl_RxParametersNcell_r16 struct {
	Sl_TDD_Configuration_r16 *TDD_UL_DL_ConfigCommon `optional`
	Sl_SyncConfigIndex_r16   int64                   `lb:0,ub:15,madatory`
}

func (ie *Sl_RxParametersNcell_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_TDD_Configuration_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_TDD_Configuration_r16 != nil {
		if err = ie.Sl_TDD_Configuration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TDD_Configuration_r16", err)
		}
	}
	if err = w.WriteInteger(ie.Sl_SyncConfigIndex_r16, &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_SyncConfigIndex_r16", err)
	}
	return nil
}

func (ie *Sl_RxParametersNcell_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_TDD_Configuration_r16Present bool
	if Sl_TDD_Configuration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_TDD_Configuration_r16Present {
		ie.Sl_TDD_Configuration_r16 = new(TDD_UL_DL_ConfigCommon)
		if err = ie.Sl_TDD_Configuration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TDD_Configuration_r16", err)
		}
	}
	var tmp_int_Sl_SyncConfigIndex_r16 int64
	if tmp_int_Sl_SyncConfigIndex_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_SyncConfigIndex_r16", err)
	}
	ie.Sl_SyncConfigIndex_r16 = tmp_int_Sl_SyncConfigIndex_r16
	return nil
}
