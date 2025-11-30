package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_QoS_Profile_r16 struct {
	Sl_PQI_r16   *SL_PQI_r16 `optional`
	Sl_GFBR_r16  *int64      `lb:0,ub:4000000000,optional`
	Sl_MFBR_r16  *int64      `lb:0,ub:4000000000,optional`
	Sl_Range_r16 *int64      `lb:1,ub:1000,optional`
}

func (ie *SL_QoS_Profile_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_PQI_r16 != nil, ie.Sl_GFBR_r16 != nil, ie.Sl_MFBR_r16 != nil, ie.Sl_Range_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_PQI_r16 != nil {
		if err = ie.Sl_PQI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PQI_r16", err)
		}
	}
	if ie.Sl_GFBR_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_GFBR_r16, &aper.Constraint{Lb: 0, Ub: 4000000000}, false); err != nil {
			return utils.WrapError("Encode Sl_GFBR_r16", err)
		}
	}
	if ie.Sl_MFBR_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_MFBR_r16, &aper.Constraint{Lb: 0, Ub: 4000000000}, false); err != nil {
			return utils.WrapError("Encode Sl_MFBR_r16", err)
		}
	}
	if ie.Sl_Range_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_Range_r16, &aper.Constraint{Lb: 1, Ub: 1000}, false); err != nil {
			return utils.WrapError("Encode Sl_Range_r16", err)
		}
	}
	return nil
}

func (ie *SL_QoS_Profile_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_PQI_r16Present bool
	if Sl_PQI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_GFBR_r16Present bool
	if Sl_GFBR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_MFBR_r16Present bool
	if Sl_MFBR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_Range_r16Present bool
	if Sl_Range_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_PQI_r16Present {
		ie.Sl_PQI_r16 = new(SL_PQI_r16)
		if err = ie.Sl_PQI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PQI_r16", err)
		}
	}
	if Sl_GFBR_r16Present {
		var tmp_int_Sl_GFBR_r16 int64
		if tmp_int_Sl_GFBR_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000}, false); err != nil {
			return utils.WrapError("Decode Sl_GFBR_r16", err)
		}
		ie.Sl_GFBR_r16 = &tmp_int_Sl_GFBR_r16
	}
	if Sl_MFBR_r16Present {
		var tmp_int_Sl_MFBR_r16 int64
		if tmp_int_Sl_MFBR_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000}, false); err != nil {
			return utils.WrapError("Decode Sl_MFBR_r16", err)
		}
		ie.Sl_MFBR_r16 = &tmp_int_Sl_MFBR_r16
	}
	if Sl_Range_r16Present {
		var tmp_int_Sl_Range_r16 int64
		if tmp_int_Sl_Range_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 1000}, false); err != nil {
			return utils.WrapError("Decode Sl_Range_r16", err)
		}
		ie.Sl_Range_r16 = &tmp_int_Sl_Range_r16
	}
	return nil
}
