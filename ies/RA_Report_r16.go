package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RA_Report_r16 struct {
	CellId_r16               RA_Report_r16_cellId_r16    `madatory`
	Ra_InformationCommon_r16 *RA_InformationCommon_r16   `optional`
	RaPurpose_r16            RA_Report_r16_raPurpose_r16 `madatory`
	SpCellID_r17             *CGI_Info_Logging_r16       `optional,ext-1`
}

func (ie *RA_Report_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.SpCellID_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Ra_InformationCommon_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CellId_r16", err)
	}
	if ie.Ra_InformationCommon_r16 != nil {
		if err = ie.Ra_InformationCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_InformationCommon_r16", err)
		}
	}
	if err = ie.RaPurpose_r16.Encode(w); err != nil {
		return utils.WrapError("Encode RaPurpose_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.SpCellID_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RA_Report_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.SpCellID_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SpCellID_r17 optional
			if ie.SpCellID_r17 != nil {
				if err = ie.SpCellID_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpCellID_r17", err)
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

func (ie *RA_Report_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_InformationCommon_r16Present bool
	if Ra_InformationCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CellId_r16", err)
	}
	if Ra_InformationCommon_r16Present {
		ie.Ra_InformationCommon_r16 = new(RA_InformationCommon_r16)
		if err = ie.Ra_InformationCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ra_InformationCommon_r16", err)
		}
	}
	if err = ie.RaPurpose_r16.Decode(r); err != nil {
		return utils.WrapError("Decode RaPurpose_r16", err)
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

			SpCellID_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SpCellID_r17 optional
			if SpCellID_r17Present {
				ie.SpCellID_r17 = new(CGI_Info_Logging_r16)
				if err = ie.SpCellID_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SpCellID_r17", err)
				}
			}
		}
	}
	return nil
}
