package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EpochTime_r17 struct {
	Sfn_r17        int64 `lb:0,ub:1023,madatory`
	SubFrameNR_r17 int64 `lb:0,ub:9,madatory`
}

func (ie *EpochTime_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Sfn_r17, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("WriteInteger Sfn_r17", err)
	}
	if err = w.WriteInteger(ie.SubFrameNR_r17, &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("WriteInteger SubFrameNR_r17", err)
	}
	return nil
}

func (ie *EpochTime_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_Sfn_r17 int64
	if tmp_int_Sfn_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("ReadInteger Sfn_r17", err)
	}
	ie.Sfn_r17 = tmp_int_Sfn_r17
	var tmp_int_SubFrameNR_r17 int64
	if tmp_int_SubFrameNR_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("ReadInteger SubFrameNR_r17", err)
	}
	ie.SubFrameNR_r17 = tmp_int_SubFrameNR_r17
	return nil
}
