package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PLMN_RAN_AreaConfigList struct {
	Value []PLMN_RAN_AreaConfig `lb:1,ub:maxPLMNIdentities,madatory`
}

func (ie *PLMN_RAN_AreaConfigList) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*PLMN_RAN_AreaConfig]([]*PLMN_RAN_AreaConfig{}, aper.Constraint{Lb: 1, Ub: maxPLMNIdentities}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PLMN_RAN_AreaConfigList", err)
	}
	return nil
}

func (ie *PLMN_RAN_AreaConfigList) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*PLMN_RAN_AreaConfig]([]*PLMN_RAN_AreaConfig{}, aper.Constraint{Lb: 1, Ub: maxPLMNIdentities}, false)
	fn := func() *PLMN_RAN_AreaConfig {
		return new(PLMN_RAN_AreaConfig)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PLMN_RAN_AreaConfigList", err)
	}
	ie.Value = []PLMN_RAN_AreaConfig{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
