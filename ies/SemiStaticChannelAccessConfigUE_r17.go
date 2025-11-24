package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SemiStaticChannelAccessConfigUE_r17 struct {
	PeriodUE_r17 SemiStaticChannelAccessConfigUE_r17_periodUE_r17 `madatory`
	OffsetUE_r17 int64                                            `lb:0,ub:559,madatory`
}

func (ie *SemiStaticChannelAccessConfigUE_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.PeriodUE_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PeriodUE_r17", err)
	}
	if err = w.WriteInteger(ie.OffsetUE_r17, &uper.Constraint{Lb: 0, Ub: 559}, false); err != nil {
		return utils.WrapError("WriteInteger OffsetUE_r17", err)
	}
	return nil
}

func (ie *SemiStaticChannelAccessConfigUE_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.PeriodUE_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PeriodUE_r17", err)
	}
	var tmp_int_OffsetUE_r17 int64
	if tmp_int_OffsetUE_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 559}, false); err != nil {
		return utils.WrapError("ReadInteger OffsetUE_r17", err)
	}
	ie.OffsetUE_r17 = tmp_int_OffsetUE_r17
	return nil
}
