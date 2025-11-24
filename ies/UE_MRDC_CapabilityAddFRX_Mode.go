package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_CapabilityAddFRX_Mode struct {
	MeasAndMobParametersMRDC_FRX_Diff MeasAndMobParametersMRDC_FRX_Diff `madatory`
}

func (ie *UE_MRDC_CapabilityAddFRX_Mode) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.MeasAndMobParametersMRDC_FRX_Diff.Encode(w); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersMRDC_FRX_Diff", err)
	}
	return nil
}

func (ie *UE_MRDC_CapabilityAddFRX_Mode) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.MeasAndMobParametersMRDC_FRX_Diff.Decode(r); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersMRDC_FRX_Diff", err)
	}
	return nil
}
