package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CarrierTypePair_r16 struct {
	CarrierForCSI_Measurement_r16 PUCCH_Grp_CarrierTypes_r16 `madatory`
	CarrierForCSI_Reporting_r16   PUCCH_Grp_CarrierTypes_r16 `madatory`
}

func (ie *CarrierTypePair_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.CarrierForCSI_Measurement_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierForCSI_Measurement_r16", err)
	}
	if err = ie.CarrierForCSI_Reporting_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierForCSI_Reporting_r16", err)
	}
	return nil
}

func (ie *CarrierTypePair_r16) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.CarrierForCSI_Measurement_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierForCSI_Measurement_r16", err)
	}
	if err = ie.CarrierForCSI_Reporting_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierForCSI_Reporting_r16", err)
	}
	return nil
}
