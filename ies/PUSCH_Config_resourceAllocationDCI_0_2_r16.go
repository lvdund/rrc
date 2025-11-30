package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_Config_resourceAllocationDCI_0_2_r16_Enum_resourceAllocationType0 aper.Enumerated = 0
	PUSCH_Config_resourceAllocationDCI_0_2_r16_Enum_resourceAllocationType1 aper.Enumerated = 1
	PUSCH_Config_resourceAllocationDCI_0_2_r16_Enum_dynamicSwitch           aper.Enumerated = 2
)

type PUSCH_Config_resourceAllocationDCI_0_2_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PUSCH_Config_resourceAllocationDCI_0_2_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PUSCH_Config_resourceAllocationDCI_0_2_r16", err)
	}
	return nil
}

func (ie *PUSCH_Config_resourceAllocationDCI_0_2_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PUSCH_Config_resourceAllocationDCI_0_2_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
