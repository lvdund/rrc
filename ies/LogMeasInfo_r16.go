package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type LogMeasInfo_r16 struct {
	LocationInfo_r16             *LocationInfo_r16                             `optional`
	RelativeTimeStamp_r16        int64                                         `lb:0,ub:7200,madatory`
	ServCellIdentity_r16         *CGI_Info_Logging_r16                         `optional`
	MeasResultServingCell_r16    *MeasResultServingCell_r16                    `optional`
	MeasResultNeighCells_r16     *LogMeasInfo_r16_measResultNeighCells_r16     `optional`
	AnyCellSelectionDetected_r16 *LogMeasInfo_r16_anyCellSelectionDetected_r16 `optional`
	InDeviceCoexDetected_r17     *LogMeasInfo_r16_inDeviceCoexDetected_r17     `optional,ext-1`
}

func (ie *LogMeasInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.InDeviceCoexDetected_r17 != nil
	preambleBits := []bool{hasExtensions, ie.LocationInfo_r16 != nil, ie.ServCellIdentity_r16 != nil, ie.MeasResultServingCell_r16 != nil, ie.MeasResultNeighCells_r16 != nil, ie.AnyCellSelectionDetected_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.LocationInfo_r16 != nil {
		if err = ie.LocationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LocationInfo_r16", err)
		}
	}
	if err = w.WriteInteger(ie.RelativeTimeStamp_r16, &aper.Constraint{Lb: 0, Ub: 7200}, false); err != nil {
		return utils.WrapError("WriteInteger RelativeTimeStamp_r16", err)
	}
	if ie.ServCellIdentity_r16 != nil {
		if err = ie.ServCellIdentity_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ServCellIdentity_r16", err)
		}
	}
	if ie.MeasResultServingCell_r16 != nil {
		if err = ie.MeasResultServingCell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultServingCell_r16", err)
		}
	}
	if ie.MeasResultNeighCells_r16 != nil {
		if err = ie.MeasResultNeighCells_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultNeighCells_r16", err)
		}
	}
	if ie.AnyCellSelectionDetected_r16 != nil {
		if err = ie.AnyCellSelectionDetected_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AnyCellSelectionDetected_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.InDeviceCoexDetected_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap LogMeasInfo_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.InDeviceCoexDetected_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode InDeviceCoexDetected_r17 optional
			if ie.InDeviceCoexDetected_r17 != nil {
				if err = ie.InDeviceCoexDetected_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InDeviceCoexDetected_r17", err)
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

func (ie *LogMeasInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var LocationInfo_r16Present bool
	if LocationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ServCellIdentity_r16Present bool
	if ServCellIdentity_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultServingCell_r16Present bool
	if MeasResultServingCell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultNeighCells_r16Present bool
	if MeasResultNeighCells_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AnyCellSelectionDetected_r16Present bool
	if AnyCellSelectionDetected_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if LocationInfo_r16Present {
		ie.LocationInfo_r16 = new(LocationInfo_r16)
		if err = ie.LocationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LocationInfo_r16", err)
		}
	}
	var tmp_int_RelativeTimeStamp_r16 int64
	if tmp_int_RelativeTimeStamp_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 7200}, false); err != nil {
		return utils.WrapError("ReadInteger RelativeTimeStamp_r16", err)
	}
	ie.RelativeTimeStamp_r16 = tmp_int_RelativeTimeStamp_r16
	if ServCellIdentity_r16Present {
		ie.ServCellIdentity_r16 = new(CGI_Info_Logging_r16)
		if err = ie.ServCellIdentity_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ServCellIdentity_r16", err)
		}
	}
	if MeasResultServingCell_r16Present {
		ie.MeasResultServingCell_r16 = new(MeasResultServingCell_r16)
		if err = ie.MeasResultServingCell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultServingCell_r16", err)
		}
	}
	if MeasResultNeighCells_r16Present {
		ie.MeasResultNeighCells_r16 = new(LogMeasInfo_r16_measResultNeighCells_r16)
		if err = ie.MeasResultNeighCells_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultNeighCells_r16", err)
		}
	}
	if AnyCellSelectionDetected_r16Present {
		ie.AnyCellSelectionDetected_r16 = new(LogMeasInfo_r16_anyCellSelectionDetected_r16)
		if err = ie.AnyCellSelectionDetected_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AnyCellSelectionDetected_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
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

			InDeviceCoexDetected_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode InDeviceCoexDetected_r17 optional
			if InDeviceCoexDetected_r17Present {
				ie.InDeviceCoexDetected_r17 = new(LogMeasInfo_r16_inDeviceCoexDetected_r17)
				if err = ie.InDeviceCoexDetected_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode InDeviceCoexDetected_r17", err)
				}
			}
		}
	}
	return nil
}
