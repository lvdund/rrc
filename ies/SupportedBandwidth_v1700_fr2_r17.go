package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz50   aper.Enumerated = 0
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz100  aper.Enumerated = 1
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz200  aper.Enumerated = 2
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz400  aper.Enumerated = 3
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz800  aper.Enumerated = 4
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz1600 aper.Enumerated = 5
	SupportedBandwidth_v1700_fr2_r17_Enum_mhz2000 aper.Enumerated = 6
)

type SupportedBandwidth_v1700_fr2_r17 struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *SupportedBandwidth_v1700_fr2_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode SupportedBandwidth_v1700_fr2_r17", err)
	}
	return nil
}

func (ie *SupportedBandwidth_v1700_fr2_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode SupportedBandwidth_v1700_fr2_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
