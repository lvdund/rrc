package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_n6     aper.Enumerated = 0
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_n15    aper.Enumerated = 1
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_n25    aper.Enumerated = 2
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_n50    aper.Enumerated = 3
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_n75    aper.Enumerated = 4
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_n100   aper.Enumerated = 5
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_spare2 aper.Enumerated = 6
	RateMatchPatternLTE_CRS_carrierBandwidthDL_Enum_spare1 aper.Enumerated = 7
)

type RateMatchPatternLTE_CRS_carrierBandwidthDL struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RateMatchPatternLTE_CRS_carrierBandwidthDL) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RateMatchPatternLTE_CRS_carrierBandwidthDL", err)
	}
	return nil
}

func (ie *RateMatchPatternLTE_CRS_carrierBandwidthDL) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RateMatchPatternLTE_CRS_carrierBandwidthDL", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
