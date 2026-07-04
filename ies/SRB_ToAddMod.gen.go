// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRB-ToAddMod (line 13133).

var sRBToAddModConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "srb-Identity"},
		{Name: "reestablishPDCP", Optional: true},
		{Name: "discardOnPDCP", Optional: true},
		{Name: "pdcp-Config", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

const (
	SRB_ToAddMod_ReestablishPDCP_True = 0
)

var sRBToAddModReestablishPDCPConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRB_ToAddMod_ReestablishPDCP_True},
}

const (
	SRB_ToAddMod_DiscardOnPDCP_True = 0
)

var sRBToAddModDiscardOnPDCPConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRB_ToAddMod_DiscardOnPDCP_True},
}

const (
	SRB_ToAddMod_Ext_N3c_BearerAssociated_r18_True = 0
)

var sRBToAddModExtN3cBearerAssociatedR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SRB_ToAddMod_Ext_N3c_BearerAssociated_r18_True},
}

type SRB_ToAddMod struct {
	Srb_Identity             SRB_Identity
	ReestablishPDCP          *int64
	DiscardOnPDCP            *int64
	Pdcp_Config              *PDCP_Config
	Srb_Identity_v1700       *SRB_Identity_v1700
	Srb_Identity_v1800       *SRB_Identity_v1800
	N3c_BearerAssociated_r18 *int64
	Srb_Identity_v1900       *SRB_Identity_v1900
}

func (ie *SRB_ToAddMod) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sRBToAddModConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Srb_Identity_v1700 != nil
	hasExtGroup1 := ie.Srb_Identity_v1800 != nil || ie.N3c_BearerAssociated_r18 != nil
	hasExtGroup2 := ie.Srb_Identity_v1900 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ReestablishPDCP != nil, ie.DiscardOnPDCP != nil, ie.Pdcp_Config != nil}); err != nil {
		return err
	}

	// 3. srb-Identity: ref
	{
		if err := ie.Srb_Identity.Encode(e); err != nil {
			return err
		}
	}

	// 4. reestablishPDCP: enumerated
	{
		if ie.ReestablishPDCP != nil {
			if err := e.EncodeEnumerated(*ie.ReestablishPDCP, sRBToAddModReestablishPDCPConstraints); err != nil {
				return err
			}
		}
	}

	// 5. discardOnPDCP: enumerated
	{
		if ie.DiscardOnPDCP != nil {
			if err := e.EncodeEnumerated(*ie.DiscardOnPDCP, sRBToAddModDiscardOnPDCPConstraints); err != nil {
				return err
			}
		}
	}

	// 6. pdcp-Config: ref
	{
		if ie.Pdcp_Config != nil {
			if err := ie.Pdcp_Config.Encode(e); err != nil {
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
					{Name: "srb-Identity-v1700", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srb_Identity_v1700 != nil}); err != nil {
				return err
			}

			if ie.Srb_Identity_v1700 != nil {
				if err := ie.Srb_Identity_v1700.Encode(ex); err != nil {
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
					{Name: "srb-Identity-v1800", Optional: true},
					{Name: "n3c-BearerAssociated-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srb_Identity_v1800 != nil, ie.N3c_BearerAssociated_r18 != nil}); err != nil {
				return err
			}

			if ie.Srb_Identity_v1800 != nil {
				if err := ie.Srb_Identity_v1800.Encode(ex); err != nil {
					return err
				}
			}

			if ie.N3c_BearerAssociated_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.N3c_BearerAssociated_r18, sRBToAddModExtN3cBearerAssociatedR18Constraints); err != nil {
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
					{Name: "srb-Identity-v1900", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srb_Identity_v1900 != nil}); err != nil {
				return err
			}

			if ie.Srb_Identity_v1900 != nil {
				if err := ie.Srb_Identity_v1900.Encode(ex); err != nil {
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

func (ie *SRB_ToAddMod) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sRBToAddModConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. srb-Identity: ref
	{
		if err := ie.Srb_Identity.Decode(d); err != nil {
			return err
		}
	}

	// 4. reestablishPDCP: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sRBToAddModReestablishPDCPConstraints)
			if err != nil {
				return err
			}
			ie.ReestablishPDCP = &idx
		}
	}

	// 5. discardOnPDCP: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sRBToAddModDiscardOnPDCPConstraints)
			if err != nil {
				return err
			}
			ie.DiscardOnPDCP = &idx
		}
	}

	// 6. pdcp-Config: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Pdcp_Config = new(PDCP_Config)
			if err := ie.Pdcp_Config.Decode(d); err != nil {
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
				{Name: "srb-Identity-v1700", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Srb_Identity_v1700 = new(SRB_Identity_v1700)
			if err := ie.Srb_Identity_v1700.Decode(dx); err != nil {
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
				{Name: "srb-Identity-v1800", Optional: true},
				{Name: "n3c-BearerAssociated-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Srb_Identity_v1800 = new(SRB_Identity_v1800)
			if err := ie.Srb_Identity_v1800.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(sRBToAddModExtN3cBearerAssociatedR18Constraints)
			if err != nil {
				return err
			}
			ie.N3c_BearerAssociated_r18 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "srb-Identity-v1900", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Srb_Identity_v1900 = new(SRB_Identity_v1900)
			if err := ie.Srb_Identity_v1900.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
