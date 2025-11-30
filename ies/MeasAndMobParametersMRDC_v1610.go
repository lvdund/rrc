package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_v1610 struct {
	MeasAndMobParametersMRDC_Common_v1610 *MeasAndMobParametersMRDC_Common_v1610                    `optional`
	InterNR_MeasEUTRA_IAB_r16             *MeasAndMobParametersMRDC_v1610_interNR_MeasEUTRA_IAB_r16 `optional`
}

func (ie *MeasAndMobParametersMRDC_v1610) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasAndMobParametersMRDC_Common_v1610 != nil, ie.InterNR_MeasEUTRA_IAB_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasAndMobParametersMRDC_Common_v1610 != nil {
		if err = ie.MeasAndMobParametersMRDC_Common_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode MeasAndMobParametersMRDC_Common_v1610", err)
		}
	}
	if ie.InterNR_MeasEUTRA_IAB_r16 != nil {
		if err = ie.InterNR_MeasEUTRA_IAB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterNR_MeasEUTRA_IAB_r16", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_v1610) Decode(r *aper.AperReader) error {
	var err error
	var MeasAndMobParametersMRDC_Common_v1610Present bool
	if MeasAndMobParametersMRDC_Common_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var InterNR_MeasEUTRA_IAB_r16Present bool
	if InterNR_MeasEUTRA_IAB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasAndMobParametersMRDC_Common_v1610Present {
		ie.MeasAndMobParametersMRDC_Common_v1610 = new(MeasAndMobParametersMRDC_Common_v1610)
		if err = ie.MeasAndMobParametersMRDC_Common_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode MeasAndMobParametersMRDC_Common_v1610", err)
		}
	}
	if InterNR_MeasEUTRA_IAB_r16Present {
		ie.InterNR_MeasEUTRA_IAB_r16 = new(MeasAndMobParametersMRDC_v1610_interNR_MeasEUTRA_IAB_r16)
		if err = ie.InterNR_MeasEUTRA_IAB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterNR_MeasEUTRA_IAB_r16", err)
		}
	}
	return nil
}
