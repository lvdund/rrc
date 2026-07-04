// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: OD-SSB-Config-r19 (line 10898).

var oDSSBConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "od-SSB-ConfigId-r19"},
		{Name: "od-SSB-SFN-Offset-r19", Optional: true},
		{Name: "od-SSB-HalfFrameIndex-r19", Optional: true},
		{Name: "od-SSB-ActivationStatus-r19", Optional: true},
		{Name: "od-SSB-Periodicity-r19"},
		{Name: "od-SSB-PositionsInBurst-r19", Optional: true},
		{Name: "od-SSB-NrofBursts-r19", Optional: true},
	},
}

var oDSSBConfigR19OdSSBSFNOffsetR19Constraints = per.Constrained(0, 15)

const (
	OD_SSB_Config_r19_Od_SSB_HalfFrameIndex_r19_Zero = 0
	OD_SSB_Config_r19_Od_SSB_HalfFrameIndex_r19_One  = 1
)

var oDSSBConfigR19OdSSBHalfFrameIndexR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OD_SSB_Config_r19_Od_SSB_HalfFrameIndex_r19_Zero, OD_SSB_Config_r19_Od_SSB_HalfFrameIndex_r19_One},
}

const (
	OD_SSB_Config_r19_Od_SSB_ActivationStatus_r19_Activated = 0
)

var oDSSBConfigR19OdSSBActivationStatusR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OD_SSB_Config_r19_Od_SSB_ActivationStatus_r19_Activated},
}

const (
	OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Ms5    = 0
	OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Ms10   = 1
	OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Ms20   = 2
	OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Ms40   = 3
	OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Ms80   = 4
	OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Ms160  = 5
	OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Spare2 = 6
	OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Spare1 = 7
)

var oDSSBConfigR19OdSSBPeriodicityR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Ms5, OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Ms10, OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Ms20, OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Ms40, OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Ms80, OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Ms160, OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Spare2, OD_SSB_Config_r19_Od_SSB_Periodicity_r19_Spare1},
}

var oD_SSB_Config_r19OdSSBPositionsInBurstR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "shortBitmap"},
		{Name: "mediumBitmap"},
		{Name: "longBitmap"},
	},
}

const (
	OD_SSB_Config_r19_Od_SSB_PositionsInBurst_r19_ShortBitmap  = 0
	OD_SSB_Config_r19_Od_SSB_PositionsInBurst_r19_MediumBitmap = 1
	OD_SSB_Config_r19_Od_SSB_PositionsInBurst_r19_LongBitmap   = 2
)

const (
	OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N5   = 0
	OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N10  = 1
	OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N15  = 2
	OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N20  = 3
	OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N25  = 4
	OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N30  = 5
	OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N40  = 6
	OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N50  = 7
	OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N75  = 8
	OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N100 = 9
	OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N150 = 10
	OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N200 = 11
)

var oDSSBConfigR19OdSSBNrofBurstsR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N5, OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N10, OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N15, OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N20, OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N25, OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N30, OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N40, OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N50, OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N75, OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N100, OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N150, OD_SSB_Config_r19_Od_SSB_NrofBursts_r19_N200},
}

type OD_SSB_Config_r19 struct {
	Od_SSB_ConfigId_r19         OD_SSB_ConfigId_r19
	Od_SSB_SFN_Offset_r19       *int64
	Od_SSB_HalfFrameIndex_r19   *int64
	Od_SSB_ActivationStatus_r19 *int64
	Od_SSB_Periodicity_r19      int64
	Od_SSB_PositionsInBurst_r19 *struct {
		Choice       int
		ShortBitmap  *per.BitString
		MediumBitmap *per.BitString
		LongBitmap   *per.BitString
	}
	Od_SSB_NrofBursts_r19 *int64
}

func (ie *OD_SSB_Config_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(oDSSBConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Od_SSB_SFN_Offset_r19 != nil, ie.Od_SSB_HalfFrameIndex_r19 != nil, ie.Od_SSB_ActivationStatus_r19 != nil, ie.Od_SSB_PositionsInBurst_r19 != nil, ie.Od_SSB_NrofBursts_r19 != nil}); err != nil {
		return err
	}

	// 3. od-SSB-ConfigId-r19: ref
	{
		if err := ie.Od_SSB_ConfigId_r19.Encode(e); err != nil {
			return err
		}
	}

	// 4. od-SSB-SFN-Offset-r19: integer
	{
		if ie.Od_SSB_SFN_Offset_r19 != nil {
			if err := e.EncodeInteger(*ie.Od_SSB_SFN_Offset_r19, oDSSBConfigR19OdSSBSFNOffsetR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. od-SSB-HalfFrameIndex-r19: enumerated
	{
		if ie.Od_SSB_HalfFrameIndex_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Od_SSB_HalfFrameIndex_r19, oDSSBConfigR19OdSSBHalfFrameIndexR19Constraints); err != nil {
				return err
			}
		}
	}

	// 6. od-SSB-ActivationStatus-r19: enumerated
	{
		if ie.Od_SSB_ActivationStatus_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Od_SSB_ActivationStatus_r19, oDSSBConfigR19OdSSBActivationStatusR19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. od-SSB-Periodicity-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.Od_SSB_Periodicity_r19, oDSSBConfigR19OdSSBPeriodicityR19Constraints); err != nil {
			return err
		}
	}

	// 8. od-SSB-PositionsInBurst-r19: choice
	{
		if ie.Od_SSB_PositionsInBurst_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(oD_SSB_Config_r19OdSSBPositionsInBurstR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Od_SSB_PositionsInBurst_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Od_SSB_PositionsInBurst_r19).Choice {
			case OD_SSB_Config_r19_Od_SSB_PositionsInBurst_r19_ShortBitmap:
				if err := e.EncodeBitString((*(*ie.Od_SSB_PositionsInBurst_r19).ShortBitmap), per.FixedSize(4)); err != nil {
					return err
				}
			case OD_SSB_Config_r19_Od_SSB_PositionsInBurst_r19_MediumBitmap:
				if err := e.EncodeBitString((*(*ie.Od_SSB_PositionsInBurst_r19).MediumBitmap), per.FixedSize(8)); err != nil {
					return err
				}
			case OD_SSB_Config_r19_Od_SSB_PositionsInBurst_r19_LongBitmap:
				if err := e.EncodeBitString((*(*ie.Od_SSB_PositionsInBurst_r19).LongBitmap), per.FixedSize(64)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Od_SSB_PositionsInBurst_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 9. od-SSB-NrofBursts-r19: enumerated
	{
		if ie.Od_SSB_NrofBursts_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Od_SSB_NrofBursts_r19, oDSSBConfigR19OdSSBNrofBurstsR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *OD_SSB_Config_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(oDSSBConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. od-SSB-ConfigId-r19: ref
	{
		if err := ie.Od_SSB_ConfigId_r19.Decode(d); err != nil {
			return err
		}
	}

	// 4. od-SSB-SFN-Offset-r19: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(oDSSBConfigR19OdSSBSFNOffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_SFN_Offset_r19 = &v
		}
	}

	// 5. od-SSB-HalfFrameIndex-r19: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(oDSSBConfigR19OdSSBHalfFrameIndexR19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_HalfFrameIndex_r19 = &idx
		}
	}

	// 6. od-SSB-ActivationStatus-r19: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(oDSSBConfigR19OdSSBActivationStatusR19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_ActivationStatus_r19 = &idx
		}
	}

	// 7. od-SSB-Periodicity-r19: enumerated
	{
		v4, err := d.DecodeEnumerated(oDSSBConfigR19OdSSBPeriodicityR19Constraints)
		if err != nil {
			return err
		}
		ie.Od_SSB_Periodicity_r19 = v4
	}

	// 8. od-SSB-PositionsInBurst-r19: choice
	{
		if seq.IsComponentPresent(5) {
			ie.Od_SSB_PositionsInBurst_r19 = &struct {
				Choice       int
				ShortBitmap  *per.BitString
				MediumBitmap *per.BitString
				LongBitmap   *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(oD_SSB_Config_r19OdSSBPositionsInBurstR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Od_SSB_PositionsInBurst_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case OD_SSB_Config_r19_Od_SSB_PositionsInBurst_r19_ShortBitmap:
				(*ie.Od_SSB_PositionsInBurst_r19).ShortBitmap = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(4))
				if err != nil {
					return err
				}
				(*(*ie.Od_SSB_PositionsInBurst_r19).ShortBitmap) = v
			case OD_SSB_Config_r19_Od_SSB_PositionsInBurst_r19_MediumBitmap:
				(*ie.Od_SSB_PositionsInBurst_r19).MediumBitmap = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*(*ie.Od_SSB_PositionsInBurst_r19).MediumBitmap) = v
			case OD_SSB_Config_r19_Od_SSB_PositionsInBurst_r19_LongBitmap:
				(*ie.Od_SSB_PositionsInBurst_r19).LongBitmap = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(64))
				if err != nil {
					return err
				}
				(*(*ie.Od_SSB_PositionsInBurst_r19).LongBitmap) = v
			}
		}
	}

	// 9. od-SSB-NrofBursts-r19: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(oDSSBConfigR19OdSSBNrofBurstsR19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_NrofBursts_r19 = &idx
		}
	}

	return nil
}
