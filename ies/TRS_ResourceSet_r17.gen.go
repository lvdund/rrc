// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TRS-ResourceSet-r17 (line 4421).

var tRSResourceSetR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "powerControlOffsetSS-r17"},
		{Name: "scramblingID-Info-r17"},
		{Name: "firstOFDMSymbolInTimeDomain-r17"},
		{Name: "startingRB-r17"},
		{Name: "nrofRBs-r17"},
		{Name: "ssb-Index-r17"},
		{Name: "periodicityAndOffset-r17"},
		{Name: "frequencyDomainAllocation-r17"},
		{Name: "indBitID-r17"},
		{Name: "nrofResources-r17"},
	},
}

const (
	TRS_ResourceSet_r17_PowerControlOffsetSS_r17_Db_3 = 0
	TRS_ResourceSet_r17_PowerControlOffsetSS_r17_Db0  = 1
	TRS_ResourceSet_r17_PowerControlOffsetSS_r17_Db3  = 2
	TRS_ResourceSet_r17_PowerControlOffsetSS_r17_Db6  = 3
)

var tRSResourceSetR17PowerControlOffsetSSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TRS_ResourceSet_r17_PowerControlOffsetSS_r17_Db_3, TRS_ResourceSet_r17_PowerControlOffsetSS_r17_Db0, TRS_ResourceSet_r17_PowerControlOffsetSS_r17_Db3, TRS_ResourceSet_r17_PowerControlOffsetSS_r17_Db6},
}

var tRS_ResourceSet_r17ScramblingIDInfoR17Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "scramblingIDforCommon-r17"},
		{Name: "scramblingIDperResourceListWith2-r17"},
		{Name: "scramblingIDperResourceListWith4-r17"},
	},
}

const (
	TRS_ResourceSet_r17_ScramblingID_Info_r17_ScramblingIDforCommon_r17            = 0
	TRS_ResourceSet_r17_ScramblingID_Info_r17_ScramblingIDperResourceListWith2_r17 = 1
	TRS_ResourceSet_r17_ScramblingID_Info_r17_ScramblingIDperResourceListWith4_r17 = 2
)

var tRSResourceSetR17FirstOFDMSymbolInTimeDomainR17Constraints = per.Constrained(0, 9)

var tRSResourceSetR17StartingRBR17Constraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

var tRSResourceSetR17NrofRBsR17Constraints = per.Constrained(24, common.MaxNrofPhysicalResourceBlocksPlus1)

var tRS_ResourceSet_r17PeriodicityAndOffsetR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "slots10"},
		{Name: "slots20"},
		{Name: "slots40"},
		{Name: "slots80"},
	},
}

const (
	TRS_ResourceSet_r17_PeriodicityAndOffset_r17_Slots10 = 0
	TRS_ResourceSet_r17_PeriodicityAndOffset_r17_Slots20 = 1
	TRS_ResourceSet_r17_PeriodicityAndOffset_r17_Slots40 = 2
	TRS_ResourceSet_r17_PeriodicityAndOffset_r17_Slots80 = 3
)

var tRSResourceSetR17FrequencyDomainAllocationR17Constraints = per.FixedSize(4)

var tRSResourceSetR17IndBitIDR17Constraints = per.Constrained(0, 5)

const (
	TRS_ResourceSet_r17_NrofResources_r17_N2 = 0
	TRS_ResourceSet_r17_NrofResources_r17_N4 = 1
)

var tRSResourceSetR17NrofResourcesR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TRS_ResourceSet_r17_NrofResources_r17_N2, TRS_ResourceSet_r17_NrofResources_r17_N4},
}

var tRSResourceSetR17ScramblingIDInfoR17ScramblingIDperResourceListWith2R17Constraints = per.FixedSize(2)

var tRSResourceSetR17ScramblingIDInfoR17ScramblingIDperResourceListWith4R17Constraints = per.FixedSize(4)

type TRS_ResourceSet_r17 struct {
	PowerControlOffsetSS_r17 int64
	ScramblingID_Info_r17    struct {
		Choice                               int
		ScramblingIDforCommon_r17            *ScramblingId
		ScramblingIDperResourceListWith2_r17 *[]ScramblingId
		ScramblingIDperResourceListWith4_r17 *[]ScramblingId
	}
	FirstOFDMSymbolInTimeDomain_r17 int64
	StartingRB_r17                  int64
	NrofRBs_r17                     int64
	Ssb_Index_r17                   SSB_Index
	PeriodicityAndOffset_r17        struct {
		Choice  int
		Slots10 *int64
		Slots20 *int64
		Slots40 *int64
		Slots80 *int64
	}
	FrequencyDomainAllocation_r17 per.BitString
	IndBitID_r17                  int64
	NrofResources_r17             int64
}

func (ie *TRS_ResourceSet_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tRSResourceSetR17Constraints)
	_ = seq

	// 1. powerControlOffsetSS-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.PowerControlOffsetSS_r17, tRSResourceSetR17PowerControlOffsetSSR17Constraints); err != nil {
			return err
		}
	}

	// 2. scramblingID-Info-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(tRS_ResourceSet_r17ScramblingIDInfoR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ScramblingID_Info_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ScramblingID_Info_r17.Choice {
		case TRS_ResourceSet_r17_ScramblingID_Info_r17_ScramblingIDforCommon_r17:
			if err := ie.ScramblingID_Info_r17.ScramblingIDforCommon_r17.Encode(e); err != nil {
				return err
			}
		case TRS_ResourceSet_r17_ScramblingID_Info_r17_ScramblingIDperResourceListWith2_r17:
			seqOf := e.NewSequenceOfEncoder(tRSResourceSetR17ScramblingIDInfoR17ScramblingIDperResourceListWith2R17Constraints)
			if err := seqOf.EncodeLength(int64(len((*ie.ScramblingID_Info_r17.ScramblingIDperResourceListWith2_r17)))); err != nil {
				return err
			}
			for i := range *ie.ScramblingID_Info_r17.ScramblingIDperResourceListWith2_r17 {
				if err := (*ie.ScramblingID_Info_r17.ScramblingIDperResourceListWith2_r17)[i].Encode(e); err != nil {
					return err
				}
			}
		case TRS_ResourceSet_r17_ScramblingID_Info_r17_ScramblingIDperResourceListWith4_r17:
			seqOf := e.NewSequenceOfEncoder(tRSResourceSetR17ScramblingIDInfoR17ScramblingIDperResourceListWith4R17Constraints)
			if err := seqOf.EncodeLength(int64(len((*ie.ScramblingID_Info_r17.ScramblingIDperResourceListWith4_r17)))); err != nil {
				return err
			}
			for i := range *ie.ScramblingID_Info_r17.ScramblingIDperResourceListWith4_r17 {
				if err := (*ie.ScramblingID_Info_r17.ScramblingIDperResourceListWith4_r17)[i].Encode(e); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ScramblingID_Info_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 3. firstOFDMSymbolInTimeDomain-r17: integer
	{
		if err := e.EncodeInteger(ie.FirstOFDMSymbolInTimeDomain_r17, tRSResourceSetR17FirstOFDMSymbolInTimeDomainR17Constraints); err != nil {
			return err
		}
	}

	// 4. startingRB-r17: integer
	{
		if err := e.EncodeInteger(ie.StartingRB_r17, tRSResourceSetR17StartingRBR17Constraints); err != nil {
			return err
		}
	}

	// 5. nrofRBs-r17: integer
	{
		if err := e.EncodeInteger(ie.NrofRBs_r17, tRSResourceSetR17NrofRBsR17Constraints); err != nil {
			return err
		}
	}

	// 6. ssb-Index-r17: ref
	{
		if err := ie.Ssb_Index_r17.Encode(e); err != nil {
			return err
		}
	}

	// 7. periodicityAndOffset-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(tRS_ResourceSet_r17PeriodicityAndOffsetR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.PeriodicityAndOffset_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.PeriodicityAndOffset_r17.Choice {
		case TRS_ResourceSet_r17_PeriodicityAndOffset_r17_Slots10:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r17.Slots10), per.Constrained(0, 9)); err != nil {
				return err
			}
		case TRS_ResourceSet_r17_PeriodicityAndOffset_r17_Slots20:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r17.Slots20), per.Constrained(0, 19)); err != nil {
				return err
			}
		case TRS_ResourceSet_r17_PeriodicityAndOffset_r17_Slots40:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r17.Slots40), per.Constrained(0, 39)); err != nil {
				return err
			}
		case TRS_ResourceSet_r17_PeriodicityAndOffset_r17_Slots80:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r17.Slots80), per.Constrained(0, 79)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.PeriodicityAndOffset_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 8. frequencyDomainAllocation-r17: bit-string
	{
		if err := e.EncodeBitString(ie.FrequencyDomainAllocation_r17, tRSResourceSetR17FrequencyDomainAllocationR17Constraints); err != nil {
			return err
		}
	}

	// 9. indBitID-r17: integer
	{
		if err := e.EncodeInteger(ie.IndBitID_r17, tRSResourceSetR17IndBitIDR17Constraints); err != nil {
			return err
		}
	}

	// 10. nrofResources-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.NrofResources_r17, tRSResourceSetR17NrofResourcesR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *TRS_ResourceSet_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tRSResourceSetR17Constraints)
	_ = seq

	// 1. powerControlOffsetSS-r17: enumerated
	{
		v0, err := d.DecodeEnumerated(tRSResourceSetR17PowerControlOffsetSSR17Constraints)
		if err != nil {
			return err
		}
		ie.PowerControlOffsetSS_r17 = v0
	}

	// 2. scramblingID-Info-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(tRS_ResourceSet_r17ScramblingIDInfoR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ScramblingID_Info_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case TRS_ResourceSet_r17_ScramblingID_Info_r17_ScramblingIDforCommon_r17:
			ie.ScramblingID_Info_r17.ScramblingIDforCommon_r17 = new(ScramblingId)
			if err := ie.ScramblingID_Info_r17.ScramblingIDforCommon_r17.Decode(d); err != nil {
				return err
			}
		case TRS_ResourceSet_r17_ScramblingID_Info_r17_ScramblingIDperResourceListWith2_r17:
			ie.ScramblingID_Info_r17.ScramblingIDperResourceListWith2_r17 = new([]ScramblingId)
			seqOf := d.NewSequenceOfDecoder(tRSResourceSetR17ScramblingIDInfoR17ScramblingIDperResourceListWith2R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*ie.ScramblingID_Info_r17.ScramblingIDperResourceListWith2_r17) = make([]ScramblingId, n)
			for i := int64(0); i < n; i++ {
				if err := (*ie.ScramblingID_Info_r17.ScramblingIDperResourceListWith2_r17)[i].Decode(d); err != nil {
					return err
				}
			}
		case TRS_ResourceSet_r17_ScramblingID_Info_r17_ScramblingIDperResourceListWith4_r17:
			ie.ScramblingID_Info_r17.ScramblingIDperResourceListWith4_r17 = new([]ScramblingId)
			seqOf := d.NewSequenceOfDecoder(tRSResourceSetR17ScramblingIDInfoR17ScramblingIDperResourceListWith4R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*ie.ScramblingID_Info_r17.ScramblingIDperResourceListWith4_r17) = make([]ScramblingId, n)
			for i := int64(0); i < n; i++ {
				if err := (*ie.ScramblingID_Info_r17.ScramblingIDperResourceListWith4_r17)[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. firstOFDMSymbolInTimeDomain-r17: integer
	{
		v2, err := d.DecodeInteger(tRSResourceSetR17FirstOFDMSymbolInTimeDomainR17Constraints)
		if err != nil {
			return err
		}
		ie.FirstOFDMSymbolInTimeDomain_r17 = v2
	}

	// 4. startingRB-r17: integer
	{
		v3, err := d.DecodeInteger(tRSResourceSetR17StartingRBR17Constraints)
		if err != nil {
			return err
		}
		ie.StartingRB_r17 = v3
	}

	// 5. nrofRBs-r17: integer
	{
		v4, err := d.DecodeInteger(tRSResourceSetR17NrofRBsR17Constraints)
		if err != nil {
			return err
		}
		ie.NrofRBs_r17 = v4
	}

	// 6. ssb-Index-r17: ref
	{
		if err := ie.Ssb_Index_r17.Decode(d); err != nil {
			return err
		}
	}

	// 7. periodicityAndOffset-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(tRS_ResourceSet_r17PeriodicityAndOffsetR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.PeriodicityAndOffset_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case TRS_ResourceSet_r17_PeriodicityAndOffset_r17_Slots10:
			ie.PeriodicityAndOffset_r17.Slots10 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r17.Slots10) = v
		case TRS_ResourceSet_r17_PeriodicityAndOffset_r17_Slots20:
			ie.PeriodicityAndOffset_r17.Slots20 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r17.Slots20) = v
		case TRS_ResourceSet_r17_PeriodicityAndOffset_r17_Slots40:
			ie.PeriodicityAndOffset_r17.Slots40 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r17.Slots40) = v
		case TRS_ResourceSet_r17_PeriodicityAndOffset_r17_Slots80:
			ie.PeriodicityAndOffset_r17.Slots80 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r17.Slots80) = v
		}
	}

	// 8. frequencyDomainAllocation-r17: bit-string
	{
		v7, err := d.DecodeBitString(tRSResourceSetR17FrequencyDomainAllocationR17Constraints)
		if err != nil {
			return err
		}
		ie.FrequencyDomainAllocation_r17 = v7
	}

	// 9. indBitID-r17: integer
	{
		v8, err := d.DecodeInteger(tRSResourceSetR17IndBitIDR17Constraints)
		if err != nil {
			return err
		}
		ie.IndBitID_r17 = v8
	}

	// 10. nrofResources-r17: enumerated
	{
		v9, err := d.DecodeEnumerated(tRSResourceSetR17NrofResourcesR17Constraints)
		if err != nil {
			return err
		}
		ie.NrofResources_r17 = v9
	}

	return nil
}
