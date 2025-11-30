package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication_Enum_supported aper.Enumerated = 0
)

type FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_v1540_dl_MCS_TableAlt_DynamicIndication", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
