package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type OffsetValue_r17 struct {
	OffsetValue_r17   int64 `lb:-20000,ub:20000,madatory`
	Shift7dot5kHz_r17 bool  `madatory`
}

func (ie *OffsetValue_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.OffsetValue_r17, &aper.Constraint{Lb: -20000, Ub: 20000}, false); err != nil {
		return utils.WrapError("WriteInteger OffsetValue_r17", err)
	}
	if err = w.WriteBoolean(ie.Shift7dot5kHz_r17); err != nil {
		return utils.WrapError("WriteBoolean Shift7dot5kHz_r17", err)
	}
	return nil
}

func (ie *OffsetValue_r17) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_OffsetValue_r17 int64
	if tmp_int_OffsetValue_r17, err = r.ReadInteger(&aper.Constraint{Lb: -20000, Ub: 20000}, false); err != nil {
		return utils.WrapError("ReadInteger OffsetValue_r17", err)
	}
	ie.OffsetValue_r17 = tmp_int_OffsetValue_r17
	var tmp_bool_Shift7dot5kHz_r17 bool
	if tmp_bool_Shift7dot5kHz_r17, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean Shift7dot5kHz_r17", err)
	}
	ie.Shift7dot5kHz_r17 = tmp_bool_Shift7dot5kHz_r17
	return nil
}
