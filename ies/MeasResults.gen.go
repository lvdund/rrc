// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResults (line 9711).

var measResultsConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "measId"},
		{Name: "measResultServingMOList"},
		{Name: "measResultNeighCells", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
	},
}

var measResultsMeasResultNeighCellsConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "measResultListNR"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "measResultListEUTRA"},
		{Name: "measResultListUTRA-FDD-r16"},
		{Name: "sl-MeasResultsCandRelay-r17"},
	},
}

const (
	MeasResults_MeasResultNeighCells_MeasResultListNR = 0
)

var measResultsSlMeasResultServingRelayR17Constraints = per.SizeConstraints{}

var measResultsCoarseLocationInfoR17Constraints = per.SizeConstraints{}

var measResultsExtCellsMetReportOnLeaveListR18Constraints = per.SizeRange(1, common.MaxCellReport)

type MeasResults struct {
	MeasId                  MeasId
	MeasResultServingMOList MeasResultServMOList
	MeasResultNeighCells    *struct {
		Choice           int
		MeasResultListNR *MeasResultListNR
	}
	MeasResultServFreqListEUTRA_SCG   *MeasResultServFreqListEUTRA_SCG
	MeasResultServFreqListNR_SCG      *MeasResultServFreqListNR_SCG
	MeasResultSFTD_EUTRA              *MeasResultSFTD_EUTRA
	MeasResultSFTD_NR                 *MeasResultCellSFTD_NR
	MeasResultCellListSFTD_NR         *MeasResultCellListSFTD_NR
	MeasResultForRSSI_r16             *MeasResultForRSSI_r16
	LocationInfo_r16                  *LocationInfo_r16
	Ul_PDCP_DelayValueResultList_r16  *UL_PDCP_DelayValueResultList_r16
	MeasResultsSL_r16                 *MeasResultsSL_r16
	MeasResultCLI_r16                 *MeasResultCLI_r16
	MeasResultRxTxTimeDiff_r17        *MeasResultRxTxTimeDiff_r17
	Sl_MeasResultServingRelay_r17     []byte
	Ul_PDCP_ExcessDelayResultList_r17 *UL_PDCP_ExcessDelayResultList_r17
	CoarseLocationInfo_r17            []byte
	AltitudeUE_r18                    *Altitude_r18
	CellsMetReportOnLeaveList_r18     []PhysCellId
}

func (ie *MeasResults) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultsConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MeasResultServFreqListEUTRA_SCG != nil || ie.MeasResultServFreqListNR_SCG != nil || ie.MeasResultSFTD_EUTRA != nil || ie.MeasResultSFTD_NR != nil
	hasExtGroup1 := ie.MeasResultCellListSFTD_NR != nil
	hasExtGroup2 := ie.MeasResultForRSSI_r16 != nil || ie.LocationInfo_r16 != nil || ie.Ul_PDCP_DelayValueResultList_r16 != nil || ie.MeasResultsSL_r16 != nil || ie.MeasResultCLI_r16 != nil
	hasExtGroup3 := ie.MeasResultRxTxTimeDiff_r17 != nil || ie.Sl_MeasResultServingRelay_r17 != nil || ie.Ul_PDCP_ExcessDelayResultList_r17 != nil || ie.CoarseLocationInfo_r17 != nil
	hasExtGroup4 := ie.AltitudeUE_r18 != nil || ie.CellsMetReportOnLeaveList_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MeasResultNeighCells != nil}); err != nil {
		return err
	}

	// 3. measId: ref
	{
		if err := ie.MeasId.Encode(e); err != nil {
			return err
		}
	}

	// 4. measResultServingMOList: ref
	{
		if err := ie.MeasResultServingMOList.Encode(e); err != nil {
			return err
		}
	}

	// 5. measResultNeighCells: choice
	{
		if ie.MeasResultNeighCells != nil {
			choiceEnc := e.NewChoiceEncoder(measResultsMeasResultNeighCellsConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.MeasResultNeighCells).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.MeasResultNeighCells).Choice {
			case MeasResults_MeasResultNeighCells_MeasResultListNR:
				if err := (*ie.MeasResultNeighCells).MeasResultListNR.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.MeasResultNeighCells).Choice), Constraint: "index out of range"}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "measResultServFreqListEUTRA-SCG", Optional: true},
					{Name: "measResultServFreqListNR-SCG", Optional: true},
					{Name: "measResultSFTD-EUTRA", Optional: true},
					{Name: "measResultSFTD-NR", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MeasResultServFreqListEUTRA_SCG != nil, ie.MeasResultServFreqListNR_SCG != nil, ie.MeasResultSFTD_EUTRA != nil, ie.MeasResultSFTD_NR != nil}); err != nil {
				return err
			}

			if ie.MeasResultServFreqListEUTRA_SCG != nil {
				if err := ie.MeasResultServFreqListEUTRA_SCG.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MeasResultServFreqListNR_SCG != nil {
				if err := ie.MeasResultServFreqListNR_SCG.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MeasResultSFTD_EUTRA != nil {
				if err := ie.MeasResultSFTD_EUTRA.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MeasResultSFTD_NR != nil {
				if err := ie.MeasResultSFTD_NR.Encode(ex); err != nil {
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
					{Name: "measResultCellListSFTD-NR", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MeasResultCellListSFTD_NR != nil}); err != nil {
				return err
			}

			if ie.MeasResultCellListSFTD_NR != nil {
				if err := ie.MeasResultCellListSFTD_NR.Encode(ex); err != nil {
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
					{Name: "measResultForRSSI-r16", Optional: true},
					{Name: "locationInfo-r16", Optional: true},
					{Name: "ul-PDCP-DelayValueResultList-r16", Optional: true},
					{Name: "measResultsSL-r16", Optional: true},
					{Name: "measResultCLI-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MeasResultForRSSI_r16 != nil, ie.LocationInfo_r16 != nil, ie.Ul_PDCP_DelayValueResultList_r16 != nil, ie.MeasResultsSL_r16 != nil, ie.MeasResultCLI_r16 != nil}); err != nil {
				return err
			}

			if ie.MeasResultForRSSI_r16 != nil {
				if err := ie.MeasResultForRSSI_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.LocationInfo_r16 != nil {
				if err := ie.LocationInfo_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ul_PDCP_DelayValueResultList_r16 != nil {
				if err := ie.Ul_PDCP_DelayValueResultList_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MeasResultsSL_r16 != nil {
				if err := ie.MeasResultsSL_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MeasResultCLI_r16 != nil {
				if err := ie.MeasResultCLI_r16.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "measResultRxTxTimeDiff-r17", Optional: true},
					{Name: "sl-MeasResultServingRelay-r17", Optional: true},
					{Name: "ul-PDCP-ExcessDelayResultList-r17", Optional: true},
					{Name: "coarseLocationInfo-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MeasResultRxTxTimeDiff_r17 != nil, ie.Sl_MeasResultServingRelay_r17 != nil, ie.Ul_PDCP_ExcessDelayResultList_r17 != nil, ie.CoarseLocationInfo_r17 != nil}); err != nil {
				return err
			}

			if ie.MeasResultRxTxTimeDiff_r17 != nil {
				if err := ie.MeasResultRxTxTimeDiff_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_MeasResultServingRelay_r17 != nil {
				if err := ex.EncodeOctetString(ie.Sl_MeasResultServingRelay_r17, measResultsSlMeasResultServingRelayR17Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_PDCP_ExcessDelayResultList_r17 != nil {
				if err := ie.Ul_PDCP_ExcessDelayResultList_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CoarseLocationInfo_r17 != nil {
				if err := ex.EncodeOctetString(ie.CoarseLocationInfo_r17, measResultsCoarseLocationInfoR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "altitudeUE-r18", Optional: true},
					{Name: "cellsMetReportOnLeaveList-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AltitudeUE_r18 != nil, ie.CellsMetReportOnLeaveList_r18 != nil}); err != nil {
				return err
			}

			if ie.AltitudeUE_r18 != nil {
				if err := ie.AltitudeUE_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CellsMetReportOnLeaveList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(measResultsExtCellsMetReportOnLeaveListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.CellsMetReportOnLeaveList_r18))); err != nil {
					return err
				}
				for i := range ie.CellsMetReportOnLeaveList_r18 {
					if err := ie.CellsMetReportOnLeaveList_r18[i].Encode(ex); err != nil {
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

func (ie *MeasResults) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultsConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. measId: ref
	{
		if err := ie.MeasId.Decode(d); err != nil {
			return err
		}
	}

	// 4. measResultServingMOList: ref
	{
		if err := ie.MeasResultServingMOList.Decode(d); err != nil {
			return err
		}
	}

	// 5. measResultNeighCells: choice
	{
		if seq.IsComponentPresent(2) {
			ie.MeasResultNeighCells = &struct {
				Choice           int
				MeasResultListNR *MeasResultListNR
			}{}
			choiceDec := d.NewChoiceDecoder(measResultsMeasResultNeighCellsConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MeasResultNeighCells).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case MeasResults_MeasResultNeighCells_MeasResultListNR:
				(*ie.MeasResultNeighCells).MeasResultListNR = new(MeasResultListNR)
				if err := (*ie.MeasResultNeighCells).MeasResultListNR.Decode(d); err != nil {
					return err
				}
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
				{Name: "measResultServFreqListEUTRA-SCG", Optional: true},
				{Name: "measResultServFreqListNR-SCG", Optional: true},
				{Name: "measResultSFTD-EUTRA", Optional: true},
				{Name: "measResultSFTD-NR", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.MeasResultServFreqListEUTRA_SCG = new(MeasResultServFreqListEUTRA_SCG)
			if err := ie.MeasResultServFreqListEUTRA_SCG.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.MeasResultServFreqListNR_SCG = new(MeasResultServFreqListNR_SCG)
			if err := ie.MeasResultServFreqListNR_SCG.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.MeasResultSFTD_EUTRA = new(MeasResultSFTD_EUTRA)
			if err := ie.MeasResultSFTD_EUTRA.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.MeasResultSFTD_NR = new(MeasResultCellSFTD_NR)
			if err := ie.MeasResultSFTD_NR.Decode(dx); err != nil {
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
				{Name: "measResultCellListSFTD-NR", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.MeasResultCellListSFTD_NR = new(MeasResultCellListSFTD_NR)
			if err := ie.MeasResultCellListSFTD_NR.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "measResultForRSSI-r16", Optional: true},
				{Name: "locationInfo-r16", Optional: true},
				{Name: "ul-PDCP-DelayValueResultList-r16", Optional: true},
				{Name: "measResultsSL-r16", Optional: true},
				{Name: "measResultCLI-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.MeasResultForRSSI_r16 = new(MeasResultForRSSI_r16)
			if err := ie.MeasResultForRSSI_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.LocationInfo_r16 = new(LocationInfo_r16)
			if err := ie.LocationInfo_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Ul_PDCP_DelayValueResultList_r16 = new(UL_PDCP_DelayValueResultList_r16)
			if err := ie.Ul_PDCP_DelayValueResultList_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.MeasResultsSL_r16 = new(MeasResultsSL_r16)
			if err := ie.MeasResultsSL_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.MeasResultCLI_r16 = new(MeasResultCLI_r16)
			if err := ie.MeasResultCLI_r16.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "measResultRxTxTimeDiff-r17", Optional: true},
				{Name: "sl-MeasResultServingRelay-r17", Optional: true},
				{Name: "ul-PDCP-ExcessDelayResultList-r17", Optional: true},
				{Name: "coarseLocationInfo-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.MeasResultRxTxTimeDiff_r17 = new(MeasResultRxTxTimeDiff_r17)
			if err := ie.MeasResultRxTxTimeDiff_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeOctetString(measResultsSlMeasResultServingRelayR17Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MeasResultServingRelay_r17 = v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Ul_PDCP_ExcessDelayResultList_r17 = new(UL_PDCP_ExcessDelayResultList_r17)
			if err := ie.Ul_PDCP_ExcessDelayResultList_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeOctetString(measResultsCoarseLocationInfoR17Constraints)
			if err != nil {
				return err
			}
			ie.CoarseLocationInfo_r17 = v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "altitudeUE-r18", Optional: true},
				{Name: "cellsMetReportOnLeaveList-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.AltitudeUE_r18 = new(Altitude_r18)
			if err := ie.AltitudeUE_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(measResultsExtCellsMetReportOnLeaveListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.CellsMetReportOnLeaveList_r18 = make([]PhysCellId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.CellsMetReportOnLeaveList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
