package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_ResourceConfigMobility struct {
	SubcarrierSpacing        SubcarrierSpacing     `madatory`
	Csi_RS_CellList_Mobility []CSI_RS_CellMobility `lb:1,ub:maxNrofCSI_RS_CellsRRM,madatory`
	RefServCellIndex         *ServCellIndex        `optional,ext-1`
}

func (ie *CSI_RS_ResourceConfigMobility) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.RefServCellIndex != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SubcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode SubcarrierSpacing", err)
	}
	tmp_Csi_RS_CellList_Mobility := utils.NewSequence[*CSI_RS_CellMobility]([]*CSI_RS_CellMobility{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_CellsRRM}, false)
	for _, i := range ie.Csi_RS_CellList_Mobility {
		tmp_Csi_RS_CellList_Mobility.Value = append(tmp_Csi_RS_CellList_Mobility.Value, &i)
	}
	if err = tmp_Csi_RS_CellList_Mobility.Encode(w); err != nil {
		return utils.WrapError("Encode Csi_RS_CellList_Mobility", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.RefServCellIndex != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_RS_ResourceConfigMobility", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.RefServCellIndex != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RefServCellIndex optional
			if ie.RefServCellIndex != nil {
				if err = ie.RefServCellIndex.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RefServCellIndex", err)
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

func (ie *CSI_RS_ResourceConfigMobility) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SubcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode SubcarrierSpacing", err)
	}
	tmp_Csi_RS_CellList_Mobility := utils.NewSequence[*CSI_RS_CellMobility]([]*CSI_RS_CellMobility{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_CellsRRM}, false)
	fn_Csi_RS_CellList_Mobility := func() *CSI_RS_CellMobility {
		return new(CSI_RS_CellMobility)
	}
	if err = tmp_Csi_RS_CellList_Mobility.Decode(r, fn_Csi_RS_CellList_Mobility); err != nil {
		return utils.WrapError("Decode Csi_RS_CellList_Mobility", err)
	}
	ie.Csi_RS_CellList_Mobility = []CSI_RS_CellMobility{}
	for _, i := range tmp_Csi_RS_CellList_Mobility.Value {
		ie.Csi_RS_CellList_Mobility = append(ie.Csi_RS_CellList_Mobility, *i)
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

			RefServCellIndexPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RefServCellIndex optional
			if RefServCellIndexPresent {
				ie.RefServCellIndex = new(ServCellIndex)
				if err = ie.RefServCellIndex.Decode(extReader); err != nil {
					return utils.WrapError("Decode RefServCellIndex", err)
				}
			}
		}
	}
	return nil
}
