// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SearchSpaceExt-v1700 (line 14479).

var searchSpaceExtV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "monitoringSlotPeriodicityAndOffset-v1710", Optional: true},
		{Name: "monitoringSlotsWithinSlotGroup-r17", Optional: true},
		{Name: "duration-r17", Optional: true},
		{Name: "searchSpaceType-r17", Optional: true},
		{Name: "searchSpaceGroupIdList-r17", Optional: true},
		{Name: "searchSpaceLinkingId-r17", Optional: true},
	},
}

var searchSpaceExt_v1700MonitoringSlotPeriodicityAndOffsetV1710Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sl32"},
		{Name: "sl64"},
		{Name: "sl128"},
		{Name: "sl5120"},
		{Name: "sl10240"},
		{Name: "sl20480"},
	},
}

const (
	SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl32    = 0
	SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl64    = 1
	SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl128   = 2
	SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl5120  = 3
	SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl10240 = 4
	SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl20480 = 5
)

var searchSpaceExt_v1700MonitoringSlotsWithinSlotGroupR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "slotGroupLength4-r17"},
		{Name: "slotGroupLength8-r17"},
	},
}

const (
	SearchSpaceExt_v1700_MonitoringSlotsWithinSlotGroup_r17_SlotGroupLength4_r17 = 0
	SearchSpaceExt_v1700_MonitoringSlotsWithinSlotGroup_r17_SlotGroupLength8_r17 = 1
)

var searchSpaceExtV1700DurationR17Constraints = per.Constrained(4, 20476)

var searchSpaceExtV1700SearchSpaceGroupIdListR17Constraints = per.SizeRange(1, 3)

var searchSpaceExtV1700SearchSpaceLinkingIdR17Constraints = per.Constrained(0, common.MaxNrofSearchSpacesLinks_1_r17)

var searchSpaceExtV1700SearchSpaceTypeR17CommonR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dci-Format4-0-r17", Optional: true},
		{Name: "dci-Format4-1-r17", Optional: true},
		{Name: "dci-Format4-2-r17", Optional: true},
		{Name: "dci-Format4-1-AndFormat4-2-r17", Optional: true},
		{Name: "dci-Format2-7-r17", Optional: true},
	},
}

var searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat40R17Constraints = per.SequenceConstraints{
	Extensible: true,
}

var searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41R17Constraints = per.SequenceConstraints{
	Extensible: true,
}

var searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat42R17Constraints = per.SequenceConstraints{
	Extensible: true,
}

var searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41AndFormat42R17Constraints = per.SequenceConstraints{
	Extensible: true,
}

var searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofCandidates-PEI-r17"},
	},
}

var searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "aggregationLevel4-r17", Optional: true},
		{Name: "aggregationLevel8-r17", Optional: true},
		{Name: "aggregationLevel16-r17", Optional: true},
	},
}

const (
	SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel4_r17_N0 = 0
	SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel4_r17_N1 = 1
	SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel4_r17_N2 = 2
	SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel4_r17_N3 = 3
	SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel4_r17_N4 = 4
)

var searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17AggregationLevel4R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel4_r17_N0, SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel4_r17_N1, SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel4_r17_N2, SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel4_r17_N3, SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel4_r17_N4},
}

const (
	SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel8_r17_N0 = 0
	SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel8_r17_N1 = 1
	SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel8_r17_N2 = 2
)

var searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17AggregationLevel8R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel8_r17_N0, SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel8_r17_N1, SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel8_r17_N2},
}

const (
	SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel16_r17_N0 = 0
	SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel16_r17_N1 = 1
)

var searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17AggregationLevel16R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel16_r17_N0, SearchSpaceExt_v1700_SearchSpaceType_r17_Common_r17_Dci_Format2_7_r17_NrofCandidates_PEI_r17_AggregationLevel16_r17_N1},
}

type SearchSpaceExt_v1700 struct {
	MonitoringSlotPeriodicityAndOffset_v1710 *struct {
		Choice  int
		Sl32    *int64
		Sl64    *int64
		Sl128   *int64
		Sl5120  *int64
		Sl10240 *int64
		Sl20480 *int64
	}
	MonitoringSlotsWithinSlotGroup_r17 *struct {
		Choice               int
		SlotGroupLength4_r17 *per.BitString
		SlotGroupLength8_r17 *per.BitString
	}
	Duration_r17        *int64
	SearchSpaceType_r17 *struct {
		Common_r17 struct {
			Dci_Format4_0_r17              *struct{}
			Dci_Format4_1_r17              *struct{}
			Dci_Format4_2_r17              *struct{}
			Dci_Format4_1_AndFormat4_2_r17 *struct{}
			Dci_Format2_7_r17              *struct {
				NrofCandidates_PEI_r17 struct {
					AggregationLevel4_r17  *int64
					AggregationLevel8_r17  *int64
					AggregationLevel16_r17 *int64
				}
			}
		}
	}
	SearchSpaceGroupIdList_r17 []int64
	SearchSpaceLinkingId_r17   *int64
}

func (ie *SearchSpaceExt_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(searchSpaceExtV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MonitoringSlotPeriodicityAndOffset_v1710 != nil, ie.MonitoringSlotsWithinSlotGroup_r17 != nil, ie.Duration_r17 != nil, ie.SearchSpaceType_r17 != nil, ie.SearchSpaceGroupIdList_r17 != nil, ie.SearchSpaceLinkingId_r17 != nil}); err != nil {
		return err
	}

	// 2. monitoringSlotPeriodicityAndOffset-v1710: choice
	{
		if ie.MonitoringSlotPeriodicityAndOffset_v1710 != nil {
			choiceEnc := e.NewChoiceEncoder(searchSpaceExt_v1700MonitoringSlotPeriodicityAndOffsetV1710Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.MonitoringSlotPeriodicityAndOffset_v1710).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.MonitoringSlotPeriodicityAndOffset_v1710).Choice {
			case SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl32:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl32), per.Constrained(0, 31)); err != nil {
					return err
				}
			case SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl64:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl64), per.Constrained(0, 63)); err != nil {
					return err
				}
			case SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl128:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl128), per.Constrained(0, 127)); err != nil {
					return err
				}
			case SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl5120:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl5120), per.Constrained(0, 5119)); err != nil {
					return err
				}
			case SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl10240:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl10240), per.Constrained(0, 10239)); err != nil {
					return err
				}
			case SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl20480:
				if err := e.EncodeInteger((*(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl20480), per.Constrained(0, 20479)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.MonitoringSlotPeriodicityAndOffset_v1710).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. monitoringSlotsWithinSlotGroup-r17: choice
	{
		if ie.MonitoringSlotsWithinSlotGroup_r17 != nil {
			choiceEnc := e.NewChoiceEncoder(searchSpaceExt_v1700MonitoringSlotsWithinSlotGroupR17Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.MonitoringSlotsWithinSlotGroup_r17).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.MonitoringSlotsWithinSlotGroup_r17).Choice {
			case SearchSpaceExt_v1700_MonitoringSlotsWithinSlotGroup_r17_SlotGroupLength4_r17:
				if err := e.EncodeBitString((*(*ie.MonitoringSlotsWithinSlotGroup_r17).SlotGroupLength4_r17), per.FixedSize(4)); err != nil {
					return err
				}
			case SearchSpaceExt_v1700_MonitoringSlotsWithinSlotGroup_r17_SlotGroupLength8_r17:
				if err := e.EncodeBitString((*(*ie.MonitoringSlotsWithinSlotGroup_r17).SlotGroupLength8_r17), per.FixedSize(8)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.MonitoringSlotsWithinSlotGroup_r17).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 4. duration-r17: integer
	{
		if ie.Duration_r17 != nil {
			if err := e.EncodeInteger(*ie.Duration_r17, searchSpaceExtV1700DurationR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. searchSpaceType-r17: sequence
	{
		if ie.SearchSpaceType_r17 != nil {
			c := ie.SearchSpaceType_r17
			{
				c := &c.Common_r17
				searchSpaceExtV1700SearchSpaceTypeR17CommonR17Seq := e.NewSequenceEncoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17Constraints)
				if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17Seq.EncodePreamble([]bool{c.Dci_Format4_0_r17 != nil, c.Dci_Format4_1_r17 != nil, c.Dci_Format4_2_r17 != nil, c.Dci_Format4_1_AndFormat4_2_r17 != nil, c.Dci_Format2_7_r17 != nil}); err != nil {
					return err
				}
				if c.Dci_Format4_0_r17 != nil {
					searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat40R17Seq := e.NewSequenceEncoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat40R17Constraints)
					if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat40R17Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
				}
				if c.Dci_Format4_1_r17 != nil {
					searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41R17Seq := e.NewSequenceEncoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41R17Constraints)
					if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41R17Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
				}
				if c.Dci_Format4_2_r17 != nil {
					searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat42R17Seq := e.NewSequenceEncoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat42R17Constraints)
					if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat42R17Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
				}
				if c.Dci_Format4_1_AndFormat4_2_r17 != nil {
					searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41AndFormat42R17Seq := e.NewSequenceEncoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41AndFormat42R17Constraints)
					if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41AndFormat42R17Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
				}
				if c.Dci_Format2_7_r17 != nil {
					c := (*c.Dci_Format2_7_r17)
					searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17Seq := e.NewSequenceEncoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17Constraints)
					if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
					{
						c := &c.NrofCandidates_PEI_r17
						searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17Seq := e.NewSequenceEncoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17Constraints)
						if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17Seq.EncodePreamble([]bool{c.AggregationLevel4_r17 != nil, c.AggregationLevel8_r17 != nil, c.AggregationLevel16_r17 != nil}); err != nil {
							return err
						}
						if c.AggregationLevel4_r17 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel4_r17), searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17AggregationLevel4R17Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel8_r17 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel8_r17), searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17AggregationLevel8R17Constraints); err != nil {
								return err
							}
						}
						if c.AggregationLevel16_r17 != nil {
							if err := e.EncodeEnumerated((*c.AggregationLevel16_r17), searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17AggregationLevel16R17Constraints); err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}

	// 6. searchSpaceGroupIdList-r17: sequence-of
	{
		if ie.SearchSpaceGroupIdList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(searchSpaceExtV1700SearchSpaceGroupIdListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SearchSpaceGroupIdList_r17))); err != nil {
				return err
			}
			for i := range ie.SearchSpaceGroupIdList_r17 {
				if err := e.EncodeInteger(ie.SearchSpaceGroupIdList_r17[i], per.Constrained(0, common.MaxNrofSearchSpaceGroups_1_r17)); err != nil {
					return err
				}
			}
		}
	}

	// 7. searchSpaceLinkingId-r17: integer
	{
		if ie.SearchSpaceLinkingId_r17 != nil {
			if err := e.EncodeInteger(*ie.SearchSpaceLinkingId_r17, searchSpaceExtV1700SearchSpaceLinkingIdR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SearchSpaceExt_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(searchSpaceExtV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. monitoringSlotPeriodicityAndOffset-v1710: choice
	{
		if seq.IsComponentPresent(0) {
			ie.MonitoringSlotPeriodicityAndOffset_v1710 = &struct {
				Choice  int
				Sl32    *int64
				Sl64    *int64
				Sl128   *int64
				Sl5120  *int64
				Sl10240 *int64
				Sl20480 *int64
			}{}
			choiceDec := d.NewChoiceDecoder(searchSpaceExt_v1700MonitoringSlotPeriodicityAndOffsetV1710Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl32:
				(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl32 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 31))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl32) = v
			case SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl64:
				(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl64 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 63))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl64) = v
			case SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl128:
				(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl128 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 127))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl128) = v
			case SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl5120:
				(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl5120 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 5119))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl5120) = v
			case SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl10240:
				(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl10240 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 10239))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl10240) = v
			case SearchSpaceExt_v1700_MonitoringSlotPeriodicityAndOffset_v1710_Sl20480:
				(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl20480 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 20479))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotPeriodicityAndOffset_v1710).Sl20480) = v
			}
		}
	}

	// 3. monitoringSlotsWithinSlotGroup-r17: choice
	{
		if seq.IsComponentPresent(1) {
			ie.MonitoringSlotsWithinSlotGroup_r17 = &struct {
				Choice               int
				SlotGroupLength4_r17 *per.BitString
				SlotGroupLength8_r17 *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(searchSpaceExt_v1700MonitoringSlotsWithinSlotGroupR17Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MonitoringSlotsWithinSlotGroup_r17).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SearchSpaceExt_v1700_MonitoringSlotsWithinSlotGroup_r17_SlotGroupLength4_r17:
				(*ie.MonitoringSlotsWithinSlotGroup_r17).SlotGroupLength4_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(4))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotsWithinSlotGroup_r17).SlotGroupLength4_r17) = v
			case SearchSpaceExt_v1700_MonitoringSlotsWithinSlotGroup_r17_SlotGroupLength8_r17:
				(*ie.MonitoringSlotsWithinSlotGroup_r17).SlotGroupLength8_r17 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*(*ie.MonitoringSlotsWithinSlotGroup_r17).SlotGroupLength8_r17) = v
			}
		}
	}

	// 4. duration-r17: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(searchSpaceExtV1700DurationR17Constraints)
			if err != nil {
				return err
			}
			ie.Duration_r17 = &v
		}
	}

	// 5. searchSpaceType-r17: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.SearchSpaceType_r17 = &struct {
				Common_r17 struct {
					Dci_Format4_0_r17              *struct{}
					Dci_Format4_1_r17              *struct{}
					Dci_Format4_2_r17              *struct{}
					Dci_Format4_1_AndFormat4_2_r17 *struct{}
					Dci_Format2_7_r17              *struct {
						NrofCandidates_PEI_r17 struct {
							AggregationLevel4_r17  *int64
							AggregationLevel8_r17  *int64
							AggregationLevel16_r17 *int64
						}
					}
				}
			}{}
			c := ie.SearchSpaceType_r17
			{
				c := &c.Common_r17
				searchSpaceExtV1700SearchSpaceTypeR17CommonR17Seq := d.NewSequenceDecoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17Constraints)
				if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17Seq.DecodePreamble(); err != nil {
					return err
				}
				if searchSpaceExtV1700SearchSpaceTypeR17CommonR17Seq.IsComponentPresent(0) {
					c.Dci_Format4_0_r17 = &struct{}{}
					searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat40R17Seq := d.NewSequenceDecoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat40R17Constraints)
					if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat40R17Seq.DecodeExtensionBit(); err != nil {
						return err
					}
				}
				if searchSpaceExtV1700SearchSpaceTypeR17CommonR17Seq.IsComponentPresent(1) {
					c.Dci_Format4_1_r17 = &struct{}{}
					searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41R17Seq := d.NewSequenceDecoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41R17Constraints)
					if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41R17Seq.DecodeExtensionBit(); err != nil {
						return err
					}
				}
				if searchSpaceExtV1700SearchSpaceTypeR17CommonR17Seq.IsComponentPresent(2) {
					c.Dci_Format4_2_r17 = &struct{}{}
					searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat42R17Seq := d.NewSequenceDecoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat42R17Constraints)
					if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat42R17Seq.DecodeExtensionBit(); err != nil {
						return err
					}
				}
				if searchSpaceExtV1700SearchSpaceTypeR17CommonR17Seq.IsComponentPresent(3) {
					c.Dci_Format4_1_AndFormat4_2_r17 = &struct{}{}
					searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41AndFormat42R17Seq := d.NewSequenceDecoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41AndFormat42R17Constraints)
					if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat41AndFormat42R17Seq.DecodeExtensionBit(); err != nil {
						return err
					}
				}
				if searchSpaceExtV1700SearchSpaceTypeR17CommonR17Seq.IsComponentPresent(4) {
					c.Dci_Format2_7_r17 = &struct {
						NrofCandidates_PEI_r17 struct {
							AggregationLevel4_r17  *int64
							AggregationLevel8_r17  *int64
							AggregationLevel16_r17 *int64
						}
					}{}
					c := (*c.Dci_Format2_7_r17)
					searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17Seq := d.NewSequenceDecoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17Constraints)
					if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17Seq.DecodeExtensionBit(); err != nil {
						return err
					}
					{
						c := &c.NrofCandidates_PEI_r17
						searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17Seq := d.NewSequenceDecoder(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17Constraints)
						if err := searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17Seq.DecodePreamble(); err != nil {
							return err
						}
						if searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17Seq.IsComponentPresent(0) {
							c.AggregationLevel4_r17 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17AggregationLevel4R17Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel4_r17) = v
						}
						if searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17Seq.IsComponentPresent(1) {
							c.AggregationLevel8_r17 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17AggregationLevel8R17Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel8_r17) = v
						}
						if searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17Seq.IsComponentPresent(2) {
							c.AggregationLevel16_r17 = new(int64)
							v, err := d.DecodeEnumerated(searchSpaceExtV1700SearchSpaceTypeR17CommonR17DciFormat27R17NrofCandidatesPEIR17AggregationLevel16R17Constraints)
							if err != nil {
								return err
							}
							(*c.AggregationLevel16_r17) = v
						}
					}
				}
			}
		}
	}

	// 6. searchSpaceGroupIdList-r17: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(searchSpaceExtV1700SearchSpaceGroupIdListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SearchSpaceGroupIdList_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofSearchSpaceGroups_1_r17))
				if err != nil {
					return err
				}
				ie.SearchSpaceGroupIdList_r17[i] = v
			}
		}
	}

	// 7. searchSpaceLinkingId-r17: integer
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeInteger(searchSpaceExtV1700SearchSpaceLinkingIdR17Constraints)
			if err != nil {
				return err
			}
			ie.SearchSpaceLinkingId_r17 = &v
		}
	}

	return nil
}
