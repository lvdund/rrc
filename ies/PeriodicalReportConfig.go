package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PeriodicalReportConfig struct {
	RsType                        NR_RS_Type                                            `madatory`
	ReportInterval                ReportInterval                                        `madatory`
	ReportAmount                  PeriodicalReportConfig_reportAmount                   `madatory`
	ReportQuantityCell            MeasReportQuantity                                    `madatory`
	MaxReportCells                int64                                                 `lb:1,ub:maxCellReport,madatory`
	ReportQuantityRS_Indexes      *MeasReportQuantity                                   `optional`
	MaxNrofRS_IndexesToReport     *int64                                                `lb:1,ub:maxNrofIndexesToReport,optional`
	IncludeBeamMeasurements       bool                                                  `madatory`
	UseAllowedCellList            bool                                                  `madatory`
	MeasRSSI_ReportConfig_r16     *MeasRSSI_ReportConfig_r16                            `optional,ext-1`
	IncludeCommonLocationInfo_r16 *PeriodicalReportConfig_includeCommonLocationInfo_r16 `optional,ext-1`
	IncludeBT_Meas_r16            *BT_NameList_r16                                      `optional,ext-1,setuprelease`
	IncludeWLAN_Meas_r16          *WLAN_NameList_r16                                    `optional,ext-1,setuprelease`
	IncludeSensor_Meas_r16        *Sensor_NameList_r16                                  `optional,ext-1,setuprelease`
	Ul_DelayValueConfig_r16       *UL_DelayValueConfig_r16                              `optional,ext-1,setuprelease`
	ReportAddNeighMeas_r16        *PeriodicalReportConfig_reportAddNeighMeas_r16        `optional,ext-1`
	Ul_ExcessDelayConfig_r17      *UL_ExcessDelayConfig_r17                             `optional,ext-2,setuprelease`
	CoarseLocationRequest_r17     *PeriodicalReportConfig_coarseLocationRequest_r17     `optional,ext-2`
	ReportQuantityRelay_r17       *SL_MeasReportQuantity_r16                            `optional,ext-2`
}

func (ie *PeriodicalReportConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.MeasRSSI_ReportConfig_r16 != nil || ie.IncludeCommonLocationInfo_r16 != nil || ie.IncludeBT_Meas_r16 != nil || ie.IncludeWLAN_Meas_r16 != nil || ie.IncludeSensor_Meas_r16 != nil || ie.Ul_DelayValueConfig_r16 != nil || ie.ReportAddNeighMeas_r16 != nil || ie.Ul_ExcessDelayConfig_r17 != nil || ie.CoarseLocationRequest_r17 != nil || ie.ReportQuantityRelay_r17 != nil
	preambleBits := []bool{hasExtensions, ie.ReportQuantityRS_Indexes != nil, ie.MaxNrofRS_IndexesToReport != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.RsType.Encode(w); err != nil {
		return utils.WrapError("Encode RsType", err)
	}
	if err = ie.ReportInterval.Encode(w); err != nil {
		return utils.WrapError("Encode ReportInterval", err)
	}
	if err = ie.ReportAmount.Encode(w); err != nil {
		return utils.WrapError("Encode ReportAmount", err)
	}
	if err = ie.ReportQuantityCell.Encode(w); err != nil {
		return utils.WrapError("Encode ReportQuantityCell", err)
	}
	if err = w.WriteInteger(ie.MaxReportCells, &uper.Constraint{Lb: 1, Ub: maxCellReport}, false); err != nil {
		return utils.WrapError("WriteInteger MaxReportCells", err)
	}
	if ie.ReportQuantityRS_Indexes != nil {
		if err = ie.ReportQuantityRS_Indexes.Encode(w); err != nil {
			return utils.WrapError("Encode ReportQuantityRS_Indexes", err)
		}
	}
	if ie.MaxNrofRS_IndexesToReport != nil {
		if err = w.WriteInteger(*ie.MaxNrofRS_IndexesToReport, &uper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport}, false); err != nil {
			return utils.WrapError("Encode MaxNrofRS_IndexesToReport", err)
		}
	}
	if err = w.WriteBoolean(ie.IncludeBeamMeasurements); err != nil {
		return utils.WrapError("WriteBoolean IncludeBeamMeasurements", err)
	}
	if err = w.WriteBoolean(ie.UseAllowedCellList); err != nil {
		return utils.WrapError("WriteBoolean UseAllowedCellList", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.MeasRSSI_ReportConfig_r16 != nil || ie.IncludeCommonLocationInfo_r16 != nil || ie.IncludeBT_Meas_r16 != nil || ie.IncludeWLAN_Meas_r16 != nil || ie.IncludeSensor_Meas_r16 != nil || ie.Ul_DelayValueConfig_r16 != nil || ie.ReportAddNeighMeas_r16 != nil, ie.Ul_ExcessDelayConfig_r17 != nil || ie.CoarseLocationRequest_r17 != nil || ie.ReportQuantityRelay_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PeriodicalReportConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.MeasRSSI_ReportConfig_r16 != nil, ie.IncludeCommonLocationInfo_r16 != nil, ie.IncludeBT_Meas_r16 != nil, ie.IncludeWLAN_Meas_r16 != nil, ie.IncludeSensor_Meas_r16 != nil, ie.Ul_DelayValueConfig_r16 != nil, ie.ReportAddNeighMeas_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MeasRSSI_ReportConfig_r16 optional
			if ie.MeasRSSI_ReportConfig_r16 != nil {
				if err = ie.MeasRSSI_ReportConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MeasRSSI_ReportConfig_r16", err)
				}
			}
			// encode IncludeCommonLocationInfo_r16 optional
			if ie.IncludeCommonLocationInfo_r16 != nil {
				if err = ie.IncludeCommonLocationInfo_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IncludeCommonLocationInfo_r16", err)
				}
			}
			// encode IncludeBT_Meas_r16 optional
			if ie.IncludeBT_Meas_r16 != nil {
				tmp_IncludeBT_Meas_r16 := utils.SetupRelease[*BT_NameList_r16]{
					Setup: ie.IncludeBT_Meas_r16,
				}
				if err = tmp_IncludeBT_Meas_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IncludeBT_Meas_r16", err)
				}
			}
			// encode IncludeWLAN_Meas_r16 optional
			if ie.IncludeWLAN_Meas_r16 != nil {
				tmp_IncludeWLAN_Meas_r16 := utils.SetupRelease[*WLAN_NameList_r16]{
					Setup: ie.IncludeWLAN_Meas_r16,
				}
				if err = tmp_IncludeWLAN_Meas_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IncludeWLAN_Meas_r16", err)
				}
			}
			// encode IncludeSensor_Meas_r16 optional
			if ie.IncludeSensor_Meas_r16 != nil {
				tmp_IncludeSensor_Meas_r16 := utils.SetupRelease[*Sensor_NameList_r16]{
					Setup: ie.IncludeSensor_Meas_r16,
				}
				if err = tmp_IncludeSensor_Meas_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IncludeSensor_Meas_r16", err)
				}
			}
			// encode Ul_DelayValueConfig_r16 optional
			if ie.Ul_DelayValueConfig_r16 != nil {
				tmp_Ul_DelayValueConfig_r16 := utils.SetupRelease[*UL_DelayValueConfig_r16]{
					Setup: ie.Ul_DelayValueConfig_r16,
				}
				if err = tmp_Ul_DelayValueConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_DelayValueConfig_r16", err)
				}
			}
			// encode ReportAddNeighMeas_r16 optional
			if ie.ReportAddNeighMeas_r16 != nil {
				if err = ie.ReportAddNeighMeas_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ReportAddNeighMeas_r16", err)
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
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.Ul_ExcessDelayConfig_r17 != nil, ie.CoarseLocationRequest_r17 != nil, ie.ReportQuantityRelay_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ul_ExcessDelayConfig_r17 optional
			if ie.Ul_ExcessDelayConfig_r17 != nil {
				tmp_Ul_ExcessDelayConfig_r17 := utils.SetupRelease[*UL_ExcessDelayConfig_r17]{
					Setup: ie.Ul_ExcessDelayConfig_r17,
				}
				if err = tmp_Ul_ExcessDelayConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_ExcessDelayConfig_r17", err)
				}
			}
			// encode CoarseLocationRequest_r17 optional
			if ie.CoarseLocationRequest_r17 != nil {
				if err = ie.CoarseLocationRequest_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CoarseLocationRequest_r17", err)
				}
			}
			// encode ReportQuantityRelay_r17 optional
			if ie.ReportQuantityRelay_r17 != nil {
				if err = ie.ReportQuantityRelay_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ReportQuantityRelay_r17", err)
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

func (ie *PeriodicalReportConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ReportQuantityRS_IndexesPresent bool
	if ReportQuantityRS_IndexesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNrofRS_IndexesToReportPresent bool
	if MaxNrofRS_IndexesToReportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.RsType.Decode(r); err != nil {
		return utils.WrapError("Decode RsType", err)
	}
	if err = ie.ReportInterval.Decode(r); err != nil {
		return utils.WrapError("Decode ReportInterval", err)
	}
	if err = ie.ReportAmount.Decode(r); err != nil {
		return utils.WrapError("Decode ReportAmount", err)
	}
	if err = ie.ReportQuantityCell.Decode(r); err != nil {
		return utils.WrapError("Decode ReportQuantityCell", err)
	}
	var tmp_int_MaxReportCells int64
	if tmp_int_MaxReportCells, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxCellReport}, false); err != nil {
		return utils.WrapError("ReadInteger MaxReportCells", err)
	}
	ie.MaxReportCells = tmp_int_MaxReportCells
	if ReportQuantityRS_IndexesPresent {
		ie.ReportQuantityRS_Indexes = new(MeasReportQuantity)
		if err = ie.ReportQuantityRS_Indexes.Decode(r); err != nil {
			return utils.WrapError("Decode ReportQuantityRS_Indexes", err)
		}
	}
	if MaxNrofRS_IndexesToReportPresent {
		var tmp_int_MaxNrofRS_IndexesToReport int64
		if tmp_int_MaxNrofRS_IndexesToReport, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport}, false); err != nil {
			return utils.WrapError("Decode MaxNrofRS_IndexesToReport", err)
		}
		ie.MaxNrofRS_IndexesToReport = &tmp_int_MaxNrofRS_IndexesToReport
	}
	var tmp_bool_IncludeBeamMeasurements bool
	if tmp_bool_IncludeBeamMeasurements, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean IncludeBeamMeasurements", err)
	}
	ie.IncludeBeamMeasurements = tmp_bool_IncludeBeamMeasurements
	var tmp_bool_UseAllowedCellList bool
	if tmp_bool_UseAllowedCellList, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean UseAllowedCellList", err)
	}
	ie.UseAllowedCellList = tmp_bool_UseAllowedCellList

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			MeasRSSI_ReportConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IncludeCommonLocationInfo_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IncludeBT_Meas_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IncludeWLAN_Meas_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IncludeSensor_Meas_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_DelayValueConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ReportAddNeighMeas_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MeasRSSI_ReportConfig_r16 optional
			if MeasRSSI_ReportConfig_r16Present {
				ie.MeasRSSI_ReportConfig_r16 = new(MeasRSSI_ReportConfig_r16)
				if err = ie.MeasRSSI_ReportConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasRSSI_ReportConfig_r16", err)
				}
			}
			// decode IncludeCommonLocationInfo_r16 optional
			if IncludeCommonLocationInfo_r16Present {
				ie.IncludeCommonLocationInfo_r16 = new(PeriodicalReportConfig_includeCommonLocationInfo_r16)
				if err = ie.IncludeCommonLocationInfo_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode IncludeCommonLocationInfo_r16", err)
				}
			}
			// decode IncludeBT_Meas_r16 optional
			if IncludeBT_Meas_r16Present {
				tmp_IncludeBT_Meas_r16 := utils.SetupRelease[*BT_NameList_r16]{}
				if err = tmp_IncludeBT_Meas_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode IncludeBT_Meas_r16", err)
				}
				ie.IncludeBT_Meas_r16 = tmp_IncludeBT_Meas_r16.Setup
			}
			// decode IncludeWLAN_Meas_r16 optional
			if IncludeWLAN_Meas_r16Present {
				tmp_IncludeWLAN_Meas_r16 := utils.SetupRelease[*WLAN_NameList_r16]{}
				if err = tmp_IncludeWLAN_Meas_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode IncludeWLAN_Meas_r16", err)
				}
				ie.IncludeWLAN_Meas_r16 = tmp_IncludeWLAN_Meas_r16.Setup
			}
			// decode IncludeSensor_Meas_r16 optional
			if IncludeSensor_Meas_r16Present {
				tmp_IncludeSensor_Meas_r16 := utils.SetupRelease[*Sensor_NameList_r16]{}
				if err = tmp_IncludeSensor_Meas_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode IncludeSensor_Meas_r16", err)
				}
				ie.IncludeSensor_Meas_r16 = tmp_IncludeSensor_Meas_r16.Setup
			}
			// decode Ul_DelayValueConfig_r16 optional
			if Ul_DelayValueConfig_r16Present {
				tmp_Ul_DelayValueConfig_r16 := utils.SetupRelease[*UL_DelayValueConfig_r16]{}
				if err = tmp_Ul_DelayValueConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_DelayValueConfig_r16", err)
				}
				ie.Ul_DelayValueConfig_r16 = tmp_Ul_DelayValueConfig_r16.Setup
			}
			// decode ReportAddNeighMeas_r16 optional
			if ReportAddNeighMeas_r16Present {
				ie.ReportAddNeighMeas_r16 = new(PeriodicalReportConfig_reportAddNeighMeas_r16)
				if err = ie.ReportAddNeighMeas_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ReportAddNeighMeas_r16", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Ul_ExcessDelayConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CoarseLocationRequest_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ReportQuantityRelay_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ul_ExcessDelayConfig_r17 optional
			if Ul_ExcessDelayConfig_r17Present {
				tmp_Ul_ExcessDelayConfig_r17 := utils.SetupRelease[*UL_ExcessDelayConfig_r17]{}
				if err = tmp_Ul_ExcessDelayConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_ExcessDelayConfig_r17", err)
				}
				ie.Ul_ExcessDelayConfig_r17 = tmp_Ul_ExcessDelayConfig_r17.Setup
			}
			// decode CoarseLocationRequest_r17 optional
			if CoarseLocationRequest_r17Present {
				ie.CoarseLocationRequest_r17 = new(PeriodicalReportConfig_coarseLocationRequest_r17)
				if err = ie.CoarseLocationRequest_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode CoarseLocationRequest_r17", err)
				}
			}
			// decode ReportQuantityRelay_r17 optional
			if ReportQuantityRelay_r17Present {
				ie.ReportQuantityRelay_r17 = new(SL_MeasReportQuantity_r16)
				if err = ie.ReportQuantityRelay_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ReportQuantityRelay_r17", err)
				}
			}
		}
	}
	return nil
}
