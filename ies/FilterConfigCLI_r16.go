package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FilterConfigCLI_r16 struct {
	FilterCoefficientSRS_RSRP_r16 FilterCoefficient `madatory`
	FilterCoefficientCLI_RSSI_r16 FilterCoefficient `madatory`
}

func (ie *FilterConfigCLI_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.FilterCoefficientSRS_RSRP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode FilterCoefficientSRS_RSRP_r16", err)
	}
	if err = ie.FilterCoefficientCLI_RSSI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode FilterCoefficientCLI_RSSI_r16", err)
	}
	return nil
}

func (ie *FilterConfigCLI_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.FilterCoefficientSRS_RSRP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode FilterCoefficientSRS_RSRP_r16", err)
	}
	if err = ie.FilterCoefficientCLI_RSSI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode FilterCoefficientCLI_RSSI_r16", err)
	}
	return nil
}
