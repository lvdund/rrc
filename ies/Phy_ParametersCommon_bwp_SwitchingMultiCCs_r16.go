package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_nothing uint64 = iota
	Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_Type1_r16
	Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_Type2_r16
)

type Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16 struct {
	Choice    uint64
	Type1_r16 *Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_type1_r16
	Type2_r16 *Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_type2_r16
}

func (ie *Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_Type1_r16:
		if err = ie.Type1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Type1_r16", err)
		}
	case Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_Type2_r16:
		if err = ie.Type2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Type2_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_Type1_r16:
		ie.Type1_r16 = new(Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_type1_r16)
		if err = ie.Type1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Type1_r16", err)
		}
	case Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_Choice_Type2_r16:
		ie.Type2_r16 = new(Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16_type2_r16)
		if err = ie.Type2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Type2_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
