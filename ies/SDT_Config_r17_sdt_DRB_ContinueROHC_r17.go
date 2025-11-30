package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SDT_Config_r17_sdt_DRB_ContinueROHC_r17_Enum_cell aper.Enumerated = 0
	SDT_Config_r17_sdt_DRB_ContinueROHC_r17_Enum_rna  aper.Enumerated = 1
)

type SDT_Config_r17_sdt_DRB_ContinueROHC_r17 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SDT_Config_r17_sdt_DRB_ContinueROHC_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SDT_Config_r17_sdt_DRB_ContinueROHC_r17", err)
	}
	return nil
}

func (ie *SDT_Config_r17_sdt_DRB_ContinueROHC_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SDT_Config_r17_sdt_DRB_ContinueROHC_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
