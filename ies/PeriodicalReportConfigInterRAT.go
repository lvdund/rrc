package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PeriodicalReportConfigInterRAT struct {
	ReportInterval                ReportInterval                                                `madatory`
	ReportAmount                  PeriodicalReportConfigInterRAT_reportAmount                   `madatory`
	ReportQuantity                MeasReportQuantity                                            `madatory`
	MaxReportCells                int64                                                         `lb:1,ub:maxCellReport,madatory`
	ReportQuantityUTRA_FDD_r16    *MeasReportQuantityUTRA_FDD_r16                               `optional,ext-1`
	IncludeCommonLocationInfo_r16 *PeriodicalReportConfigInterRAT_includeCommonLocationInfo_r16 `optional,ext-2`
	IncludeBT_Meas_r16            *BT_NameList_r16                                              `optional,ext-2,setuprelease`
	IncludeWLAN_Meas_r16          *WLAN_NameList_r16                                            `optional,ext-2,setuprelease`
	IncludeSensor_Meas_r16        *Sensor_NameList_r16                                          `optional,ext-2,setuprelease`
	ReportQuantityRelay_r17       *SL_MeasReportQuantity_r16                                    `optional,ext-3`
}

func (ie *PeriodicalReportConfigInterRAT) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.ReportQuantityUTRA_FDD_r16 != nil || ie.IncludeCommonLocationInfo_r16 != nil || ie.IncludeBT_Meas_r16 != nil || ie.IncludeWLAN_Meas_r16 != nil || ie.IncludeSensor_Meas_r16 != nil || ie.ReportQuantityRelay_r17 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ReportInterval.Encode(w); err != nil {
		return utils.WrapError("Encode ReportInterval", err)
	}
	if err = ie.ReportAmount.Encode(w); err != nil {
		return utils.WrapError("Encode ReportAmount", err)
	}
	if err = ie.ReportQuantity.Encode(w); err != nil {
		return utils.WrapError("Encode ReportQuantity", err)
	}
	if err = w.WriteInteger(ie.MaxReportCells, &aper.Constraint{Lb: 1, Ub: maxCellReport}, false); err != nil {
		return utils.WrapError("WriteInteger MaxReportCells", err)
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.ReportQuantityUTRA_FDD_r16 != nil, ie.IncludeCommonLocationInfo_r16 != nil || ie.IncludeBT_Meas_r16 != nil || ie.IncludeWLAN_Meas_r16 != nil || ie.IncludeSensor_Meas_r16 != nil, ie.ReportQuantityRelay_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PeriodicalReportConfigInterRAT", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ReportQuantityUTRA_FDD_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ReportQuantityUTRA_FDD_r16 optional
			if ie.ReportQuantityUTRA_FDD_r16 != nil {
				if err = ie.ReportQuantityUTRA_FDD_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ReportQuantityUTRA_FDD_r16", err)
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
			optionals_ext_2 := []bool{ie.IncludeCommonLocationInfo_r16 != nil, ie.IncludeBT_Meas_r16 != nil, ie.IncludeWLAN_Meas_r16 != nil, ie.IncludeSensor_Meas_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
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
			optionals_ext_3 := []bool{ie.ReportQuantityRelay_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
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

func (ie *PeriodicalReportConfigInterRAT) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ReportInterval.Decode(r); err != nil {
		return utils.WrapError("Decode ReportInterval", err)
	}
	if err = ie.ReportAmount.Decode(r); err != nil {
		return utils.WrapError("Decode ReportAmount", err)
	}
	if err = ie.ReportQuantity.Decode(r); err != nil {
		return utils.WrapError("Decode ReportQuantity", err)
	}
	var tmp_int_MaxReportCells int64
	if tmp_int_MaxReportCells, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxCellReport}, false); err != nil {
		return utils.WrapError("ReadInteger MaxReportCells", err)
	}
	ie.MaxReportCells = tmp_int_MaxReportCells

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			ReportQuantityUTRA_FDD_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ReportQuantityUTRA_FDD_r16 optional
			if ReportQuantityUTRA_FDD_r16Present {
				ie.ReportQuantityUTRA_FDD_r16 = new(MeasReportQuantityUTRA_FDD_r16)
				if err = ie.ReportQuantityUTRA_FDD_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ReportQuantityUTRA_FDD_r16", err)
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
			// decode IncludeCommonLocationInfo_r16 optional
			if IncludeCommonLocationInfo_r16Present {
				ie.IncludeCommonLocationInfo_r16 = new(PeriodicalReportConfigInterRAT_includeCommonLocationInfo_r16)
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
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			ReportQuantityRelay_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
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
