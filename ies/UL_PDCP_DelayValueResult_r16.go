package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UL_PDCP_DelayValueResult_r16 struct {
	Drb_Id_r16       DRB_Identity `madatory`
	AverageDelay_r16 int64        `lb:0,ub:10000,madatory`
}

func (ie *UL_PDCP_DelayValueResult_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Drb_Id_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Drb_Id_r16", err)
	}
	if err = w.WriteInteger(ie.AverageDelay_r16, &uper.Constraint{Lb: 0, Ub: 10000}, false); err != nil {
		return utils.WrapError("WriteInteger AverageDelay_r16", err)
	}
	return nil
}

func (ie *UL_PDCP_DelayValueResult_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Drb_Id_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Drb_Id_r16", err)
	}
	var tmp_int_AverageDelay_r16 int64
	if tmp_int_AverageDelay_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10000}, false); err != nil {
		return utils.WrapError("ReadInteger AverageDelay_r16", err)
	}
	ie.AverageDelay_r16 = tmp_int_AverageDelay_r16
	return nil
}
