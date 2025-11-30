package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfigCommon_discoveryBurstWindowLength_r17_Enum_ms0dot125 aper.Enumerated = 0
	ServingCellConfigCommon_discoveryBurstWindowLength_r17_Enum_ms0dot25  aper.Enumerated = 1
	ServingCellConfigCommon_discoveryBurstWindowLength_r17_Enum_ms0dot5   aper.Enumerated = 2
	ServingCellConfigCommon_discoveryBurstWindowLength_r17_Enum_ms0dot75  aper.Enumerated = 3
	ServingCellConfigCommon_discoveryBurstWindowLength_r17_Enum_ms1       aper.Enumerated = 4
	ServingCellConfigCommon_discoveryBurstWindowLength_r17_Enum_ms1dot25  aper.Enumerated = 5
)

type ServingCellConfigCommon_discoveryBurstWindowLength_r17 struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *ServingCellConfigCommon_discoveryBurstWindowLength_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfigCommon_discoveryBurstWindowLength_r17", err)
	}
	return nil
}

func (ie *ServingCellConfigCommon_discoveryBurstWindowLength_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfigCommon_discoveryBurstWindowLength_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
