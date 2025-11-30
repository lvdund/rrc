package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SupportedBandwidth_fr1_Enum_mhz5   aper.Enumerated = 0
	SupportedBandwidth_fr1_Enum_mhz10  aper.Enumerated = 1
	SupportedBandwidth_fr1_Enum_mhz15  aper.Enumerated = 2
	SupportedBandwidth_fr1_Enum_mhz20  aper.Enumerated = 3
	SupportedBandwidth_fr1_Enum_mhz25  aper.Enumerated = 4
	SupportedBandwidth_fr1_Enum_mhz30  aper.Enumerated = 5
	SupportedBandwidth_fr1_Enum_mhz40  aper.Enumerated = 6
	SupportedBandwidth_fr1_Enum_mhz50  aper.Enumerated = 7
	SupportedBandwidth_fr1_Enum_mhz60  aper.Enumerated = 8
	SupportedBandwidth_fr1_Enum_mhz80  aper.Enumerated = 9
	SupportedBandwidth_fr1_Enum_mhz100 aper.Enumerated = 10
)

type SupportedBandwidth_fr1 struct {
	Value aper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *SupportedBandwidth_fr1) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode SupportedBandwidth_fr1", err)
	}
	return nil
}

func (ie *SupportedBandwidth_fr1) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode SupportedBandwidth_fr1", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
