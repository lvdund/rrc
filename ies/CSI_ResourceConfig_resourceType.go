package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ResourceConfig_resourceType_Enum_aperiodic      aper.Enumerated = 0
	CSI_ResourceConfig_resourceType_Enum_semiPersistent aper.Enumerated = 1
	CSI_ResourceConfig_resourceType_Enum_periodic       aper.Enumerated = 2
)

type CSI_ResourceConfig_resourceType struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *CSI_ResourceConfig_resourceType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode CSI_ResourceConfig_resourceType", err)
	}
	return nil
}

func (ie *CSI_ResourceConfig_resourceType) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode CSI_ResourceConfig_resourceType", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
