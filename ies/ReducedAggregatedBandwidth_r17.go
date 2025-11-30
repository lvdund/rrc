package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReducedAggregatedBandwidth_r17_Enum_mhz0    aper.Enumerated = 0
	ReducedAggregatedBandwidth_r17_Enum_mhz100  aper.Enumerated = 1
	ReducedAggregatedBandwidth_r17_Enum_mhz200  aper.Enumerated = 2
	ReducedAggregatedBandwidth_r17_Enum_mhz400  aper.Enumerated = 3
	ReducedAggregatedBandwidth_r17_Enum_mhz800  aper.Enumerated = 4
	ReducedAggregatedBandwidth_r17_Enum_mhz1200 aper.Enumerated = 5
	ReducedAggregatedBandwidth_r17_Enum_mhz1600 aper.Enumerated = 6
	ReducedAggregatedBandwidth_r17_Enum_mhz2000 aper.Enumerated = 7
)

type ReducedAggregatedBandwidth_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *ReducedAggregatedBandwidth_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode ReducedAggregatedBandwidth_r17", err)
	}
	return nil
}

func (ie *ReducedAggregatedBandwidth_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode ReducedAggregatedBandwidth_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
