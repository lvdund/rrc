package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PowerControl_r16 struct {
	Sl_MaxTransPower_r16     int64                                         `lb:-30,ub:33,madatory`
	Sl_Alpha_PSSCH_PSCCH_r16 *SL_PowerControl_r16_sl_Alpha_PSSCH_PSCCH_r16 `optional`
	Dl_Alpha_PSSCH_PSCCH_r16 *SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16 `optional`
	Sl_P0_PSSCH_PSCCH_r16    *int64                                        `lb:-16,ub:15,optional`
	Dl_P0_PSSCH_PSCCH_r16    *int64                                        `lb:-16,ub:15,optional`
	Dl_Alpha_PSFCH_r16       *SL_PowerControl_r16_dl_Alpha_PSFCH_r16       `optional`
	Dl_P0_PSFCH_r16          *int64                                        `lb:-16,ub:15,optional`
	Dl_P0_PSSCH_PSCCH_r17    *int64                                        `lb:-202,ub:24,optional,ext-1`
	Sl_P0_PSSCH_PSCCH_r17    *int64                                        `lb:-202,ub:24,optional,ext-1`
	Dl_P0_PSFCH_r17          *int64                                        `lb:-202,ub:24,optional,ext-1`
}

func (ie *SL_PowerControl_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Dl_P0_PSSCH_PSCCH_r17 != nil || ie.Sl_P0_PSSCH_PSCCH_r17 != nil || ie.Dl_P0_PSFCH_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Sl_Alpha_PSSCH_PSCCH_r16 != nil, ie.Dl_Alpha_PSSCH_PSCCH_r16 != nil, ie.Sl_P0_PSSCH_PSCCH_r16 != nil, ie.Dl_P0_PSSCH_PSCCH_r16 != nil, ie.Dl_Alpha_PSFCH_r16 != nil, ie.Dl_P0_PSFCH_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.Sl_MaxTransPower_r16, &uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_MaxTransPower_r16", err)
	}
	if ie.Sl_Alpha_PSSCH_PSCCH_r16 != nil {
		if err = ie.Sl_Alpha_PSSCH_PSCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_Alpha_PSSCH_PSCCH_r16", err)
		}
	}
	if ie.Dl_Alpha_PSSCH_PSCCH_r16 != nil {
		if err = ie.Dl_Alpha_PSSCH_PSCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_Alpha_PSSCH_PSCCH_r16", err)
		}
	}
	if ie.Sl_P0_PSSCH_PSCCH_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_P0_PSSCH_PSCCH_r16, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Sl_P0_PSSCH_PSCCH_r16", err)
		}
	}
	if ie.Dl_P0_PSSCH_PSCCH_r16 != nil {
		if err = w.WriteInteger(*ie.Dl_P0_PSSCH_PSCCH_r16, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Dl_P0_PSSCH_PSCCH_r16", err)
		}
	}
	if ie.Dl_Alpha_PSFCH_r16 != nil {
		if err = ie.Dl_Alpha_PSFCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_Alpha_PSFCH_r16", err)
		}
	}
	if ie.Dl_P0_PSFCH_r16 != nil {
		if err = w.WriteInteger(*ie.Dl_P0_PSFCH_r16, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Dl_P0_PSFCH_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Dl_P0_PSSCH_PSCCH_r17 != nil || ie.Sl_P0_PSSCH_PSCCH_r17 != nil || ie.Dl_P0_PSFCH_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_PowerControl_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Dl_P0_PSSCH_PSCCH_r17 != nil, ie.Sl_P0_PSSCH_PSCCH_r17 != nil, ie.Dl_P0_PSFCH_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Dl_P0_PSSCH_PSCCH_r17 optional
			if ie.Dl_P0_PSSCH_PSCCH_r17 != nil {
				if err = extWriter.WriteInteger(*ie.Dl_P0_PSSCH_PSCCH_r17, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Encode Dl_P0_PSSCH_PSCCH_r17", err)
				}
			}
			// encode Sl_P0_PSSCH_PSCCH_r17 optional
			if ie.Sl_P0_PSSCH_PSCCH_r17 != nil {
				if err = extWriter.WriteInteger(*ie.Sl_P0_PSSCH_PSCCH_r17, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Encode Sl_P0_PSSCH_PSCCH_r17", err)
				}
			}
			// encode Dl_P0_PSFCH_r17 optional
			if ie.Dl_P0_PSFCH_r17 != nil {
				if err = extWriter.WriteInteger(*ie.Dl_P0_PSFCH_r17, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Encode Dl_P0_PSFCH_r17", err)
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

func (ie *SL_PowerControl_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_Alpha_PSSCH_PSCCH_r16Present bool
	if Sl_Alpha_PSSCH_PSCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_Alpha_PSSCH_PSCCH_r16Present bool
	if Dl_Alpha_PSSCH_PSCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_P0_PSSCH_PSCCH_r16Present bool
	if Sl_P0_PSSCH_PSCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_P0_PSSCH_PSCCH_r16Present bool
	if Dl_P0_PSSCH_PSCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_Alpha_PSFCH_r16Present bool
	if Dl_Alpha_PSFCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dl_P0_PSFCH_r16Present bool
	if Dl_P0_PSFCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_Sl_MaxTransPower_r16 int64
	if tmp_int_Sl_MaxTransPower_r16, err = r.ReadInteger(&uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_MaxTransPower_r16", err)
	}
	ie.Sl_MaxTransPower_r16 = tmp_int_Sl_MaxTransPower_r16
	if Sl_Alpha_PSSCH_PSCCH_r16Present {
		ie.Sl_Alpha_PSSCH_PSCCH_r16 = new(SL_PowerControl_r16_sl_Alpha_PSSCH_PSCCH_r16)
		if err = ie.Sl_Alpha_PSSCH_PSCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_Alpha_PSSCH_PSCCH_r16", err)
		}
	}
	if Dl_Alpha_PSSCH_PSCCH_r16Present {
		ie.Dl_Alpha_PSSCH_PSCCH_r16 = new(SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16)
		if err = ie.Dl_Alpha_PSSCH_PSCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_Alpha_PSSCH_PSCCH_r16", err)
		}
	}
	if Sl_P0_PSSCH_PSCCH_r16Present {
		var tmp_int_Sl_P0_PSSCH_PSCCH_r16 int64
		if tmp_int_Sl_P0_PSSCH_PSCCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Sl_P0_PSSCH_PSCCH_r16", err)
		}
		ie.Sl_P0_PSSCH_PSCCH_r16 = &tmp_int_Sl_P0_PSSCH_PSCCH_r16
	}
	if Dl_P0_PSSCH_PSCCH_r16Present {
		var tmp_int_Dl_P0_PSSCH_PSCCH_r16 int64
		if tmp_int_Dl_P0_PSSCH_PSCCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Dl_P0_PSSCH_PSCCH_r16", err)
		}
		ie.Dl_P0_PSSCH_PSCCH_r16 = &tmp_int_Dl_P0_PSSCH_PSCCH_r16
	}
	if Dl_Alpha_PSFCH_r16Present {
		ie.Dl_Alpha_PSFCH_r16 = new(SL_PowerControl_r16_dl_Alpha_PSFCH_r16)
		if err = ie.Dl_Alpha_PSFCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_Alpha_PSFCH_r16", err)
		}
	}
	if Dl_P0_PSFCH_r16Present {
		var tmp_int_Dl_P0_PSFCH_r16 int64
		if tmp_int_Dl_P0_PSFCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Dl_P0_PSFCH_r16", err)
		}
		ie.Dl_P0_PSFCH_r16 = &tmp_int_Dl_P0_PSFCH_r16
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

			Dl_P0_PSSCH_PSCCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sl_P0_PSSCH_PSCCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_P0_PSFCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Dl_P0_PSSCH_PSCCH_r17 optional
			if Dl_P0_PSSCH_PSCCH_r17Present {
				var tmp_int_Dl_P0_PSSCH_PSCCH_r17 int64
				if tmp_int_Dl_P0_PSSCH_PSCCH_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Decode Dl_P0_PSSCH_PSCCH_r17", err)
				}
				ie.Dl_P0_PSSCH_PSCCH_r17 = &tmp_int_Dl_P0_PSSCH_PSCCH_r17
			}
			// decode Sl_P0_PSSCH_PSCCH_r17 optional
			if Sl_P0_PSSCH_PSCCH_r17Present {
				var tmp_int_Sl_P0_PSSCH_PSCCH_r17 int64
				if tmp_int_Sl_P0_PSSCH_PSCCH_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Decode Sl_P0_PSSCH_PSCCH_r17", err)
				}
				ie.Sl_P0_PSSCH_PSCCH_r17 = &tmp_int_Sl_P0_PSSCH_PSCCH_r17
			}
			// decode Dl_P0_PSFCH_r17 optional
			if Dl_P0_PSFCH_r17Present {
				var tmp_int_Dl_P0_PSFCH_r17 int64
				if tmp_int_Dl_P0_PSFCH_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Decode Dl_P0_PSFCH_r17", err)
				}
				ie.Dl_P0_PSFCH_r17 = &tmp_int_Dl_P0_PSFCH_r17
			}
		}
	}
	return nil
}
