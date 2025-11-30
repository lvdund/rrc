package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16 struct {
	CondPSCellChangeFDD_TDD_r16 *MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16_condPSCellChangeFDD_TDD_r16 `optional`
	CondPSCellChangeFR1_FR2_r16 *MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16_condPSCellChangeFR1_FR2_r16 `optional`
}

func (ie *MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.CondPSCellChangeFDD_TDD_r16 != nil, ie.CondPSCellChangeFR1_FR2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CondPSCellChangeFDD_TDD_r16 != nil {
		if err = ie.CondPSCellChangeFDD_TDD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CondPSCellChangeFDD_TDD_r16", err)
		}
	}
	if ie.CondPSCellChangeFR1_FR2_r16 != nil {
		if err = ie.CondPSCellChangeFR1_FR2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CondPSCellChangeFR1_FR2_r16", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16) Decode(r *aper.AperReader) error {
	var err error
	var CondPSCellChangeFDD_TDD_r16Present bool
	if CondPSCellChangeFDD_TDD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CondPSCellChangeFR1_FR2_r16Present bool
	if CondPSCellChangeFR1_FR2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if CondPSCellChangeFDD_TDD_r16Present {
		ie.CondPSCellChangeFDD_TDD_r16 = new(MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16_condPSCellChangeFDD_TDD_r16)
		if err = ie.CondPSCellChangeFDD_TDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CondPSCellChangeFDD_TDD_r16", err)
		}
	}
	if CondPSCellChangeFR1_FR2_r16Present {
		ie.CondPSCellChangeFR1_FR2_r16 = new(MeasAndMobParametersMRDC_Common_v1610_condPSCellChangeParametersCommon_r16_condPSCellChangeFR1_FR2_r16)
		if err = ie.CondPSCellChangeFR1_FR2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CondPSCellChangeFR1_FR2_r16", err)
		}
	}
	return nil
}
