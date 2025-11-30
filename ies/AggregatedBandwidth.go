package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	AggregatedBandwidth_Enum_mhz50  aper.Enumerated = 0
	AggregatedBandwidth_Enum_mhz100 aper.Enumerated = 1
	AggregatedBandwidth_Enum_mhz150 aper.Enumerated = 2
	AggregatedBandwidth_Enum_mhz200 aper.Enumerated = 3
	AggregatedBandwidth_Enum_mhz250 aper.Enumerated = 4
	AggregatedBandwidth_Enum_mhz300 aper.Enumerated = 5
	AggregatedBandwidth_Enum_mhz350 aper.Enumerated = 6
	AggregatedBandwidth_Enum_mhz400 aper.Enumerated = 7
	AggregatedBandwidth_Enum_mhz450 aper.Enumerated = 8
	AggregatedBandwidth_Enum_mhz500 aper.Enumerated = 9
	AggregatedBandwidth_Enum_mhz550 aper.Enumerated = 10
	AggregatedBandwidth_Enum_mhz600 aper.Enumerated = 11
	AggregatedBandwidth_Enum_mhz650 aper.Enumerated = 12
	AggregatedBandwidth_Enum_mhz700 aper.Enumerated = 13
	AggregatedBandwidth_Enum_mhz750 aper.Enumerated = 14
	AggregatedBandwidth_Enum_mhz800 aper.Enumerated = 15
)

type AggregatedBandwidth struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *AggregatedBandwidth) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode AggregatedBandwidth", err)
	}
	return nil
}

func (ie *AggregatedBandwidth) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode AggregatedBandwidth", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
