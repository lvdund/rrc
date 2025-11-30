package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PLMN_RAN_AreaConfig struct {
	Plmn_Identity *PLMN_Identity   `optional`
	Ran_Area      []RAN_AreaConfig `lb:1,ub:16,madatory`
}

func (ie *PLMN_RAN_AreaConfig) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Plmn_Identity != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Plmn_Identity != nil {
		if err = ie.Plmn_Identity.Encode(w); err != nil {
			return utils.WrapError("Encode Plmn_Identity", err)
		}
	}
	tmp_Ran_Area := utils.NewSequence[*RAN_AreaConfig]([]*RAN_AreaConfig{}, aper.Constraint{Lb: 1, Ub: 16}, false)
	for _, i := range ie.Ran_Area {
		tmp_Ran_Area.Value = append(tmp_Ran_Area.Value, &i)
	}
	if err = tmp_Ran_Area.Encode(w); err != nil {
		return utils.WrapError("Encode Ran_Area", err)
	}
	return nil
}

func (ie *PLMN_RAN_AreaConfig) Decode(r *aper.AperReader) error {
	var err error
	var Plmn_IdentityPresent bool
	if Plmn_IdentityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Plmn_IdentityPresent {
		ie.Plmn_Identity = new(PLMN_Identity)
		if err = ie.Plmn_Identity.Decode(r); err != nil {
			return utils.WrapError("Decode Plmn_Identity", err)
		}
	}
	tmp_Ran_Area := utils.NewSequence[*RAN_AreaConfig]([]*RAN_AreaConfig{}, aper.Constraint{Lb: 1, Ub: 16}, false)
	fn_Ran_Area := func() *RAN_AreaConfig {
		return new(RAN_AreaConfig)
	}
	if err = tmp_Ran_Area.Decode(r, fn_Ran_Area); err != nil {
		return utils.WrapError("Decode Ran_Area", err)
	}
	ie.Ran_Area = []RAN_AreaConfig{}
	for _, i := range tmp_Ran_Area.Value {
		ie.Ran_Area = append(ie.Ran_Area, *i)
	}
	return nil
}
