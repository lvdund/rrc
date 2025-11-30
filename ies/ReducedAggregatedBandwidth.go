package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReducedAggregatedBandwidth_Enum_mhz0   aper.Enumerated = 0
	ReducedAggregatedBandwidth_Enum_mhz10  aper.Enumerated = 1
	ReducedAggregatedBandwidth_Enum_mhz20  aper.Enumerated = 2
	ReducedAggregatedBandwidth_Enum_mhz30  aper.Enumerated = 3
	ReducedAggregatedBandwidth_Enum_mhz40  aper.Enumerated = 4
	ReducedAggregatedBandwidth_Enum_mhz50  aper.Enumerated = 5
	ReducedAggregatedBandwidth_Enum_mhz60  aper.Enumerated = 6
	ReducedAggregatedBandwidth_Enum_mhz80  aper.Enumerated = 7
	ReducedAggregatedBandwidth_Enum_mhz100 aper.Enumerated = 8
	ReducedAggregatedBandwidth_Enum_mhz200 aper.Enumerated = 9
	ReducedAggregatedBandwidth_Enum_mhz300 aper.Enumerated = 10
	ReducedAggregatedBandwidth_Enum_mhz400 aper.Enumerated = 11
)

type ReducedAggregatedBandwidth struct {
	Value aper.Enumerated `lb:0,ub:11,madatory`
}

func (ie *ReducedAggregatedBandwidth) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("Encode ReducedAggregatedBandwidth", err)
	}
	return nil
}

func (ie *ReducedAggregatedBandwidth) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
		return utils.WrapError("Decode ReducedAggregatedBandwidth", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
