// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-Group-Config-r17 (line 18343).

var pUCCHGroupConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fr1-FR1-NonSharedTDD-r17", Optional: true},
		{Name: "fr2-FR2-NonSharedTDD-r17", Optional: true},
		{Name: "fr1-FR2-NonSharedTDD-r17", Optional: true},
	},
}

const (
	PUCCH_Group_Config_r17_Fr1_FR1_NonSharedTDD_r17_Supported = 0
)

var pUCCHGroupConfigR17Fr1FR1NonSharedTDDR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Group_Config_r17_Fr1_FR1_NonSharedTDD_r17_Supported},
}

const (
	PUCCH_Group_Config_r17_Fr2_FR2_NonSharedTDD_r17_Supported = 0
)

var pUCCHGroupConfigR17Fr2FR2NonSharedTDDR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Group_Config_r17_Fr2_FR2_NonSharedTDD_r17_Supported},
}

const (
	PUCCH_Group_Config_r17_Fr1_FR2_NonSharedTDD_r17_Supported = 0
)

var pUCCHGroupConfigR17Fr1FR2NonSharedTDDR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_Group_Config_r17_Fr1_FR2_NonSharedTDD_r17_Supported},
}

type PUCCH_Group_Config_r17 struct {
	Fr1_FR1_NonSharedTDD_r17 *int64
	Fr2_FR2_NonSharedTDD_r17 *int64
	Fr1_FR2_NonSharedTDD_r17 *int64
}

func (ie *PUCCH_Group_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHGroupConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Fr1_FR1_NonSharedTDD_r17 != nil, ie.Fr2_FR2_NonSharedTDD_r17 != nil, ie.Fr1_FR2_NonSharedTDD_r17 != nil}); err != nil {
		return err
	}

	// 2. fr1-FR1-NonSharedTDD-r17: enumerated
	{
		if ie.Fr1_FR1_NonSharedTDD_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Fr1_FR1_NonSharedTDD_r17, pUCCHGroupConfigR17Fr1FR1NonSharedTDDR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. fr2-FR2-NonSharedTDD-r17: enumerated
	{
		if ie.Fr2_FR2_NonSharedTDD_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Fr2_FR2_NonSharedTDD_r17, pUCCHGroupConfigR17Fr2FR2NonSharedTDDR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. fr1-FR2-NonSharedTDD-r17: enumerated
	{
		if ie.Fr1_FR2_NonSharedTDD_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Fr1_FR2_NonSharedTDD_r17, pUCCHGroupConfigR17Fr1FR2NonSharedTDDR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUCCH_Group_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHGroupConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. fr1-FR1-NonSharedTDD-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pUCCHGroupConfigR17Fr1FR1NonSharedTDDR17Constraints)
			if err != nil {
				return err
			}
			ie.Fr1_FR1_NonSharedTDD_r17 = &idx
		}
	}

	// 3. fr2-FR2-NonSharedTDD-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(pUCCHGroupConfigR17Fr2FR2NonSharedTDDR17Constraints)
			if err != nil {
				return err
			}
			ie.Fr2_FR2_NonSharedTDD_r17 = &idx
		}
	}

	// 4. fr1-FR2-NonSharedTDD-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(pUCCHGroupConfigR17Fr1FR2NonSharedTDDR17Constraints)
			if err != nil {
				return err
			}
			ie.Fr1_FR2_NonSharedTDD_r17 = &idx
		}
	}

	return nil
}
