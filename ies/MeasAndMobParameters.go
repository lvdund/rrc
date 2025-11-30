package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParameters struct {
	MeasAndMobParametersCommon   *MeasAndMobParametersCommon   `optional`
	MeasAndMobParametersXDD_Diff *MeasAndMobParametersXDD_Diff `optional`
	MeasAndMobParametersFRX_Diff *MeasAndMobParametersFRX_Diff `optional`
}

func (ie *MeasAndMobParameters) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasAndMobParametersCommon != nil, ie.MeasAndMobParametersXDD_Diff != nil, ie.MeasAndMobParametersFRX_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasAndMobParametersCommon != nil {
		if err = ie.MeasAndMobParametersCommon.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersCommon", err)
		}
	}
	if ie.MeasAndMobParametersXDD_Diff != nil {
		if err = ie.MeasAndMobParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersXDD_Diff", err)
		}
	}
	if ie.MeasAndMobParametersFRX_Diff != nil {
		if err = ie.MeasAndMobParametersFRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersFRX_Diff", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParameters) Decode(r *aper.AperReader) error {
	var err error
	var MeasAndMobParametersCommonPresent bool
	if MeasAndMobParametersCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasAndMobParametersXDD_DiffPresent bool
	if MeasAndMobParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasAndMobParametersFRX_DiffPresent bool
	if MeasAndMobParametersFRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasAndMobParametersCommonPresent {
		ie.MeasAndMobParametersCommon = new(MeasAndMobParametersCommon)
		if err = ie.MeasAndMobParametersCommon.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersCommon", err)
		}
	}
	if MeasAndMobParametersXDD_DiffPresent {
		ie.MeasAndMobParametersXDD_Diff = new(MeasAndMobParametersXDD_Diff)
		if err = ie.MeasAndMobParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersXDD_Diff", err)
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
