// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DMRS-UplinkConfig (line 7781).

var dMRSUplinkConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dmrs-Type", Optional: true},
		{Name: "dmrs-AdditionalPosition", Optional: true},
		{Name: "phaseTrackingRS", Optional: true},
		{Name: "maxLength", Optional: true},
		{Name: "transformPrecodingDisabled", Optional: true},
		{Name: "transformPrecodingEnabled", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	DMRS_UplinkConfig_Dmrs_Type_Type2 = 0
)

var dMRSUplinkConfigDmrsTypeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_UplinkConfig_Dmrs_Type_Type2},
}

const (
	DMRS_UplinkConfig_Dmrs_AdditionalPosition_Pos0 = 0
	DMRS_UplinkConfig_Dmrs_AdditionalPosition_Pos1 = 1
	DMRS_UplinkConfig_Dmrs_AdditionalPosition_Pos3 = 2
)

var dMRSUplinkConfigDmrsAdditionalPositionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_UplinkConfig_Dmrs_AdditionalPosition_Pos0, DMRS_UplinkConfig_Dmrs_AdditionalPosition_Pos1, DMRS_UplinkConfig_Dmrs_AdditionalPosition_Pos3},
}

var dMRS_UplinkConfigPhaseTrackingRSConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	DMRS_UplinkConfig_PhaseTrackingRS_Release = 0
	DMRS_UplinkConfig_PhaseTrackingRS_Setup   = 1
)

const (
	DMRS_UplinkConfig_MaxLength_Len2 = 0
)

var dMRSUplinkConfigMaxLengthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_UplinkConfig_MaxLength_Len2},
}

const (
	DMRS_UplinkConfig_Ext_Dmrs_TypeEnh_r18_Enabled = 0
)

var dMRSUplinkConfigExtDmrsTypeEnhR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_UplinkConfig_Ext_Dmrs_TypeEnh_r18_Enabled},
}

var dMRSUplinkConfigTransformPrecodingDisabledConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "scramblingID0", Optional: true},
		{Name: "scramblingID1", Optional: true},
	},
}

var dMRSUplinkConfigTransformPrecodingEnabledConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nPUSCH-Identity", Optional: true},
		{Name: "sequenceGroupHopping", Optional: true},
		{Name: "sequenceHopping", Optional: true},
	},
}

const (
	DMRS_UplinkConfig_TransformPrecodingEnabled_SequenceGroupHopping_Disabled = 0
)

var dMRSUplinkConfigTransformPrecodingEnabledSequenceGroupHoppingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_UplinkConfig_TransformPrecodingEnabled_SequenceGroupHopping_Disabled},
}

const (
	DMRS_UplinkConfig_TransformPrecodingEnabled_SequenceHopping_Enabled = 0
)

var dMRSUplinkConfigTransformPrecodingEnabledSequenceHoppingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{DMRS_UplinkConfig_TransformPrecodingEnabled_SequenceHopping_Enabled},
}

type DMRS_UplinkConfig struct {
	Dmrs_Type               *int64
	Dmrs_AdditionalPosition *int64
	PhaseTrackingRS         *struct {
		Choice  int
		Release *struct{}
		Setup   *PTRS_UplinkConfig
	}
	MaxLength                  *int64
	TransformPrecodingDisabled *struct {
		ScramblingID0 *int64
		ScramblingID1 *int64
	}
	TransformPrecodingEnabled *struct {
		NPUSCH_Identity      *int64
		SequenceGroupHopping *int64
		SequenceHopping      *int64
	}
	Dmrs_TypeEnh_r18 *int64
}

func (ie *DMRS_UplinkConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(dMRSUplinkConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Dmrs_TypeEnh_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dmrs_Type != nil, ie.Dmrs_AdditionalPosition != nil, ie.PhaseTrackingRS != nil, ie.MaxLength != nil, ie.TransformPrecodingDisabled != nil, ie.TransformPrecodingEnabled != nil}); err != nil {
		return err
	}

	// 3. dmrs-Type: enumerated
	{
		if ie.Dmrs_Type != nil {
			if err := e.EncodeEnumerated(*ie.Dmrs_Type, dMRSUplinkConfigDmrsTypeConstraints); err != nil {
				return err
			}
		}
	}

	// 4. dmrs-AdditionalPosition: enumerated
	{
		if ie.Dmrs_AdditionalPosition != nil {
			if err := e.EncodeEnumerated(*ie.Dmrs_AdditionalPosition, dMRSUplinkConfigDmrsAdditionalPositionConstraints); err != nil {
				return err
			}
		}
	}

	// 5. phaseTrackingRS: choice
	{
		if ie.PhaseTrackingRS != nil {
			choiceEnc := e.NewChoiceEncoder(dMRS_UplinkConfigPhaseTrackingRSConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.PhaseTrackingRS).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.PhaseTrackingRS).Choice {
			case DMRS_UplinkConfig_PhaseTrackingRS_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case DMRS_UplinkConfig_PhaseTrackingRS_Setup:
				if err := (*ie.PhaseTrackingRS).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.PhaseTrackingRS).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. maxLength: enumerated
	{
		if ie.MaxLength != nil {
			if err := e.EncodeEnumerated(*ie.MaxLength, dMRSUplinkConfigMaxLengthConstraints); err != nil {
				return err
			}
		}
	}

	// 7. transformPrecodingDisabled: sequence
	{
		if ie.TransformPrecodingDisabled != nil {
			c := ie.TransformPrecodingDisabled
			dMRSUplinkConfigTransformPrecodingDisabledSeq := e.NewSequenceEncoder(dMRSUplinkConfigTransformPrecodingDisabledConstraints)
			if err := dMRSUplinkConfigTransformPrecodingDisabledSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := dMRSUplinkConfigTransformPrecodingDisabledSeq.EncodePreamble([]bool{c.ScramblingID0 != nil, c.ScramblingID1 != nil}); err != nil {
				return err
			}
			if c.ScramblingID0 != nil {
				if err := e.EncodeInteger((*c.ScramblingID0), per.Constrained(0, 65535)); err != nil {
					return err
				}
			}
			if c.ScramblingID1 != nil {
				if err := e.EncodeInteger((*c.ScramblingID1), per.Constrained(0, 65535)); err != nil {
					return err
				}
			}
		}
	}

	// 8. transformPrecodingEnabled: sequence
	{
		if ie.TransformPrecodingEnabled != nil {
			c := ie.TransformPrecodingEnabled
			dMRSUplinkConfigTransformPrecodingEnabledSeq := e.NewSequenceEncoder(dMRSUplinkConfigTransformPrecodingEnabledConstraints)
			if err := dMRSUplinkConfigTransformPrecodingEnabledSeq.EncodeExtensionBit(false); err != nil {
				return err
			}
			if err := dMRSUplinkConfigTransformPrecodingEnabledSeq.EncodePreamble([]bool{c.NPUSCH_Identity != nil, c.SequenceGroupHopping != nil, c.SequenceHopping != nil}); err != nil {
				return err
			}
			if c.NPUSCH_Identity != nil {
				if err := e.EncodeInteger((*c.NPUSCH_Identity), per.Constrained(0, 1007)); err != nil {
					return err
				}
			}
			if c.SequenceGroupHopping != nil {
				if err := e.EncodeEnumerated((*c.SequenceGroupHopping), dMRSUplinkConfigTransformPrecodingEnabledSequenceGroupHoppingConstraints); err != nil {
					return err
				}
			}
			if c.SequenceHopping != nil {
				if err := e.EncodeEnumerated((*c.SequenceHopping), dMRSUplinkConfigTransformPrecodingEnabledSequenceHoppingConstraints); err != nil {
					return err
				}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
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
				if err := ex.EncodeEnumerated(*ie.Dmrs_TypeEnh_r18, dMRSUplinkConfigExtDmrsTypeEnhR18Constraints); err != nil {
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

func (ie *DMRS_UplinkConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(dMRSUplinkConfigConstraints)

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
			idx, err := d.DecodeEnumerated(dMRSUplinkConfigDmrsTypeConstraints)
			if err != nil {
				return err
			}
			ie.Dmrs_Type = &idx
		}
	}

	// 4. dmrs-AdditionalPosition: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(dMRSUplinkConfigDmrsAdditionalPositionConstraints)
			if err != nil {
				return err
			}
			ie.Dmrs_AdditionalPosition = &idx
		}
	}

	// 5. phaseTrackingRS: choice
	{
		if seq.IsComponentPresent(2) {
			ie.PhaseTrackingRS = &struct {
				Choice  int
				Release *struct{}
				Setup   *PTRS_UplinkConfig
			}{}
			choiceDec := d.NewChoiceDecoder(dMRS_UplinkConfigPhaseTrackingRSConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PhaseTrackingRS).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case DMRS_UplinkConfig_PhaseTrackingRS_Release:
				(*ie.PhaseTrackingRS).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case DMRS_UplinkConfig_PhaseTrackingRS_Setup:
				(*ie.PhaseTrackingRS).Setup = new(PTRS_UplinkConfig)
				if err := (*ie.PhaseTrackingRS).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. maxLength: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(dMRSUplinkConfigMaxLengthConstraints)
			if err != nil {
				return err
			}
			ie.MaxLength = &idx
		}
	}

	// 7. transformPrecodingDisabled: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.TransformPrecodingDisabled = &struct {
				ScramblingID0 *int64
				ScramblingID1 *int64
			}{}
			c := ie.TransformPrecodingDisabled
			dMRSUplinkConfigTransformPrecodingDisabledSeq := d.NewSequenceDecoder(dMRSUplinkConfigTransformPrecodingDisabledConstraints)
			if err := dMRSUplinkConfigTransformPrecodingDisabledSeq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := dMRSUplinkConfigTransformPrecodingDisabledSeq.DecodePreamble(); err != nil {
				return err
			}
			if dMRSUplinkConfigTransformPrecodingDisabledSeq.IsComponentPresent(0) {
				c.ScramblingID0 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 65535))
				if err != nil {
					return err
				}
				(*c.ScramblingID0) = v
			}
			if dMRSUplinkConfigTransformPrecodingDisabledSeq.IsComponentPresent(1) {
				c.ScramblingID1 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 65535))
				if err != nil {
					return err
				}
				(*c.ScramblingID1) = v
			}
		}
	}

	// 8. transformPrecodingEnabled: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.TransformPrecodingEnabled = &struct {
				NPUSCH_Identity      *int64
				SequenceGroupHopping *int64
				SequenceHopping      *int64
			}{}
			c := ie.TransformPrecodingEnabled
			dMRSUplinkConfigTransformPrecodingEnabledSeq := d.NewSequenceDecoder(dMRSUplinkConfigTransformPrecodingEnabledConstraints)
			if err := dMRSUplinkConfigTransformPrecodingEnabledSeq.DecodeExtensionBit(); err != nil {
				return err
			}
			if err := dMRSUplinkConfigTransformPrecodingEnabledSeq.DecodePreamble(); err != nil {
				return err
			}
			if dMRSUplinkConfigTransformPrecodingEnabledSeq.IsComponentPresent(0) {
				c.NPUSCH_Identity = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 1007))
				if err != nil {
					return err
				}
				(*c.NPUSCH_Identity) = v
			}
			if dMRSUplinkConfigTransformPrecodingEnabledSeq.IsComponentPresent(1) {
				c.SequenceGroupHopping = new(int64)
				v, err := d.DecodeEnumerated(dMRSUplinkConfigTransformPrecodingEnabledSequenceGroupHoppingConstraints)
				if err != nil {
					return err
				}
				(*c.SequenceGroupHopping) = v
			}
			if dMRSUplinkConfigTransformPrecodingEnabledSeq.IsComponentPresent(2) {
				c.SequenceHopping = new(int64)
				v, err := d.DecodeEnumerated(dMRSUplinkConfigTransformPrecodingEnabledSequenceHoppingConstraints)
				if err != nil {
					return err
				}
				(*c.SequenceHopping) = v
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
				{Name: "dmrs-TypeEnh-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(dMRSUplinkConfigExtDmrsTypeEnhR18Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_TypeEnh_r18 = &v
		}
	}

	return nil
}
