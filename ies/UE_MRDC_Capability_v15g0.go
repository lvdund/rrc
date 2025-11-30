package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_Capability_v15g0 struct {
	Rf_ParametersMRDC_v15g0 *RF_ParametersMRDC_v15g0 `optional`
	NonCriticalExtension    interface{}              `optional`
}

func (ie *UE_MRDC_Capability_v15g0) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Rf_ParametersMRDC_v15g0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rf_ParametersMRDC_v15g0 != nil {
		if err = ie.Rf_ParametersMRDC_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode Rf_ParametersMRDC_v15g0", err)
		}
	}
	return nil
}

func (ie *UE_MRDC_Capability_v15g0) Decode(r *aper.AperReader) error {
	var err error
	var Rf_ParametersMRDC_v15g0Present bool
	if Rf_ParametersMRDC_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Rf_ParametersMRDC_v15g0Present {
		ie.Rf_ParametersMRDC_v15g0 = new(RF_ParametersMRDC_v15g0)
		if err = ie.Rf_ParametersMRDC_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode Rf_ParametersMRDC_v15g0", err)
		}
	}
	return nil
}
