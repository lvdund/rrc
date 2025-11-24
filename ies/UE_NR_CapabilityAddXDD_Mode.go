package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_CapabilityAddXDD_Mode struct {
	Phy_ParametersXDD_Diff       *Phy_ParametersXDD_Diff       `optional`
	Mac_ParametersXDD_Diff       *MAC_ParametersXDD_Diff       `optional`
	MeasAndMobParametersXDD_Diff *MeasAndMobParametersXDD_Diff `optional`
}

func (ie *UE_NR_CapabilityAddXDD_Mode) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Phy_ParametersXDD_Diff != nil, ie.Mac_ParametersXDD_Diff != nil, ie.MeasAndMobParametersXDD_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Phy_ParametersXDD_Diff != nil {
		if err = ie.Phy_ParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_ParametersXDD_Diff", err)
		}
	}
	if ie.Mac_ParametersXDD_Diff != nil {
		if err = ie.Mac_ParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_ParametersXDD_Diff", err)
		}
	}
	if ie.MeasAndMobParametersXDD_Diff != nil {
		if err = ie.MeasAndMobParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersXDD_Diff", err)
		}
	}
	return nil
}

func (ie *UE_NR_CapabilityAddXDD_Mode) Decode(r *uper.UperReader) error {
	var err error
	var Phy_ParametersXDD_DiffPresent bool
	if Phy_ParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mac_ParametersXDD_DiffPresent bool
	if Mac_ParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasAndMobParametersXDD_DiffPresent bool
	if MeasAndMobParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Phy_ParametersXDD_DiffPresent {
		ie.Phy_ParametersXDD_Diff = new(Phy_ParametersXDD_Diff)
		if err = ie.Phy_ParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_ParametersXDD_Diff", err)
		}
	}
	if Mac_ParametersXDD_DiffPresent {
		ie.Mac_ParametersXDD_Diff = new(MAC_ParametersXDD_Diff)
		if err = ie.Mac_ParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_ParametersXDD_Diff", err)
		}
	}
	if MeasAndMobParametersXDD_DiffPresent {
		ie.MeasAndMobParametersXDD_Diff = new(MeasAndMobParametersXDD_Diff)
		if err = ie.MeasAndMobParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersXDD_Diff", err)
		}
	}
	return nil
}
