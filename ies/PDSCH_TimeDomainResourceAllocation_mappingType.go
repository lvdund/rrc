package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_TimeDomainResourceAllocation_mappingType_Enum_typeA aper.Enumerated = 0
	PDSCH_TimeDomainResourceAllocation_mappingType_Enum_typeB aper.Enumerated = 1
)

type PDSCH_TimeDomainResourceAllocation_mappingType struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PDSCH_TimeDomainResourceAllocation_mappingType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PDSCH_TimeDomainResourceAllocation_mappingType", err)
	}
	return nil
}

func (ie *PDSCH_TimeDomainResourceAllocation_mappingType) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PDSCH_TimeDomainResourceAllocation_mappingType", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
