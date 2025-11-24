package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SlotFormatCombinationsPerCell struct {
	ServingCellId          ServCellIndex                                         `madatory`
	SubcarrierSpacing      SubcarrierSpacing                                     `madatory`
	SubcarrierSpacing2     *SubcarrierSpacing                                    `optional`
	SlotFormatCombinations []SlotFormatCombination                               `lb:1,ub:maxNrofSlotFormatCombinationsPerSet,optional`
	PositionInDCI          *int64                                                `lb:0,ub:maxSFI_DCI_PayloadSize_1,optional`
	EnableConfiguredUL_r16 *SlotFormatCombinationsPerCell_enableConfiguredUL_r16 `optional,ext-1`
}

func (ie *SlotFormatCombinationsPerCell) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.EnableConfiguredUL_r16 != nil
	preambleBits := []bool{hasExtensions, ie.SubcarrierSpacing2 != nil, len(ie.SlotFormatCombinations) > 0, ie.PositionInDCI != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ServingCellId.Encode(w); err != nil {
		return utils.WrapError("Encode ServingCellId", err)
	}
	if err = ie.SubcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode SubcarrierSpacing", err)
	}
	if ie.SubcarrierSpacing2 != nil {
		if err = ie.SubcarrierSpacing2.Encode(w); err != nil {
			return utils.WrapError("Encode SubcarrierSpacing2", err)
		}
	}
	if len(ie.SlotFormatCombinations) > 0 {
		tmp_SlotFormatCombinations := utils.NewSequence[*SlotFormatCombination]([]*SlotFormatCombination{}, uper.Constraint{Lb: 1, Ub: maxNrofSlotFormatCombinationsPerSet}, false)
		for _, i := range ie.SlotFormatCombinations {
			tmp_SlotFormatCombinations.Value = append(tmp_SlotFormatCombinations.Value, &i)
		}
		if err = tmp_SlotFormatCombinations.Encode(w); err != nil {
			return utils.WrapError("Encode SlotFormatCombinations", err)
		}
	}
	if ie.PositionInDCI != nil {
		if err = w.WriteInteger(*ie.PositionInDCI, &uper.Constraint{Lb: 0, Ub: maxSFI_DCI_PayloadSize_1}, false); err != nil {
			return utils.WrapError("Encode PositionInDCI", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.EnableConfiguredUL_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SlotFormatCombinationsPerCell", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.EnableConfiguredUL_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode EnableConfiguredUL_r16 optional
			if ie.EnableConfiguredUL_r16 != nil {
				if err = ie.EnableConfiguredUL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnableConfiguredUL_r16", err)
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

func (ie *SlotFormatCombinationsPerCell) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var SubcarrierSpacing2Present bool
	if SubcarrierSpacing2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SlotFormatCombinationsPresent bool
	if SlotFormatCombinationsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PositionInDCIPresent bool
	if PositionInDCIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ServingCellId.Decode(r); err != nil {
		return utils.WrapError("Decode ServingCellId", err)
	}
	if err = ie.SubcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode SubcarrierSpacing", err)
	}
	if SubcarrierSpacing2Present {
		ie.SubcarrierSpacing2 = new(SubcarrierSpacing)
		if err = ie.SubcarrierSpacing2.Decode(r); err != nil {
			return utils.WrapError("Decode SubcarrierSpacing2", err)
		}
	}
	if SlotFormatCombinationsPresent {
		tmp_SlotFormatCombinations := utils.NewSequence[*SlotFormatCombination]([]*SlotFormatCombination{}, uper.Constraint{Lb: 1, Ub: maxNrofSlotFormatCombinationsPerSet}, false)
		fn_SlotFormatCombinations := func() *SlotFormatCombination {
			return new(SlotFormatCombination)
		}
		if err = tmp_SlotFormatCombinations.Decode(r, fn_SlotFormatCombinations); err != nil {
			return utils.WrapError("Decode SlotFormatCombinations", err)
		}
		ie.SlotFormatCombinations = []SlotFormatCombination{}
		for _, i := range tmp_SlotFormatCombinations.Value {
			ie.SlotFormatCombinations = append(ie.SlotFormatCombinations, *i)
		}
	}
	if PositionInDCIPresent {
		var tmp_int_PositionInDCI int64
		if tmp_int_PositionInDCI, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxSFI_DCI_PayloadSize_1}, false); err != nil {
			return utils.WrapError("Decode PositionInDCI", err)
		}
		ie.PositionInDCI = &tmp_int_PositionInDCI
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

			EnableConfiguredUL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode EnableConfiguredUL_r16 optional
			if EnableConfiguredUL_r16Present {
				ie.EnableConfiguredUL_r16 = new(SlotFormatCombinationsPerCell_enableConfiguredUL_r16)
				if err = ie.EnableConfiguredUL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnableConfiguredUL_r16", err)
				}
			}
		}
	}
	return nil
}
