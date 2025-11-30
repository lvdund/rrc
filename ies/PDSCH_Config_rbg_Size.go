package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_Config_rbg_Size_Enum_config1 aper.Enumerated = 0
	PDSCH_Config_rbg_Size_Enum_config2 aper.Enumerated = 1
)

type PDSCH_Config_rbg_Size struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PDSCH_Config_rbg_Size) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PDSCH_Config_rbg_Size", err)
	}
	return nil
}

func (ie *PDSCH_Config_rbg_Size) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PDSCH_Config_rbg_Size", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
