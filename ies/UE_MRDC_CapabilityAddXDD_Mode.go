package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_CapabilityAddXDD_Mode struct {
	MeasAndMobParametersMRDC_XDD_Diff *MeasAndMobParametersMRDC_XDD_Diff `optional`
	GeneralParametersMRDC_XDD_Diff    *GeneralParametersMRDC_XDD_Diff    `optional`
}

func (ie *UE_MRDC_CapabilityAddXDD_Mode) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasAndMobParametersMRDC_XDD_Diff != nil, ie.GeneralParametersMRDC_XDD_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasAndMobParametersMRDC_XDD_Diff != nil {
		if err = ie.MeasAndMobParametersMRDC_XDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersMRDC_XDD_Diff", err)
		}
	}
	if ie.GeneralParametersMRDC_XDD_Diff != nil {
		if err = ie.GeneralParametersMRDC_XDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode GeneralParametersMRDC_XDD_Diff", err)
		}
	}
	return nil
}

func (ie *UE_MRDC_CapabilityAddXDD_Mode) Decode(r *aper.AperReader) error {
	var err error
	var MeasAndMobParametersMRDC_XDD_DiffPresent bool
	if MeasAndMobParametersMRDC_XDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var GeneralParametersMRDC_XDD_DiffPresent bool
	if GeneralParametersMRDC_XDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasAndMobParametersMRDC_XDD_DiffPresent {
		ie.MeasAndMobParametersMRDC_XDD_Diff = new(MeasAndMobParametersMRDC_XDD_Diff)
		if err = ie.MeasAndMobParametersMRDC_XDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersMRDC_XDD_Diff", err)
		}
	}
	if GeneralParametersMRDC_XDD_DiffPresent {
		ie.GeneralParametersMRDC_XDD_Diff = new(GeneralParametersMRDC_XDD_Diff)
		if err = ie.GeneralParametersMRDC_XDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode GeneralParametersMRDC_XDD_Diff", err)
		}
	}
	return nil
}
