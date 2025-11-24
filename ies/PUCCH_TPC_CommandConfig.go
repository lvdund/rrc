package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_TPC_CommandConfig struct {
	Tpc_IndexPCell                               *int64 `lb:1,ub:15,optional`
	Tpc_IndexPUCCH_SCell                         *int64 `lb:1,ub:15,optional`
	Tpc_IndexPUCCH_sSCell_r17                    *int64 `lb:1,ub:15,optional,ext-1`
	Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 *int64 `lb:1,ub:15,optional,ext-1`
}

func (ie *PUCCH_TPC_CommandConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Tpc_IndexPUCCH_sSCell_r17 != nil || ie.Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Tpc_IndexPCell != nil, ie.Tpc_IndexPUCCH_SCell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Tpc_IndexPCell != nil {
		if err = w.WriteInteger(*ie.Tpc_IndexPCell, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Tpc_IndexPCell", err)
		}
	}
	if ie.Tpc_IndexPUCCH_SCell != nil {
		if err = w.WriteInteger(*ie.Tpc_IndexPUCCH_SCell, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Tpc_IndexPUCCH_SCell", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Tpc_IndexPUCCH_sSCell_r17 != nil || ie.Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUCCH_TPC_CommandConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Tpc_IndexPUCCH_sSCell_r17 != nil, ie.Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Tpc_IndexPUCCH_sSCell_r17 optional
			if ie.Tpc_IndexPUCCH_sSCell_r17 != nil {
				if err = extWriter.WriteInteger(*ie.Tpc_IndexPUCCH_sSCell_r17, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode Tpc_IndexPUCCH_sSCell_r17", err)
				}
			}
			// encode Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 optional
			if ie.Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 != nil {
				if err = extWriter.WriteInteger(*ie.Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17", err)
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

func (ie *PUCCH_TPC_CommandConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Tpc_IndexPCellPresent bool
	if Tpc_IndexPCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tpc_IndexPUCCH_SCellPresent bool
	if Tpc_IndexPUCCH_SCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Tpc_IndexPCellPresent {
		var tmp_int_Tpc_IndexPCell int64
		if tmp_int_Tpc_IndexPCell, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Tpc_IndexPCell", err)
		}
		ie.Tpc_IndexPCell = &tmp_int_Tpc_IndexPCell
	}
	if Tpc_IndexPUCCH_SCellPresent {
		var tmp_int_Tpc_IndexPUCCH_SCell int64
		if tmp_int_Tpc_IndexPUCCH_SCell, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Tpc_IndexPUCCH_SCell", err)
		}
		ie.Tpc_IndexPUCCH_SCell = &tmp_int_Tpc_IndexPUCCH_SCell
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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Tpc_IndexPUCCH_sSCell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Tpc_IndexPUCCH_sSCell_r17 optional
			if Tpc_IndexPUCCH_sSCell_r17Present {
				var tmp_int_Tpc_IndexPUCCH_sSCell_r17 int64
				if tmp_int_Tpc_IndexPUCCH_sSCell_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode Tpc_IndexPUCCH_sSCell_r17", err)
				}
				ie.Tpc_IndexPUCCH_sSCell_r17 = &tmp_int_Tpc_IndexPUCCH_sSCell_r17
			}
			// decode Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 optional
			if Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17Present {
				var tmp_int_Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 int64
				if tmp_int_Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17", err)
				}
				ie.Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17 = &tmp_int_Tpc_IndexPUCCH_sScellSecondaryPUCCHgroup_r17
			}
		}
	}
	return nil
}
