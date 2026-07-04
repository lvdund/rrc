// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUSCH-ConfigDCI-0-3-r18 (line 12530).

var pUSCHConfigDCI03R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceAllocationDCI-0-3-r18", Optional: true},
		{Name: "rbg-SizeDCI-0-3-r18", Optional: true},
		{Name: "resourceAllocationType1GranularityDCI-0-3-r18", Optional: true},
		{Name: "numberOfBitsForRV-DCI-0-3-r18", Optional: true},
		{Name: "harq-ProcessNumberSizeDCI-0-3-r18", Optional: true},
		{Name: "uci-OnPUSCH-ListDCI-0-3-r18", Optional: true},
	},
}

const (
	PUSCH_ConfigDCI_0_3_r18_ResourceAllocationDCI_0_3_r18_ResourceAllocationType0 = 0
	PUSCH_ConfigDCI_0_3_r18_ResourceAllocationDCI_0_3_r18_ResourceAllocationType1 = 1
	PUSCH_ConfigDCI_0_3_r18_ResourceAllocationDCI_0_3_r18_DynamicSwitch           = 2
)

var pUSCHConfigDCI03R18ResourceAllocationDCI03R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_ConfigDCI_0_3_r18_ResourceAllocationDCI_0_3_r18_ResourceAllocationType0, PUSCH_ConfigDCI_0_3_r18_ResourceAllocationDCI_0_3_r18_ResourceAllocationType1, PUSCH_ConfigDCI_0_3_r18_ResourceAllocationDCI_0_3_r18_DynamicSwitch},
}

const (
	PUSCH_ConfigDCI_0_3_r18_Rbg_SizeDCI_0_3_r18_Config2 = 0
	PUSCH_ConfigDCI_0_3_r18_Rbg_SizeDCI_0_3_r18_Config3 = 1
)

var pUSCHConfigDCI03R18RbgSizeDCI03R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_ConfigDCI_0_3_r18_Rbg_SizeDCI_0_3_r18_Config2, PUSCH_ConfigDCI_0_3_r18_Rbg_SizeDCI_0_3_r18_Config3},
}

const (
	PUSCH_ConfigDCI_0_3_r18_ResourceAllocationType1GranularityDCI_0_3_r18_N2  = 0
	PUSCH_ConfigDCI_0_3_r18_ResourceAllocationType1GranularityDCI_0_3_r18_N4  = 1
	PUSCH_ConfigDCI_0_3_r18_ResourceAllocationType1GranularityDCI_0_3_r18_N8  = 2
	PUSCH_ConfigDCI_0_3_r18_ResourceAllocationType1GranularityDCI_0_3_r18_N16 = 3
)

var pUSCHConfigDCI03R18ResourceAllocationType1GranularityDCI03R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_ConfigDCI_0_3_r18_ResourceAllocationType1GranularityDCI_0_3_r18_N2, PUSCH_ConfigDCI_0_3_r18_ResourceAllocationType1GranularityDCI_0_3_r18_N4, PUSCH_ConfigDCI_0_3_r18_ResourceAllocationType1GranularityDCI_0_3_r18_N8, PUSCH_ConfigDCI_0_3_r18_ResourceAllocationType1GranularityDCI_0_3_r18_N16},
}

var pUSCHConfigDCI03R18NumberOfBitsForRVDCI03R18Constraints = per.Constrained(0, 2)

var pUSCHConfigDCI03R18HarqProcessNumberSizeDCI03R18Constraints = per.Constrained(0, 5)

var pUSCH_ConfigDCI_0_3_r18UciOnPUSCHListDCI03R18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_ConfigDCI_0_3_r18_Uci_OnPUSCH_ListDCI_0_3_r18_Release = 0
	PUSCH_ConfigDCI_0_3_r18_Uci_OnPUSCH_ListDCI_0_3_r18_Setup   = 1
)

type PUSCH_ConfigDCI_0_3_r18 struct {
	ResourceAllocationDCI_0_3_r18                 *int64
	Rbg_SizeDCI_0_3_r18                           *int64
	ResourceAllocationType1GranularityDCI_0_3_r18 *int64
	NumberOfBitsForRV_DCI_0_3_r18                 *int64
	Harq_ProcessNumberSizeDCI_0_3_r18             *int64
	Uci_OnPUSCH_ListDCI_0_3_r18                   *struct {
		Choice  int
		Release *struct{}
		Setup   *UCI_OnPUSCH_ListDCI_0_1_r16
	}
}

func (ie *PUSCH_ConfigDCI_0_3_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHConfigDCI03R18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ResourceAllocationDCI_0_3_r18 != nil, ie.Rbg_SizeDCI_0_3_r18 != nil, ie.ResourceAllocationType1GranularityDCI_0_3_r18 != nil, ie.NumberOfBitsForRV_DCI_0_3_r18 != nil, ie.Harq_ProcessNumberSizeDCI_0_3_r18 != nil, ie.Uci_OnPUSCH_ListDCI_0_3_r18 != nil}); err != nil {
		return err
	}

	// 2. resourceAllocationDCI-0-3-r18: enumerated
	{
		if ie.ResourceAllocationDCI_0_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ResourceAllocationDCI_0_3_r18, pUSCHConfigDCI03R18ResourceAllocationDCI03R18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. rbg-SizeDCI-0-3-r18: enumerated
	{
		if ie.Rbg_SizeDCI_0_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Rbg_SizeDCI_0_3_r18, pUSCHConfigDCI03R18RbgSizeDCI03R18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. resourceAllocationType1GranularityDCI-0-3-r18: enumerated
	{
		if ie.ResourceAllocationType1GranularityDCI_0_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ResourceAllocationType1GranularityDCI_0_3_r18, pUSCHConfigDCI03R18ResourceAllocationType1GranularityDCI03R18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. numberOfBitsForRV-DCI-0-3-r18: integer
	{
		if ie.NumberOfBitsForRV_DCI_0_3_r18 != nil {
			if err := e.EncodeInteger(*ie.NumberOfBitsForRV_DCI_0_3_r18, pUSCHConfigDCI03R18NumberOfBitsForRVDCI03R18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. harq-ProcessNumberSizeDCI-0-3-r18: integer
	{
		if ie.Harq_ProcessNumberSizeDCI_0_3_r18 != nil {
			if err := e.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_0_3_r18, pUSCHConfigDCI03R18HarqProcessNumberSizeDCI03R18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. uci-OnPUSCH-ListDCI-0-3-r18: choice
	{
		if ie.Uci_OnPUSCH_ListDCI_0_3_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(pUSCH_ConfigDCI_0_3_r18UciOnPUSCHListDCI03R18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Uci_OnPUSCH_ListDCI_0_3_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Uci_OnPUSCH_ListDCI_0_3_r18).Choice {
			case PUSCH_ConfigDCI_0_3_r18_Uci_OnPUSCH_ListDCI_0_3_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PUSCH_ConfigDCI_0_3_r18_Uci_OnPUSCH_ListDCI_0_3_r18_Setup:
				if err := (*ie.Uci_OnPUSCH_ListDCI_0_3_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Uci_OnPUSCH_ListDCI_0_3_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *PUSCH_ConfigDCI_0_3_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHConfigDCI03R18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. resourceAllocationDCI-0-3-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pUSCHConfigDCI03R18ResourceAllocationDCI03R18Constraints)
			if err != nil {
				return err
			}
			ie.ResourceAllocationDCI_0_3_r18 = &idx
		}
	}

	// 3. rbg-SizeDCI-0-3-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(pUSCHConfigDCI03R18RbgSizeDCI03R18Constraints)
			if err != nil {
				return err
			}
			ie.Rbg_SizeDCI_0_3_r18 = &idx
		}
	}

	// 4. resourceAllocationType1GranularityDCI-0-3-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(pUSCHConfigDCI03R18ResourceAllocationType1GranularityDCI03R18Constraints)
			if err != nil {
				return err
			}
			ie.ResourceAllocationType1GranularityDCI_0_3_r18 = &idx
		}
	}

	// 5. numberOfBitsForRV-DCI-0-3-r18: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(pUSCHConfigDCI03R18NumberOfBitsForRVDCI03R18Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfBitsForRV_DCI_0_3_r18 = &v
		}
	}

	// 6. harq-ProcessNumberSizeDCI-0-3-r18: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(pUSCHConfigDCI03R18HarqProcessNumberSizeDCI03R18Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_0_3_r18 = &v
		}
	}

	// 7. uci-OnPUSCH-ListDCI-0-3-r18: choice
	{
		if seq.IsComponentPresent(5) {
			ie.Uci_OnPUSCH_ListDCI_0_3_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UCI_OnPUSCH_ListDCI_0_1_r16
			}{}
			choiceDec := d.NewChoiceDecoder(pUSCH_ConfigDCI_0_3_r18UciOnPUSCHListDCI03R18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Uci_OnPUSCH_ListDCI_0_3_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PUSCH_ConfigDCI_0_3_r18_Uci_OnPUSCH_ListDCI_0_3_r18_Release:
				(*ie.Uci_OnPUSCH_ListDCI_0_3_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_ConfigDCI_0_3_r18_Uci_OnPUSCH_ListDCI_0_3_r18_Setup:
				(*ie.Uci_OnPUSCH_ListDCI_0_3_r18).Setup = new(UCI_OnPUSCH_ListDCI_0_1_r16)
				if err := (*ie.Uci_OnPUSCH_ListDCI_0_3_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
