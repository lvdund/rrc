package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz5   aper.Enumerated = 0
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz10  aper.Enumerated = 1
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz15  aper.Enumerated = 2
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz20  aper.Enumerated = 3
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz25  aper.Enumerated = 4
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz30  aper.Enumerated = 5
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz35  aper.Enumerated = 6
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz40  aper.Enumerated = 7
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz45  aper.Enumerated = 8
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz50  aper.Enumerated = 9
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz60  aper.Enumerated = 10
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz70  aper.Enumerated = 11
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz80  aper.Enumerated = 12
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz90  aper.Enumerated = 13
	SupportedBandwidth_v1700_fr1_r17_Enum_mhz100 aper.Enumerated = 14
)

type SupportedBandwidth_v1700_fr1_r17 struct {
	Value aper.Enumerated `lb:0,ub:14,madatory`
}

func (ie *SupportedBandwidth_v1700_fr1_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 14}, false); err != nil {
		return utils.WrapError("Encode SupportedBandwidth_v1700_fr1_r17", err)
	}
	return nil
}

func (ie *SupportedBandwidth_v1700_fr1_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 14}, false); err != nil {
		return utils.WrapError("Decode SupportedBandwidth_v1700_fr1_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
