package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MPE_Config_FR2_r17 struct {
	Mpe_ProhibitTimer_r17 MPE_Config_FR2_r17_mpe_ProhibitTimer_r17 `madatory`
	Mpe_Threshold_r17     MPE_Config_FR2_r17_mpe_Threshold_r17     `madatory`
	NumberOfN_r17         int64                                    `lb:1,ub:4,madatory`
}

func (ie *MPE_Config_FR2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Mpe_ProhibitTimer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mpe_ProhibitTimer_r17", err)
	}
	if err = ie.Mpe_Threshold_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mpe_Threshold_r17", err)
	}
	if err = w.WriteInteger(ie.NumberOfN_r17, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger NumberOfN_r17", err)
	}
	return nil
}

func (ie *MPE_Config_FR2_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Mpe_ProhibitTimer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mpe_ProhibitTimer_r17", err)
	}
	if err = ie.Mpe_Threshold_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mpe_Threshold_r17", err)
	}
	var tmp_int_NumberOfN_r17 int64
	if tmp_int_NumberOfN_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger NumberOfN_r17", err)
	}
	ie.NumberOfN_r17 = tmp_int_NumberOfN_r17
	return nil
}
