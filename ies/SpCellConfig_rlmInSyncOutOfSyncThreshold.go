package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpCellConfig_rlmInSyncOutOfSyncThreshold_Enum_n1 aper.Enumerated = 0
)

type SpCellConfig_rlmInSyncOutOfSyncThreshold struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *SpCellConfig_rlmInSyncOutOfSyncThreshold) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode SpCellConfig_rlmInSyncOutOfSyncThreshold", err)
	}
	return nil
}

func (ie *SpCellConfig_rlmInSyncOutOfSyncThreshold) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode SpCellConfig_rlmInSyncOutOfSyncThreshold", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
