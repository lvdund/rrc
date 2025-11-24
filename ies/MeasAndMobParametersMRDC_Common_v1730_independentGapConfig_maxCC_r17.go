package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_Common_v1730_independentGapConfig_maxCC_r17 struct {
	Fr1_Only_r17   *int64 `lb:1,ub:32,optional`
	Fr2_Only_r17   *int64 `lb:1,ub:32,optional`
	Fr1_AndFR2_r17 *int64 `lb:1,ub:32,optional`
}

func (ie *MeasAndMobParametersMRDC_Common_v1730_independentGapConfig_maxCC_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Fr1_Only_r17 != nil, ie.Fr2_Only_r17 != nil, ie.Fr1_AndFR2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Fr1_Only_r17 != nil {
		if err = w.WriteInteger(*ie.Fr1_Only_r17, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode Fr1_Only_r17", err)
		}
	}
	if ie.Fr2_Only_r17 != nil {
		if err = w.WriteInteger(*ie.Fr2_Only_r17, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode Fr2_Only_r17", err)
		}
	}
	if ie.Fr1_AndFR2_r17 != nil {
		if err = w.WriteInteger(*ie.Fr1_AndFR2_r17, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode Fr1_AndFR2_r17", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1730_independentGapConfig_maxCC_r17) Decode(r *uper.UperReader) error {
	var err error
	var Fr1_Only_r17Present bool
	if Fr1_Only_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr2_Only_r17Present bool
	if Fr2_Only_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fr1_AndFR2_r17Present bool
	if Fr1_AndFR2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Fr1_Only_r17Present {
		var tmp_int_Fr1_Only_r17 int64
		if tmp_int_Fr1_Only_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode Fr1_Only_r17", err)
		}
		ie.Fr1_Only_r17 = &tmp_int_Fr1_Only_r17
	}
	if Fr2_Only_r17Present {
		var tmp_int_Fr2_Only_r17 int64
		if tmp_int_Fr2_Only_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode Fr2_Only_r17", err)
		}
		ie.Fr2_Only_r17 = &tmp_int_Fr2_Only_r17
	}
	if Fr1_AndFR2_r17Present {
		var tmp_int_Fr1_AndFR2_r17 int64
		if tmp_int_Fr1_AndFR2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode Fr1_AndFR2_r17", err)
		}
		ie.Fr1_AndFR2_r17 = &tmp_int_Fr1_AndFR2_r17
	}
	return nil
}
