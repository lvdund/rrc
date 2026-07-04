// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TRS-ResourceSet-r18 (line 4461).

var tRSResourceSetR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "powerControlOffsetSS-r18"},
		{Name: "scramblingID-Info-r18"},
		{Name: "firstOFDMSymbolInTimeDomain-r18"},
		{Name: "startingRB-r18"},
		{Name: "nrofRBs-r18"},
		{Name: "ssb-Index-r18"},
		{Name: "periodicityAndOffset-r18"},
		{Name: "frequencyDomainAllocation-r18"},
		{Name: "indBitID-r18"},
		{Name: "nrofResources-r18"},
	},
}

const (
	TRS_ResourceSet_r18_PowerControlOffsetSS_r18_Db_3 = 0
	TRS_ResourceSet_r18_PowerControlOffsetSS_r18_Db0  = 1
	TRS_ResourceSet_r18_PowerControlOffsetSS_r18_Db3  = 2
	TRS_ResourceSet_r18_PowerControlOffsetSS_r18_Db6  = 3
)

var tRSResourceSetR18PowerControlOffsetSSR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TRS_ResourceSet_r18_PowerControlOffsetSS_r18_Db_3, TRS_ResourceSet_r18_PowerControlOffsetSS_r18_Db0, TRS_ResourceSet_r18_PowerControlOffsetSS_r18_Db3, TRS_ResourceSet_r18_PowerControlOffsetSS_r18_Db6},
}

var tRS_ResourceSet_r18ScramblingIDInfoR18Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "scramblingIDforCommon-r18"},
		{Name: "scramblingIDperResourceListWith2-r18"},
		{Name: "scramblingIDperResourceListWith4-r18"},
	},
}

const (
	TRS_ResourceSet_r18_ScramblingID_Info_r18_ScramblingIDforCommon_r18            = 0
	TRS_ResourceSet_r18_ScramblingID_Info_r18_ScramblingIDperResourceListWith2_r18 = 1
	TRS_ResourceSet_r18_ScramblingID_Info_r18_ScramblingIDperResourceListWith4_r18 = 2
)

var tRSResourceSetR18FirstOFDMSymbolInTimeDomainR18Constraints = per.Constrained(0, 9)

var tRSResourceSetR18StartingRBR18Constraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

var tRSResourceSetR18NrofRBsR18Constraints = per.Constrained(24, common.MaxNrofPhysicalResourceBlocksPlus1)

var tRS_ResourceSet_r18PeriodicityAndOffsetR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "slots10"},
		{Name: "slots20"},
		{Name: "slots40"},
		{Name: "slots80"},
		{Name: "slots160"},
		{Name: "slots320"},
		{Name: "slots640"},
	},
}

const (
	TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots10  = 0
	TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots20  = 1
	TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots40  = 2
	TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots80  = 3
	TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots160 = 4
	TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots320 = 5
	TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots640 = 6
)

var tRSResourceSetR18FrequencyDomainAllocationR18Constraints = per.FixedSize(4)

var tRSResourceSetR18IndBitIDR18Constraints = per.Constrained(0, 5)

const (
	TRS_ResourceSet_r18_NrofResources_r18_N2 = 0
	TRS_ResourceSet_r18_NrofResources_r18_N4 = 1
)

var tRSResourceSetR18NrofResourcesR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{TRS_ResourceSet_r18_NrofResources_r18_N2, TRS_ResourceSet_r18_NrofResources_r18_N4},
}

var tRSResourceSetR18ScramblingIDInfoR18ScramblingIDperResourceListWith2R18Constraints = per.FixedSize(2)

var tRSResourceSetR18ScramblingIDInfoR18ScramblingIDperResourceListWith4R18Constraints = per.FixedSize(4)

type TRS_ResourceSet_r18 struct {
	PowerControlOffsetSS_r18 int64
	ScramblingID_Info_r18    struct {
		Choice                               int
		ScramblingIDforCommon_r18            *ScramblingId
		ScramblingIDperResourceListWith2_r18 *[]ScramblingId
		ScramblingIDperResourceListWith4_r18 *[]ScramblingId
	}
	FirstOFDMSymbolInTimeDomain_r18 int64
	StartingRB_r18                  int64
	NrofRBs_r18                     int64
	Ssb_Index_r18                   SSB_Index
	PeriodicityAndOffset_r18        struct {
		Choice   int
		Slots10  *int64
		Slots20  *int64
		Slots40  *int64
		Slots80  *int64
		Slots160 *int64
		Slots320 *int64
		Slots640 *int64
	}
	FrequencyDomainAllocation_r18 per.BitString
	IndBitID_r18                  int64
	NrofResources_r18             int64
}

func (ie *TRS_ResourceSet_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(tRSResourceSetR18Constraints)
	_ = seq

	// 1. powerControlOffsetSS-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.PowerControlOffsetSS_r18, tRSResourceSetR18PowerControlOffsetSSR18Constraints); err != nil {
			return err
		}
	}

	// 2. scramblingID-Info-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(tRS_ResourceSet_r18ScramblingIDInfoR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ScramblingID_Info_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ScramblingID_Info_r18.Choice {
		case TRS_ResourceSet_r18_ScramblingID_Info_r18_ScramblingIDforCommon_r18:
			if err := ie.ScramblingID_Info_r18.ScramblingIDforCommon_r18.Encode(e); err != nil {
				return err
			}
		case TRS_ResourceSet_r18_ScramblingID_Info_r18_ScramblingIDperResourceListWith2_r18:
			seqOf := e.NewSequenceOfEncoder(tRSResourceSetR18ScramblingIDInfoR18ScramblingIDperResourceListWith2R18Constraints)
			if err := seqOf.EncodeLength(int64(len((*ie.ScramblingID_Info_r18.ScramblingIDperResourceListWith2_r18)))); err != nil {
				return err
			}
			for i := range *ie.ScramblingID_Info_r18.ScramblingIDperResourceListWith2_r18 {
				if err := (*ie.ScramblingID_Info_r18.ScramblingIDperResourceListWith2_r18)[i].Encode(e); err != nil {
					return err
				}
			}
		case TRS_ResourceSet_r18_ScramblingID_Info_r18_ScramblingIDperResourceListWith4_r18:
			seqOf := e.NewSequenceOfEncoder(tRSResourceSetR18ScramblingIDInfoR18ScramblingIDperResourceListWith4R18Constraints)
			if err := seqOf.EncodeLength(int64(len((*ie.ScramblingID_Info_r18.ScramblingIDperResourceListWith4_r18)))); err != nil {
				return err
			}
			for i := range *ie.ScramblingID_Info_r18.ScramblingIDperResourceListWith4_r18 {
				if err := (*ie.ScramblingID_Info_r18.ScramblingIDperResourceListWith4_r18)[i].Encode(e); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ScramblingID_Info_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 3. firstOFDMSymbolInTimeDomain-r18: integer
	{
		if err := e.EncodeInteger(ie.FirstOFDMSymbolInTimeDomain_r18, tRSResourceSetR18FirstOFDMSymbolInTimeDomainR18Constraints); err != nil {
			return err
		}
	}

	// 4. startingRB-r18: integer
	{
		if err := e.EncodeInteger(ie.StartingRB_r18, tRSResourceSetR18StartingRBR18Constraints); err != nil {
			return err
		}
	}

	// 5. nrofRBs-r18: integer
	{
		if err := e.EncodeInteger(ie.NrofRBs_r18, tRSResourceSetR18NrofRBsR18Constraints); err != nil {
			return err
		}
	}

	// 6. ssb-Index-r18: ref
	{
		if err := ie.Ssb_Index_r18.Encode(e); err != nil {
			return err
		}
	}

	// 7. periodicityAndOffset-r18: choice
	{
		choiceEnc := e.NewChoiceEncoder(tRS_ResourceSet_r18PeriodicityAndOffsetR18Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.PeriodicityAndOffset_r18.Choice), false, nil); err != nil {
			return err
		}
		switch ie.PeriodicityAndOffset_r18.Choice {
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots10:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r18.Slots10), per.Constrained(0, 9)); err != nil {
				return err
			}
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots20:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r18.Slots20), per.Constrained(0, 19)); err != nil {
				return err
			}
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots40:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r18.Slots40), per.Constrained(0, 39)); err != nil {
				return err
			}
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots80:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r18.Slots80), per.Constrained(0, 79)); err != nil {
				return err
			}
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots160:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r18.Slots160), per.Constrained(0, 159)); err != nil {
				return err
			}
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots320:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r18.Slots320), per.Constrained(0, 319)); err != nil {
				return err
			}
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots640:
			if err := e.EncodeInteger((*ie.PeriodicityAndOffset_r18.Slots640), per.Constrained(0, 639)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.PeriodicityAndOffset_r18.Choice), Constraint: "index out of range"}
		}
	}

	// 8. frequencyDomainAllocation-r18: bit-string
	{
		if err := e.EncodeBitString(ie.FrequencyDomainAllocation_r18, tRSResourceSetR18FrequencyDomainAllocationR18Constraints); err != nil {
			return err
		}
	}

	// 9. indBitID-r18: integer
	{
		if err := e.EncodeInteger(ie.IndBitID_r18, tRSResourceSetR18IndBitIDR18Constraints); err != nil {
			return err
		}
	}

	// 10. nrofResources-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.NrofResources_r18, tRSResourceSetR18NrofResourcesR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *TRS_ResourceSet_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(tRSResourceSetR18Constraints)
	_ = seq

	// 1. powerControlOffsetSS-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(tRSResourceSetR18PowerControlOffsetSSR18Constraints)
		if err != nil {
			return err
		}
		ie.PowerControlOffsetSS_r18 = v0
	}

	// 2. scramblingID-Info-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(tRS_ResourceSet_r18ScramblingIDInfoR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ScramblingID_Info_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case TRS_ResourceSet_r18_ScramblingID_Info_r18_ScramblingIDforCommon_r18:
			ie.ScramblingID_Info_r18.ScramblingIDforCommon_r18 = new(ScramblingId)
			if err := ie.ScramblingID_Info_r18.ScramblingIDforCommon_r18.Decode(d); err != nil {
				return err
			}
		case TRS_ResourceSet_r18_ScramblingID_Info_r18_ScramblingIDperResourceListWith2_r18:
			ie.ScramblingID_Info_r18.ScramblingIDperResourceListWith2_r18 = new([]ScramblingId)
			seqOf := d.NewSequenceOfDecoder(tRSResourceSetR18ScramblingIDInfoR18ScramblingIDperResourceListWith2R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*ie.ScramblingID_Info_r18.ScramblingIDperResourceListWith2_r18) = make([]ScramblingId, n)
			for i := int64(0); i < n; i++ {
				if err := (*ie.ScramblingID_Info_r18.ScramblingIDperResourceListWith2_r18)[i].Decode(d); err != nil {
					return err
				}
			}
		case TRS_ResourceSet_r18_ScramblingID_Info_r18_ScramblingIDperResourceListWith4_r18:
			ie.ScramblingID_Info_r18.ScramblingIDperResourceListWith4_r18 = new([]ScramblingId)
			seqOf := d.NewSequenceOfDecoder(tRSResourceSetR18ScramblingIDInfoR18ScramblingIDperResourceListWith4R18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			(*ie.ScramblingID_Info_r18.ScramblingIDperResourceListWith4_r18) = make([]ScramblingId, n)
			for i := int64(0); i < n; i++ {
				if err := (*ie.ScramblingID_Info_r18.ScramblingIDperResourceListWith4_r18)[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. firstOFDMSymbolInTimeDomain-r18: integer
	{
		v2, err := d.DecodeInteger(tRSResourceSetR18FirstOFDMSymbolInTimeDomainR18Constraints)
		if err != nil {
			return err
		}
		ie.FirstOFDMSymbolInTimeDomain_r18 = v2
	}

	// 4. startingRB-r18: integer
	{
		v3, err := d.DecodeInteger(tRSResourceSetR18StartingRBR18Constraints)
		if err != nil {
			return err
		}
		ie.StartingRB_r18 = v3
	}

	// 5. nrofRBs-r18: integer
	{
		v4, err := d.DecodeInteger(tRSResourceSetR18NrofRBsR18Constraints)
		if err != nil {
			return err
		}
		ie.NrofRBs_r18 = v4
	}

	// 6. ssb-Index-r18: ref
	{
		if err := ie.Ssb_Index_r18.Decode(d); err != nil {
			return err
		}
	}

	// 7. periodicityAndOffset-r18: choice
	{
		choiceDec := d.NewChoiceDecoder(tRS_ResourceSet_r18PeriodicityAndOffsetR18Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.PeriodicityAndOffset_r18.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots10:
			ie.PeriodicityAndOffset_r18.Slots10 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 9))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r18.Slots10) = v
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots20:
			ie.PeriodicityAndOffset_r18.Slots20 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 19))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r18.Slots20) = v
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots40:
			ie.PeriodicityAndOffset_r18.Slots40 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 39))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r18.Slots40) = v
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots80:
			ie.PeriodicityAndOffset_r18.Slots80 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 79))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r18.Slots80) = v
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots160:
			ie.PeriodicityAndOffset_r18.Slots160 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 159))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r18.Slots160) = v
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots320:
			ie.PeriodicityAndOffset_r18.Slots320 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 319))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r18.Slots320) = v
		case TRS_ResourceSet_r18_PeriodicityAndOffset_r18_Slots640:
			ie.PeriodicityAndOffset_r18.Slots640 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 639))
			if err != nil {
				return err
			}
			(*ie.PeriodicityAndOffset_r18.Slots640) = v
		}
	}

	// 8. frequencyDomainAllocation-r18: bit-string
	{
		v7, err := d.DecodeBitString(tRSResourceSetR18FrequencyDomainAllocationR18Constraints)
		if err != nil {
			return err
		}
		ie.FrequencyDomainAllocation_r18 = v7
	}

	// 9. indBitID-r18: integer
	{
		v8, err := d.DecodeInteger(tRSResourceSetR18IndBitIDR18Constraints)
		if err != nil {
			return err
		}
		ie.IndBitID_r18 = v8
	}

	// 10. nrofResources-r18: enumerated
	{
		v9, err := d.DecodeEnumerated(tRSResourceSetR18NrofResourcesR18Constraints)
		if err != nil {
			return err
		}
		ie.NrofResources_r18 = v9
	}

	return nil
}
