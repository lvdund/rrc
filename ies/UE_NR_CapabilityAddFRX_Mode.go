package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_CapabilityAddFRX_Mode struct {
	Phy_ParametersFRX_Diff       *Phy_ParametersFRX_Diff       `optional`
	MeasAndMobParametersFRX_Diff *MeasAndMobParametersFRX_Diff `optional`
}

func (ie *UE_NR_CapabilityAddFRX_Mode) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Phy_ParametersFRX_Diff != nil, ie.MeasAndMobParametersFRX_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Phy_ParametersFRX_Diff != nil {
		if err = ie.Phy_ParametersFRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_ParametersFRX_Diff", err)
		}
	}
	if ie.MeasAndMobParametersFRX_Diff != nil {
		if err = ie.MeasAndMobParametersFRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersFRX_Diff", err)
		}
	}
	return nil
}

func (ie *UE_NR_CapabilityAddFRX_Mode) Decode(r *aper.AperReader) error {
	var err error
	var Phy_ParametersFRX_DiffPresent bool
	if Phy_ParametersFRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasAndMobParametersFRX_DiffPresent bool
	if MeasAndMobParametersFRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Phy_ParametersFRX_DiffPresent {
		ie.Phy_ParametersFRX_Diff = new(Phy_ParametersFRX_Diff)
		if err = ie.Phy_ParametersFRX_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_ParametersFRX_Diff", err)
		}
	}
	if MeasAndMobParametersFRX_DiffPresent {
		ie.MeasAndMobParametersFRX_Diff = new(MeasAndMobParametersFRX_Diff)
		if err = ie.MeasAndMobParametersFRX_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersFRX_Diff", err)
		}
	}
	return nil
}
