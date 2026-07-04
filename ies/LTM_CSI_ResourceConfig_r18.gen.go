// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LTM-CSI-ResourceConfig-r18 (line 8914).

var lTMCSIResourceConfigR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ltm-CSI-ResourceConfigId-r18"},
		{Name: "ltm-SSB-ResourceSet-r18"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	LTM_CSI_ResourceConfig_r18_Ext_ResourceType_r19_Periodic       = 0
	LTM_CSI_ResourceConfig_r18_Ext_ResourceType_r19_SemiPersistent = 1
)

var lTMCSIResourceConfigR18ExtResourceTypeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTM_CSI_ResourceConfig_r18_Ext_ResourceType_r19_Periodic, LTM_CSI_ResourceConfig_r18_Ext_ResourceType_r19_SemiPersistent},
}

type LTM_CSI_ResourceConfig_r18 struct {
	Ltm_CSI_ResourceConfigId_r18   LTM_CSI_ResourceConfigId_r18
	Ltm_SSB_ResourceSet_r18        LTM_SSB_ResourceSet_r18
	Ltm_NZP_CSI_RS_ResourceSet_r19 *LTM_NZP_CSI_RS_ResourceSet_r19
	Ltm_CSI_IM_ResourceSet_r19     *LTM_CSI_IM_ResourceSet_r19
	ResourceType_r19               *int64
}

func (ie *LTM_CSI_ResourceConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTMCSIResourceConfigR18Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ltm_NZP_CSI_RS_ResourceSet_r19 != nil || ie.Ltm_CSI_IM_ResourceSet_r19 != nil || ie.ResourceType_r19 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. ltm-CSI-ResourceConfigId-r18: ref
	{
		if err := ie.Ltm_CSI_ResourceConfigId_r18.Encode(e); err != nil {
			return err
		}
	}

	// 3. ltm-SSB-ResourceSet-r18: ref
	{
		if err := ie.Ltm_SSB_ResourceSet_r18.Encode(e); err != nil {
			return err
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
					{Name: "ltm-NZP-CSI-RS-ResourceSet-r19", Optional: true},
					{Name: "ltm-CSI-IM-ResourceSet-r19", Optional: true},
					{Name: "resourceType-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ltm_NZP_CSI_RS_ResourceSet_r19 != nil, ie.Ltm_CSI_IM_ResourceSet_r19 != nil, ie.ResourceType_r19 != nil}); err != nil {
				return err
			}

			if ie.Ltm_NZP_CSI_RS_ResourceSet_r19 != nil {
				if err := ie.Ltm_NZP_CSI_RS_ResourceSet_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ltm_CSI_IM_ResourceSet_r19 != nil {
				if err := ie.Ltm_CSI_IM_ResourceSet_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ResourceType_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.ResourceType_r19, lTMCSIResourceConfigR18ExtResourceTypeR19Constraints); err != nil {
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

func (ie *LTM_CSI_ResourceConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTMCSIResourceConfigR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. ltm-CSI-ResourceConfigId-r18: ref
	{
		if err := ie.Ltm_CSI_ResourceConfigId_r18.Decode(d); err != nil {
			return err
		}
	}

	// 3. ltm-SSB-ResourceSet-r18: ref
	{
		if err := ie.Ltm_SSB_ResourceSet_r18.Decode(d); err != nil {
			return err
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
				{Name: "ltm-NZP-CSI-RS-ResourceSet-r19", Optional: true},
				{Name: "ltm-CSI-IM-ResourceSet-r19", Optional: true},
				{Name: "resourceType-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ltm_NZP_CSI_RS_ResourceSet_r19 = new(LTM_NZP_CSI_RS_ResourceSet_r19)
			if err := ie.Ltm_NZP_CSI_RS_ResourceSet_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Ltm_CSI_IM_ResourceSet_r19 = new(LTM_CSI_IM_ResourceSet_r19)
			if err := ie.Ltm_CSI_IM_ResourceSet_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(lTMCSIResourceConfigR18ExtResourceTypeR19Constraints)
			if err != nil {
				return err
			}
			ie.ResourceType_r19 = &v
		}
	}

	return nil
}
