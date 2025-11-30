package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfig_pathlossReferenceLinking_Enum_spCell aper.Enumerated = 0
	ServingCellConfig_pathlossReferenceLinking_Enum_sCell  aper.Enumerated = 1
)

type ServingCellConfig_pathlossReferenceLinking struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *ServingCellConfig_pathlossReferenceLinking) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode ServingCellConfig_pathlossReferenceLinking", err)
	}
	return nil
}

func (ie *ServingCellConfig_pathlossReferenceLinking) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode ServingCellConfig_pathlossReferenceLinking", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
