package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DL_PRS_QCL_Info_r17_ssb_r17 struct {
	Ssb_Index_r17 int64                                   `lb:0,ub:63,madatory`
	Rs_Type_r17   DL_PRS_QCL_Info_r17_ssb_r17_rs_Type_r17 `madatory`
}

func (ie *DL_PRS_QCL_Info_r17_ssb_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Ssb_Index_r17, &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger Ssb_Index_r17", err)
	}
	if err = ie.Rs_Type_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Rs_Type_r17", err)
	}
	return nil
}

func (ie *DL_PRS_QCL_Info_r17_ssb_r17) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_Ssb_Index_r17 int64
	if tmp_int_Ssb_Index_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger Ssb_Index_r17", err)
	}
	ie.Ssb_Index_r17 = tmp_int_Ssb_Index_r17
	if err = ie.Rs_Type_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Rs_Type_r17", err)
	}
	return nil
}
