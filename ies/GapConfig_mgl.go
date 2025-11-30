package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	GapConfig_mgl_Enum_ms1dot5 aper.Enumerated = 0
	GapConfig_mgl_Enum_ms3     aper.Enumerated = 1
	GapConfig_mgl_Enum_ms3dot5 aper.Enumerated = 2
	GapConfig_mgl_Enum_ms4     aper.Enumerated = 3
	GapConfig_mgl_Enum_ms5dot5 aper.Enumerated = 4
	GapConfig_mgl_Enum_ms6     aper.Enumerated = 5
)

type GapConfig_mgl struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *GapConfig_mgl) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode GapConfig_mgl", err)
	}
	return nil
}

func (ie *GapConfig_mgl) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode GapConfig_mgl", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
