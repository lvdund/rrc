package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17_Enum_pbfd1  aper.Enumerated = 0
	BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17_Enum_pbfd2  aper.Enumerated = 1
	BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17_Enum_pbfd3  aper.Enumerated = 2
	BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17_Enum_pbfd4  aper.Enumerated = 3
	BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17_Enum_pbfd5  aper.Enumerated = 4
	BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17_Enum_pbfd6  aper.Enumerated = 5
	BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17_Enum_pbfd8  aper.Enumerated = 6
	BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17_Enum_pbfd10 aper.Enumerated = 7
)

type BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17", err)
	}
	return nil
}

func (ie *BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode BeamFailureDetectionSet_r17_beamFailureDetectionTimer_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
