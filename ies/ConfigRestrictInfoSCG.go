package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ConfigRestrictInfoSCG struct {
	AllowedBC_ListMRDC                     *BandCombinationInfoList                         `optional`
	PowerCoordination_FR1                  *ConfigRestrictInfoSCG_powerCoordination_FR1     `optional`
	ServCellIndexRangeSCG                  *ConfigRestrictInfoSCG_servCellIndexRangeSCG     `optional`
	MaxMeasFreqsSCG                        *int64                                           `lb:1,ub:maxMeasFreqsMN,optional`
	Dummy                                  *int64                                           `lb:1,ub:maxMeasIdentitiesMN,optional`
	SelectedBandEntriesMNList              []SelectedBandEntriesMN                          `lb:1,ub:maxBandComb,optional,ext-1`
	Pdcch_BlindDetectionSCG                *int64                                           `lb:1,ub:15,optional,ext-1`
	MaxNumberROHC_ContextSessionsSN        *int64                                           `lb:0,ub:16384,optional,ext-1`
	MaxIntraFreqMeasIdentitiesSCG          *int64                                           `lb:1,ub:maxMeasIdentitiesMN,optional,ext-2`
	MaxInterFreqMeasIdentitiesSCG          *int64                                           `lb:1,ub:maxMeasIdentitiesMN,optional,ext-2`
	P_maxNR_FR1_MCG_r16                    *P_Max                                           `optional,ext-3`
	PowerCoordination_FR2_r16              *ConfigRestrictInfoSCG_powerCoordination_FR2_r16 `optional,ext-3`
	Nrdc_PC_mode_FR1_r16                   *ConfigRestrictInfoSCG_nrdc_PC_mode_FR1_r16      `optional,ext-3`
	Nrdc_PC_mode_FR2_r16                   *ConfigRestrictInfoSCG_nrdc_PC_mode_FR2_r16      `optional,ext-3`
	MaxMeasSRS_ResourceSCG_r16             *int64                                           `lb:0,ub:maxNrofCLI_SRS_Resources_r16,optional,ext-3`
	MaxMeasCLI_ResourceSCG_r16             *int64                                           `lb:0,ub:maxNrofCLI_RSSI_Resources_r16,optional,ext-3`
	MaxNumberEHC_ContextsSN_r16            *int64                                           `lb:0,ub:65536,optional,ext-3`
	AllowedReducedConfigForOverheating_r16 *OverheatingAssistance                           `optional,ext-3`
	MaxToffset_r16                         *T_Offset_r16                                    `optional,ext-3`
	AllowedReducedConfigForOverheating_r17 *OverheatingAssistance_r17                       `optional,ext-4`
	MaxNumberUDC_DRB_r17                   *int64                                           `lb:0,ub:2,optional,ext-4`
	MaxNumberCPCCandidates_r17             *int64                                           `lb:0,ub:maxNrofCondCells_1_r17,optional,ext-4`
}

func (ie *ConfigRestrictInfoSCG) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := len(ie.SelectedBandEntriesMNList) > 0 || ie.Pdcch_BlindDetectionSCG != nil || ie.MaxNumberROHC_ContextSessionsSN != nil || ie.MaxIntraFreqMeasIdentitiesSCG != nil || ie.MaxInterFreqMeasIdentitiesSCG != nil || ie.P_maxNR_FR1_MCG_r16 != nil || ie.PowerCoordination_FR2_r16 != nil || ie.Nrdc_PC_mode_FR1_r16 != nil || ie.Nrdc_PC_mode_FR2_r16 != nil || ie.MaxMeasSRS_ResourceSCG_r16 != nil || ie.MaxMeasCLI_ResourceSCG_r16 != nil || ie.MaxNumberEHC_ContextsSN_r16 != nil || ie.AllowedReducedConfigForOverheating_r16 != nil || ie.MaxToffset_r16 != nil || ie.AllowedReducedConfigForOverheating_r17 != nil || ie.MaxNumberUDC_DRB_r17 != nil || ie.MaxNumberCPCCandidates_r17 != nil
	preambleBits := []bool{hasExtensions, ie.AllowedBC_ListMRDC != nil, ie.PowerCoordination_FR1 != nil, ie.ServCellIndexRangeSCG != nil, ie.MaxMeasFreqsSCG != nil, ie.Dummy != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.AllowedBC_ListMRDC != nil {
		if err = ie.AllowedBC_ListMRDC.Encode(w); err != nil {
			return utils.WrapError("Encode AllowedBC_ListMRDC", err)
		}
	}
	if ie.PowerCoordination_FR1 != nil {
		if err = ie.PowerCoordination_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode PowerCoordination_FR1", err)
		}
	}
	if ie.ServCellIndexRangeSCG != nil {
		if err = ie.ServCellIndexRangeSCG.Encode(w); err != nil {
			return utils.WrapError("Encode ServCellIndexRangeSCG", err)
		}
	}
	if ie.MaxMeasFreqsSCG != nil {
		if err = w.WriteInteger(*ie.MaxMeasFreqsSCG, &aper.Constraint{Lb: 1, Ub: maxMeasFreqsMN}, false); err != nil {
			return utils.WrapError("Encode MaxMeasFreqsSCG", err)
		}
	}
	if ie.Dummy != nil {
		if err = w.WriteInteger(*ie.Dummy, &aper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
			return utils.WrapError("Encode Dummy", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{len(ie.SelectedBandEntriesMNList) > 0 || ie.Pdcch_BlindDetectionSCG != nil || ie.MaxNumberROHC_ContextSessionsSN != nil, ie.MaxIntraFreqMeasIdentitiesSCG != nil || ie.MaxInterFreqMeasIdentitiesSCG != nil, ie.P_maxNR_FR1_MCG_r16 != nil || ie.PowerCoordination_FR2_r16 != nil || ie.Nrdc_PC_mode_FR1_r16 != nil || ie.Nrdc_PC_mode_FR2_r16 != nil || ie.MaxMeasSRS_ResourceSCG_r16 != nil || ie.MaxMeasCLI_ResourceSCG_r16 != nil || ie.MaxNumberEHC_ContextsSN_r16 != nil || ie.AllowedReducedConfigForOverheating_r16 != nil || ie.MaxToffset_r16 != nil, ie.AllowedReducedConfigForOverheating_r17 != nil || ie.MaxNumberUDC_DRB_r17 != nil || ie.MaxNumberCPCCandidates_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ConfigRestrictInfoSCG", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.SelectedBandEntriesMNList) > 0, ie.Pdcch_BlindDetectionSCG != nil, ie.MaxNumberROHC_ContextSessionsSN != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SelectedBandEntriesMNList optional
			if len(ie.SelectedBandEntriesMNList) > 0 {
				tmp_SelectedBandEntriesMNList := utils.NewSequence[*SelectedBandEntriesMN]([]*SelectedBandEntriesMN{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
				for _, i := range ie.SelectedBandEntriesMNList {
					tmp_SelectedBandEntriesMNList.Value = append(tmp_SelectedBandEntriesMNList.Value, &i)
				}
				if err = tmp_SelectedBandEntriesMNList.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SelectedBandEntriesMNList", err)
				}
			}
			// encode Pdcch_BlindDetectionSCG optional
			if ie.Pdcch_BlindDetectionSCG != nil {
				if err = extWriter.WriteInteger(*ie.Pdcch_BlindDetectionSCG, &aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode Pdcch_BlindDetectionSCG", err)
				}
			}
			// encode MaxNumberROHC_ContextSessionsSN optional
			if ie.MaxNumberROHC_ContextSessionsSN != nil {
				if err = extWriter.WriteInteger(*ie.MaxNumberROHC_ContextSessionsSN, &aper.Constraint{Lb: 0, Ub: 16384}, false); err != nil {
					return utils.WrapError("Encode MaxNumberROHC_ContextSessionsSN", err)
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
			optionals_ext_2 := []bool{ie.MaxIntraFreqMeasIdentitiesSCG != nil, ie.MaxInterFreqMeasIdentitiesSCG != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MaxIntraFreqMeasIdentitiesSCG optional
			if ie.MaxIntraFreqMeasIdentitiesSCG != nil {
				if err = extWriter.WriteInteger(*ie.MaxIntraFreqMeasIdentitiesSCG, &aper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Encode MaxIntraFreqMeasIdentitiesSCG", err)
				}
			}
			// encode MaxInterFreqMeasIdentitiesSCG optional
			if ie.MaxInterFreqMeasIdentitiesSCG != nil {
				if err = extWriter.WriteInteger(*ie.MaxInterFreqMeasIdentitiesSCG, &aper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Encode MaxInterFreqMeasIdentitiesSCG", err)
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
			optionals_ext_3 := []bool{ie.P_maxNR_FR1_MCG_r16 != nil, ie.PowerCoordination_FR2_r16 != nil, ie.Nrdc_PC_mode_FR1_r16 != nil, ie.Nrdc_PC_mode_FR2_r16 != nil, ie.MaxMeasSRS_ResourceSCG_r16 != nil, ie.MaxMeasCLI_ResourceSCG_r16 != nil, ie.MaxNumberEHC_ContextsSN_r16 != nil, ie.AllowedReducedConfigForOverheating_r16 != nil, ie.MaxToffset_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode P_maxNR_FR1_MCG_r16 optional
			if ie.P_maxNR_FR1_MCG_r16 != nil {
				if err = ie.P_maxNR_FR1_MCG_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode P_maxNR_FR1_MCG_r16", err)
				}
			}
			// encode PowerCoordination_FR2_r16 optional
			if ie.PowerCoordination_FR2_r16 != nil {
				if err = ie.PowerCoordination_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PowerCoordination_FR2_r16", err)
				}
			}
			// encode Nrdc_PC_mode_FR1_r16 optional
			if ie.Nrdc_PC_mode_FR1_r16 != nil {
				if err = ie.Nrdc_PC_mode_FR1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nrdc_PC_mode_FR1_r16", err)
				}
			}
			// encode Nrdc_PC_mode_FR2_r16 optional
			if ie.Nrdc_PC_mode_FR2_r16 != nil {
				if err = ie.Nrdc_PC_mode_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nrdc_PC_mode_FR2_r16", err)
				}
			}
			// encode MaxMeasSRS_ResourceSCG_r16 optional
			if ie.MaxMeasSRS_ResourceSCG_r16 != nil {
				if err = extWriter.WriteInteger(*ie.MaxMeasSRS_ResourceSCG_r16, &aper.Constraint{Lb: 0, Ub: maxNrofCLI_SRS_Resources_r16}, false); err != nil {
					return utils.WrapError("Encode MaxMeasSRS_ResourceSCG_r16", err)
				}
			}
			// encode MaxMeasCLI_ResourceSCG_r16 optional
			if ie.MaxMeasCLI_ResourceSCG_r16 != nil {
				if err = extWriter.WriteInteger(*ie.MaxMeasCLI_ResourceSCG_r16, &aper.Constraint{Lb: 0, Ub: maxNrofCLI_RSSI_Resources_r16}, false); err != nil {
					return utils.WrapError("Encode MaxMeasCLI_ResourceSCG_r16", err)
				}
			}
			// encode MaxNumberEHC_ContextsSN_r16 optional
			if ie.MaxNumberEHC_ContextsSN_r16 != nil {
				if err = extWriter.WriteInteger(*ie.MaxNumberEHC_ContextsSN_r16, &aper.Constraint{Lb: 0, Ub: 65536}, false); err != nil {
					return utils.WrapError("Encode MaxNumberEHC_ContextsSN_r16", err)
				}
			}
			// encode AllowedReducedConfigForOverheating_r16 optional
			if ie.AllowedReducedConfigForOverheating_r16 != nil {
				if err = ie.AllowedReducedConfigForOverheating_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AllowedReducedConfigForOverheating_r16", err)
				}
			}
			// encode MaxToffset_r16 optional
			if ie.MaxToffset_r16 != nil {
				if err = ie.MaxToffset_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxToffset_r16", err)
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
			optionals_ext_4 := []bool{ie.AllowedReducedConfigForOverheating_r17 != nil, ie.MaxNumberUDC_DRB_r17 != nil, ie.MaxNumberCPCCandidates_r17 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode AllowedReducedConfigForOverheating_r17 optional
			if ie.AllowedReducedConfigForOverheating_r17 != nil {
				if err = ie.AllowedReducedConfigForOverheating_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AllowedReducedConfigForOverheating_r17", err)
				}
			}
			// encode MaxNumberUDC_DRB_r17 optional
			if ie.MaxNumberUDC_DRB_r17 != nil {
				if err = extWriter.WriteInteger(*ie.MaxNumberUDC_DRB_r17, &aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode MaxNumberUDC_DRB_r17", err)
				}
			}
			// encode MaxNumberCPCCandidates_r17 optional
			if ie.MaxNumberCPCCandidates_r17 != nil {
				if err = extWriter.WriteInteger(*ie.MaxNumberCPCCandidates_r17, &aper.Constraint{Lb: 0, Ub: maxNrofCondCells_1_r17}, false); err != nil {
					return utils.WrapError("Encode MaxNumberCPCCandidates_r17", err)
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

func (ie *ConfigRestrictInfoSCG) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var AllowedBC_ListMRDCPresent bool
	if AllowedBC_ListMRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PowerCoordination_FR1Present bool
	if PowerCoordination_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ServCellIndexRangeSCGPresent bool
	if ServCellIndexRangeSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxMeasFreqsSCGPresent bool
	if MaxMeasFreqsSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DummyPresent bool
	if DummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if AllowedBC_ListMRDCPresent {
		ie.AllowedBC_ListMRDC = new(BandCombinationInfoList)
		if err = ie.AllowedBC_ListMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode AllowedBC_ListMRDC", err)
		}
	}
	if PowerCoordination_FR1Present {
		ie.PowerCoordination_FR1 = new(ConfigRestrictInfoSCG_powerCoordination_FR1)
		if err = ie.PowerCoordination_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode PowerCoordination_FR1", err)
		}
	}
	if ServCellIndexRangeSCGPresent {
		ie.ServCellIndexRangeSCG = new(ConfigRestrictInfoSCG_servCellIndexRangeSCG)
		if err = ie.ServCellIndexRangeSCG.Decode(r); err != nil {
			return utils.WrapError("Decode ServCellIndexRangeSCG", err)
		}
	}
	if MaxMeasFreqsSCGPresent {
		var tmp_int_MaxMeasFreqsSCG int64
		if tmp_int_MaxMeasFreqsSCG, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxMeasFreqsMN}, false); err != nil {
			return utils.WrapError("Decode MaxMeasFreqsSCG", err)
		}
		ie.MaxMeasFreqsSCG = &tmp_int_MaxMeasFreqsSCG
	}
	if DummyPresent {
		var tmp_int_Dummy int64
		if tmp_int_Dummy, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
			return utils.WrapError("Decode Dummy", err)
		}
		ie.Dummy = &tmp_int_Dummy
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

			SelectedBandEntriesMNListPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdcch_BlindDetectionSCGPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberROHC_ContextSessionsSNPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SelectedBandEntriesMNList optional
			if SelectedBandEntriesMNListPresent {
				tmp_SelectedBandEntriesMNList := utils.NewSequence[*SelectedBandEntriesMN]([]*SelectedBandEntriesMN{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
				fn_SelectedBandEntriesMNList := func() *SelectedBandEntriesMN {
					return new(SelectedBandEntriesMN)
				}
				if err = tmp_SelectedBandEntriesMNList.Decode(extReader, fn_SelectedBandEntriesMNList); err != nil {
					return utils.WrapError("Decode SelectedBandEntriesMNList", err)
				}
				ie.SelectedBandEntriesMNList = []SelectedBandEntriesMN{}
				for _, i := range tmp_SelectedBandEntriesMNList.Value {
					ie.SelectedBandEntriesMNList = append(ie.SelectedBandEntriesMNList, *i)
				}
			}
			// decode Pdcch_BlindDetectionSCG optional
			if Pdcch_BlindDetectionSCGPresent {
				var tmp_int_Pdcch_BlindDetectionSCG int64
				if tmp_int_Pdcch_BlindDetectionSCG, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode Pdcch_BlindDetectionSCG", err)
				}
				ie.Pdcch_BlindDetectionSCG = &tmp_int_Pdcch_BlindDetectionSCG
			}
			// decode MaxNumberROHC_ContextSessionsSN optional
			if MaxNumberROHC_ContextSessionsSNPresent {
				var tmp_int_MaxNumberROHC_ContextSessionsSN int64
				if tmp_int_MaxNumberROHC_ContextSessionsSN, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: 16384}, false); err != nil {
					return utils.WrapError("Decode MaxNumberROHC_ContextSessionsSN", err)
				}
				ie.MaxNumberROHC_ContextSessionsSN = &tmp_int_MaxNumberROHC_ContextSessionsSN
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			MaxIntraFreqMeasIdentitiesSCGPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxInterFreqMeasIdentitiesSCGPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MaxIntraFreqMeasIdentitiesSCG optional
			if MaxIntraFreqMeasIdentitiesSCGPresent {
				var tmp_int_MaxIntraFreqMeasIdentitiesSCG int64
				if tmp_int_MaxIntraFreqMeasIdentitiesSCG, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Decode MaxIntraFreqMeasIdentitiesSCG", err)
				}
				ie.MaxIntraFreqMeasIdentitiesSCG = &tmp_int_MaxIntraFreqMeasIdentitiesSCG
			}
			// decode MaxInterFreqMeasIdentitiesSCG optional
			if MaxInterFreqMeasIdentitiesSCGPresent {
				var tmp_int_MaxInterFreqMeasIdentitiesSCG int64
				if tmp_int_MaxInterFreqMeasIdentitiesSCG, err = extReader.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Decode MaxInterFreqMeasIdentitiesSCG", err)
				}
				ie.MaxInterFreqMeasIdentitiesSCG = &tmp_int_MaxInterFreqMeasIdentitiesSCG
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			P_maxNR_FR1_MCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PowerCoordination_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nrdc_PC_mode_FR1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nrdc_PC_mode_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxMeasSRS_ResourceSCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxMeasCLI_ResourceSCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberEHC_ContextsSN_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AllowedReducedConfigForOverheating_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxToffset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode P_maxNR_FR1_MCG_r16 optional
			if P_maxNR_FR1_MCG_r16Present {
				ie.P_maxNR_FR1_MCG_r16 = new(P_Max)
				if err = ie.P_maxNR_FR1_MCG_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode P_maxNR_FR1_MCG_r16", err)
				}
			}
			// decode PowerCoordination_FR2_r16 optional
			if PowerCoordination_FR2_r16Present {
				ie.PowerCoordination_FR2_r16 = new(ConfigRestrictInfoSCG_powerCoordination_FR2_r16)
				if err = ie.PowerCoordination_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode PowerCoordination_FR2_r16", err)
				}
			}
			// decode Nrdc_PC_mode_FR1_r16 optional
			if Nrdc_PC_mode_FR1_r16Present {
				ie.Nrdc_PC_mode_FR1_r16 = new(ConfigRestrictInfoSCG_nrdc_PC_mode_FR1_r16)
				if err = ie.Nrdc_PC_mode_FR1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nrdc_PC_mode_FR1_r16", err)
				}
			}
			// decode Nrdc_PC_mode_FR2_r16 optional
			if Nrdc_PC_mode_FR2_r16Present {
				ie.Nrdc_PC_mode_FR2_r16 = new(ConfigRestrictInfoSCG_nrdc_PC_mode_FR2_r16)
				if err = ie.Nrdc_PC_mode_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nrdc_PC_mode_FR2_r16", err)
				}
			}
			// decode MaxMeasSRS_ResourceSCG_r16 optional
			if MaxMeasSRS_ResourceSCG_r16Present {
				var tmp_int_MaxMeasSRS_ResourceSCG_r16 int64
				if tmp_int_MaxMeasSRS_ResourceSCG_r16, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofCLI_SRS_Resources_r16}, false); err != nil {
					return utils.WrapError("Decode MaxMeasSRS_ResourceSCG_r16", err)
				}
				ie.MaxMeasSRS_ResourceSCG_r16 = &tmp_int_MaxMeasSRS_ResourceSCG_r16
			}
			// decode MaxMeasCLI_ResourceSCG_r16 optional
			if MaxMeasCLI_ResourceSCG_r16Present {
				var tmp_int_MaxMeasCLI_ResourceSCG_r16 int64
				if tmp_int_MaxMeasCLI_ResourceSCG_r16, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofCLI_RSSI_Resources_r16}, false); err != nil {
					return utils.WrapError("Decode MaxMeasCLI_ResourceSCG_r16", err)
				}
				ie.MaxMeasCLI_ResourceSCG_r16 = &tmp_int_MaxMeasCLI_ResourceSCG_r16
			}
			// decode MaxNumberEHC_ContextsSN_r16 optional
			if MaxNumberEHC_ContextsSN_r16Present {
				var tmp_int_MaxNumberEHC_ContextsSN_r16 int64
				if tmp_int_MaxNumberEHC_ContextsSN_r16, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65536}, false); err != nil {
					return utils.WrapError("Decode MaxNumberEHC_ContextsSN_r16", err)
				}
				ie.MaxNumberEHC_ContextsSN_r16 = &tmp_int_MaxNumberEHC_ContextsSN_r16
			}
			// decode AllowedReducedConfigForOverheating_r16 optional
			if AllowedReducedConfigForOverheating_r16Present {
				ie.AllowedReducedConfigForOverheating_r16 = new(OverheatingAssistance)
				if err = ie.AllowedReducedConfigForOverheating_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode AllowedReducedConfigForOverheating_r16", err)
				}
			}
			// decode MaxToffset_r16 optional
			if MaxToffset_r16Present {
				ie.MaxToffset_r16 = new(T_Offset_r16)
				if err = ie.MaxToffset_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxToffset_r16", err)
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

			AllowedReducedConfigForOverheating_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberUDC_DRB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberCPCCandidates_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode AllowedReducedConfigForOverheating_r17 optional
			if AllowedReducedConfigForOverheating_r17Present {
				ie.AllowedReducedConfigForOverheating_r17 = new(OverheatingAssistance_r17)
				if err = ie.AllowedReducedConfigForOverheating_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode AllowedReducedConfigForOverheating_r17", err)
				}
			}
			// decode MaxNumberUDC_DRB_r17 optional
			if MaxNumberUDC_DRB_r17Present {
				var tmp_int_MaxNumberUDC_DRB_r17 int64
				if tmp_int_MaxNumberUDC_DRB_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode MaxNumberUDC_DRB_r17", err)
				}
				ie.MaxNumberUDC_DRB_r17 = &tmp_int_MaxNumberUDC_DRB_r17
			}
			// decode MaxNumberCPCCandidates_r17 optional
			if MaxNumberCPCCandidates_r17Present {
				var tmp_int_MaxNumberCPCCandidates_r17 int64
				if tmp_int_MaxNumberCPCCandidates_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofCondCells_1_r17}, false); err != nil {
					return utils.WrapError("Decode MaxNumberCPCCandidates_r17", err)
				}
				ie.MaxNumberCPCCandidates_r17 = &tmp_int_MaxNumberCPCCandidates_r17
			}
		}
	}
	return nil
}
