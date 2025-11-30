package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type IAB_IP_AddressNumReq_r16 struct {
	All_Traffic_NumReq_r16    *int64 `lb:1,ub:8,optional`
	F1_C_Traffic_NumReq_r16   *int64 `lb:1,ub:8,optional`
	F1_U_Traffic_NumReq_r16   *int64 `lb:1,ub:8,optional`
	Non_F1_Traffic_NumReq_r16 *int64 `lb:1,ub:8,optional`
}

func (ie *IAB_IP_AddressNumReq_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.All_Traffic_NumReq_r16 != nil, ie.F1_C_Traffic_NumReq_r16 != nil, ie.F1_U_Traffic_NumReq_r16 != nil, ie.Non_F1_Traffic_NumReq_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.All_Traffic_NumReq_r16 != nil {
		if err = w.WriteInteger(*ie.All_Traffic_NumReq_r16, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode All_Traffic_NumReq_r16", err)
		}
	}
	if ie.F1_C_Traffic_NumReq_r16 != nil {
		if err = w.WriteInteger(*ie.F1_C_Traffic_NumReq_r16, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode F1_C_Traffic_NumReq_r16", err)
		}
	}
	if ie.F1_U_Traffic_NumReq_r16 != nil {
		if err = w.WriteInteger(*ie.F1_U_Traffic_NumReq_r16, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode F1_U_Traffic_NumReq_r16", err)
		}
	}
	if ie.Non_F1_Traffic_NumReq_r16 != nil {
		if err = w.WriteInteger(*ie.Non_F1_Traffic_NumReq_r16, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Non_F1_Traffic_NumReq_r16", err)
		}
	}
	return nil
}

func (ie *IAB_IP_AddressNumReq_r16) Decode(r *aper.AperReader) error {
	var err error
	var All_Traffic_NumReq_r16Present bool
	if All_Traffic_NumReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var F1_C_Traffic_NumReq_r16Present bool
	if F1_C_Traffic_NumReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var F1_U_Traffic_NumReq_r16Present bool
	if F1_U_Traffic_NumReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Non_F1_Traffic_NumReq_r16Present bool
	if Non_F1_Traffic_NumReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if All_Traffic_NumReq_r16Present {
		var tmp_int_All_Traffic_NumReq_r16 int64
		if tmp_int_All_Traffic_NumReq_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode All_Traffic_NumReq_r16", err)
		}
		ie.All_Traffic_NumReq_r16 = &tmp_int_All_Traffic_NumReq_r16
	}
	if F1_C_Traffic_NumReq_r16Present {
		var tmp_int_F1_C_Traffic_NumReq_r16 int64
		if tmp_int_F1_C_Traffic_NumReq_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode F1_C_Traffic_NumReq_r16", err)
		}
		ie.F1_C_Traffic_NumReq_r16 = &tmp_int_F1_C_Traffic_NumReq_r16
	}
	if F1_U_Traffic_NumReq_r16Present {
		var tmp_int_F1_U_Traffic_NumReq_r16 int64
		if tmp_int_F1_U_Traffic_NumReq_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode F1_U_Traffic_NumReq_r16", err)
		}
		ie.F1_U_Traffic_NumReq_r16 = &tmp_int_F1_U_Traffic_NumReq_r16
	}
	if Non_F1_Traffic_NumReq_r16Present {
		var tmp_int_Non_F1_Traffic_NumReq_r16 int64
		if tmp_int_Non_F1_Traffic_NumReq_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Non_F1_Traffic_NumReq_r16", err)
		}
		ie.Non_F1_Traffic_NumReq_r16 = &tmp_int_Non_F1_Traffic_NumReq_r16
	}
	return nil
}
