// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: TCI-State (line 15993).

var tCIStateConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tci-StateId"},
		{Name: "qcl-Type1"},
		{Name: "qcl-Type2", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

const (
	TCI_State_Ext_PathlossOffset_r19_DB_12 = 0
	TCI_State_Ext_PathlossOffset_r19_DB_8  = 1
	TCI_State_Ext_PathlossOffset_r19_DB_4  = 2
	TCI_State_Ext_PathlossOffset_r19_DB0   = 3
	TCI_State_Ext_PathlossOffset_r19_DB4   = 4
	TCI_State_Ext_PathlossOffset_r19_DB8   = 5
	TCI_State_Ext_PathlossOffset_r19_DB12  = 6
	TCI_State_Ext_PathlossOffset_r19_DB16  = 7
	TCI_State_Ext_PathlossOffset_r19_DB20  = 8
	TCI_State_Ext_PathlossOffset_r19_DB24  = 9
	TCI_State_Ext_PathlossOffset_r19_DB28  = 10
	TCI_State_Ext_PathlossOffset_r19_DB32  = 11
	TCI_State_Ext_PathlossOffset_r19_DB36  = 12
	TCI_State_Ext_PathlossOffset_r19_DB40  = 13
	TCI_State_Ext_PathlossOffset_r19_DB44  = 14
	TCI_State_Ext_PathlossOffset_r19_DB48  = 15
	TCI_State_Ext_PathlossOffset_r19_DB52  = 16
	TCI_State_Ext_PathlossOffset_r19_DB56  = 17
	TCI_State_Ext_PathlossOffset_r19_DB60  = 18
)

var tCIStateExtPathlossOffsetR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TCI_State_Ext_PathlossOffset_r19_DB_12, TCI_State_Ext_PathlossOffset_r19_DB_8, TCI_State_Ext_PathlossOffset_r19_DB_4, TCI_State_Ext_PathlossOffset_r19_DB0, TCI_State_Ext_PathlossOffset_r19_DB4, TCI_State_Ext_PathlossOffset_r19_DB8, TCI_State_Ext_PathlossOffset_r19_DB12, TCI_State_Ext_PathlossOffset_r19_DB16, TCI_State_Ext_PathlossOffset_r19_DB20, TCI_State_Ext_PathlossOffset_r19_DB24, TCI_State_Ext_PathlossOffset_r19_DB28, TCI_State_Ext_PathlossOffset_r19_DB32, TCI_State_Ext_PathlossOffset_r19_DB36, TCI_State_Ext_PathlossOffset_r19_DB40, TCI_State_Ext_PathlossOffset_r19_DB44, TCI_State_Ext_PathlossOffset_r19_DB48, TCI_State_Ext_PathlossOffset_r19_DB52, TCI_State_Ext_PathlossOffset_r19_DB56, TCI_State_Ext_PathlossOffset_r19_DB60},
}

type TCI_State struct {
	Tci_StateId                TCI_StateId
	Qcl_Type1                  QCL_Info
	Qcl_Type2                  *QCL_Info
	AdditionalPCI_r17          *AdditionalPCIIndex_r17
	PathlossReferenceRS_Id_r17 *PathlossReferenceRS_Id_r17
	Ul_PowerControl_r17        *Uplink_PowerControlId_r17
	Tag_Id_Ptr_r18             *Tag_Id_Ptr_r18
	PathlossOffset_r19         *int64
}

func (ie *TCI_State) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tCIStateConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.AdditionalPCI_r17 != nil || ie.PathlossReferenceRS_Id_r17 != nil || ie.Ul_PowerControl_r17 != nil
	hasExtGroup1 := ie.Tag_Id_Ptr_r18 != nil
	hasExtGroup2 := ie.PathlossOffset_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Qcl_Type2 != nil}); err != nil {
		return err
	}

	// 3. tci-StateId: ref
	{
		if err := ie.Tci_StateId.Encode(e); err != nil {
			return err
		}
	}

	// 4. qcl-Type1: ref
	{
		if err := ie.Qcl_Type1.Encode(e); err != nil {
			return err
		}
	}

	// 5. qcl-Type2: ref
	{
		if ie.Qcl_Type2 != nil {
			if err := ie.Qcl_Type2.Encode(e); err != nil {
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
					{Name: "additionalPCI-r17", Optional: true},
					{Name: "pathlossReferenceRS-Id-r17", Optional: true},
					{Name: "ul-powerControl-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AdditionalPCI_r17 != nil, ie.PathlossReferenceRS_Id_r17 != nil, ie.Ul_PowerControl_r17 != nil}); err != nil {
				return err
			}

			if ie.AdditionalPCI_r17 != nil {
				if err := ie.AdditionalPCI_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PathlossReferenceRS_Id_r17 != nil {
				if err := ie.PathlossReferenceRS_Id_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ul_PowerControl_r17 != nil {
				if err := ie.Ul_PowerControl_r17.Encode(ex); err != nil {
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
					{Name: "tag-Id-ptr-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Tag_Id_Ptr_r18 != nil}); err != nil {
				return err
			}

			if ie.Tag_Id_Ptr_r18 != nil {
				if err := ie.Tag_Id_Ptr_r18.Encode(ex); err != nil {
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
					{Name: "pathlossOffset-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PathlossOffset_r19 != nil}); err != nil {
				return err
			}

			if ie.PathlossOffset_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PathlossOffset_r19, tCIStateExtPathlossOffsetR19Constraints); err != nil {
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

func (ie *TCI_State) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tCIStateConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. tci-StateId: ref
	{
		if err := ie.Tci_StateId.Decode(d); err != nil {
			return err
		}
	}

	// 4. qcl-Type1: ref
	{
		if err := ie.Qcl_Type1.Decode(d); err != nil {
			return err
		}
	}

	// 5. qcl-Type2: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Qcl_Type2 = new(QCL_Info)
			if err := ie.Qcl_Type2.Decode(d); err != nil {
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
				{Name: "additionalPCI-r17", Optional: true},
				{Name: "pathlossReferenceRS-Id-r17", Optional: true},
				{Name: "ul-powerControl-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.AdditionalPCI_r17 = new(AdditionalPCIIndex_r17)
			if err := ie.AdditionalPCI_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.PathlossReferenceRS_Id_r17 = new(PathlossReferenceRS_Id_r17)
			if err := ie.PathlossReferenceRS_Id_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Ul_PowerControl_r17 = new(Uplink_PowerControlId_r17)
			if err := ie.Ul_PowerControl_r17.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "tag-Id-ptr-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Tag_Id_Ptr_r18 = new(Tag_Id_Ptr_r18)
			if err := ie.Tag_Id_Ptr_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pathlossOffset-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(tCIStateExtPathlossOffsetR19Constraints)
			if err != nil {
				return err
			}
			ie.PathlossOffset_r19 = &v
		}
	}

	return nil
}
