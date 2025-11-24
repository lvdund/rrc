package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type QuantityConfigUTRA_FDD_r16 struct {
	FilterCoefficientRSCP_r16 FilterCoefficient `madatory`
	FilterCoefficientEcNO_r16 FilterCoefficient `madatory`
}

func (ie *QuantityConfigUTRA_FDD_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.FilterCoefficientRSCP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode FilterCoefficientRSCP_r16", err)
	}
	if err = ie.FilterCoefficientEcNO_r16.Encode(w); err != nil {
		return utils.WrapError("Encode FilterCoefficientEcNO_r16", err)
	}
	return nil
}

func (ie *QuantityConfigUTRA_FDD_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.FilterCoefficientRSCP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode FilterCoefficientRSCP_r16", err)
	}
	if err = ie.FilterCoefficientEcNO_r16.Decode(r); err != nil {
		return utils.WrapError("Decode FilterCoefficientEcNO_r16", err)
	}
	return nil
}
