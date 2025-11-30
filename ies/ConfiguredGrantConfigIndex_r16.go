package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ConfiguredGrantConfigIndex_r16 struct {
	Value uint64 `lb:0,ub:maxNrofConfiguredGrantConfig_1_r16,madatory`
}

func (ie *ConfiguredGrantConfigIndex_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: maxNrofConfiguredGrantConfig_1_r16}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfigIndex_r16", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfigIndex_r16) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofConfiguredGrantConfig_1_r16}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfigIndex_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
