package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_Parameters struct {
	SupportedBandListEUTRA   []FreqBandIndicatorEUTRA  `lb:1,ub:maxBandsEUTRA,madatory`
	Eutra_ParametersCommon   *EUTRA_ParametersCommon   `optional`
	Eutra_ParametersXDD_Diff *EUTRA_ParametersXDD_Diff `optional`
}

func (ie *EUTRA_Parameters) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Eutra_ParametersCommon != nil, ie.Eutra_ParametersXDD_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_SupportedBandListEUTRA := utils.NewSequence[*FreqBandIndicatorEUTRA]([]*FreqBandIndicatorEUTRA{}, uper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
	for _, i := range ie.SupportedBandListEUTRA {
		tmp_SupportedBandListEUTRA.Value = append(tmp_SupportedBandListEUTRA.Value, &i)
	}
	if err = tmp_SupportedBandListEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedBandListEUTRA", err)
	}
	if ie.Eutra_ParametersCommon != nil {
		if err = ie.Eutra_ParametersCommon.Encode(w); err != nil {
			return utils.WrapError("Encode Eutra_ParametersCommon", err)
		}
	}
	if ie.Eutra_ParametersXDD_Diff != nil {
		if err = ie.Eutra_ParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode Eutra_ParametersXDD_Diff", err)
		}
	}
	return nil
}

func (ie *EUTRA_Parameters) Decode(r *uper.UperReader) error {
	var err error
	var Eutra_ParametersCommonPresent bool
	if Eutra_ParametersCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Eutra_ParametersXDD_DiffPresent bool
	if Eutra_ParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_SupportedBandListEUTRA := utils.NewSequence[*FreqBandIndicatorEUTRA]([]*FreqBandIndicatorEUTRA{}, uper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
	fn_SupportedBandListEUTRA := func() *FreqBandIndicatorEUTRA {
		return new(FreqBandIndicatorEUTRA)
	}
	if err = tmp_SupportedBandListEUTRA.Decode(r, fn_SupportedBandListEUTRA); err != nil {
		return utils.WrapError("Decode SupportedBandListEUTRA", err)
	}
	ie.SupportedBandListEUTRA = []FreqBandIndicatorEUTRA{}
	for _, i := range tmp_SupportedBandListEUTRA.Value {
		ie.SupportedBandListEUTRA = append(ie.SupportedBandListEUTRA, *i)
	}
	if Eutra_ParametersCommonPresent {
		ie.Eutra_ParametersCommon = new(EUTRA_ParametersCommon)
		if err = ie.Eutra_ParametersCommon.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra_ParametersCommon", err)
		}
	}
	if Eutra_ParametersXDD_DiffPresent {
		ie.Eutra_ParametersXDD_Diff = new(EUTRA_ParametersXDD_Diff)
		if err = ie.Eutra_ParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra_ParametersXDD_Diff", err)
		}
	}
	return nil
}
