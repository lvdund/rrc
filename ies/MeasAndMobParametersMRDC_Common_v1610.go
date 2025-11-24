package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_Common_v1610 struct {
	CondPSCellChangeParametersCommon_r16 *MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16 `optional`
	PscellT312_r16                       *MeasAndMobParametersMRDC_Common_v1610_pscellT312_r16                       `optional`
}

func (ie *MeasAndMobParametersMRDC_Common_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CondPSCellChangeParametersCommon_r16 != nil, ie.PscellT312_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CondPSCellChangeParametersCommon_r16 != nil {
		if err = ie.CondPSCellChangeParametersCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CondPSCellChangeParametersCommon_r16", err)
		}
	}
	if ie.PscellT312_r16 != nil {
		if err = ie.PscellT312_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PscellT312_r16", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1610) Decode(r *uper.UperReader) error {
	var err error
	var CondPSCellChangeParametersCommon_r16Present bool
	if CondPSCellChangeParametersCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PscellT312_r16Present bool
	if PscellT312_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if CondPSCellChangeParametersCommon_r16Present {
		ie.CondPSCellChangeParametersCommon_r16 = new(MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16)
		if err = ie.CondPSCellChangeParametersCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CondPSCellChangeParametersCommon_r16", err)
		}
	}
	if PscellT312_r16Present {
		ie.PscellT312_r16 = new(MeasAndMobParametersMRDC_Common_v1610_pscellT312_r16)
		if err = ie.PscellT312_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PscellT312_r16", err)
		}
	}
	return nil
}
