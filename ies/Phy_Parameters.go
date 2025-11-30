package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Phy_Parameters struct {
	Phy_ParametersCommon   *Phy_ParametersCommon   `optional`
	Phy_ParametersXDD_Diff *Phy_ParametersXDD_Diff `optional`
	Phy_ParametersFRX_Diff *Phy_ParametersFRX_Diff `optional`
	Phy_ParametersFR1      *Phy_ParametersFR1      `optional`
	Phy_ParametersFR2      *Phy_ParametersFR2      `optional`
}

func (ie *Phy_Parameters) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Phy_ParametersCommon != nil, ie.Phy_ParametersXDD_Diff != nil, ie.Phy_ParametersFRX_Diff != nil, ie.Phy_ParametersFR1 != nil, ie.Phy_ParametersFR2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Phy_ParametersCommon != nil {
		if err = ie.Phy_ParametersCommon.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_ParametersCommon", err)
		}
	}
	if ie.Phy_ParametersXDD_Diff != nil {
		if err = ie.Phy_ParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_ParametersXDD_Diff", err)
		}
	}
	if ie.Phy_ParametersFRX_Diff != nil {
		if err = ie.Phy_ParametersFRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_ParametersFRX_Diff", err)
		}
	}
	if ie.Phy_ParametersFR1 != nil {
		if err = ie.Phy_ParametersFR1.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_ParametersFR1", err)
		}
	}
	if ie.Phy_ParametersFR2 != nil {
		if err = ie.Phy_ParametersFR2.Encode(w); err != nil {
			return utils.WrapError("Encode Phy_ParametersFR2", err)
		}
	}
	return nil
}

func (ie *Phy_Parameters) Decode(r *aper.AperReader) error {
	var err error
	var Phy_ParametersCommonPresent bool
	if Phy_ParametersCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Phy_ParametersXDD_DiffPresent bool
	if Phy_ParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Phy_ParametersFRX_DiffPresent bool
	if Phy_ParametersFRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Phy_ParametersFR1Present bool
	if Phy_ParametersFR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Phy_ParametersFR2Present bool
	if Phy_ParametersFR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Phy_ParametersCommonPresent {
		ie.Phy_ParametersCommon = new(Phy_ParametersCommon)
		if err = ie.Phy_ParametersCommon.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_ParametersCommon", err)
		}
	}
	if Phy_ParametersXDD_DiffPresent {
		ie.Phy_ParametersXDD_Diff = new(Phy_ParametersXDD_Diff)
		if err = ie.Phy_ParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_ParametersXDD_Diff", err)
		}
	}
	if Phy_ParametersFRX_DiffPresent {
		ie.Phy_ParametersFRX_Diff = new(Phy_ParametersFRX_Diff)
		if err = ie.Phy_ParametersFRX_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_ParametersFRX_Diff", err)
		}
	}
	if Phy_ParametersFR1Present {
		ie.Phy_ParametersFR1 = new(Phy_ParametersFR1)
		if err = ie.Phy_ParametersFR1.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_ParametersFR1", err)
		}
	}
	if Phy_ParametersFR2Present {
		ie.Phy_ParametersFR2 = new(Phy_ParametersFR2)
		if err = ie.Phy_ParametersFR2.Decode(r); err != nil {
			return utils.WrapError("Decode Phy_ParametersFR2", err)
		}
	}
	return nil
}
