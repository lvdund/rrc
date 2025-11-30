package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectNR struct {
	SsbFrequency                    *ARFCN_ValueNR                    `optional`
	SsbSubcarrierSpacing            *SubcarrierSpacing                `optional`
	Smtc1                           *SSB_MTC                          `optional`
	Smtc2                           *SSB_MTC2                         `optional`
	RefFreqCSI_RS                   *ARFCN_ValueNR                    `optional`
	ReferenceSignalConfig           ReferenceSignalConfig             `madatory`
	AbsThreshSS_BlocksConsolidation *ThresholdNR                      `optional`
	AbsThreshCSI_RS_Consolidation   *ThresholdNR                      `optional`
	NrofSS_BlocksToAverage          *int64                            `lb:2,ub:maxNrofSS_BlocksToAverage,optional`
	NrofCSI_RS_ResourcesToAverage   *int64                            `lb:2,ub:maxNrofCSI_RS_ResourcesToAverage,optional`
	QuantityConfigIndex             int64                             `lb:1,ub:maxNrofQuantityConfig,madatory`
	OffsetMO                        Q_OffsetRangeList                 `madatory`
	CellsToRemoveList               *PCI_List                         `optional`
	CellsToAddModList               *CellsToAddModList                `optional`
	ExcludedCellsToRemoveList       *PCI_RangeIndexList               `optional`
	ExcludedCellsToAddModList       []PCI_RangeElement                `lb:1,ub:maxNrofPCI_Ranges,optional`
	AllowedCellsToRemoveList        *PCI_RangeIndexList               `optional`
	AllowedCellsToAddModList        []PCI_RangeElement                `lb:1,ub:maxNrofPCI_Ranges,optional`
	FreqBandIndicatorNR             *FreqBandIndicatorNR              `optional,ext-1`
	MeasCycleSCell                  *MeasObjectNR_measCycleSCell      `optional,ext-1`
	Smtc3list_r16                   *SSB_MTC3List_r16                 `optional,ext-2`
	Rmtc_Config_r16                 *RMTC_Config_r16                  `optional,ext-2,setuprelease`
	T312_r16                        *T312_r16                         `optional,ext-2,setuprelease`
	AssociatedMeasGapSSB_r17        *MeasGapId_r17                    `optional,ext-3`
	AssociatedMeasGapCSIRS_r17      *MeasGapId_r17                    `optional,ext-3`
	Smtc4list_r17                   *SSB_MTC4List_r17                 `optional,ext-3`
	MeasCyclePSCell_r17             *MeasObjectNR_measCyclePSCell_r17 `optional,ext-3`
	CellsToAddModListExt_v1710      *CellsToAddModListExt_v1710       `optional,ext-3`
	AssociatedMeasGapSSB2_v1720     *MeasGapId_r17                    `optional,ext-4`
	AssociatedMeasGapCSIRS2_v1720   *MeasGapId_r17                    `optional,ext-4`
}

func (ie *MeasObjectNR) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.FreqBandIndicatorNR != nil || ie.MeasCycleSCell != nil || ie.Smtc3list_r16 != nil || ie.Rmtc_Config_r16 != nil || ie.T312_r16 != nil || ie.AssociatedMeasGapSSB_r17 != nil || ie.AssociatedMeasGapCSIRS_r17 != nil || ie.Smtc4list_r17 != nil || ie.MeasCyclePSCell_r17 != nil || ie.CellsToAddModListExt_v1710 != nil || ie.AssociatedMeasGapSSB2_v1720 != nil || ie.AssociatedMeasGapCSIRS2_v1720 != nil
	preambleBits := []bool{hasExtensions, ie.SsbFrequency != nil, ie.SsbSubcarrierSpacing != nil, ie.Smtc1 != nil, ie.Smtc2 != nil, ie.RefFreqCSI_RS != nil, ie.AbsThreshSS_BlocksConsolidation != nil, ie.AbsThreshCSI_RS_Consolidation != nil, ie.NrofSS_BlocksToAverage != nil, ie.NrofCSI_RS_ResourcesToAverage != nil, ie.CellsToRemoveList != nil, ie.CellsToAddModList != nil, ie.ExcludedCellsToRemoveList != nil, len(ie.ExcludedCellsToAddModList) > 0, ie.AllowedCellsToRemoveList != nil, len(ie.AllowedCellsToAddModList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SsbFrequency != nil {
		if err = ie.SsbFrequency.Encode(w); err != nil {
			return utils.WrapError("Encode SsbFrequency", err)
		}
	}
	if ie.SsbSubcarrierSpacing != nil {
		if err = ie.SsbSubcarrierSpacing.Encode(w); err != nil {
			return utils.WrapError("Encode SsbSubcarrierSpacing", err)
		}
	}
	if ie.Smtc1 != nil {
		if err = ie.Smtc1.Encode(w); err != nil {
			return utils.WrapError("Encode Smtc1", err)
		}
	}
	if ie.Smtc2 != nil {
		if err = ie.Smtc2.Encode(w); err != nil {
			return utils.WrapError("Encode Smtc2", err)
		}
	}
	if ie.RefFreqCSI_RS != nil {
		if err = ie.RefFreqCSI_RS.Encode(w); err != nil {
			return utils.WrapError("Encode RefFreqCSI_RS", err)
		}
	}
	if err = ie.ReferenceSignalConfig.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceSignalConfig", err)
	}
	if ie.AbsThreshSS_BlocksConsolidation != nil {
		if err = ie.AbsThreshSS_BlocksConsolidation.Encode(w); err != nil {
			return utils.WrapError("Encode AbsThreshSS_BlocksConsolidation", err)
		}
	}
	if ie.AbsThreshCSI_RS_Consolidation != nil {
		if err = ie.AbsThreshCSI_RS_Consolidation.Encode(w); err != nil {
			return utils.WrapError("Encode AbsThreshCSI_RS_Consolidation", err)
		}
	}
	if ie.NrofSS_BlocksToAverage != nil {
		if err = w.WriteInteger(*ie.NrofSS_BlocksToAverage, &aper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Encode NrofSS_BlocksToAverage", err)
		}
	}
	if ie.NrofCSI_RS_ResourcesToAverage != nil {
		if err = w.WriteInteger(*ie.NrofCSI_RS_ResourcesToAverage, &aper.Constraint{Lb: 2, Ub: maxNrofCSI_RS_ResourcesToAverage}, false); err != nil {
			return utils.WrapError("Encode NrofCSI_RS_ResourcesToAverage", err)
		}
	}
	if err = w.WriteInteger(ie.QuantityConfigIndex, &aper.Constraint{Lb: 1, Ub: maxNrofQuantityConfig}, false); err != nil {
		return utils.WrapError("WriteInteger QuantityConfigIndex", err)
	}
	if err = ie.OffsetMO.Encode(w); err != nil {
		return utils.WrapError("Encode OffsetMO", err)
	}
	if ie.CellsToRemoveList != nil {
		if err = ie.CellsToRemoveList.Encode(w); err != nil {
			return utils.WrapError("Encode CellsToRemoveList", err)
		}
	}
	if ie.CellsToAddModList != nil {
		if err = ie.CellsToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode CellsToAddModList", err)
		}
	}
	if ie.ExcludedCellsToRemoveList != nil {
		if err = ie.ExcludedCellsToRemoveList.Encode(w); err != nil {
			return utils.WrapError("Encode ExcludedCellsToRemoveList", err)
		}
	}
	if len(ie.ExcludedCellsToAddModList) > 0 {
		tmp_ExcludedCellsToAddModList := utils.NewSequence[*PCI_RangeElement]([]*PCI_RangeElement{}, aper.Constraint{Lb: 1, Ub: maxNrofPCI_Ranges}, false)
		for _, i := range ie.ExcludedCellsToAddModList {
			tmp_ExcludedCellsToAddModList.Value = append(tmp_ExcludedCellsToAddModList.Value, &i)
		}
		if err = tmp_ExcludedCellsToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode ExcludedCellsToAddModList", err)
		}
	}
	if ie.AllowedCellsToRemoveList != nil {
		if err = ie.AllowedCellsToRemoveList.Encode(w); err != nil {
			return utils.WrapError("Encode AllowedCellsToRemoveList", err)
		}
	}
	if len(ie.AllowedCellsToAddModList) > 0 {
		tmp_AllowedCellsToAddModList := utils.NewSequence[*PCI_RangeElement]([]*PCI_RangeElement{}, aper.Constraint{Lb: 1, Ub: maxNrofPCI_Ranges}, false)
		for _, i := range ie.AllowedCellsToAddModList {
			tmp_AllowedCellsToAddModList.Value = append(tmp_AllowedCellsToAddModList.Value, &i)
		}
		if err = tmp_AllowedCellsToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode AllowedCellsToAddModList", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{ie.FreqBandIndicatorNR != nil || ie.MeasCycleSCell != nil, ie.Smtc3list_r16 != nil || ie.Rmtc_Config_r16 != nil || ie.T312_r16 != nil, ie.AssociatedMeasGapSSB_r17 != nil || ie.AssociatedMeasGapCSIRS_r17 != nil || ie.Smtc4list_r17 != nil || ie.MeasCyclePSCell_r17 != nil || ie.CellsToAddModListExt_v1710 != nil, ie.AssociatedMeasGapSSB2_v1720 != nil || ie.AssociatedMeasGapCSIRS2_v1720 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasObjectNR", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.FreqBandIndicatorNR != nil, ie.MeasCycleSCell != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FreqBandIndicatorNR optional
			if ie.FreqBandIndicatorNR != nil {
				if err = ie.FreqBandIndicatorNR.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FreqBandIndicatorNR", err)
				}
			}
			// encode MeasCycleSCell optional
			if ie.MeasCycleSCell != nil {
				if err = ie.MeasCycleSCell.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasCycleSCell", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.Smtc3list_r16 != nil, ie.Rmtc_Config_r16 != nil, ie.T312_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Smtc3list_r16 optional
			if ie.Smtc3list_r16 != nil {
				if err = ie.Smtc3list_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Smtc3list_r16", err)
				}
			}
			// encode Rmtc_Config_r16 optional
			if ie.Rmtc_Config_r16 != nil {
				tmp_Rmtc_Config_r16 := utils.SetupRelease[*RMTC_Config_r16]{
					Setup: ie.Rmtc_Config_r16,
				}
				if err = tmp_Rmtc_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rmtc_Config_r16", err)
				}
			}
			// encode T312_r16 optional
			if ie.T312_r16 != nil {
				tmp_T312_r16 := utils.SetupRelease[*T312_r16]{
					Setup: ie.T312_r16,
				}
				if err = tmp_T312_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode T312_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.AssociatedMeasGapSSB_r17 != nil, ie.AssociatedMeasGapCSIRS_r17 != nil, ie.Smtc4list_r17 != nil, ie.MeasCyclePSCell_r17 != nil, ie.CellsToAddModListExt_v1710 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode AssociatedMeasGapSSB_r17 optional
			if ie.AssociatedMeasGapSSB_r17 != nil {
				if err = ie.AssociatedMeasGapSSB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AssociatedMeasGapSSB_r17", err)
				}
			}
			// encode AssociatedMeasGapCSIRS_r17 optional
			if ie.AssociatedMeasGapCSIRS_r17 != nil {
				if err = ie.AssociatedMeasGapCSIRS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AssociatedMeasGapCSIRS_r17", err)
				}
			}
			// encode Smtc4list_r17 optional
			if ie.Smtc4list_r17 != nil {
				if err = ie.Smtc4list_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Smtc4list_r17", err)
				}
			}
			// encode MeasCyclePSCell_r17 optional
			if ie.MeasCyclePSCell_r17 != nil {
				if err = ie.MeasCyclePSCell_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasCyclePSCell_r17", err)
				}
			}
			// encode CellsToAddModListExt_v1710 optional
			if ie.CellsToAddModListExt_v1710 != nil {
				if err = ie.CellsToAddModListExt_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CellsToAddModListExt_v1710", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.AssociatedMeasGapSSB2_v1720 != nil, ie.AssociatedMeasGapCSIRS2_v1720 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode AssociatedMeasGapSSB2_v1720 optional
			if ie.AssociatedMeasGapSSB2_v1720 != nil {
				if err = ie.AssociatedMeasGapSSB2_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AssociatedMeasGapSSB2_v1720", err)
				}
			}
			// encode AssociatedMeasGapCSIRS2_v1720 optional
			if ie.AssociatedMeasGapCSIRS2_v1720 != nil {
				if err = ie.AssociatedMeasGapCSIRS2_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AssociatedMeasGapCSIRS2_v1720", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *MeasObjectNR) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var SsbFrequencyPresent bool
	if SsbFrequencyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SsbSubcarrierSpacingPresent bool
	if SsbSubcarrierSpacingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Smtc1Present bool
	if Smtc1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Smtc2Present bool
	if Smtc2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RefFreqCSI_RSPresent bool
	if RefFreqCSI_RSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AbsThreshSS_BlocksConsolidationPresent bool
	if AbsThreshSS_BlocksConsolidationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AbsThreshCSI_RS_ConsolidationPresent bool
	if AbsThreshCSI_RS_ConsolidationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NrofSS_BlocksToAveragePresent bool
	if NrofSS_BlocksToAveragePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NrofCSI_RS_ResourcesToAveragePresent bool
	if NrofCSI_RS_ResourcesToAveragePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CellsToRemoveListPresent bool
	if CellsToRemoveListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CellsToAddModListPresent bool
	if CellsToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ExcludedCellsToRemoveListPresent bool
	if ExcludedCellsToRemoveListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ExcludedCellsToAddModListPresent bool
	if ExcludedCellsToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AllowedCellsToRemoveListPresent bool
	if AllowedCellsToRemoveListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AllowedCellsToAddModListPresent bool
	if AllowedCellsToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SsbFrequencyPresent {
		ie.SsbFrequency = new(ARFCN_ValueNR)
		if err = ie.SsbFrequency.Decode(r); err != nil {
			return utils.WrapError("Decode SsbFrequency", err)
		}
	}
	if SsbSubcarrierSpacingPresent {
		ie.SsbSubcarrierSpacing = new(SubcarrierSpacing)
		if err = ie.SsbSubcarrierSpacing.Decode(r); err != nil {
			return utils.WrapError("Decode SsbSubcarrierSpacing", err)
		}
	}
	if Smtc1Present {
		ie.Smtc1 = new(SSB_MTC)
		if err = ie.Smtc1.Decode(r); err != nil {
			return utils.WrapError("Decode Smtc1", err)
		}
	}
	if Smtc2Present {
		ie.Smtc2 = new(SSB_MTC2)
		if err = ie.Smtc2.Decode(r); err != nil {
			return utils.WrapError("Decode Smtc2", err)
		}
	}
	if RefFreqCSI_RSPresent {
		ie.RefFreqCSI_RS = new(ARFCN_ValueNR)
		if err = ie.RefFreqCSI_RS.Decode(r); err != nil {
			return utils.WrapError("Decode RefFreqCSI_RS", err)
		}
	}
	if err = ie.ReferenceSignalConfig.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceSignalConfig", err)
	}
	if AbsThreshSS_BlocksConsolidationPresent {
		ie.AbsThreshSS_BlocksConsolidation = new(ThresholdNR)
		if err = ie.AbsThreshSS_BlocksConsolidation.Decode(r); err != nil {
			return utils.WrapError("Decode AbsThreshSS_BlocksConsolidation", err)
		}
	}
	if AbsThreshCSI_RS_ConsolidationPresent {
		ie.AbsThreshCSI_RS_Consolidation = new(ThresholdNR)
		if err = ie.AbsThreshCSI_RS_Consolidation.Decode(r); err != nil {
			return utils.WrapError("Decode AbsThreshCSI_RS_Consolidation", err)
		}
	}
	if NrofSS_BlocksToAveragePresent {
		var tmp_int_NrofSS_BlocksToAverage int64
		if tmp_int_NrofSS_BlocksToAverage, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Decode NrofSS_BlocksToAverage", err)
		}
		ie.NrofSS_BlocksToAverage = &tmp_int_NrofSS_BlocksToAverage
	}
	if NrofCSI_RS_ResourcesToAveragePresent {
		var tmp_int_NrofCSI_RS_ResourcesToAverage int64
		if tmp_int_NrofCSI_RS_ResourcesToAverage, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: maxNrofCSI_RS_ResourcesToAverage}, false); err != nil {
			return utils.WrapError("Decode NrofCSI_RS_ResourcesToAverage", err)
		}
		ie.NrofCSI_RS_ResourcesToAverage = &tmp_int_NrofCSI_RS_ResourcesToAverage
	}
	var tmp_int_QuantityConfigIndex int64
	if tmp_int_QuantityConfigIndex, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxNrofQuantityConfig}, false); err != nil {
		return utils.WrapError("ReadInteger QuantityConfigIndex", err)
	}
	ie.QuantityConfigIndex = tmp_int_QuantityConfigIndex
	if err = ie.OffsetMO.Decode(r); err != nil {
		return utils.WrapError("Decode OffsetMO", err)
	}
	if CellsToRemoveListPresent {
		ie.CellsToRemoveList = new(PCI_List)
		if err = ie.CellsToRemoveList.Decode(r); err != nil {
			return utils.WrapError("Decode CellsToRemoveList", err)
		}
	}
	if CellsToAddModListPresent {
		ie.CellsToAddModList = new(CellsToAddModList)
		if err = ie.CellsToAddModList.Decode(r); err != nil {
			return utils.WrapError("Decode CellsToAddModList", err)
		}
	}
	if ExcludedCellsToRemoveListPresent {
		ie.ExcludedCellsToRemoveList = new(PCI_RangeIndexList)
		if err = ie.ExcludedCellsToRemoveList.Decode(r); err != nil {
			return utils.WrapError("Decode ExcludedCellsToRemoveList", err)
		}
	}
	if ExcludedCellsToAddModListPresent {
		tmp_ExcludedCellsToAddModList := utils.NewSequence[*PCI_RangeElement]([]*PCI_RangeElement{}, aper.Constraint{Lb: 1, Ub: maxNrofPCI_Ranges}, false)
		fn_ExcludedCellsToAddModList := func() *PCI_RangeElement {
			return new(PCI_RangeElement)
		}
		if err = tmp_ExcludedCellsToAddModList.Decode(r, fn_ExcludedCellsToAddModList); err != nil {
			return utils.WrapError("Decode ExcludedCellsToAddModList", err)
		}
		ie.ExcludedCellsToAddModList = []PCI_RangeElement{}
		for _, i := range tmp_ExcludedCellsToAddModList.Value {
			ie.ExcludedCellsToAddModList = append(ie.ExcludedCellsToAddModList, *i)
		}
	}
	if AllowedCellsToRemoveListPresent {
		ie.AllowedCellsToRemoveList = new(PCI_RangeIndexList)
		if err = ie.AllowedCellsToRemoveList.Decode(r); err != nil {
			return utils.WrapError("Decode AllowedCellsToRemoveList", err)
		}
	}
	if AllowedCellsToAddModListPresent {
		tmp_AllowedCellsToAddModList := utils.NewSequence[*PCI_RangeElement]([]*PCI_RangeElement{}, aper.Constraint{Lb: 1, Ub: maxNrofPCI_Ranges}, false)
		fn_AllowedCellsToAddModList := func() *PCI_RangeElement {
			return new(PCI_RangeElement)
		}
		if err = tmp_AllowedCellsToAddModList.Decode(r, fn_AllowedCellsToAddModList); err != nil {
			return utils.WrapError("Decode AllowedCellsToAddModList", err)
		}
		ie.AllowedCellsToAddModList = []PCI_RangeElement{}
		for _, i := range tmp_AllowedCellsToAddModList.Value {
			ie.AllowedCellsToAddModList = append(ie.AllowedCellsToAddModList, *i)
		}
	}

	if extensionBit {
		// Read extension bitmap: 4 bits for 4 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			FreqBandIndicatorNRPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MeasCycleSCellPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FreqBandIndicatorNR optional
			if FreqBandIndicatorNRPresent {
				ie.FreqBandIndicatorNR = new(FreqBandIndicatorNR)
				if err = ie.FreqBandIndicatorNR.Decode(extReader); err != nil {
					return utils.WrapError("Decode FreqBandIndicatorNR", err)
				}
			}
			// decode MeasCycleSCell optional
			if MeasCycleSCellPresent {
				ie.MeasCycleSCell = new(MeasObjectNR_measCycleSCell)
				if err = ie.MeasCycleSCell.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasCycleSCell", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Smtc3list_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Rmtc_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			T312_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Smtc3list_r16 optional
			if Smtc3list_r16Present {
				ie.Smtc3list_r16 = new(SSB_MTC3List_r16)
				if err = ie.Smtc3list_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Smtc3list_r16", err)
				}
			}
			// decode Rmtc_Config_r16 optional
			if Rmtc_Config_r16Present {
				tmp_Rmtc_Config_r16 := utils.SetupRelease[*RMTC_Config_r16]{}
				if err = tmp_Rmtc_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rmtc_Config_r16", err)
				}
				ie.Rmtc_Config_r16 = tmp_Rmtc_Config_r16.Setup
			}
			// decode T312_r16 optional
			if T312_r16Present {
				tmp_T312_r16 := utils.SetupRelease[*T312_r16]{}
				if err = tmp_T312_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode T312_r16", err)
				}
				ie.T312_r16 = tmp_T312_r16.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			AssociatedMeasGapSSB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AssociatedMeasGapCSIRS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Smtc4list_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MeasCyclePSCell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CellsToAddModListExt_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode AssociatedMeasGapSSB_r17 optional
			if AssociatedMeasGapSSB_r17Present {
				ie.AssociatedMeasGapSSB_r17 = new(MeasGapId_r17)
				if err = ie.AssociatedMeasGapSSB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode AssociatedMeasGapSSB_r17", err)
				}
			}
			// decode AssociatedMeasGapCSIRS_r17 optional
			if AssociatedMeasGapCSIRS_r17Present {
				ie.AssociatedMeasGapCSIRS_r17 = new(MeasGapId_r17)
				if err = ie.AssociatedMeasGapCSIRS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode AssociatedMeasGapCSIRS_r17", err)
				}
			}
			// decode Smtc4list_r17 optional
			if Smtc4list_r17Present {
				ie.Smtc4list_r17 = new(SSB_MTC4List_r17)
				if err = ie.Smtc4list_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Smtc4list_r17", err)
				}
			}
			// decode MeasCyclePSCell_r17 optional
			if MeasCyclePSCell_r17Present {
				ie.MeasCyclePSCell_r17 = new(MeasObjectNR_measCyclePSCell_r17)
				if err = ie.MeasCyclePSCell_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasCyclePSCell_r17", err)
				}
			}
			// decode CellsToAddModListExt_v1710 optional
			if CellsToAddModListExt_v1710Present {
				ie.CellsToAddModListExt_v1710 = new(CellsToAddModListExt_v1710)
				if err = ie.CellsToAddModListExt_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode CellsToAddModListExt_v1710", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			AssociatedMeasGapSSB2_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AssociatedMeasGapCSIRS2_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode AssociatedMeasGapSSB2_v1720 optional
			if AssociatedMeasGapSSB2_v1720Present {
				ie.AssociatedMeasGapSSB2_v1720 = new(MeasGapId_r17)
				if err = ie.AssociatedMeasGapSSB2_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode AssociatedMeasGapSSB2_v1720", err)
				}
			}
			// decode AssociatedMeasGapCSIRS2_v1720 optional
			if AssociatedMeasGapCSIRS2_v1720Present {
				ie.AssociatedMeasGapCSIRS2_v1720 = new(MeasGapId_r17)
				if err = ie.AssociatedMeasGapCSIRS2_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode AssociatedMeasGapCSIRS2_v1720", err)
				}
			}
		}
	}
	return nil
}
