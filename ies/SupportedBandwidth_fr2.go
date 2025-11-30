package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SupportedBandwidth_fr2_Enum_mhz50  aper.Enumerated = 0
	SupportedBandwidth_fr2_Enum_mhz100 aper.Enumerated = 1
	SupportedBandwidth_fr2_Enum_mhz200 aper.Enumerated = 2
	SupportedBandwidth_fr2_Enum_mhz400 aper.Enumerated = 3
)

type SupportedBandwidth_fr2 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SupportedBandwidth_fr2) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SupportedBandwidth_fr2", err)
	}
	return nil
}

func (ie *SupportedBandwidth_fr2) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SupportedBandwidth_fr2", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
