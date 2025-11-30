package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ControlResourceSet_precoderGranularity_Enum_sameAsREG_bundle aper.Enumerated = 0
	ControlResourceSet_precoderGranularity_Enum_allContiguousRBs aper.Enumerated = 1
)

type ControlResourceSet_precoderGranularity struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *ControlResourceSet_precoderGranularity) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode ControlResourceSet_precoderGranularity", err)
	}
	return nil
}

func (ie *ControlResourceSet_precoderGranularity) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode ControlResourceSet_precoderGranularity", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
