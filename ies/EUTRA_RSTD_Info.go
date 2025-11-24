package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_RSTD_Info struct {
	CarrierFreq    ARFCN_ValueEUTRA `madatory`
	MeasPRS_Offset int64            `lb:0,ub:39,madatory`
}

func (ie *EUTRA_RSTD_Info) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.CarrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq", err)
	}
	if err = w.WriteInteger(ie.MeasPRS_Offset, &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
		return utils.WrapError("WriteInteger MeasPRS_Offset", err)
	}
	return nil
}

func (ie *EUTRA_RSTD_Info) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.CarrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq", err)
	}
	var tmp_int_MeasPRS_Offset int64
	if tmp_int_MeasPRS_Offset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
		return utils.WrapError("ReadInteger MeasPRS_Offset", err)
	}
	ie.MeasPRS_Offset = tmp_int_MeasPRS_Offset
	return nil
}
