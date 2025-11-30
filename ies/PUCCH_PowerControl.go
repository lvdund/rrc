package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_PowerControl struct {
	DeltaF_PUCCH_f0              *int64                                           `lb:-16,ub:15,optional`
	DeltaF_PUCCH_f1              *int64                                           `lb:-16,ub:15,optional`
	DeltaF_PUCCH_f2              *int64                                           `lb:-16,ub:15,optional`
	DeltaF_PUCCH_f3              *int64                                           `lb:-16,ub:15,optional`
	DeltaF_PUCCH_f4              *int64                                           `lb:-16,ub:15,optional`
	P0_Set                       []P0_PUCCH                                       `lb:1,ub:maxNrofPUCCH_P0_PerSet,optional`
	PathlossReferenceRSs         []PUCCH_PathlossReferenceRS                      `lb:1,ub:maxNrofPUCCH_PathlossReferenceRSs,optional`
	TwoPUCCH_PC_AdjustmentStates *PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates `optional`
	PathlossReferenceRSs_v1610   *PathlossReferenceRSs_v1610                      `optional,ext-1,setuprelease`
}

func (ie *PUCCH_PowerControl) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.PathlossReferenceRSs_v1610 != nil
	preambleBits := []bool{hasExtensions, ie.DeltaF_PUCCH_f0 != nil, ie.DeltaF_PUCCH_f1 != nil, ie.DeltaF_PUCCH_f2 != nil, ie.DeltaF_PUCCH_f3 != nil, ie.DeltaF_PUCCH_f4 != nil, len(ie.P0_Set) > 0, len(ie.PathlossReferenceRSs) > 0, ie.TwoPUCCH_PC_AdjustmentStates != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DeltaF_PUCCH_f0 != nil {
		if err = w.WriteInteger(*ie.DeltaF_PUCCH_f0, &aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode DeltaF_PUCCH_f0", err)
		}
	}
	if ie.DeltaF_PUCCH_f1 != nil {
		if err = w.WriteInteger(*ie.DeltaF_PUCCH_f1, &aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode DeltaF_PUCCH_f1", err)
		}
	}
	if ie.DeltaF_PUCCH_f2 != nil {
		if err = w.WriteInteger(*ie.DeltaF_PUCCH_f2, &aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode DeltaF_PUCCH_f2", err)
		}
	}
	if ie.DeltaF_PUCCH_f3 != nil {
		if err = w.WriteInteger(*ie.DeltaF_PUCCH_f3, &aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode DeltaF_PUCCH_f3", err)
		}
	}
	if ie.DeltaF_PUCCH_f4 != nil {
		if err = w.WriteInteger(*ie.DeltaF_PUCCH_f4, &aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode DeltaF_PUCCH_f4", err)
		}
	}
	if len(ie.P0_Set) > 0 {
		tmp_P0_Set := utils.NewSequence[*P0_PUCCH]([]*P0_PUCCH{}, aper.Constraint{Lb: 1, Ub: maxNrofPUCCH_P0_PerSet}, false)
		for _, i := range ie.P0_Set {
			tmp_P0_Set.Value = append(tmp_P0_Set.Value, &i)
		}
		if err = tmp_P0_Set.Encode(w); err != nil {
			return utils.WrapError("Encode P0_Set", err)
		}
	}
	if len(ie.PathlossReferenceRSs) > 0 {
		tmp_PathlossReferenceRSs := utils.NewSequence[*PUCCH_PathlossReferenceRS]([]*PUCCH_PathlossReferenceRS{}, aper.Constraint{Lb: 1, Ub: maxNrofPUCCH_PathlossReferenceRSs}, false)
		for _, i := range ie.PathlossReferenceRSs {
			tmp_PathlossReferenceRSs.Value = append(tmp_PathlossReferenceRSs.Value, &i)
		}
		if err = tmp_PathlossReferenceRSs.Encode(w); err != nil {
			return utils.WrapError("Encode PathlossReferenceRSs", err)
		}
	}
	if ie.TwoPUCCH_PC_AdjustmentStates != nil {
		if err = ie.TwoPUCCH_PC_AdjustmentStates.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_PC_AdjustmentStates", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.PathlossReferenceRSs_v1610 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUCCH_PowerControl", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.PathlossReferenceRSs_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode PathlossReferenceRSs_v1610 optional
			if ie.PathlossReferenceRSs_v1610 != nil {
				tmp_PathlossReferenceRSs_v1610 := utils.SetupRelease[*PathlossReferenceRSs_v1610]{
					Setup: ie.PathlossReferenceRSs_v1610,
				}
				if err = tmp_PathlossReferenceRSs_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PathlossReferenceRSs_v1610", err)
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

func (ie *PUCCH_PowerControl) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var DeltaF_PUCCH_f0Present bool
	if DeltaF_PUCCH_f0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DeltaF_PUCCH_f1Present bool
	if DeltaF_PUCCH_f1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DeltaF_PUCCH_f2Present bool
	if DeltaF_PUCCH_f2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DeltaF_PUCCH_f3Present bool
	if DeltaF_PUCCH_f3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DeltaF_PUCCH_f4Present bool
	if DeltaF_PUCCH_f4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var P0_SetPresent bool
	if P0_SetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PathlossReferenceRSsPresent bool
	if PathlossReferenceRSsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_PC_AdjustmentStatesPresent bool
	if TwoPUCCH_PC_AdjustmentStatesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if DeltaF_PUCCH_f0Present {
		var tmp_int_DeltaF_PUCCH_f0 int64
		if tmp_int_DeltaF_PUCCH_f0, err = r.ReadInteger(&aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode DeltaF_PUCCH_f0", err)
		}
		ie.DeltaF_PUCCH_f0 = &tmp_int_DeltaF_PUCCH_f0
	}
	if DeltaF_PUCCH_f1Present {
		var tmp_int_DeltaF_PUCCH_f1 int64
		if tmp_int_DeltaF_PUCCH_f1, err = r.ReadInteger(&aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode DeltaF_PUCCH_f1", err)
		}
		ie.DeltaF_PUCCH_f1 = &tmp_int_DeltaF_PUCCH_f1
	}
	if DeltaF_PUCCH_f2Present {
		var tmp_int_DeltaF_PUCCH_f2 int64
		if tmp_int_DeltaF_PUCCH_f2, err = r.ReadInteger(&aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode DeltaF_PUCCH_f2", err)
		}
		ie.DeltaF_PUCCH_f2 = &tmp_int_DeltaF_PUCCH_f2
	}
	if DeltaF_PUCCH_f3Present {
		var tmp_int_DeltaF_PUCCH_f3 int64
		if tmp_int_DeltaF_PUCCH_f3, err = r.ReadInteger(&aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode DeltaF_PUCCH_f3", err)
		}
		ie.DeltaF_PUCCH_f3 = &tmp_int_DeltaF_PUCCH_f3
	}
	if DeltaF_PUCCH_f4Present {
		var tmp_int_DeltaF_PUCCH_f4 int64
		if tmp_int_DeltaF_PUCCH_f4, err = r.ReadInteger(&aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode DeltaF_PUCCH_f4", err)
		}
		ie.DeltaF_PUCCH_f4 = &tmp_int_DeltaF_PUCCH_f4
	}
	if P0_SetPresent {
		tmp_P0_Set := utils.NewSequence[*P0_PUCCH]([]*P0_PUCCH{}, aper.Constraint{Lb: 1, Ub: maxNrofPUCCH_P0_PerSet}, false)
		fn_P0_Set := func() *P0_PUCCH {
			return new(P0_PUCCH)
		}
		if err = tmp_P0_Set.Decode(r, fn_P0_Set); err != nil {
			return utils.WrapError("Decode P0_Set", err)
		}
		ie.P0_Set = []P0_PUCCH{}
		for _, i := range tmp_P0_Set.Value {
			ie.P0_Set = append(ie.P0_Set, *i)
		}
	}
	if PathlossReferenceRSsPresent {
		tmp_PathlossReferenceRSs := utils.NewSequence[*PUCCH_PathlossReferenceRS]([]*PUCCH_PathlossReferenceRS{}, aper.Constraint{Lb: 1, Ub: maxNrofPUCCH_PathlossReferenceRSs}, false)
		fn_PathlossReferenceRSs := func() *PUCCH_PathlossReferenceRS {
			return new(PUCCH_PathlossReferenceRS)
		}
		if err = tmp_PathlossReferenceRSs.Decode(r, fn_PathlossReferenceRSs); err != nil {
			return utils.WrapError("Decode PathlossReferenceRSs", err)
		}
		ie.PathlossReferenceRSs = []PUCCH_PathlossReferenceRS{}
		for _, i := range tmp_PathlossReferenceRSs.Value {
			ie.PathlossReferenceRSs = append(ie.PathlossReferenceRSs, *i)
		}
	}
	if TwoPUCCH_PC_AdjustmentStatesPresent {
		ie.TwoPUCCH_PC_AdjustmentStates = new(PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates)
		if err = ie.TwoPUCCH_PC_AdjustmentStates.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_PC_AdjustmentStates", err)
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

			PathlossReferenceRSs_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode PathlossReferenceRSs_v1610 optional
			if PathlossReferenceRSs_v1610Present {
				tmp_PathlossReferenceRSs_v1610 := utils.SetupRelease[*PathlossReferenceRSs_v1610]{}
				if err = tmp_PathlossReferenceRSs_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode PathlossReferenceRSs_v1610", err)
				}
				ie.PathlossReferenceRSs_v1610 = tmp_PathlossReferenceRSs_v1610.Setup
			}
		}
	}
	return nil
}
