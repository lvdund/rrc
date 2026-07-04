// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUSCH-ConfigCommon (line 12547).

var pUSCHConfigCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "groupHoppingEnabledTransformPrecoding", Optional: true},
		{Name: "pusch-TimeDomainAllocationList", Optional: true},
		{Name: "msg3-DeltaPreamble", Optional: true},
		{Name: "p0-NominalWithGrant", Optional: true},
	},
}

const (
	PUSCH_ConfigCommon_GroupHoppingEnabledTransformPrecoding_Enabled = 0
)

var pUSCHConfigCommonGroupHoppingEnabledTransformPrecodingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_ConfigCommon_GroupHoppingEnabledTransformPrecoding_Enabled},
}

var pUSCHConfigCommonMsg3DeltaPreambleConstraints = per.Constrained(-1, 6)

var pUSCHConfigCommonP0NominalWithGrantConstraints = per.Constrained(-202, 24)

type PUSCH_ConfigCommon struct {
	GroupHoppingEnabledTransformPrecoding *int64
	Pusch_TimeDomainAllocationList        *PUSCH_TimeDomainResourceAllocationList
	Msg3_DeltaPreamble                    *int64
	P0_NominalWithGrant                   *int64
}

func (ie *PUSCH_ConfigCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHConfigCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.GroupHoppingEnabledTransformPrecoding != nil, ie.Pusch_TimeDomainAllocationList != nil, ie.Msg3_DeltaPreamble != nil, ie.P0_NominalWithGrant != nil}); err != nil {
		return err
	}

	// 3. groupHoppingEnabledTransformPrecoding: enumerated
	{
		if ie.GroupHoppingEnabledTransformPrecoding != nil {
			if err := e.EncodeEnumerated(*ie.GroupHoppingEnabledTransformPrecoding, pUSCHConfigCommonGroupHoppingEnabledTransformPrecodingConstraints); err != nil {
				return err
			}
		}
	}

	// 4. pusch-TimeDomainAllocationList: ref
	{
		if ie.Pusch_TimeDomainAllocationList != nil {
			if err := ie.Pusch_TimeDomainAllocationList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. msg3-DeltaPreamble: integer
	{
		if ie.Msg3_DeltaPreamble != nil {
			if err := e.EncodeInteger(*ie.Msg3_DeltaPreamble, pUSCHConfigCommonMsg3DeltaPreambleConstraints); err != nil {
				return err
			}
		}
	}

	// 6. p0-NominalWithGrant: integer
	{
		if ie.P0_NominalWithGrant != nil {
			if err := e.EncodeInteger(*ie.P0_NominalWithGrant, pUSCHConfigCommonP0NominalWithGrantConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUSCH_ConfigCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHConfigCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. groupHoppingEnabledTransformPrecoding: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pUSCHConfigCommonGroupHoppingEnabledTransformPrecodingConstraints)
			if err != nil {
				return err
			}
			ie.GroupHoppingEnabledTransformPrecoding = &idx
		}
	}

	// 4. pusch-TimeDomainAllocationList: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Pusch_TimeDomainAllocationList = new(PUSCH_TimeDomainResourceAllocationList)
			if err := ie.Pusch_TimeDomainAllocationList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. msg3-DeltaPreamble: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(pUSCHConfigCommonMsg3DeltaPreambleConstraints)
			if err != nil {
				return err
			}
			ie.Msg3_DeltaPreamble = &v
		}
	}

	// 6. p0-NominalWithGrant: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(pUSCHConfigCommonP0NominalWithGrantConstraints)
			if err != nil {
				return err
			}
			ie.P0_NominalWithGrant = &v
		}
	}

	return nil
}
