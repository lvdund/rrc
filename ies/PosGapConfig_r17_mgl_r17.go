package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosGapConfig_r17_mgl_r17_Enum_ms1dot5 aper.Enumerated = 0
	PosGapConfig_r17_mgl_r17_Enum_ms3     aper.Enumerated = 1
	PosGapConfig_r17_mgl_r17_Enum_ms3dot5 aper.Enumerated = 2
	PosGapConfig_r17_mgl_r17_Enum_ms4     aper.Enumerated = 3
	PosGapConfig_r17_mgl_r17_Enum_ms5dot5 aper.Enumerated = 4
	PosGapConfig_r17_mgl_r17_Enum_ms6     aper.Enumerated = 5
	PosGapConfig_r17_mgl_r17_Enum_ms10    aper.Enumerated = 6
	PosGapConfig_r17_mgl_r17_Enum_ms20    aper.Enumerated = 7
)

type PosGapConfig_r17_mgl_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *PosGapConfig_r17_mgl_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PosGapConfig_r17_mgl_r17", err)
	}
	return nil
}

func (ie *PosGapConfig_r17_mgl_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PosGapConfig_r17_mgl_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
