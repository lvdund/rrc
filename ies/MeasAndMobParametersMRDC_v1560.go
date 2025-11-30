package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_v1560 struct {
	MeasAndMobParametersMRDC_XDD_Diff_v1560 *MeasAndMobParametersMRDC_XDD_Diff_v1560 `optional`
}

func (ie *MeasAndMobParametersMRDC_v1560) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasAndMobParametersMRDC_XDD_Diff_v1560 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasAndMobParametersMRDC_XDD_Diff_v1560 != nil {
		if err = ie.MeasAndMobParametersMRDC_XDD_Diff_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersMRDC_XDD_Diff_v1560", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_v1560) Decode(r *aper.AperReader) error {
	var err error
	var MeasAndMobParametersMRDC_XDD_Diff_v1560Present bool
	if MeasAndMobParametersMRDC_XDD_Diff_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasAndMobParametersMRDC_XDD_Diff_v1560Present {
		ie.MeasAndMobParametersMRDC_XDD_Diff_v1560 = new(MeasAndMobParametersMRDC_XDD_Diff_v1560)
		if err = ie.MeasAndMobParametersMRDC_XDD_Diff_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersMRDC_XDD_Diff_v1560", err)
		}
	}
	return nil
}
