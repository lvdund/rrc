// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RadioBearerConfig (line 13108).

var radioBearerConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srb-ToAddModList", Optional: true},
		{Name: "srb3-ToRelease", Optional: true},
		{Name: "drb-ToAddModList", Optional: true},
		{Name: "drb-ToReleaseList", Optional: true},
		{Name: "securityConfig", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

const (
	RadioBearerConfig_Srb3_ToRelease_True = 0
)

var radioBearerConfigSrb3ToReleaseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RadioBearerConfig_Srb3_ToRelease_True},
}

const (
	RadioBearerConfig_Ext_Srb4_ToRelease_r17_True = 0
)

var radioBearerConfigExtSrb4ToReleaseR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RadioBearerConfig_Ext_Srb4_ToRelease_r17_True},
}

const (
	RadioBearerConfig_Ext_Srb5_ToRelease_r18_True = 0
)

var radioBearerConfigExtSrb5ToReleaseR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RadioBearerConfig_Ext_Srb5_ToRelease_r18_True},
}

const (
	RadioBearerConfig_Ext_Srb6_ToRelease_r19_True = 0
)

var radioBearerConfigExtSrb6ToReleaseR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RadioBearerConfig_Ext_Srb6_ToRelease_r19_True},
}

type RadioBearerConfig struct {
	Srb_ToAddModList      *SRB_ToAddModList
	Srb3_ToRelease        *int64
	Drb_ToAddModList      *DRB_ToAddModList
	Drb_ToReleaseList     *DRB_ToReleaseList
	SecurityConfig        *SecurityConfig
	Mrb_ToAddModList_r17  *MRB_ToAddModList_r17
	Mrb_ToReleaseList_r17 *MRB_ToReleaseList_r17
	Srb4_ToAddMod_r17     *SRB_ToAddMod
	Srb4_ToRelease_r17    *int64
	Srb5_ToAddMod_r18     *SRB_ToAddMod
	Srb5_ToRelease_r18    *int64
	Srb6_ToAddMod_r19     *SRB_ToAddMod
	Srb6_ToRelease_r19    *int64
}

func (ie *RadioBearerConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(radioBearerConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Mrb_ToAddModList_r17 != nil || ie.Mrb_ToReleaseList_r17 != nil || ie.Srb4_ToAddMod_r17 != nil || ie.Srb4_ToRelease_r17 != nil
	hasExtGroup1 := ie.Srb5_ToAddMod_r18 != nil || ie.Srb5_ToRelease_r18 != nil
	hasExtGroup2 := ie.Srb6_ToAddMod_r19 != nil || ie.Srb6_ToRelease_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srb_ToAddModList != nil, ie.Srb3_ToRelease != nil, ie.Drb_ToAddModList != nil, ie.Drb_ToReleaseList != nil, ie.SecurityConfig != nil}); err != nil {
		return err
	}

	// 3. srb-ToAddModList: ref
	{
		if ie.Srb_ToAddModList != nil {
			if err := ie.Srb_ToAddModList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. srb3-ToRelease: enumerated
	{
		if ie.Srb3_ToRelease != nil {
			if err := e.EncodeEnumerated(*ie.Srb3_ToRelease, radioBearerConfigSrb3ToReleaseConstraints); err != nil {
				return err
			}
		}
	}

	// 5. drb-ToAddModList: ref
	{
		if ie.Drb_ToAddModList != nil {
			if err := ie.Drb_ToAddModList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. drb-ToReleaseList: ref
	{
		if ie.Drb_ToReleaseList != nil {
			if err := ie.Drb_ToReleaseList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. securityConfig: ref
	{
		if ie.SecurityConfig != nil {
			if err := ie.SecurityConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "mrb-ToAddModList-r17", Optional: true},
					{Name: "mrb-ToReleaseList-r17", Optional: true},
					{Name: "srb4-ToAddMod-r17", Optional: true},
					{Name: "srb4-ToRelease-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Mrb_ToAddModList_r17 != nil, ie.Mrb_ToReleaseList_r17 != nil, ie.Srb4_ToAddMod_r17 != nil, ie.Srb4_ToRelease_r17 != nil}); err != nil {
				return err
			}

			if ie.Mrb_ToAddModList_r17 != nil {
				if err := ie.Mrb_ToAddModList_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Mrb_ToReleaseList_r17 != nil {
				if err := ie.Mrb_ToReleaseList_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Srb4_ToAddMod_r17 != nil {
				if err := ie.Srb4_ToAddMod_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Srb4_ToRelease_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Srb4_ToRelease_r17, radioBearerConfigExtSrb4ToReleaseR17Constraints); err != nil {
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
					{Name: "srb5-ToAddMod-r18", Optional: true},
					{Name: "srb5-ToRelease-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srb5_ToAddMod_r18 != nil, ie.Srb5_ToRelease_r18 != nil}); err != nil {
				return err
			}

			if ie.Srb5_ToAddMod_r18 != nil {
				if err := ie.Srb5_ToAddMod_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Srb5_ToRelease_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Srb5_ToRelease_r18, radioBearerConfigExtSrb5ToReleaseR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "srb6-ToAddMod-r19", Optional: true},
					{Name: "srb6-ToRelease-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srb6_ToAddMod_r19 != nil, ie.Srb6_ToRelease_r19 != nil}); err != nil {
				return err
			}

			if ie.Srb6_ToAddMod_r19 != nil {
				if err := ie.Srb6_ToAddMod_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Srb6_ToRelease_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Srb6_ToRelease_r19, radioBearerConfigExtSrb6ToReleaseR19Constraints); err != nil {
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

func (ie *RadioBearerConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(radioBearerConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. srb-ToAddModList: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Srb_ToAddModList = new(SRB_ToAddModList)
			if err := ie.Srb_ToAddModList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. srb3-ToRelease: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(radioBearerConfigSrb3ToReleaseConstraints)
			if err != nil {
				return err
			}
			ie.Srb3_ToRelease = &idx
		}
	}

	// 5. drb-ToAddModList: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Drb_ToAddModList = new(DRB_ToAddModList)
			if err := ie.Drb_ToAddModList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. drb-ToReleaseList: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Drb_ToReleaseList = new(DRB_ToReleaseList)
			if err := ie.Drb_ToReleaseList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. securityConfig: ref
	{
		if seq.IsComponentPresent(4) {
			ie.SecurityConfig = new(SecurityConfig)
			if err := ie.SecurityConfig.Decode(d); err != nil {
				return err
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
				{Name: "mrb-ToAddModList-r17", Optional: true},
				{Name: "mrb-ToReleaseList-r17", Optional: true},
				{Name: "srb4-ToAddMod-r17", Optional: true},
				{Name: "srb4-ToRelease-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Mrb_ToAddModList_r17 = new(MRB_ToAddModList_r17)
			if err := ie.Mrb_ToAddModList_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Mrb_ToReleaseList_r17 = new(MRB_ToReleaseList_r17)
			if err := ie.Mrb_ToReleaseList_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Srb4_ToAddMod_r17 = new(SRB_ToAddMod)
			if err := ie.Srb4_ToAddMod_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(radioBearerConfigExtSrb4ToReleaseR17Constraints)
			if err != nil {
				return err
			}
			ie.Srb4_ToRelease_r17 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "srb5-ToAddMod-r18", Optional: true},
				{Name: "srb5-ToRelease-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Srb5_ToAddMod_r18 = new(SRB_ToAddMod)
			if err := ie.Srb5_ToAddMod_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(radioBearerConfigExtSrb5ToReleaseR18Constraints)
			if err != nil {
				return err
			}
			ie.Srb5_ToRelease_r18 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "srb6-ToAddMod-r19", Optional: true},
				{Name: "srb6-ToRelease-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Srb6_ToAddMod_r19 = new(SRB_ToAddMod)
			if err := ie.Srb6_ToAddMod_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(radioBearerConfigExtSrb6ToReleaseR19Constraints)
			if err != nil {
				return err
			}
			ie.Srb6_ToRelease_r19 = &v
		}
	}

	return nil
}
