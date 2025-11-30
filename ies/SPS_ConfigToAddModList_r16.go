package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SPS_ConfigToAddModList_r16 struct {
	Value []SPS_Config `lb:1,ub:maxNrofSPS_Config_r16,madatory`
}

func (ie *SPS_ConfigToAddModList_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*SPS_Config]([]*SPS_Config{}, aper.Constraint{Lb: 1, Ub: maxNrofSPS_Config_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SPS_ConfigToAddModList_r16", err)
	}
	return nil
}

func (ie *SPS_ConfigToAddModList_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*SPS_Config]([]*SPS_Config{}, aper.Constraint{Lb: 1, Ub: maxNrofSPS_Config_r16}, false)
	fn := func() *SPS_Config {
		return new(SPS_Config)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SPS_ConfigToAddModList_r16", err)
	}
	ie.Value = []SPS_Config{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
