package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_SSB_ResourceSetId struct {
	Value uint64 `lb:0,ub:maxNrofCSI_SSB_ResourceSets_1,madatory`
}

func (ie *CSI_SSB_ResourceSetId) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: maxNrofCSI_SSB_ResourceSets_1}, false); err != nil {
		return utils.WrapError("Encode CSI_SSB_ResourceSetId", err)
	}
	return nil
}

func (ie *CSI_SSB_ResourceSetId) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofCSI_SSB_ResourceSets_1}, false); err != nil {
		return utils.WrapError("Decode CSI_SSB_ResourceSetId", err)
	}
	ie.Value = uint64(v)
	return nil
}
