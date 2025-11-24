package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfig struct {
	EventId                       EventTriggerConfig_eventId                        `lb:1,ub:65525,madatory`
	RsType                        NR_RS_Type                                        `madatory,ext`
	ReportInterval                ReportInterval                                    `madatory,ext`
	ReportAmount                  EventTriggerConfig_reportAmount                   `madatory,ext`
	ReportQuantityCell            MeasReportQuantity                                `madatory,ext`
	MaxReportCells                int64                                             `lb:1,ub:maxCellReport,madatory,ext`
	ReportQuantityRS_Indexes      *MeasReportQuantity                               `optional,ext`
	MaxNrofRS_IndexesToReport     *int64                                            `lb:1,ub:maxNrofIndexesToReport,optional,ext`
	IncludeBeamMeasurements       bool                                              `madatory,ext`
	ReportAddNeighMeas            *EventTriggerConfig_reportAddNeighMeas            `optional,ext`
	MeasRSSI_ReportConfig_r16     *MeasRSSI_ReportConfig_r16                        `optional,ext-2`
	UseT312_r16                   *bool                                             `optional,ext-2`
	IncludeCommonLocationInfo_r16 *EventTriggerConfig_includeCommonLocationInfo_r16 `optional,ext-2`
	IncludeBT_Meas_r16            *BT_NameList_r16                                  `optional,ext-2,setuprelease`
	IncludeWLAN_Meas_r16          *WLAN_NameList_r16                                `optional,ext-2,setuprelease`
	IncludeSensor_Meas_r16        *Sensor_NameList_r16                              `optional,ext-2,setuprelease`
	CoarseLocationRequest_r17     *EventTriggerConfig_coarseLocationRequest_r17     `optional,ext-3`
	ReportQuantityRelay_r17       *SL_MeasReportQuantity_r16                        `optional,ext-3`
}

func (ie *EventTriggerConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.MeasRSSI_ReportConfig_r16 != nil || ie.UseT312_r16 != nil || ie.IncludeCommonLocationInfo_r16 != nil || ie.IncludeBT_Meas_r16 != nil || ie.IncludeWLAN_Meas_r16 != nil || ie.IncludeSensor_Meas_r16 != nil || ie.CoarseLocationRequest_r17 != nil || ie.ReportQuantityRelay_r17 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.EventId.Encode(w); err != nil {
		return utils.WrapError("Encode EventId", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.MeasRSSI_ReportConfig_r16 != nil || ie.UseT312_r16 != nil || ie.IncludeCommonLocationInfo_r16 != nil || ie.IncludeBT_Meas_r16 != nil || ie.IncludeWLAN_Meas_r16 != nil || ie.IncludeSensor_Meas_r16 != nil, ie.CoarseLocationRequest_r17 != nil || ie.ReportQuantityRelay_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap EventTriggerConfig", err)
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.MeasRSSI_ReportConfig_r16 != nil, ie.UseT312_r16 != nil, ie.IncludeCommonLocationInfo_r16 != nil, ie.IncludeBT_Meas_r16 != nil, ie.IncludeWLAN_Meas_r16 != nil, ie.IncludeSensor_Meas_r16 != nil}
			for _, bit := range optionals_ext_2 {
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
			// encode UseT312_r16 optional
			if ie.UseT312_r16 != nil {
				if err = extWriter.WriteBoolean(*ie.UseT312_r16); err != nil {
					return utils.WrapError("Encode UseT312_r16", err)
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
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.CoarseLocationRequest_r17 != nil, ie.ReportQuantityRelay_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
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

func (ie *EventTriggerConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.EventId.Decode(r); err != nil {
		return utils.WrapError("Decode EventId", err)
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			MeasRSSI_ReportConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UseT312_r16Present, err := extReader.ReadBool()
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
			// decode MeasRSSI_ReportConfig_r16 optional
			if MeasRSSI_ReportConfig_r16Present {
				ie.MeasRSSI_ReportConfig_r16 = new(MeasRSSI_ReportConfig_r16)
				if err = ie.MeasRSSI_ReportConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MeasRSSI_ReportConfig_r16", err)
				}
			}
			// decode UseT312_r16 optional
			if UseT312_r16Present {
				var tmp_bool_UseT312_r16 bool
				if tmp_bool_UseT312_r16, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode UseT312_r16", err)
				}
				ie.UseT312_r16 = &tmp_bool_UseT312_r16
			}
			// decode IncludeCommonLocationInfo_r16 optional
			if IncludeCommonLocationInfo_r16Present {
				ie.IncludeCommonLocationInfo_r16 = new(EventTriggerConfig_includeCommonLocationInfo_r16)
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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			CoarseLocationRequest_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ReportQuantityRelay_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode CoarseLocationRequest_r17 optional
			if CoarseLocationRequest_r17Present {
				ie.CoarseLocationRequest_r17 = new(EventTriggerConfig_coarseLocationRequest_r17)
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
