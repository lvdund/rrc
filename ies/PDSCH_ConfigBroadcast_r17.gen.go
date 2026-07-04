// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PDSCH-ConfigBroadcast-r17 (line 28626).

var pDSCHConfigBroadcastR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdschConfigList-r17"},
		{Name: "pdsch-TimeDomainAllocationList-r17", Optional: true},
		{Name: "rateMatchPatternToAddModList-r17", Optional: true},
		{Name: "lte-CRS-ToMatchAround-r17", Optional: true},
		{Name: "mcs-Table-r17", Optional: true},
		{Name: "xOverhead-r17", Optional: true},
	},
}

var pDSCHConfigBroadcastR17PdschConfigListR17Constraints = per.SizeRange(1, common.MaxNrofPDSCH_ConfigPTM_r17)

var pDSCHConfigBroadcastR17RateMatchPatternToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofRateMatchPatterns)

const (
	PDSCH_ConfigBroadcast_r17_Mcs_Table_r17_Qam256     = 0
	PDSCH_ConfigBroadcast_r17_Mcs_Table_r17_Qam64LowSE = 1
)

var pDSCHConfigBroadcastR17McsTableR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_ConfigBroadcast_r17_Mcs_Table_r17_Qam256, PDSCH_ConfigBroadcast_r17_Mcs_Table_r17_Qam64LowSE},
}

const (
	PDSCH_ConfigBroadcast_r17_XOverhead_r17_XOh6  = 0
	PDSCH_ConfigBroadcast_r17_XOverhead_r17_XOh12 = 1
	PDSCH_ConfigBroadcast_r17_XOverhead_r17_XOh18 = 2
)

var pDSCHConfigBroadcastR17XOverheadR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_ConfigBroadcast_r17_XOverhead_r17_XOh6, PDSCH_ConfigBroadcast_r17_XOverhead_r17_XOh12, PDSCH_ConfigBroadcast_r17_XOverhead_r17_XOh18},
}

type PDSCH_ConfigBroadcast_r17 struct {
	PdschConfigList_r17                []PDSCH_ConfigPTM_r17
	Pdsch_TimeDomainAllocationList_r17 *PDSCH_TimeDomainResourceAllocationList_r16
	RateMatchPatternToAddModList_r17   []RateMatchPattern
	Lte_CRS_ToMatchAround_r17          *RateMatchPatternLTE_CRS
	Mcs_Table_r17                      *int64
	XOverhead_r17                      *int64
}

func (ie *PDSCH_ConfigBroadcast_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDSCHConfigBroadcastR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pdsch_TimeDomainAllocationList_r17 != nil, ie.RateMatchPatternToAddModList_r17 != nil, ie.Lte_CRS_ToMatchAround_r17 != nil, ie.Mcs_Table_r17 != nil, ie.XOverhead_r17 != nil}); err != nil {
		return err
	}

	// 2. pdschConfigList-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(pDSCHConfigBroadcastR17PdschConfigListR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.PdschConfigList_r17))); err != nil {
			return err
		}
		for i := range ie.PdschConfigList_r17 {
			if err := ie.PdschConfigList_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. pdsch-TimeDomainAllocationList-r17: ref
	{
		if ie.Pdsch_TimeDomainAllocationList_r17 != nil {
			if err := ie.Pdsch_TimeDomainAllocationList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. rateMatchPatternToAddModList-r17: sequence-of
	{
		if ie.RateMatchPatternToAddModList_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(pDSCHConfigBroadcastR17RateMatchPatternToAddModListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.RateMatchPatternToAddModList_r17))); err != nil {
				return err
			}
			for i := range ie.RateMatchPatternToAddModList_r17 {
				if err := ie.RateMatchPatternToAddModList_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. lte-CRS-ToMatchAround-r17: ref
	{
		if ie.Lte_CRS_ToMatchAround_r17 != nil {
			if err := ie.Lte_CRS_ToMatchAround_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. mcs-Table-r17: enumerated
	{
		if ie.Mcs_Table_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Mcs_Table_r17, pDSCHConfigBroadcastR17McsTableR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. xOverhead-r17: enumerated
	{
		if ie.XOverhead_r17 != nil {
			if err := e.EncodeEnumerated(*ie.XOverhead_r17, pDSCHConfigBroadcastR17XOverheadR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDSCH_ConfigBroadcast_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDSCHConfigBroadcastR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pdschConfigList-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(pDSCHConfigBroadcastR17PdschConfigListR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.PdschConfigList_r17 = make([]PDSCH_ConfigPTM_r17, n)
		for i := int64(0); i < n; i++ {
			if err := ie.PdschConfigList_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. pdsch-TimeDomainAllocationList-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Pdsch_TimeDomainAllocationList_r17 = new(PDSCH_TimeDomainResourceAllocationList_r16)
			if err := ie.Pdsch_TimeDomainAllocationList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. rateMatchPatternToAddModList-r17: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(pDSCHConfigBroadcastR17RateMatchPatternToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RateMatchPatternToAddModList_r17 = make([]RateMatchPattern, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RateMatchPatternToAddModList_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. lte-CRS-ToMatchAround-r17: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Lte_CRS_ToMatchAround_r17 = new(RateMatchPatternLTE_CRS)
			if err := ie.Lte_CRS_ToMatchAround_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. mcs-Table-r17: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(pDSCHConfigBroadcastR17McsTableR17Constraints)
			if err != nil {
				return err
			}
			ie.Mcs_Table_r17 = &idx
		}
	}

	// 7. xOverhead-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(pDSCHConfigBroadcastR17XOverheadR17Constraints)
			if err != nil {
				return err
			}
			ie.XOverhead_r17 = &idx
		}
	}

	return nil
}
