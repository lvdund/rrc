package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16 struct {
	MaxNumberPerCORESET_Pool_r16         MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16_maxNumberPerCORESET_Pool_r16         `madatory`
	MaxTotalNumberAcrossCORESET_Pool_r16 MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16_maxTotalNumberAcrossCORESET_Pool_r16 `madatory`
}

func (ie *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MaxNumberPerCORESET_Pool_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberPerCORESET_Pool_r16", err)
	}
	if err = ie.MaxTotalNumberAcrossCORESET_Pool_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxTotalNumberAcrossCORESET_Pool_r16", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MaxNumberPerCORESET_Pool_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberPerCORESET_Pool_r16", err)
	}
	if err = ie.MaxTotalNumberAcrossCORESET_Pool_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxTotalNumberAcrossCORESET_Pool_r16", err)
	}
	return nil
}
