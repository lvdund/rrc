package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfiguredGrantConfig_resourceAllocation_Enum_resourceAllocationType0 aper.Enumerated = 0
	ConfiguredGrantConfig_resourceAllocation_Enum_resourceAllocationType1 aper.Enumerated = 1
	ConfiguredGrantConfig_resourceAllocation_Enum_dynamicSwitch           aper.Enumerated = 2
)

type ConfiguredGrantConfig_resourceAllocation struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *ConfiguredGrantConfig_resourceAllocation) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfig_resourceAllocation", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfig_resourceAllocation) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfig_resourceAllocation", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
