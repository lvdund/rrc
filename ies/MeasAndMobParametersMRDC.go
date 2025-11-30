package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC struct {
	MeasAndMobParametersMRDC_Common   *MeasAndMobParametersMRDC_Common   `optional`
	MeasAndMobParametersMRDC_XDD_Diff *MeasAndMobParametersMRDC_XDD_Diff `optional`
	MeasAndMobParametersMRDC_FRX_Diff *MeasAndMobParametersMRDC_FRX_Diff `optional`
}

func (ie *MeasAndMobParametersMRDC) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasAndMobParametersMRDC_Common != nil, ie.MeasAndMobParametersMRDC_XDD_Diff != nil, ie.MeasAndMobParametersMRDC_FRX_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasAndMobParametersMRDC_Common != nil {
		if err = ie.MeasAndMobParametersMRDC_Common.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersMRDC_Common", err)
		}
	}
	if ie.MeasAndMobParametersMRDC_XDD_Diff != nil {
		if err = ie.MeasAndMobParametersMRDC_XDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersMRDC_XDD_Diff", err)
		}
	}
	if ie.MeasAndMobParametersMRDC_FRX_Diff != nil {
		if err = ie.MeasAndMobParametersMRDC_FRX_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersMRDC_FRX_Diff", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC) Decode(r *aper.AperReader) error {
	var err error
	var MeasAndMobParametersMRDC_CommonPresent bool
	if MeasAndMobParametersMRDC_CommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasAndMobParametersMRDC_XDD_DiffPresent bool
	if MeasAndMobParametersMRDC_XDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasAndMobParametersMRDC_FRX_DiffPresent bool
	if MeasAndMobParametersMRDC_FRX_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasAndMobParametersMRDC_CommonPresent {
		ie.MeasAndMobParametersMRDC_Common = new(MeasAndMobParametersMRDC_Common)
		if err = ie.MeasAndMobParametersMRDC_Common.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersMRDC_Common", err)
		}
	}
	if MeasAndMobParametersMRDC_XDD_DiffPresent {
		ie.MeasAndMobParametersMRDC_XDD_Diff = new(MeasAndMobParametersMRDC_XDD_Diff)
		if err = ie.MeasAndMobParametersMRDC_XDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersMRDC_XDD_Diff", err)
		}
	}
	if MeasAndMobParametersMRDC_FRX_DiffPresent {
		ie.MeasAndMobParametersMRDC_FRX_Diff = new(MeasAndMobParametersMRDC_FRX_Diff)
		if err = ie.MeasAndMobParametersMRDC_FRX_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersMRDC_FRX_Diff", err)
		}
	}
	return nil
}
