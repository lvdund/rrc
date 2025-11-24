package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PerRASSBInfo_r16 struct {
	Ssb_Index_r16                  SSB_Index                `madatory`
	NumberOfPreamblesSentOnSSB_r16 int64                    `lb:1,ub:200,madatory`
	PerRAAttemptInfoList_r16       PerRAAttemptInfoList_r16 `madatory`
}

func (ie *PerRASSBInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Ssb_Index_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_Index_r16", err)
	}
	if err = w.WriteInteger(ie.NumberOfPreamblesSentOnSSB_r16, &uper.Constraint{Lb: 1, Ub: 200}, false); err != nil {
		return utils.WrapError("WriteInteger NumberOfPreamblesSentOnSSB_r16", err)
	}
	if err = ie.PerRAAttemptInfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PerRAAttemptInfoList_r16", err)
	}
	return nil
}

func (ie *PerRASSBInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Ssb_Index_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb_Index_r16", err)
	}
	var tmp_int_NumberOfPreamblesSentOnSSB_r16 int64
	if tmp_int_NumberOfPreamblesSentOnSSB_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 200}, false); err != nil {
		return utils.WrapError("ReadInteger NumberOfPreamblesSentOnSSB_r16", err)
	}
	ie.NumberOfPreamblesSentOnSSB_r16 = tmp_int_NumberOfPreamblesSentOnSSB_r16
	if err = ie.PerRAAttemptInfoList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PerRAAttemptInfoList_r16", err)
	}
	return nil
}
