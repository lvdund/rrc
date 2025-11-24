package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FilterConfig struct {
	FilterCoefficientRSRP    FilterCoefficient `madatory`
	FilterCoefficientRSRQ    FilterCoefficient `madatory`
	FilterCoefficientRS_SINR FilterCoefficient `madatory`
}

func (ie *FilterConfig) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.FilterCoefficientRSRP.Encode(w); err != nil {
		return utils.WrapError("Encode FilterCoefficientRSRP", err)
	}
	if err = ie.FilterCoefficientRSRQ.Encode(w); err != nil {
		return utils.WrapError("Encode FilterCoefficientRSRQ", err)
	}
	if err = ie.FilterCoefficientRS_SINR.Encode(w); err != nil {
		return utils.WrapError("Encode FilterCoefficientRS_SINR", err)
	}
	return nil
}

func (ie *FilterConfig) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.FilterCoefficientRSRP.Decode(r); err != nil {
		return utils.WrapError("Decode FilterCoefficientRSRP", err)
	}
	if err = ie.FilterCoefficientRSRQ.Decode(r); err != nil {
		return utils.WrapError("Decode FilterCoefficientRSRQ", err)
	}
	if err = ie.FilterCoefficientRS_SINR.Decode(r); err != nil {
		return utils.WrapError("Decode FilterCoefficientRS_SINR", err)
	}
	return nil
}
