package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersCommon_condHandoverParametersCommon_r16 struct {
	CondHandoverFDD_TDD_r16 *MeasAndMobParametersCommon_condHandoverParametersCommon_r16_condHandoverFDD_TDD_r16 `optional`
	CondHandoverFR1_FR2_r16 *MeasAndMobParametersCommon_condHandoverParametersCommon_r16_condHandoverFR1_FR2_r16 `optional`
}

func (ie *MeasAndMobParametersCommon_condHandoverParametersCommon_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CondHandoverFDD_TDD_r16 != nil, ie.CondHandoverFR1_FR2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CondHandoverFDD_TDD_r16 != nil {
		if err = ie.CondHandoverFDD_TDD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CondHandoverFDD_TDD_r16", err)
		}
	}
	if ie.CondHandoverFR1_FR2_r16 != nil {
		if err = ie.CondHandoverFR1_FR2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CondHandoverFR1_FR2_r16", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersCommon_condHandoverParametersCommon_r16) Decode(r *uper.UperReader) error {
	var err error
	var CondHandoverFDD_TDD_r16Present bool
	if CondHandoverFDD_TDD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CondHandoverFR1_FR2_r16Present bool
	if CondHandoverFR1_FR2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if CondHandoverFDD_TDD_r16Present {
		ie.CondHandoverFDD_TDD_r16 = new(MeasAndMobParametersCommon_condHandoverParametersCommon_r16_condHandoverFDD_TDD_r16)
		if err = ie.CondHandoverFDD_TDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CondHandoverFDD_TDD_r16", err)
		}
	}
	if CondHandoverFR1_FR2_r16Present {
		ie.CondHandoverFR1_FR2_r16 = new(MeasAndMobParametersCommon_condHandoverParametersCommon_r16_condHandoverFR1_FR2_r16)
		if err = ie.CondHandoverFR1_FR2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CondHandoverFR1_FR2_r16", err)
		}
	}
	return nil
}
