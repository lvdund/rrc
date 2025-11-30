package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_v1730 struct {
	MeasAndMobParametersMRDC_Common_v1730 *MeasAndMobParametersMRDC_Common_v1730 `optional`
}

func (ie *MeasAndMobParametersMRDC_v1730) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasAndMobParametersMRDC_Common_v1730 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasAndMobParametersMRDC_Common_v1730 != nil {
		if err = ie.MeasAndMobParametersMRDC_Common_v1730.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersMRDC_Common_v1730", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_v1730) Decode(r *aper.AperReader) error {
	var err error
	var MeasAndMobParametersMRDC_Common_v1730Present bool
	if MeasAndMobParametersMRDC_Common_v1730Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasAndMobParametersMRDC_Common_v1730Present {
		ie.MeasAndMobParametersMRDC_Common_v1730 = new(MeasAndMobParametersMRDC_Common_v1730)
		if err = ie.MeasAndMobParametersMRDC_Common_v1730.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersMRDC_Common_v1730", err)
		}
	}
	return nil
}
