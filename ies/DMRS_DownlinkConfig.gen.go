// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DMRS-DownlinkConfig (line 7762).

var dMRSDownlinkConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dmrs-Type", Optional: true},
		{Name: "dmrs-AdditionalPosition", Optional: true},
		{Name: "maxLength", Optional: true},
		{Name: "scramblingID0", Optional: true},
		{Name: "scramblingID1", Optional: true},
		{Name: "phaseTrackingRS", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	DMRS_DownlinkConfig_Dmrs_Type_Type2 = 0
)

var dMRSDownlinkConfigDmrsTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_DownlinkConfig_Dmrs_Type_Type2},
}

const (
	DMRS_DownlinkConfig_Dmrs_AdditionalPosition_Pos0 = 0
	DMRS_DownlinkConfig_Dmrs_AdditionalPosition_Pos1 = 1
	DMRS_DownlinkConfig_Dmrs_AdditionalPosition_Pos3 = 2
)

var dMRSDownlinkConfigDmrsAdditionalPositionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_DownlinkConfig_Dmrs_AdditionalPosition_Pos0, DMRS_DownlinkConfig_Dmrs_AdditionalPosition_Pos1, DMRS_DownlinkConfig_Dmrs_AdditionalPosition_Pos3},
}

const (
	DMRS_DownlinkConfig_MaxLength_Len2 = 0
)

var dMRSDownlinkConfigMaxLengthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_DownlinkConfig_MaxLength_Len2},
}

var dMRSDownlinkConfigScramblingID0Constraints = per.Constrained(0, 65535)

var dMRSDownlinkConfigScramblingID1Constraints = per.Constrained(0, 65535)

var dMRS_DownlinkConfigPhaseTrackingRSConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	DMRS_DownlinkConfig_PhaseTrackingRS_Release = 0
	DMRS_DownlinkConfig_PhaseTrackingRS_Setup   = 1
)

const (
	DMRS_DownlinkConfig_Ext_Dmrs_Downlink_r16_Enabled = 0
)

var dMRSDownlinkConfigExtDmrsDownlinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_DownlinkConfig_Ext_Dmrs_Downlink_r16_Enabled},
}

const (
	DMRS_DownlinkConfig_Ext_Dmrs_TypeEnh_r18_Enabled = 0
)

var dMRSDownlinkConfigExtDmrsTypeEnhR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_DownlinkConfig_Ext_Dmrs_TypeEnh_r18_Enabled},
}

type DMRS_DownlinkConfig struct {
	Dmrs_Type               *int64
	Dmrs_AdditionalPosition *int64
	MaxLength               *int64
	ScramblingID0           *int64
	ScramblingID1           *int64
	PhaseTrackingRS         *struct {
		Choice  int
		Release *struct{}
		Setup   *PTRS_DownlinkConfig
	}
	Dmrs_Downlink_r16 *int64
	Dmrs_TypeEnh_r18  *int64
}

func (ie *DMRS_DownlinkConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dMRSDownlinkConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Dmrs_Downlink_r16 != nil
	hasExtGroup1 := ie.Dmrs_TypeEnh_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dmrs_Type != nil, ie.Dmrs_AdditionalPosition != nil, ie.MaxLength != nil, ie.ScramblingID0 != nil, ie.ScramblingID1 != nil, ie.PhaseTrackingRS != nil}); err != nil {
		return err
	}

	// 3. dmrs-Type: enumerated
	{
		if ie.Dmrs_Type != nil {
			if err := e.EncodeEnumerated(*ie.Dmrs_Type, dMRSDownlinkConfigDmrsTypeConstraints); err != nil {
				return err
			}
		}
	}

	// 4. dmrs-AdditionalPosition: enumerated
	{
		if ie.Dmrs_AdditionalPosition != nil {
			if err := e.EncodeEnumerated(*ie.Dmrs_AdditionalPosition, dMRSDownlinkConfigDmrsAdditionalPositionConstraints); err != nil {
				return err
			}
		}
	}

	// 5. maxLength: enumerated
	{
		if ie.MaxLength != nil {
			if err := e.EncodeEnumerated(*ie.MaxLength, dMRSDownlinkConfigMaxLengthConstraints); err != nil {
				return err
			}
		}
	}

	// 6. scramblingID0: integer
	{
		if ie.ScramblingID0 != nil {
			if err := e.EncodeInteger(*ie.ScramblingID0, dMRSDownlinkConfigScramblingID0Constraints); err != nil {
				return err
			}
		}
	}

	// 7. scramblingID1: integer
	{
		if ie.ScramblingID1 != nil {
			if err := e.EncodeInteger(*ie.ScramblingID1, dMRSDownlinkConfigScramblingID1Constraints); err != nil {
				return err
			}
		}
	}

	// 8. phaseTrackingRS: choice
	{
		if ie.PhaseTrackingRS != nil {
			choiceEnc := e.NewChoiceEncoder(dMRS_DownlinkConfigPhaseTrackingRSConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.PhaseTrackingRS).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.PhaseTrackingRS).Choice {
			case DMRS_DownlinkConfig_PhaseTrackingRS_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case DMRS_DownlinkConfig_PhaseTrackingRS_Setup:
				if err := (*ie.PhaseTrackingRS).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.PhaseTrackingRS).Choice), Constraint: "index out of range"}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "dmrs-Downlink-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dmrs_Downlink_r16 != nil}); err != nil {
				return err
			}

			if ie.Dmrs_Downlink_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Dmrs_Downlink_r16, dMRSDownlinkConfigExtDmrsDownlinkR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "dmrs-TypeEnh-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dmrs_TypeEnh_r18 != nil}); err != nil {
				return err
			}

			if ie.Dmrs_TypeEnh_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Dmrs_TypeEnh_r18, dMRSDownlinkConfigExtDmrsTypeEnhR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DMRS_DownlinkConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dMRSDownlinkConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dmrs-Type: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(dMRSDownlinkConfigDmrsTypeConstraints)
			if err != nil {
				return err
			}
			ie.Dmrs_Type = &idx
		}
	}

	// 4. dmrs-AdditionalPosition: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(dMRSDownlinkConfigDmrsAdditionalPositionConstraints)
			if err != nil {
				return err
			}
			ie.Dmrs_AdditionalPosition = &idx
		}
	}

	// 5. maxLength: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(dMRSDownlinkConfigMaxLengthConstraints)
			if err != nil {
				return err
			}
			ie.MaxLength = &idx
		}
	}

	// 6. scramblingID0: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(dMRSDownlinkConfigScramblingID0Constraints)
			if err != nil {
				return err
			}
			ie.ScramblingID0 = &v
		}
	}

	// 7. scramblingID1: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(dMRSDownlinkConfigScramblingID1Constraints)
			if err != nil {
				return err
			}
			ie.ScramblingID1 = &v
		}
	}

	// 8. phaseTrackingRS: choice
	{
		if seq.IsComponentPresent(5) {
			ie.PhaseTrackingRS = &struct {
				Choice  int
				Release *struct{}
				Setup   *PTRS_DownlinkConfig
			}{}
			choiceDec := d.NewChoiceDecoder(dMRS_DownlinkConfigPhaseTrackingRSConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PhaseTrackingRS).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case DMRS_DownlinkConfig_PhaseTrackingRS_Release:
				(*ie.PhaseTrackingRS).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case DMRS_DownlinkConfig_PhaseTrackingRS_Setup:
				(*ie.PhaseTrackingRS).Setup = new(PTRS_DownlinkConfig)
				if err := (*ie.PhaseTrackingRS).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "dmrs-Downlink-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(dMRSDownlinkConfigExtDmrsDownlinkR16Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_Downlink_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "dmrs-TypeEnh-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(dMRSDownlinkConfigExtDmrsTypeEnhR18Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_TypeEnh_r18 = &v
		}
	}

	return nil
}
