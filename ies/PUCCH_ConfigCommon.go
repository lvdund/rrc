package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ConfigCommon struct {
	Pucch_ResourceCommon           *int64                                      `lb:0,ub:15,optional`
	Pucch_GroupHopping             PUCCH_ConfigCommon_pucch_GroupHopping       `madatory`
	HoppingId                      *int64                                      `lb:0,ub:1023,optional`
	P0_nominal                     *int64                                      `lb:-202,ub:24,optional`
	NrofPRBs                       *int64                                      `lb:1,ub:16,optional,ext-1`
	Intra_SlotFH_r17               *PUCCH_ConfigCommon_intra_SlotFH_r17        `optional,ext-1`
	Pucch_ResourceCommonRedCap_r17 *int64                                      `lb:0,ub:15,optional,ext-1`
	AdditionalPRBOffset_r17        *PUCCH_ConfigCommon_additionalPRBOffset_r17 `optional,ext-1`
}

func (ie *PUCCH_ConfigCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.NrofPRBs != nil || ie.Intra_SlotFH_r17 != nil || ie.Pucch_ResourceCommonRedCap_r17 != nil || ie.AdditionalPRBOffset_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Pucch_ResourceCommon != nil, ie.HoppingId != nil, ie.P0_nominal != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pucch_ResourceCommon != nil {
		if err = w.WriteInteger(*ie.Pucch_ResourceCommon, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Pucch_ResourceCommon", err)
		}
	}
	if err = ie.Pucch_GroupHopping.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_GroupHopping", err)
	}
	if ie.HoppingId != nil {
		if err = w.WriteInteger(*ie.HoppingId, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode HoppingId", err)
		}
	}
	if ie.P0_nominal != nil {
		if err = w.WriteInteger(*ie.P0_nominal, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Encode P0_nominal", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.NrofPRBs != nil || ie.Intra_SlotFH_r17 != nil || ie.Pucch_ResourceCommonRedCap_r17 != nil || ie.AdditionalPRBOffset_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUCCH_ConfigCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.NrofPRBs != nil, ie.Intra_SlotFH_r17 != nil, ie.Pucch_ResourceCommonRedCap_r17 != nil, ie.AdditionalPRBOffset_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode NrofPRBs optional
			if ie.NrofPRBs != nil {
				if err = extWriter.WriteInteger(*ie.NrofPRBs, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
					return utils.WrapError("Encode NrofPRBs", err)
				}
			}
			// encode Intra_SlotFH_r17 optional
			if ie.Intra_SlotFH_r17 != nil {
				if err = ie.Intra_SlotFH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Intra_SlotFH_r17", err)
				}
			}
			// encode Pucch_ResourceCommonRedCap_r17 optional
			if ie.Pucch_ResourceCommonRedCap_r17 != nil {
				if err = extWriter.WriteInteger(*ie.Pucch_ResourceCommonRedCap_r17, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode Pucch_ResourceCommonRedCap_r17", err)
				}
			}
			// encode AdditionalPRBOffset_r17 optional
			if ie.AdditionalPRBOffset_r17 != nil {
				if err = ie.AdditionalPRBOffset_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AdditionalPRBOffset_r17", err)
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

func (ie *PUCCH_ConfigCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_ResourceCommonPresent bool
	if Pucch_ResourceCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var HoppingIdPresent bool
	if HoppingIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var P0_nominalPresent bool
	if P0_nominalPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Pucch_ResourceCommonPresent {
		var tmp_int_Pucch_ResourceCommon int64
		if tmp_int_Pucch_ResourceCommon, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Pucch_ResourceCommon", err)
		}
		ie.Pucch_ResourceCommon = &tmp_int_Pucch_ResourceCommon
	}
	if err = ie.Pucch_GroupHopping.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_GroupHopping", err)
	}
	if HoppingIdPresent {
		var tmp_int_HoppingId int64
		if tmp_int_HoppingId, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode HoppingId", err)
		}
		ie.HoppingId = &tmp_int_HoppingId
	}
	if P0_nominalPresent {
		var tmp_int_P0_nominal int64
		if tmp_int_P0_nominal, err = r.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode P0_nominal", err)
		}
		ie.P0_nominal = &tmp_int_P0_nominal
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

			NrofPRBsPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Intra_SlotFH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_ResourceCommonRedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AdditionalPRBOffset_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode NrofPRBs optional
			if NrofPRBsPresent {
				var tmp_int_NrofPRBs int64
				if tmp_int_NrofPRBs, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
					return utils.WrapError("Decode NrofPRBs", err)
				}
				ie.NrofPRBs = &tmp_int_NrofPRBs
			}
			// decode Intra_SlotFH_r17 optional
			if Intra_SlotFH_r17Present {
				ie.Intra_SlotFH_r17 = new(PUCCH_ConfigCommon_intra_SlotFH_r17)
				if err = ie.Intra_SlotFH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Intra_SlotFH_r17", err)
				}
			}
			// decode Pucch_ResourceCommonRedCap_r17 optional
			if Pucch_ResourceCommonRedCap_r17Present {
				var tmp_int_Pucch_ResourceCommonRedCap_r17 int64
				if tmp_int_Pucch_ResourceCommonRedCap_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode Pucch_ResourceCommonRedCap_r17", err)
				}
				ie.Pucch_ResourceCommonRedCap_r17 = &tmp_int_Pucch_ResourceCommonRedCap_r17
			}
			// decode AdditionalPRBOffset_r17 optional
			if AdditionalPRBOffset_r17Present {
				ie.AdditionalPRBOffset_r17 = new(PUCCH_ConfigCommon_additionalPRBOffset_r17)
				if err = ie.AdditionalPRBOffset_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode AdditionalPRBOffset_r17", err)
				}
			}
		}
	}
	return nil
}
