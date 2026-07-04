// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RACH-LessHO-r18 (line 5807).

var rACHLessHOR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "targetNTA-r18", Optional: true},
		{Name: "beamIndication-r18", Optional: true},
	},
}

const (
	RACH_LessHO_r18_TargetNTA_r18_Zero   = 0
	RACH_LessHO_r18_TargetNTA_r18_Source = 1
)

var rACHLessHOR18TargetNTAR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_LessHO_r18_TargetNTA_r18_Zero, RACH_LessHO_r18_TargetNTA_r18_Source},
}

var rACH_LessHO_r18BeamIndicationR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "tci-StateID-r18"},
		{Name: "ssb-Index-r18"},
	},
}

const (
	RACH_LessHO_r18_BeamIndication_r18_Tci_StateID_r18 = 0
	RACH_LessHO_r18_BeamIndication_r18_Ssb_Index_r18   = 1
)

type RACH_LessHO_r18 struct {
	TargetNTA_r18      *int64
	BeamIndication_r18 *struct {
		Choice          int
		Tci_StateID_r18 *TCI_StateId
		Ssb_Index_r18   *SSB_Index
	}
}

func (ie *RACH_LessHO_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rACHLessHOR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TargetNTA_r18 != nil, ie.BeamIndication_r18 != nil}); err != nil {
		return err
	}

	// 3. targetNTA-r18: enumerated
	{
		if ie.TargetNTA_r18 != nil {
			if err := e.EncodeEnumerated(*ie.TargetNTA_r18, rACHLessHOR18TargetNTAR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. beamIndication-r18: choice
	{
		if ie.BeamIndication_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(rACH_LessHO_r18BeamIndicationR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.BeamIndication_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.BeamIndication_r18).Choice {
			case RACH_LessHO_r18_BeamIndication_r18_Tci_StateID_r18:
				if err := (*ie.BeamIndication_r18).Tci_StateID_r18.Encode(e); err != nil {
					return err
				}
			case RACH_LessHO_r18_BeamIndication_r18_Ssb_Index_r18:
				if err := (*ie.BeamIndication_r18).Ssb_Index_r18.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.BeamIndication_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *RACH_LessHO_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rACHLessHOR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. targetNTA-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rACHLessHOR18TargetNTAR18Constraints)
			if err != nil {
				return err
			}
			ie.TargetNTA_r18 = &idx
		}
	}

	// 4. beamIndication-r18: choice
	{
		if seq.IsComponentPresent(1) {
			ie.BeamIndication_r18 = &struct {
				Choice          int
				Tci_StateID_r18 *TCI_StateId
				Ssb_Index_r18   *SSB_Index
			}{}
			choiceDec := d.NewChoiceDecoder(rACH_LessHO_r18BeamIndicationR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.BeamIndication_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RACH_LessHO_r18_BeamIndication_r18_Tci_StateID_r18:
				(*ie.BeamIndication_r18).Tci_StateID_r18 = new(TCI_StateId)
				if err := (*ie.BeamIndication_r18).Tci_StateID_r18.Decode(d); err != nil {
					return err
				}
			case RACH_LessHO_r18_BeamIndication_r18_Ssb_Index_r18:
				(*ie.BeamIndication_r18).Ssb_Index_r18 = new(SSB_Index)
				if err := (*ie.BeamIndication_r18).Ssb_Index_r18.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
