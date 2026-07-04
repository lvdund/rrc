// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: EUTRA-ParametersCommon (line 20908).

var eUTRAParametersCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "mfbi-EUTRA", Optional: true},
		{Name: "modifiedMPR-BehaviorEUTRA", Optional: true},
		{Name: "multiNS-Pmax-EUTRA", Optional: true},
		{Name: "rs-SINR-MeasEUTRA", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	EUTRA_ParametersCommon_Mfbi_EUTRA_Supported = 0
)

var eUTRAParametersCommonMfbiEUTRAConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EUTRA_ParametersCommon_Mfbi_EUTRA_Supported},
}

var eUTRAParametersCommonModifiedMPRBehaviorEUTRAConstraints = per.FixedSize(32)

const (
	EUTRA_ParametersCommon_MultiNS_Pmax_EUTRA_Supported = 0
)

var eUTRAParametersCommonMultiNSPmaxEUTRAConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EUTRA_ParametersCommon_MultiNS_Pmax_EUTRA_Supported},
}

const (
	EUTRA_ParametersCommon_Rs_SINR_MeasEUTRA_Supported = 0
)

var eUTRAParametersCommonRsSINRMeasEUTRAConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EUTRA_ParametersCommon_Rs_SINR_MeasEUTRA_Supported},
}

const (
	EUTRA_ParametersCommon_Ext_Ne_DC_Supported = 0
)

var eUTRAParametersCommonExtNeDCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EUTRA_ParametersCommon_Ext_Ne_DC_Supported},
}

const (
	EUTRA_ParametersCommon_Ext_Nr_HO_ToEN_DC_r16_Supported = 0
)

var eUTRAParametersCommonExtNrHOToENDCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{EUTRA_ParametersCommon_Ext_Nr_HO_ToEN_DC_r16_Supported},
}

type EUTRA_ParametersCommon struct {
	Mfbi_EUTRA                *int64
	ModifiedMPR_BehaviorEUTRA *per.BitString
	MultiNS_Pmax_EUTRA        *int64
	Rs_SINR_MeasEUTRA         *int64
	Ne_DC                     *int64
	Nr_HO_ToEN_DC_r16         *int64
}

func (ie *EUTRA_ParametersCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eUTRAParametersCommonConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ne_DC != nil
	hasExtGroup1 := ie.Nr_HO_ToEN_DC_r16 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mfbi_EUTRA != nil, ie.ModifiedMPR_BehaviorEUTRA != nil, ie.MultiNS_Pmax_EUTRA != nil, ie.Rs_SINR_MeasEUTRA != nil}); err != nil {
		return err
	}

	// 3. mfbi-EUTRA: enumerated
	{
		if ie.Mfbi_EUTRA != nil {
			if err := e.EncodeEnumerated(*ie.Mfbi_EUTRA, eUTRAParametersCommonMfbiEUTRAConstraints); err != nil {
				return err
			}
		}
	}

	// 4. modifiedMPR-BehaviorEUTRA: bit-string
	{
		if ie.ModifiedMPR_BehaviorEUTRA != nil {
			if err := e.EncodeBitString(*ie.ModifiedMPR_BehaviorEUTRA, eUTRAParametersCommonModifiedMPRBehaviorEUTRAConstraints); err != nil {
				return err
			}
		}
	}

	// 5. multiNS-Pmax-EUTRA: enumerated
	{
		if ie.MultiNS_Pmax_EUTRA != nil {
			if err := e.EncodeEnumerated(*ie.MultiNS_Pmax_EUTRA, eUTRAParametersCommonMultiNSPmaxEUTRAConstraints); err != nil {
				return err
			}
		}
	}

	// 6. rs-SINR-MeasEUTRA: enumerated
	{
		if ie.Rs_SINR_MeasEUTRA != nil {
			if err := e.EncodeEnumerated(*ie.Rs_SINR_MeasEUTRA, eUTRAParametersCommonRsSINRMeasEUTRAConstraints); err != nil {
				return err
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
					{Name: "ne-DC", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ne_DC != nil}); err != nil {
				return err
			}

			if ie.Ne_DC != nil {
				if err := ex.EncodeEnumerated(*ie.Ne_DC, eUTRAParametersCommonExtNeDCConstraints); err != nil {
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
					{Name: "nr-HO-ToEN-DC-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Nr_HO_ToEN_DC_r16 != nil}); err != nil {
				return err
			}

			if ie.Nr_HO_ToEN_DC_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_HO_ToEN_DC_r16, eUTRAParametersCommonExtNrHOToENDCR16Constraints); err != nil {
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

func (ie *EUTRA_ParametersCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eUTRAParametersCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. mfbi-EUTRA: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(eUTRAParametersCommonMfbiEUTRAConstraints)
			if err != nil {
				return err
			}
			ie.Mfbi_EUTRA = &idx
		}
	}

	// 4. modifiedMPR-BehaviorEUTRA: bit-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeBitString(eUTRAParametersCommonModifiedMPRBehaviorEUTRAConstraints)
			if err != nil {
				return err
			}
			ie.ModifiedMPR_BehaviorEUTRA = &v
		}
	}

	// 5. multiNS-Pmax-EUTRA: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(eUTRAParametersCommonMultiNSPmaxEUTRAConstraints)
			if err != nil {
				return err
			}
			ie.MultiNS_Pmax_EUTRA = &idx
		}
	}

	// 6. rs-SINR-MeasEUTRA: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(eUTRAParametersCommonRsSINRMeasEUTRAConstraints)
			if err != nil {
				return err
			}
			ie.Rs_SINR_MeasEUTRA = &idx
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
				{Name: "ne-DC", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(eUTRAParametersCommonExtNeDCConstraints)
			if err != nil {
				return err
			}
			ie.Ne_DC = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "nr-HO-ToEN-DC-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(eUTRAParametersCommonExtNrHOToENDCR16Constraints)
			if err != nil {
				return err
			}
			ie.Nr_HO_ToEN_DC_r16 = &v
		}
	}

	return nil
}
