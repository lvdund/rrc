package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosResource_r16_freqHopping_r16 struct {
	C_SRS_r16 int64 `lb:0,ub:63,madatory`
}

func (ie *SRS_PosResource_r16_freqHopping_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.C_SRS_r16, &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger C_SRS_r16", err)
	}
	return nil
}

func (ie *SRS_PosResource_r16_freqHopping_r16) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_C_SRS_r16 int64
	if tmp_int_C_SRS_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger C_SRS_r16", err)
	}
	ie.C_SRS_r16 = tmp_int_C_SRS_r16
	return nil
}
