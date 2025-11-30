package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resource_ptrs_PortIndex_Enum_n0 aper.Enumerated = 0
	SRS_Resource_ptrs_PortIndex_Enum_n1 aper.Enumerated = 1
)

type SRS_Resource_ptrs_PortIndex struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SRS_Resource_ptrs_PortIndex) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SRS_Resource_ptrs_PortIndex", err)
	}
	return nil
}

func (ie *SRS_Resource_ptrs_PortIndex) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SRS_Resource_ptrs_PortIndex", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
