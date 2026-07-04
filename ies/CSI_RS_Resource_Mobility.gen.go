// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-RS-Resource-Mobility (line 7537).

var cSIRSResourceMobilityConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-RS-Index"},
		{Name: "slotConfig"},
		{Name: "associatedSSB", Optional: true},
		{Name: "frequencyDomainAllocation"},
		{Name: "firstOFDMSymbolInTimeDomain"},
		{Name: "sequenceGenerationConfig"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var cSI_RS_Resource_MobilitySlotConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ms4"},
		{Name: "ms5"},
		{Name: "ms10"},
		{Name: "ms20"},
		{Name: "ms40"},
	},
}

const (
	CSI_RS_Resource_Mobility_SlotConfig_Ms4  = 0
	CSI_RS_Resource_Mobility_SlotConfig_Ms5  = 1
	CSI_RS_Resource_Mobility_SlotConfig_Ms10 = 2
	CSI_RS_Resource_Mobility_SlotConfig_Ms20 = 3
	CSI_RS_Resource_Mobility_SlotConfig_Ms40 = 4
)

var cSI_RS_Resource_MobilityFrequencyDomainAllocationConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "row1"},
		{Name: "row2"},
	},
}

const (
	CSI_RS_Resource_Mobility_FrequencyDomainAllocation_Row1 = 0
	CSI_RS_Resource_Mobility_FrequencyDomainAllocation_Row2 = 1
)

var cSIRSResourceMobilityFirstOFDMSymbolInTimeDomainConstraints = per.Constrained(0, 13)

var cSIRSResourceMobilitySequenceGenerationConfigConstraints = per.Constrained(0, 1023)

var cSIRSResourceMobilityExtSlotConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "ms4"},
		{Name: "ms5"},
		{Name: "ms10"},
		{Name: "ms20"},
		{Name: "ms40"},
	},
}

const (
	CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms4  = 0
	CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms5  = 1
	CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms10 = 2
	CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms20 = 3
	CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms40 = 4
)

type CSI_RS_Resource_Mobility struct {
	Csi_RS_Index CSI_RS_Index
	SlotConfig   struct {
		Choice int
		Ms4    *int64
		Ms5    *int64
		Ms10   *int64
		Ms20   *int64
		Ms40   *int64
	}
	AssociatedSSB *struct {
		Ssb_Index        SSB_Index
		IsQuasiColocated bool
	}
	FrequencyDomainAllocation struct {
		Choice int
		Row1   *per.BitString
		Row2   *per.BitString
	}
	FirstOFDMSymbolInTimeDomain int64
	SequenceGenerationConfig    int64
	SlotConfig_r17              *struct {
		Choice int
		Ms4    *int64
		Ms5    *int64
		Ms10   *int64
		Ms20   *int64
		Ms40   *int64
	}
}

func (ie *CSI_RS_Resource_Mobility) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIRSResourceMobilityConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.SlotConfig_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AssociatedSSB != nil}); err != nil {
		return err
	}

	// 3. csi-RS-Index: ref
	{
		if err := ie.Csi_RS_Index.Encode(e); err != nil {
			return err
		}
	}

	// 4. slotConfig: choice
	{
		choiceEnc := e.NewChoiceEncoder(cSI_RS_Resource_MobilitySlotConfigConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.SlotConfig.Choice), false, nil); err != nil {
			return err
		}
		switch ie.SlotConfig.Choice {
		case CSI_RS_Resource_Mobility_SlotConfig_Ms4:
			if err := e.EncodeInteger((*ie.SlotConfig.Ms4), per.Constrained(0, 31)); err != nil {
				return err
			}
		case CSI_RS_Resource_Mobility_SlotConfig_Ms5:
			if err := e.EncodeInteger((*ie.SlotConfig.Ms5), per.Constrained(0, 39)); err != nil {
				return err
			}
		case CSI_RS_Resource_Mobility_SlotConfig_Ms10:
			if err := e.EncodeInteger((*ie.SlotConfig.Ms10), per.Constrained(0, 79)); err != nil {
				return err
			}
		case CSI_RS_Resource_Mobility_SlotConfig_Ms20:
			if err := e.EncodeInteger((*ie.SlotConfig.Ms20), per.Constrained(0, 159)); err != nil {
				return err
			}
		case CSI_RS_Resource_Mobility_SlotConfig_Ms40:
			if err := e.EncodeInteger((*ie.SlotConfig.Ms40), per.Constrained(0, 319)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.SlotConfig.Choice), Constraint: "index out of range"}
		}
	}

	// 5. associatedSSB: sequence
	{
		if ie.AssociatedSSB != nil {
			c := ie.AssociatedSSB
			if err := c.Ssb_Index.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeBoolean(c.IsQuasiColocated); err != nil {
				return err
			}
		}
	}

	// 6. frequencyDomainAllocation: choice
	{
		choiceEnc := e.NewChoiceEncoder(cSI_RS_Resource_MobilityFrequencyDomainAllocationConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.FrequencyDomainAllocation.Choice), false, nil); err != nil {
			return err
		}
		switch ie.FrequencyDomainAllocation.Choice {
		case CSI_RS_Resource_Mobility_FrequencyDomainAllocation_Row1:
			if err := e.EncodeBitString((*ie.FrequencyDomainAllocation.Row1), per.FixedSize(4)); err != nil {
				return err
			}
		case CSI_RS_Resource_Mobility_FrequencyDomainAllocation_Row2:
			if err := e.EncodeBitString((*ie.FrequencyDomainAllocation.Row2), per.FixedSize(12)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.FrequencyDomainAllocation.Choice), Constraint: "index out of range"}
		}
	}

	// 7. firstOFDMSymbolInTimeDomain: integer
	{
		if err := e.EncodeInteger(ie.FirstOFDMSymbolInTimeDomain, cSIRSResourceMobilityFirstOFDMSymbolInTimeDomainConstraints); err != nil {
			return err
		}
	}

	// 8. sequenceGenerationConfig: integer
	{
		if err := e.EncodeInteger(ie.SequenceGenerationConfig, cSIRSResourceMobilitySequenceGenerationConfigConstraints); err != nil {
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
					{Name: "slotConfig-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SlotConfig_r17 != nil}); err != nil {
				return err
			}

			if ie.SlotConfig_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(cSIRSResourceMobilityExtSlotConfigR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.SlotConfig_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.SlotConfig_r17).Choice {
				case CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms4:
					if err := ex.EncodeInteger((*(*ie.SlotConfig_r17).Ms4), per.Constrained(0, 255)); err != nil {
						return err
					}
				case CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms5:
					if err := ex.EncodeInteger((*(*ie.SlotConfig_r17).Ms5), per.Constrained(0, 319)); err != nil {
						return err
					}
				case CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms10:
					if err := ex.EncodeInteger((*(*ie.SlotConfig_r17).Ms10), per.Constrained(0, 639)); err != nil {
						return err
					}
				case CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms20:
					if err := ex.EncodeInteger((*(*ie.SlotConfig_r17).Ms20), per.Constrained(0, 1279)); err != nil {
						return err
					}
				case CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms40:
					if err := ex.EncodeInteger((*(*ie.SlotConfig_r17).Ms40), per.Constrained(0, 2559)); err != nil {
						return err
					}
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

func (ie *CSI_RS_Resource_Mobility) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIRSResourceMobilityConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. csi-RS-Index: ref
	{
		if err := ie.Csi_RS_Index.Decode(d); err != nil {
			return err
		}
	}

	// 4. slotConfig: choice
	{
		choiceDec := d.NewChoiceDecoder(cSI_RS_Resource_MobilitySlotConfigConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.SlotConfig.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CSI_RS_Resource_Mobility_SlotConfig_Ms4:
			ie.SlotConfig.Ms4 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 31))
			if err != nil {
				return err
			}
			(*ie.SlotConfig.Ms4) = v
		case CSI_RS_Resource_Mobility_SlotConfig_Ms5:
			ie.SlotConfig.Ms5 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*ie.SlotConfig.Ms5) = v
		case CSI_RS_Resource_Mobility_SlotConfig_Ms10:
			ie.SlotConfig.Ms10 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*ie.SlotConfig.Ms10) = v
		case CSI_RS_Resource_Mobility_SlotConfig_Ms20:
			ie.SlotConfig.Ms20 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*ie.SlotConfig.Ms20) = v
		case CSI_RS_Resource_Mobility_SlotConfig_Ms40:
			ie.SlotConfig.Ms40 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*ie.SlotConfig.Ms40) = v
		}
	}

	// 5. associatedSSB: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.AssociatedSSB = &struct {
				Ssb_Index        SSB_Index
				IsQuasiColocated bool
			}{}
			c := ie.AssociatedSSB
			{
				if err := c.Ssb_Index.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeBoolean()
				if err != nil {
					return err
				}
				c.IsQuasiColocated = v
			}
		}
	}

	// 6. frequencyDomainAllocation: choice
	{
		choiceDec := d.NewChoiceDecoder(cSI_RS_Resource_MobilityFrequencyDomainAllocationConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.FrequencyDomainAllocation.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CSI_RS_Resource_Mobility_FrequencyDomainAllocation_Row1:
			ie.FrequencyDomainAllocation.Row1 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(4))
			if err != nil {
				return err
			}
			(*ie.FrequencyDomainAllocation.Row1) = v
		case CSI_RS_Resource_Mobility_FrequencyDomainAllocation_Row2:
			ie.FrequencyDomainAllocation.Row2 = new(per.BitString)
			v, err := d.DecodeBitString(per.FixedSize(12))
			if err != nil {
				return err
			}
			(*ie.FrequencyDomainAllocation.Row2) = v
		}
	}

	// 7. firstOFDMSymbolInTimeDomain: integer
	{
		v4, err := d.DecodeInteger(cSIRSResourceMobilityFirstOFDMSymbolInTimeDomainConstraints)
		if err != nil {
			return err
		}
		ie.FirstOFDMSymbolInTimeDomain = v4
	}

	// 8. sequenceGenerationConfig: integer
	{
		v5, err := d.DecodeInteger(cSIRSResourceMobilitySequenceGenerationConfigConstraints)
		if err != nil {
			return err
		}
		ie.SequenceGenerationConfig = v5
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
				{Name: "slotConfig-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SlotConfig_r17 = &struct {
				Choice int
				Ms4    *int64
				Ms5    *int64
				Ms10   *int64
				Ms20   *int64
				Ms40   *int64
			}{}
			choiceDec := dx.NewChoiceDecoder(cSIRSResourceMobilityExtSlotConfigR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SlotConfig_r17).Choice = int(index)
			switch index {
			case CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms4:
				(*ie.SlotConfig_r17).Ms4 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(0, 255))
				if err != nil {
					return err
				}
				(*(*ie.SlotConfig_r17).Ms4) = v
			case CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms5:
				(*ie.SlotConfig_r17).Ms5 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(0, 319))
				if err != nil {
					return err
				}
				(*(*ie.SlotConfig_r17).Ms5) = v
			case CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms10:
				(*ie.SlotConfig_r17).Ms10 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(0, 639))
				if err != nil {
					return err
				}
				(*(*ie.SlotConfig_r17).Ms10) = v
			case CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms20:
				(*ie.SlotConfig_r17).Ms20 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(0, 1279))
				if err != nil {
					return err
				}
				(*(*ie.SlotConfig_r17).Ms20) = v
			case CSI_RS_Resource_Mobility_Ext_SlotConfig_r17_Ms40:
				(*ie.SlotConfig_r17).Ms40 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(0, 2559))
				if err != nil {
					return err
				}
				(*(*ie.SlotConfig_r17).Ms40) = v
			}
		}
	}

	return nil
}
