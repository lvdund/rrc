package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Fr1_100mhz_scs_15kHz_Enum_supported aper.Enumerated = 0
)

type Fr1_100mhz_scs_15kHz struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *Fr1_100mhz_scs_15kHz) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode Fr1_100mhz_scs_15kHz", err)
	}
	return nil
}

func (ie *Fr1_100mhz_scs_15kHz) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode Fr1_100mhz_scs_15kHz", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
