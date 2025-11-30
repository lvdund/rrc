package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_XDD_Diff_v1560 struct {
	Sftd_MeasPSCell_NEDC *MeasAndMobParametersMRDC_XDD_Diff_v1560_sftd_MeasPSCell_NEDC `optional`
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff_v1560) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sftd_MeasPSCell_NEDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sftd_MeasPSCell_NEDC != nil {
		if err = ie.Sftd_MeasPSCell_NEDC.Encode(w); err != nil {
			return utils.WrapError("Encode Sftd_MeasPSCell_NEDC", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff_v1560) Decode(r *aper.AperReader) error {
	var err error
	var Sftd_MeasPSCell_NEDCPresent bool
	if Sftd_MeasPSCell_NEDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Sftd_MeasPSCell_NEDCPresent {
		ie.Sftd_MeasPSCell_NEDC = new(MeasAndMobParametersMRDC_XDD_Diff_v1560_sftd_MeasPSCell_NEDC)
		if err = ie.Sftd_MeasPSCell_NEDC.Decode(r); err != nil {
			return utils.WrapError("Decode Sftd_MeasPSCell_NEDC", err)
		}
	}
	return nil
}
