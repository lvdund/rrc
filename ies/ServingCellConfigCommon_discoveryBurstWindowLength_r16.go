package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfigCommon_discoveryBurstWindowLength_r16_Enum_ms0dot5 aper.Enumerated = 0
	ServingCellConfigCommon_discoveryBurstWindowLength_r16_Enum_ms1     aper.Enumerated = 1
	ServingCellConfigCommon_discoveryBurstWindowLength_r16_Enum_ms2     aper.Enumerated = 2
	ServingCellConfigCommon_discoveryBurstWindowLength_r16_Enum_ms3     aper.Enumerated = 3
	ServingCellConfigCommon_discoveryBurstWindowLength_r16_Enum_ms4     aper.Enumerated = 4
	ServingCellConfigCommon_discoveryBurstWindowLength_r16_Enum_ms5     aper.Enumerated = 5
)

type ServingCellConfigCommon_discoveryBurstWindowLength_r16 struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *ServingCellConfigCommon_discoveryBurstWindowLength_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfigCommon_discoveryBurstWindowLength_r16", err)
	}
	return nil
}

func (ie *ServingCellConfigCommon_discoveryBurstWindowLength_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfigCommon_discoveryBurstWindowLength_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
