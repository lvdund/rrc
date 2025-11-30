package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_Capability_v1730 struct {
	MeasAndMobParametersMRDC_v1730 *MeasAndMobParametersMRDC_v1730 `optional`
	NonCriticalExtension           interface{}                     `optional`
}

func (ie *UE_MRDC_Capability_v1730) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasAndMobParametersMRDC_v1730 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasAndMobParametersMRDC_v1730 != nil {
		if err = ie.MeasAndMobParametersMRDC_v1730.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersMRDC_v1730", err)
		}
	}
	return nil
}

func (ie *UE_MRDC_Capability_v1730) Decode(r *aper.AperReader) error {
	var err error
	var MeasAndMobParametersMRDC_v1730Present bool
	if MeasAndMobParametersMRDC_v1730Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasAndMobParametersMRDC_v1730Present {
		ie.MeasAndMobParametersMRDC_v1730 = new(MeasAndMobParametersMRDC_v1730)
		if err = ie.MeasAndMobParametersMRDC_v1730.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersMRDC_v1730", err)
		}
	}
	return nil
}
