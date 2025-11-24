package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TA_Info_r17 struct {
	Ta_Common_r17             int64  `lb:0,ub:66485757,madatory`
	Ta_CommonDrift_r17        *int64 `lb:-257303,ub:257303,optional`
	Ta_CommonDriftVariant_r17 *int64 `lb:0,ub:28949,optional`
}

func (ie *TA_Info_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ta_CommonDrift_r17 != nil, ie.Ta_CommonDriftVariant_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.Ta_Common_r17, &uper.Constraint{Lb: 0, Ub: 66485757}, false); err != nil {
		return utils.WrapError("WriteInteger Ta_Common_r17", err)
	}
	if ie.Ta_CommonDrift_r17 != nil {
		if err = w.WriteInteger(*ie.Ta_CommonDrift_r17, &uper.Constraint{Lb: -257303, Ub: 257303}, false); err != nil {
			return utils.WrapError("Encode Ta_CommonDrift_r17", err)
		}
	}
	if ie.Ta_CommonDriftVariant_r17 != nil {
		if err = w.WriteInteger(*ie.Ta_CommonDriftVariant_r17, &uper.Constraint{Lb: 0, Ub: 28949}, false); err != nil {
			return utils.WrapError("Encode Ta_CommonDriftVariant_r17", err)
		}
	}
	return nil
}

func (ie *TA_Info_r17) Decode(r *uper.UperReader) error {
	var err error
	var Ta_CommonDrift_r17Present bool
	if Ta_CommonDrift_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ta_CommonDriftVariant_r17Present bool
	if Ta_CommonDriftVariant_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_Ta_Common_r17 int64
	if tmp_int_Ta_Common_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 66485757}, false); err != nil {
		return utils.WrapError("ReadInteger Ta_Common_r17", err)
	}
	ie.Ta_Common_r17 = tmp_int_Ta_Common_r17
	if Ta_CommonDrift_r17Present {
		var tmp_int_Ta_CommonDrift_r17 int64
		if tmp_int_Ta_CommonDrift_r17, err = r.ReadInteger(&uper.Constraint{Lb: -257303, Ub: 257303}, false); err != nil {
			return utils.WrapError("Decode Ta_CommonDrift_r17", err)
		}
		ie.Ta_CommonDrift_r17 = &tmp_int_Ta_CommonDrift_r17
	}
	if Ta_CommonDriftVariant_r17Present {
		var tmp_int_Ta_CommonDriftVariant_r17 int64
		if tmp_int_Ta_CommonDriftVariant_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 28949}, false); err != nil {
			return utils.WrapError("Decode Ta_CommonDriftVariant_r17", err)
		}
		ie.Ta_CommonDriftVariant_r17 = &tmp_int_Ta_CommonDriftVariant_r17
	}
	return nil
}
