// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LTE-NeighCellsCRS-AssistInfo-r17 (line 8646).

var lTENeighCellsCRSAssistInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "neighCarrierBandwidthDL-r17", Optional: true},
		{Name: "neighCarrierFreqDL-r17", Optional: true},
		{Name: "neighCellId-r17", Optional: true},
		{Name: "neighCRS-muting-r17", Optional: true},
		{Name: "neighMBSFN-SubframeConfigList-r17", Optional: true},
		{Name: "neighNrofCRS-Ports-r17", Optional: true},
		{Name: "neighV-Shift-r17", Optional: true},
	},
}

const (
	LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_N6     = 0
	LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_N15    = 1
	LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_N25    = 2
	LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_N50    = 3
	LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_N75    = 4
	LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_N100   = 5
	LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_Spare2 = 6
	LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_Spare1 = 7
)

var lTENeighCellsCRSAssistInfoR17NeighCarrierBandwidthDLR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_N6, LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_N15, LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_N25, LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_N50, LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_N75, LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_N100, LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_Spare2, LTE_NeighCellsCRS_AssistInfo_r17_NeighCarrierBandwidthDL_r17_Spare1},
}

var lTENeighCellsCRSAssistInfoR17NeighCarrierFreqDLR17Constraints = per.Constrained(0, 16383)

const (
	LTE_NeighCellsCRS_AssistInfo_r17_NeighCRS_Muting_r17_Enabled = 0
)

var lTENeighCellsCRSAssistInfoR17NeighCRSMutingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTE_NeighCellsCRS_AssistInfo_r17_NeighCRS_Muting_r17_Enabled},
}

const (
	LTE_NeighCellsCRS_AssistInfo_r17_NeighNrofCRS_Ports_r17_N1 = 0
	LTE_NeighCellsCRS_AssistInfo_r17_NeighNrofCRS_Ports_r17_N2 = 1
	LTE_NeighCellsCRS_AssistInfo_r17_NeighNrofCRS_Ports_r17_N4 = 2
)

var lTENeighCellsCRSAssistInfoR17NeighNrofCRSPortsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTE_NeighCellsCRS_AssistInfo_r17_NeighNrofCRS_Ports_r17_N1, LTE_NeighCellsCRS_AssistInfo_r17_NeighNrofCRS_Ports_r17_N2, LTE_NeighCellsCRS_AssistInfo_r17_NeighNrofCRS_Ports_r17_N4},
}

const (
	LTE_NeighCellsCRS_AssistInfo_r17_NeighV_Shift_r17_N0 = 0
	LTE_NeighCellsCRS_AssistInfo_r17_NeighV_Shift_r17_N1 = 1
	LTE_NeighCellsCRS_AssistInfo_r17_NeighV_Shift_r17_N2 = 2
	LTE_NeighCellsCRS_AssistInfo_r17_NeighV_Shift_r17_N3 = 3
	LTE_NeighCellsCRS_AssistInfo_r17_NeighV_Shift_r17_N4 = 4
	LTE_NeighCellsCRS_AssistInfo_r17_NeighV_Shift_r17_N5 = 5
)

var lTENeighCellsCRSAssistInfoR17NeighVShiftR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{LTE_NeighCellsCRS_AssistInfo_r17_NeighV_Shift_r17_N0, LTE_NeighCellsCRS_AssistInfo_r17_NeighV_Shift_r17_N1, LTE_NeighCellsCRS_AssistInfo_r17_NeighV_Shift_r17_N2, LTE_NeighCellsCRS_AssistInfo_r17_NeighV_Shift_r17_N3, LTE_NeighCellsCRS_AssistInfo_r17_NeighV_Shift_r17_N4, LTE_NeighCellsCRS_AssistInfo_r17_NeighV_Shift_r17_N5},
}

type LTE_NeighCellsCRS_AssistInfo_r17 struct {
	NeighCarrierBandwidthDL_r17       *int64
	NeighCarrierFreqDL_r17            *int64
	NeighCellId_r17                   *EUTRA_PhysCellId
	NeighCRS_Muting_r17               *int64
	NeighMBSFN_SubframeConfigList_r17 *EUTRA_MBSFN_SubframeConfigList
	NeighNrofCRS_Ports_r17            *int64
	NeighV_Shift_r17                  *int64
}

func (ie *LTE_NeighCellsCRS_AssistInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lTENeighCellsCRSAssistInfoR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.NeighCarrierBandwidthDL_r17 != nil, ie.NeighCarrierFreqDL_r17 != nil, ie.NeighCellId_r17 != nil, ie.NeighCRS_Muting_r17 != nil, ie.NeighMBSFN_SubframeConfigList_r17 != nil, ie.NeighNrofCRS_Ports_r17 != nil, ie.NeighV_Shift_r17 != nil}); err != nil {
		return err
	}

	// 2. neighCarrierBandwidthDL-r17: enumerated
	{
		if ie.NeighCarrierBandwidthDL_r17 != nil {
			if err := e.EncodeEnumerated(*ie.NeighCarrierBandwidthDL_r17, lTENeighCellsCRSAssistInfoR17NeighCarrierBandwidthDLR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. neighCarrierFreqDL-r17: integer
	{
		if ie.NeighCarrierFreqDL_r17 != nil {
			if err := e.EncodeInteger(*ie.NeighCarrierFreqDL_r17, lTENeighCellsCRSAssistInfoR17NeighCarrierFreqDLR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. neighCellId-r17: ref
	{
		if ie.NeighCellId_r17 != nil {
			if err := ie.NeighCellId_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. neighCRS-muting-r17: enumerated
	{
		if ie.NeighCRS_Muting_r17 != nil {
			if err := e.EncodeEnumerated(*ie.NeighCRS_Muting_r17, lTENeighCellsCRSAssistInfoR17NeighCRSMutingR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. neighMBSFN-SubframeConfigList-r17: ref
	{
		if ie.NeighMBSFN_SubframeConfigList_r17 != nil {
			if err := ie.NeighMBSFN_SubframeConfigList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. neighNrofCRS-Ports-r17: enumerated
	{
		if ie.NeighNrofCRS_Ports_r17 != nil {
			if err := e.EncodeEnumerated(*ie.NeighNrofCRS_Ports_r17, lTENeighCellsCRSAssistInfoR17NeighNrofCRSPortsR17Constraints); err != nil {
				return err
			}
		}
	}

	// 8. neighV-Shift-r17: enumerated
	{
		if ie.NeighV_Shift_r17 != nil {
			if err := e.EncodeEnumerated(*ie.NeighV_Shift_r17, lTENeighCellsCRSAssistInfoR17NeighVShiftR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LTE_NeighCellsCRS_AssistInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lTENeighCellsCRSAssistInfoR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. neighCarrierBandwidthDL-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(lTENeighCellsCRSAssistInfoR17NeighCarrierBandwidthDLR17Constraints)
			if err != nil {
				return err
			}
			ie.NeighCarrierBandwidthDL_r17 = &idx
		}
	}

	// 3. neighCarrierFreqDL-r17: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(lTENeighCellsCRSAssistInfoR17NeighCarrierFreqDLR17Constraints)
			if err != nil {
				return err
			}
			ie.NeighCarrierFreqDL_r17 = &v
		}
	}

	// 4. neighCellId-r17: ref
	{
		if seq.IsComponentPresent(2) {
			ie.NeighCellId_r17 = new(EUTRA_PhysCellId)
			if err := ie.NeighCellId_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. neighCRS-muting-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(lTENeighCellsCRSAssistInfoR17NeighCRSMutingR17Constraints)
			if err != nil {
				return err
			}
			ie.NeighCRS_Muting_r17 = &idx
		}
	}

	// 6. neighMBSFN-SubframeConfigList-r17: ref
	{
		if seq.IsComponentPresent(4) {
			ie.NeighMBSFN_SubframeConfigList_r17 = new(EUTRA_MBSFN_SubframeConfigList)
			if err := ie.NeighMBSFN_SubframeConfigList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. neighNrofCRS-Ports-r17: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(lTENeighCellsCRSAssistInfoR17NeighNrofCRSPortsR17Constraints)
			if err != nil {
				return err
			}
			ie.NeighNrofCRS_Ports_r17 = &idx
		}
	}

	// 8. neighV-Shift-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(lTENeighCellsCRSAssistInfoR17NeighVShiftR17Constraints)
			if err != nil {
				return err
			}
			ie.NeighV_Shift_r17 = &idx
		}
	}

	return nil
}
