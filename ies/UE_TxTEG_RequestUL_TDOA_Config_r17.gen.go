// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UE-TxTEG-RequestUL-TDOA-Config-r17 (line 1122).

var uETxTEGRequestULTDOAConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "oneShot-r17"},
		{Name: "periodicReporting-r17"},
	},
}

const (
	UE_TxTEG_RequestUL_TDOA_Config_r17_OneShot_r17           = 0
	UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17 = 1
)

const (
	UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms160    = 0
	UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms320    = 1
	UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms1280   = 2
	UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms2560   = 3
	UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms61440  = 4
	UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms81920  = 5
	UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms368640 = 6
	UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms737280 = 7
)

var uETxTEGRequestULTDOAConfigR17PeriodicReportingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms160, UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms320, UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms1280, UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms2560, UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms61440, UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms81920, UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms368640, UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17_Ms737280},
}

type UE_TxTEG_RequestUL_TDOA_Config_r17 struct {
	Choice                int
	OneShot_r17           *struct{}
	PeriodicReporting_r17 *int64
}

func (ie *UE_TxTEG_RequestUL_TDOA_Config_r17) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(uETxTEGRequestULTDOAConfigR17Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case UE_TxTEG_RequestUL_TDOA_Config_r17_OneShot_r17:
		if err := e.EncodeNull(); err != nil {
			return err
		}
	case UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17:
		if err := e.EncodeEnumerated((*ie.PeriodicReporting_r17), uETxTEGRequestULTDOAConfigR17PeriodicReportingR17Constraints); err != nil {
			return err
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "UE-TxTEG-RequestUL-TDOA-Config-r17",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *UE_TxTEG_RequestUL_TDOA_Config_r17) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(uETxTEGRequestULTDOAConfigR17Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "UE-TxTEG-RequestUL-TDOA-Config-r17", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case UE_TxTEG_RequestUL_TDOA_Config_r17_OneShot_r17:
		ie.OneShot_r17 = &struct{}{}
		if err := d.DecodeNull(); err != nil {
			return err
		}
	case UE_TxTEG_RequestUL_TDOA_Config_r17_PeriodicReporting_r17:
		ie.PeriodicReporting_r17 = new(int64)
		v, err := d.DecodeEnumerated(uETxTEGRequestULTDOAConfigR17PeriodicReportingR17Constraints)
		if err != nil {
			return err
		}
		(*ie.PeriodicReporting_r17) = v
	default:
		return &per.DecodeError{TypeName: "UE-TxTEG-RequestUL-TDOA-Config-r17", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
