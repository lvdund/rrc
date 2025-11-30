package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_v1700 struct {
	MeasAndMobParametersMRDC_Common_v1700 *MeasAndMobParametersMRDC_Common_v1700 `optional`
}

func (ie *MeasAndMobParametersMRDC_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasAndMobParametersMRDC_Common_v1700 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasAndMobParametersMRDC_Common_v1700 != nil {
		if err = ie.MeasAndMobParametersMRDC_Common_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersMRDC_Common_v1700", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_v1700) Decode(r *aper.AperReader) error {
	var err error
	var MeasAndMobParametersMRDC_Common_v1700Present bool
	if MeasAndMobParametersMRDC_Common_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasAndMobParametersMRDC_Common_v1700Present {
		ie.MeasAndMobParametersMRDC_Common_v1700 = new(MeasAndMobParametersMRDC_Common_v1700)
		if err = ie.MeasAndMobParametersMRDC_Common_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersMRDC_Common_v1700", err)
		}
	}
	return nil
}
