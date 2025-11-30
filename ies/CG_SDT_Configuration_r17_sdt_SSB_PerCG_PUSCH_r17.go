package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_oneEighth aper.Enumerated = 0
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_oneFourth aper.Enumerated = 1
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_half      aper.Enumerated = 2
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_one       aper.Enumerated = 3
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_two       aper.Enumerated = 4
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_four      aper.Enumerated = 5
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_eight     aper.Enumerated = 6
	CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17_Enum_sixteen   aper.Enumerated = 7
)

type CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17", err)
	}
	return nil
}

func (ie *CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
