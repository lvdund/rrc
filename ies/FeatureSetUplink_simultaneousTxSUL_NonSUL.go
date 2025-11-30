package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplink_simultaneousTxSUL_NonSUL_Enum_supported aper.Enumerated = 0
)

type FeatureSetUplink_simultaneousTxSUL_NonSUL struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *FeatureSetUplink_simultaneousTxSUL_NonSUL) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplink_simultaneousTxSUL_NonSUL", err)
	}
	return nil
}

func (ie *FeatureSetUplink_simultaneousTxSUL_NonSUL) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplink_simultaneousTxSUL_NonSUL", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
