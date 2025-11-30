package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_ConfigBroadcast_r17_xOverhead_r17_Enum_xOh6  aper.Enumerated = 0
	PDSCH_ConfigBroadcast_r17_xOverhead_r17_Enum_xOh12 aper.Enumerated = 1
	PDSCH_ConfigBroadcast_r17_xOverhead_r17_Enum_xOh18 aper.Enumerated = 2
)

type PDSCH_ConfigBroadcast_r17_xOverhead_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PDSCH_ConfigBroadcast_r17_xOverhead_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PDSCH_ConfigBroadcast_r17_xOverhead_r17", err)
	}
	return nil
}

func (ie *PDSCH_ConfigBroadcast_r17_xOverhead_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PDSCH_ConfigBroadcast_r17_xOverhead_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
