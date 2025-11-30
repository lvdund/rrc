package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_Capability_v1700 struct {
	MeasAndMobParametersMRDC_v1700 MeasAndMobParametersMRDC_v1700 `madatory`
	NonCriticalExtension           *UE_MRDC_Capability_v1730      `optional`
}

func (ie *UE_MRDC_Capability_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MeasAndMobParametersMRDC_v1700.Encode(w); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersMRDC_v1700", err)
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_MRDC_Capability_v1700) Decode(r *aper.AperReader) error {
	var err error
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MeasAndMobParametersMRDC_v1700.Decode(r); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersMRDC_v1700", err)
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_MRDC_Capability_v1730)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
