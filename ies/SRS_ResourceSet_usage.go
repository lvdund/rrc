package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_ResourceSet_usage_Enum_beamManagement   aper.Enumerated = 0
	SRS_ResourceSet_usage_Enum_codebook         aper.Enumerated = 1
	SRS_ResourceSet_usage_Enum_nonCodebook      aper.Enumerated = 2
	SRS_ResourceSet_usage_Enum_antennaSwitching aper.Enumerated = 3
)

type SRS_ResourceSet_usage struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SRS_ResourceSet_usage) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SRS_ResourceSet_usage", err)
	}
	return nil
}

func (ie *SRS_ResourceSet_usage) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SRS_ResourceSet_usage", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
