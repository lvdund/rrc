// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SDT-Config-r17 (line 1369).

var sDTConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sdt-DRB-List-r17", Optional: true},
		{Name: "sdt-SRB2-Indication-r17", Optional: true},
		{Name: "sdt-MAC-PHY-CG-Config-r17", Optional: true},
		{Name: "sdt-DRB-ContinueROHC-r17", Optional: true},
	},
}

var sDTConfigR17SdtDRBListR17Constraints = per.SizeRange(0, common.MaxDRB)

const (
	SDT_Config_r17_Sdt_SRB2_Indication_r17_Allowed = 0
)

var sDTConfigR17SdtSRB2IndicationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SDT_Config_r17_Sdt_SRB2_Indication_r17_Allowed},
}

var sDT_Config_r17SdtMACPHYCGConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SDT_Config_r17_Sdt_MAC_PHY_CG_Config_r17_Release = 0
	SDT_Config_r17_Sdt_MAC_PHY_CG_Config_r17_Setup   = 1
)

const (
	SDT_Config_r17_Sdt_DRB_ContinueROHC_r17_Cell = 0
	SDT_Config_r17_Sdt_DRB_ContinueROHC_r17_Rna  = 1
)

var sDTConfigR17SdtDRBContinueROHCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SDT_Config_r17_Sdt_DRB_ContinueROHC_r17_Cell, SDT_Config_r17_Sdt_DRB_ContinueROHC_r17_Rna},
}

type SDT_Config_r17 struct {
	Sdt_DRB_List_r17          []DRB_Identity
	Sdt_SRB2_Indication_r17   *int64
	Sdt_MAC_PHY_CG_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SDT_CG_Config_r17
	}
	Sdt_DRB_ContinueROHC_r17 *int64
}

func (ie *SDT_Config_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sDTConfigR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sdt_DRB_List_r17 != nil, ie.Sdt_SRB2_Indication_r17 != nil, ie.Sdt_MAC_PHY_CG_Config_r17 != nil, ie.Sdt_DRB_ContinueROHC_r17 != nil}); err != nil {
		return err
	}

	// 2. sdt-DRB-List-r17: sequence-of
	{
		if ie.Sdt_DRB_List_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(sDTConfigR17SdtDRBListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sdt_DRB_List_r17))); err != nil {
				return err
			}
			for i := range ie.Sdt_DRB_List_r17 {
				if err := ie.Sdt_DRB_List_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. sdt-SRB2-Indication-r17: enumerated
	{
		if ie.Sdt_SRB2_Indication_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sdt_SRB2_Indication_r17, sDTConfigR17SdtSRB2IndicationR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sdt-MAC-PHY-CG-Config-r17: choice
	{
		if ie.Sdt_MAC_PHY_CG_Config_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(sDT_Config_r17SdtMACPHYCGConfigR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sdt_MAC_PHY_CG_Config_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sdt_MAC_PHY_CG_Config_r17).Choice {
			case SDT_Config_r17_Sdt_MAC_PHY_CG_Config_r17_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SDT_Config_r17_Sdt_MAC_PHY_CG_Config_r17_Setup:
				if err := (*ie.Sdt_MAC_PHY_CG_Config_r17).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sdt_MAC_PHY_CG_Config_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. sdt-DRB-ContinueROHC-r17: enumerated
	{
		if ie.Sdt_DRB_ContinueROHC_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Sdt_DRB_ContinueROHC_r17, sDTConfigR17SdtDRBContinueROHCR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SDT_Config_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sDTConfigR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sdt-DRB-List-r17: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(sDTConfigR17SdtDRBListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sdt_DRB_List_r17 = make([]DRB_Identity, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sdt_DRB_List_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sdt-SRB2-Indication-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sDTConfigR17SdtSRB2IndicationR17Constraints)
			if err != nil {
				return err
			}
			ie.Sdt_SRB2_Indication_r17 = &idx
		}
	}

	// 4. sdt-MAC-PHY-CG-Config-r17: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Sdt_MAC_PHY_CG_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SDT_CG_Config_r17
			}{}
			choiceDec := d.NewChoiceDecoder(sDT_Config_r17SdtMACPHYCGConfigR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sdt_MAC_PHY_CG_Config_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SDT_Config_r17_Sdt_MAC_PHY_CG_Config_r17_Release:
				(*ie.Sdt_MAC_PHY_CG_Config_r17).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SDT_Config_r17_Sdt_MAC_PHY_CG_Config_r17_Setup:
				(*ie.Sdt_MAC_PHY_CG_Config_r17).Setup = new(SDT_CG_Config_r17)
				if err := (*ie.Sdt_MAC_PHY_CG_Config_r17).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. sdt-DRB-ContinueROHC-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sDTConfigR17SdtDRBContinueROHCR17Constraints)
			if err != nil {
				return err
			}
			ie.Sdt_DRB_ContinueROHC_r17 = &idx
		}
	}

	return nil
}
