package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_ServingCellConfig struct {
	SlotFormatIndicator          *SlotFormatIndicator       `optional,setuprelease`
	AvailabilityIndicator_r16    *AvailabilityIndicator_r16 `optional,ext-1,setuprelease`
	SearchSpaceSwitchTimer_r16   *int64                     `lb:1,ub:80,optional,ext-1`
	SearchSpaceSwitchTimer_v1710 *int64                     `lb:81,ub:1280,optional,ext-2`
}

func (ie *PDCCH_ServingCellConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.AvailabilityIndicator_r16 != nil || ie.SearchSpaceSwitchTimer_r16 != nil || ie.SearchSpaceSwitchTimer_v1710 != nil
	preambleBits := []bool{hasExtensions, ie.SlotFormatIndicator != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SlotFormatIndicator != nil {
		tmp_SlotFormatIndicator := utils.SetupRelease[*SlotFormatIndicator]{
			Setup: ie.SlotFormatIndicator,
		}
		if err = tmp_SlotFormatIndicator.Encode(w); err != nil {
			return utils.WrapError("Encode SlotFormatIndicator", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.AvailabilityIndicator_r16 != nil || ie.SearchSpaceSwitchTimer_r16 != nil, ie.SearchSpaceSwitchTimer_v1710 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDCCH_ServingCellConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.AvailabilityIndicator_r16 != nil, ie.SearchSpaceSwitchTimer_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode AvailabilityIndicator_r16 optional
			if ie.AvailabilityIndicator_r16 != nil {
				tmp_AvailabilityIndicator_r16 := utils.SetupRelease[*AvailabilityIndicator_r16]{
					Setup: ie.AvailabilityIndicator_r16,
				}
				if err = tmp_AvailabilityIndicator_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AvailabilityIndicator_r16", err)
				}
			}
			// encode SearchSpaceSwitchTimer_r16 optional
			if ie.SearchSpaceSwitchTimer_r16 != nil {
				if err = extWriter.WriteInteger(*ie.SearchSpaceSwitchTimer_r16, &uper.Constraint{Lb: 1, Ub: 80}, false); err != nil {
					return utils.WrapError("Encode SearchSpaceSwitchTimer_r16", err)
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
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.SearchSpaceSwitchTimer_v1710 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SearchSpaceSwitchTimer_v1710 optional
			if ie.SearchSpaceSwitchTimer_v1710 != nil {
				if err = extWriter.WriteInteger(*ie.SearchSpaceSwitchTimer_v1710, &uper.Constraint{Lb: 81, Ub: 1280}, false); err != nil {
					return utils.WrapError("Encode SearchSpaceSwitchTimer_v1710", err)
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

func (ie *PDCCH_ServingCellConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var SlotFormatIndicatorPresent bool
	if SlotFormatIndicatorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SlotFormatIndicatorPresent {
		tmp_SlotFormatIndicator := utils.SetupRelease[*SlotFormatIndicator]{}
		if err = tmp_SlotFormatIndicator.Decode(r); err != nil {
			return utils.WrapError("Decode SlotFormatIndicator", err)
		}
		ie.SlotFormatIndicator = tmp_SlotFormatIndicator.Setup
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			AvailabilityIndicator_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SearchSpaceSwitchTimer_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode AvailabilityIndicator_r16 optional
			if AvailabilityIndicator_r16Present {
				tmp_AvailabilityIndicator_r16 := utils.SetupRelease[*AvailabilityIndicator_r16]{}
				if err = tmp_AvailabilityIndicator_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode AvailabilityIndicator_r16", err)
				}
				ie.AvailabilityIndicator_r16 = tmp_AvailabilityIndicator_r16.Setup
			}
			// decode SearchSpaceSwitchTimer_r16 optional
			if SearchSpaceSwitchTimer_r16Present {
				var tmp_int_SearchSpaceSwitchTimer_r16 int64
				if tmp_int_SearchSpaceSwitchTimer_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 80}, false); err != nil {
					return utils.WrapError("Decode SearchSpaceSwitchTimer_r16", err)
				}
				ie.SearchSpaceSwitchTimer_r16 = &tmp_int_SearchSpaceSwitchTimer_r16
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			SearchSpaceSwitchTimer_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SearchSpaceSwitchTimer_v1710 optional
			if SearchSpaceSwitchTimer_v1710Present {
				var tmp_int_SearchSpaceSwitchTimer_v1710 int64
				if tmp_int_SearchSpaceSwitchTimer_v1710, err = extReader.ReadInteger(&uper.Constraint{Lb: 81, Ub: 1280}, false); err != nil {
					return utils.WrapError("Decode SearchSpaceSwitchTimer_v1710", err)
				}
				ie.SearchSpaceSwitchTimer_v1710 = &tmp_int_SearchSpaceSwitchTimer_v1710
			}
		}
	}
	return nil
}
