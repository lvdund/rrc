package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SSB_ConfigMobility struct {
	Ssb_ToMeasure                         *SSB_ToMeasure                         `optional,setuprelease`
	DeriveSSB_IndexFromCell               bool                                   `madatory`
	Ss_RSSI_Measurement                   *SS_RSSI_Measurement                   `optional`
	Ssb_PositionQCL_Common_r16            *SSB_PositionQCL_Relation_r16          `optional,ext-1`
	Ssb_PositionQCL_CellsToAddModList_r16 *SSB_PositionQCL_CellsToAddModList_r16 `optional,ext-1`
	Ssb_PositionQCL_CellsToRemoveList_r16 *PCI_List                              `optional,ext-1`
	DeriveSSB_IndexFromCellInter_r17      *ServCellIndex                         `optional,ext-2`
	Ssb_PositionQCL_Common_r17            *SSB_PositionQCL_Relation_r17          `optional,ext-2`
	Ssb_PositionQCL_Cells_r17             *SSB_PositionQCL_CellList_r17          `optional,ext-2,setuprelease`
	Cca_CellsToAddModList_r17             *PCI_List                              `optional,ext-3`
	Cca_CellsToRemoveList_r17             *PCI_List                              `optional,ext-3`
}

func (ie *SSB_ConfigMobility) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Ssb_PositionQCL_Common_r16 != nil || ie.Ssb_PositionQCL_CellsToAddModList_r16 != nil || ie.Ssb_PositionQCL_CellsToRemoveList_r16 != nil || ie.DeriveSSB_IndexFromCellInter_r17 != nil || ie.Ssb_PositionQCL_Common_r17 != nil || ie.Ssb_PositionQCL_Cells_r17 != nil || ie.Cca_CellsToAddModList_r17 != nil || ie.Cca_CellsToRemoveList_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Ssb_ToMeasure != nil, ie.Ss_RSSI_Measurement != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ssb_ToMeasure != nil {
		tmp_Ssb_ToMeasure := utils.SetupRelease[*SSB_ToMeasure]{
			Setup: ie.Ssb_ToMeasure,
		}
		if err = tmp_Ssb_ToMeasure.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_ToMeasure", err)
		}
	}
	if err = w.WriteBoolean(ie.DeriveSSB_IndexFromCell); err != nil {
		return utils.WrapError("WriteBoolean DeriveSSB_IndexFromCell", err)
	}
	if ie.Ss_RSSI_Measurement != nil {
		if err = ie.Ss_RSSI_Measurement.Encode(w); err != nil {
			return utils.WrapError("Encode Ss_RSSI_Measurement", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.Ssb_PositionQCL_Common_r16 != nil || ie.Ssb_PositionQCL_CellsToAddModList_r16 != nil || ie.Ssb_PositionQCL_CellsToRemoveList_r16 != nil, ie.DeriveSSB_IndexFromCellInter_r17 != nil || ie.Ssb_PositionQCL_Common_r17 != nil || ie.Ssb_PositionQCL_Cells_r17 != nil, ie.Cca_CellsToAddModList_r17 != nil || ie.Cca_CellsToRemoveList_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SSB_ConfigMobility", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Ssb_PositionQCL_Common_r16 != nil, ie.Ssb_PositionQCL_CellsToAddModList_r16 != nil, ie.Ssb_PositionQCL_CellsToRemoveList_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ssb_PositionQCL_Common_r16 optional
			if ie.Ssb_PositionQCL_Common_r16 != nil {
				if err = ie.Ssb_PositionQCL_Common_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ssb_PositionQCL_Common_r16", err)
				}
			}
			// encode Ssb_PositionQCL_CellsToAddModList_r16 optional
			if ie.Ssb_PositionQCL_CellsToAddModList_r16 != nil {
				if err = ie.Ssb_PositionQCL_CellsToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ssb_PositionQCL_CellsToAddModList_r16", err)
				}
			}
			// encode Ssb_PositionQCL_CellsToRemoveList_r16 optional
			if ie.Ssb_PositionQCL_CellsToRemoveList_r16 != nil {
				if err = ie.Ssb_PositionQCL_CellsToRemoveList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ssb_PositionQCL_CellsToRemoveList_r16", err)
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
			optionals_ext_2 := []bool{ie.DeriveSSB_IndexFromCellInter_r17 != nil, ie.Ssb_PositionQCL_Common_r17 != nil, ie.Ssb_PositionQCL_Cells_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode DeriveSSB_IndexFromCellInter_r17 optional
			if ie.DeriveSSB_IndexFromCellInter_r17 != nil {
				if err = ie.DeriveSSB_IndexFromCellInter_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DeriveSSB_IndexFromCellInter_r17", err)
				}
			}
			// encode Ssb_PositionQCL_Common_r17 optional
			if ie.Ssb_PositionQCL_Common_r17 != nil {
				if err = ie.Ssb_PositionQCL_Common_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ssb_PositionQCL_Common_r17", err)
				}
			}
			// encode Ssb_PositionQCL_Cells_r17 optional
			if ie.Ssb_PositionQCL_Cells_r17 != nil {
				tmp_Ssb_PositionQCL_Cells_r17 := utils.SetupRelease[*SSB_PositionQCL_CellList_r17]{
					Setup: ie.Ssb_PositionQCL_Cells_r17,
				}
				if err = tmp_Ssb_PositionQCL_Cells_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ssb_PositionQCL_Cells_r17", err)
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
			optionals_ext_3 := []bool{ie.Cca_CellsToAddModList_r17 != nil, ie.Cca_CellsToRemoveList_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Cca_CellsToAddModList_r17 optional
			if ie.Cca_CellsToAddModList_r17 != nil {
				if err = ie.Cca_CellsToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cca_CellsToAddModList_r17", err)
				}
			}
			// encode Cca_CellsToRemoveList_r17 optional
			if ie.Cca_CellsToRemoveList_r17 != nil {
				if err = ie.Cca_CellsToRemoveList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cca_CellsToRemoveList_r17", err)
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

func (ie *SSB_ConfigMobility) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_ToMeasurePresent bool
	if Ssb_ToMeasurePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ss_RSSI_MeasurementPresent bool
	if Ss_RSSI_MeasurementPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ssb_ToMeasurePresent {
		tmp_Ssb_ToMeasure := utils.SetupRelease[*SSB_ToMeasure]{}
		if err = tmp_Ssb_ToMeasure.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_ToMeasure", err)
		}
		ie.Ssb_ToMeasure = tmp_Ssb_ToMeasure.Setup
	}
	var tmp_bool_DeriveSSB_IndexFromCell bool
	if tmp_bool_DeriveSSB_IndexFromCell, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean DeriveSSB_IndexFromCell", err)
	}
	ie.DeriveSSB_IndexFromCell = tmp_bool_DeriveSSB_IndexFromCell
	if Ss_RSSI_MeasurementPresent {
		ie.Ss_RSSI_Measurement = new(SS_RSSI_Measurement)
		if err = ie.Ss_RSSI_Measurement.Decode(r); err != nil {
			return utils.WrapError("Decode Ss_RSSI_Measurement", err)
		}
	}

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

			Ssb_PositionQCL_Common_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ssb_PositionQCL_CellsToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ssb_PositionQCL_CellsToRemoveList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ssb_PositionQCL_Common_r16 optional
			if Ssb_PositionQCL_Common_r16Present {
				ie.Ssb_PositionQCL_Common_r16 = new(SSB_PositionQCL_Relation_r16)
				if err = ie.Ssb_PositionQCL_Common_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ssb_PositionQCL_Common_r16", err)
				}
			}
			// decode Ssb_PositionQCL_CellsToAddModList_r16 optional
			if Ssb_PositionQCL_CellsToAddModList_r16Present {
				ie.Ssb_PositionQCL_CellsToAddModList_r16 = new(SSB_PositionQCL_CellsToAddModList_r16)
				if err = ie.Ssb_PositionQCL_CellsToAddModList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ssb_PositionQCL_CellsToAddModList_r16", err)
				}
			}
			// decode Ssb_PositionQCL_CellsToRemoveList_r16 optional
			if Ssb_PositionQCL_CellsToRemoveList_r16Present {
				ie.Ssb_PositionQCL_CellsToRemoveList_r16 = new(PCI_List)
				if err = ie.Ssb_PositionQCL_CellsToRemoveList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ssb_PositionQCL_CellsToRemoveList_r16", err)
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

			DeriveSSB_IndexFromCellInter_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ssb_PositionQCL_Common_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ssb_PositionQCL_Cells_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode DeriveSSB_IndexFromCellInter_r17 optional
			if DeriveSSB_IndexFromCellInter_r17Present {
				ie.DeriveSSB_IndexFromCellInter_r17 = new(ServCellIndex)
				if err = ie.DeriveSSB_IndexFromCellInter_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode DeriveSSB_IndexFromCellInter_r17", err)
				}
			}
			// decode Ssb_PositionQCL_Common_r17 optional
			if Ssb_PositionQCL_Common_r17Present {
				ie.Ssb_PositionQCL_Common_r17 = new(SSB_PositionQCL_Relation_r17)
				if err = ie.Ssb_PositionQCL_Common_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ssb_PositionQCL_Common_r17", err)
				}
			}
			// decode Ssb_PositionQCL_Cells_r17 optional
			if Ssb_PositionQCL_Cells_r17Present {
				tmp_Ssb_PositionQCL_Cells_r17 := utils.SetupRelease[*SSB_PositionQCL_CellList_r17]{}
				if err = tmp_Ssb_PositionQCL_Cells_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ssb_PositionQCL_Cells_r17", err)
				}
				ie.Ssb_PositionQCL_Cells_r17 = tmp_Ssb_PositionQCL_Cells_r17.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Cca_CellsToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cca_CellsToRemoveList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Cca_CellsToAddModList_r17 optional
			if Cca_CellsToAddModList_r17Present {
				ie.Cca_CellsToAddModList_r17 = new(PCI_List)
				if err = ie.Cca_CellsToAddModList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cca_CellsToAddModList_r17", err)
				}
			}
			// decode Cca_CellsToRemoveList_r17 optional
			if Cca_CellsToRemoveList_r17Present {
				ie.Cca_CellsToRemoveList_r17 = new(PCI_List)
				if err = ie.Cca_CellsToRemoveList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cca_CellsToRemoveList_r17", err)
				}
			}
		}
	}
	return nil
}
