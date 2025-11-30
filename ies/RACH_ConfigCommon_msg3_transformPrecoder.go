package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_msg3_transformPrecoder_Enum_enabled aper.Enumerated = 0
)

type RACH_ConfigCommon_msg3_transformPrecoder struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *RACH_ConfigCommon_msg3_transformPrecoder) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigCommon_msg3_transformPrecoder", err)
	}
	return nil
}

func (ie *RACH_ConfigCommon_msg3_transformPrecoder) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigCommon_msg3_transformPrecoder", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
