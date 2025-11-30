package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosGapConfig_r17_mgrp_r17_Enum_ms20  aper.Enumerated = 0
	PosGapConfig_r17_mgrp_r17_Enum_ms40  aper.Enumerated = 1
	PosGapConfig_r17_mgrp_r17_Enum_ms80  aper.Enumerated = 2
	PosGapConfig_r17_mgrp_r17_Enum_ms160 aper.Enumerated = 3
)

type PosGapConfig_r17_mgrp_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PosGapConfig_r17_mgrp_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PosGapConfig_r17_mgrp_r17", err)
	}
	return nil
}

func (ie *PosGapConfig_r17_mgrp_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PosGapConfig_r17_mgrp_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
