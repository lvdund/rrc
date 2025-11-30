package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_ConfiguredGrantConfig_r16 struct {
	Sl_ConfigIndexCG_r16            SL_ConfigIndexCG_r16                                          `madatory`
	Sl_PeriodCG_r16                 *SL_PeriodCG_r16                                              `optional`
	Sl_NrOfHARQ_Processes_r16       *int64                                                        `lb:1,ub:16,optional`
	Sl_HARQ_ProcID_offset_r16       *int64                                                        `lb:0,ub:15,optional`
	Sl_CG_MaxTransNumList_r16       *SL_CG_MaxTransNumList_r16                                    `optional`
	Rrc_ConfiguredSidelinkGrant_r16 *SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16 `lb:0,ub:496,optional`
	Sl_N1PUCCH_AN_Type2_r16         *PUCCH_ResourceId                                             `optional,ext-1`
}

func (ie *SL_ConfiguredGrantConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Sl_N1PUCCH_AN_Type2_r16 != nil
	preambleBits := []bool{hasExtensions, ie.Sl_PeriodCG_r16 != nil, ie.Sl_NrOfHARQ_Processes_r16 != nil, ie.Sl_HARQ_ProcID_offset_r16 != nil, ie.Sl_CG_MaxTransNumList_r16 != nil, ie.Rrc_ConfiguredSidelinkGrant_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_ConfigIndexCG_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_ConfigIndexCG_r16", err)
	}
	if ie.Sl_PeriodCG_r16 != nil {
		if err = ie.Sl_PeriodCG_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_PeriodCG_r16", err)
		}
	}
	if ie.Sl_NrOfHARQ_Processes_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_NrOfHARQ_Processes_r16, &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Sl_NrOfHARQ_Processes_r16", err)
		}
	}
	if ie.Sl_HARQ_ProcID_offset_r16 != nil {
		if err = w.WriteInteger(*ie.Sl_HARQ_ProcID_offset_r16, &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Sl_HARQ_ProcID_offset_r16", err)
		}
	}
	if ie.Sl_CG_MaxTransNumList_r16 != nil {
		if err = ie.Sl_CG_MaxTransNumList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_CG_MaxTransNumList_r16", err)
		}
	}
	if ie.Rrc_ConfiguredSidelinkGrant_r16 != nil {
		if err = ie.Rrc_ConfiguredSidelinkGrant_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rrc_ConfiguredSidelinkGrant_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Sl_N1PUCCH_AN_Type2_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_ConfiguredGrantConfig_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Sl_N1PUCCH_AN_Type2_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Sl_N1PUCCH_AN_Type2_r16 optional
			if ie.Sl_N1PUCCH_AN_Type2_r16 != nil {
				if err = ie.Sl_N1PUCCH_AN_Type2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sl_N1PUCCH_AN_Type2_r16", err)
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

func (ie *SL_ConfiguredGrantConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_PeriodCG_r16Present bool
	if Sl_PeriodCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_NrOfHARQ_Processes_r16Present bool
	if Sl_NrOfHARQ_Processes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_HARQ_ProcID_offset_r16Present bool
	if Sl_HARQ_ProcID_offset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_CG_MaxTransNumList_r16Present bool
	if Sl_CG_MaxTransNumList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rrc_ConfiguredSidelinkGrant_r16Present bool
	if Rrc_ConfiguredSidelinkGrant_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_ConfigIndexCG_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_ConfigIndexCG_r16", err)
	}
	if Sl_PeriodCG_r16Present {
		ie.Sl_PeriodCG_r16 = new(SL_PeriodCG_r16)
		if err = ie.Sl_PeriodCG_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_PeriodCG_r16", err)
		}
	}
	if Sl_NrOfHARQ_Processes_r16Present {
		var tmp_int_Sl_NrOfHARQ_Processes_r16 int64
		if tmp_int_Sl_NrOfHARQ_Processes_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Sl_NrOfHARQ_Processes_r16", err)
		}
		ie.Sl_NrOfHARQ_Processes_r16 = &tmp_int_Sl_NrOfHARQ_Processes_r16
	}
	if Sl_HARQ_ProcID_offset_r16Present {
		var tmp_int_Sl_HARQ_ProcID_offset_r16 int64
		if tmp_int_Sl_HARQ_ProcID_offset_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Sl_HARQ_ProcID_offset_r16", err)
		}
		ie.Sl_HARQ_ProcID_offset_r16 = &tmp_int_Sl_HARQ_ProcID_offset_r16
	}
	if Sl_CG_MaxTransNumList_r16Present {
		ie.Sl_CG_MaxTransNumList_r16 = new(SL_CG_MaxTransNumList_r16)
		if err = ie.Sl_CG_MaxTransNumList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_CG_MaxTransNumList_r16", err)
		}
	}
	if Rrc_ConfiguredSidelinkGrant_r16Present {
		ie.Rrc_ConfiguredSidelinkGrant_r16 = new(SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16)
		if err = ie.Rrc_ConfiguredSidelinkGrant_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rrc_ConfiguredSidelinkGrant_r16", err)
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

			Sl_N1PUCCH_AN_Type2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Sl_N1PUCCH_AN_Type2_r16 optional
			if Sl_N1PUCCH_AN_Type2_r16Present {
				ie.Sl_N1PUCCH_AN_Type2_r16 = new(PUCCH_ResourceId)
				if err = ie.Sl_N1PUCCH_AN_Type2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sl_N1PUCCH_AN_Type2_r16", err)
				}
			}
		}
	}
	return nil
}
