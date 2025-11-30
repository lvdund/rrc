package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_Config_resourceAllocation_Enum_resourceAllocationType0 aper.Enumerated = 0
	PDSCH_Config_resourceAllocation_Enum_resourceAllocationType1 aper.Enumerated = 1
	PDSCH_Config_resourceAllocation_Enum_dynamicSwitch           aper.Enumerated = 2
)

type PDSCH_Config_resourceAllocation struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PDSCH_Config_resourceAllocation) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PDSCH_Config_resourceAllocation", err)
	}
	return nil
}

func (ie *PDSCH_Config_resourceAllocation) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PDSCH_Config_resourceAllocation", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
