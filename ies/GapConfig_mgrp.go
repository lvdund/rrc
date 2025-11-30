package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	GapConfig_mgrp_Enum_ms20  aper.Enumerated = 0
	GapConfig_mgrp_Enum_ms40  aper.Enumerated = 1
	GapConfig_mgrp_Enum_ms80  aper.Enumerated = 2
	GapConfig_mgrp_Enum_ms160 aper.Enumerated = 3
)

type GapConfig_mgrp struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *GapConfig_mgrp) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode GapConfig_mgrp", err)
	}
	return nil
}

func (ie *GapConfig_mgrp) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode GapConfig_mgrp", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
