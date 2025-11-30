package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_activeConfiguredGrant_r16 struct {
	MaxNumberConfigsPerBWP_r16 BandNR_activeConfiguredGrant_r16_maxNumberConfigsPerBWP_r16 `madatory`
	MaxNumberConfigsAllCC_r16  int64                                                       `lb:2,ub:32,madatory`
}

func (ie *BandNR_activeConfiguredGrant_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.MaxNumberConfigsPerBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberConfigsPerBWP_r16", err)
	}
	if err = w.WriteInteger(ie.MaxNumberConfigsAllCC_r16, &aper.Constraint{Lb: 2, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberConfigsAllCC_r16", err)
	}
	return nil
}

func (ie *BandNR_activeConfiguredGrant_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.MaxNumberConfigsPerBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberConfigsPerBWP_r16", err)
	}
	var tmp_int_MaxNumberConfigsAllCC_r16 int64
	if tmp_int_MaxNumberConfigsAllCC_r16, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberConfigsAllCC_r16", err)
	}
	ie.MaxNumberConfigsAllCC_r16 = tmp_int_MaxNumberConfigsAllCC_r16
	return nil
}
