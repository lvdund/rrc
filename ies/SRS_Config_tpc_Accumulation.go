package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Config_tpc_Accumulation_Enum_disabled aper.Enumerated = 0
)

type SRS_Config_tpc_Accumulation struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *SRS_Config_tpc_Accumulation) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode SRS_Config_tpc_Accumulation", err)
	}
	return nil
}

func (ie *SRS_Config_tpc_Accumulation) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode SRS_Config_tpc_Accumulation", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
