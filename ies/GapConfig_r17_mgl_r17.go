package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	GapConfig_r17_mgl_r17_Enum_ms1     aper.Enumerated = 0
	GapConfig_r17_mgl_r17_Enum_ms1dot5 aper.Enumerated = 1
	GapConfig_r17_mgl_r17_Enum_ms2     aper.Enumerated = 2
	GapConfig_r17_mgl_r17_Enum_ms3     aper.Enumerated = 3
	GapConfig_r17_mgl_r17_Enum_ms3dot5 aper.Enumerated = 4
	GapConfig_r17_mgl_r17_Enum_ms4     aper.Enumerated = 5
	GapConfig_r17_mgl_r17_Enum_ms5     aper.Enumerated = 6
	GapConfig_r17_mgl_r17_Enum_ms5dot5 aper.Enumerated = 7
	GapConfig_r17_mgl_r17_Enum_ms6     aper.Enumerated = 8
	GapConfig_r17_mgl_r17_Enum_ms10    aper.Enumerated = 9
	GapConfig_r17_mgl_r17_Enum_ms20    aper.Enumerated = 10
)

type GapConfig_r17_mgl_r17 struct {
	Value aper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *GapConfig_r17_mgl_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode GapConfig_r17_mgl_r17", err)
	}
	return nil
}

func (ie *GapConfig_r17_mgl_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode GapConfig_r17_mgl_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
