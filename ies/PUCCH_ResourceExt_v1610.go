package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ResourceExt_v1610 struct {
	InterlaceAllocation_r16       *PUCCH_ResourceExt_v1610_interlaceAllocation_r16       `lb:0,ub:4,optional`
	Format_v1610                  *PUCCH_ResourceExt_v1610_format_v1610                  `lb:0,ub:9,optional`
	FormatExt_v1700               *PUCCH_ResourceExt_v1610_formatExt_v1700               `lb:1,ub:16,optional,ext-1`
	Pucch_RepetitionNrofSlots_r17 *PUCCH_ResourceExt_v1610_pucch_RepetitionNrofSlots_r17 `optional,ext-1`
}

func (ie *PUCCH_ResourceExt_v1610) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.FormatExt_v1700 != nil || ie.Pucch_RepetitionNrofSlots_r17 != nil
	preambleBits := []bool{hasExtensions, ie.InterlaceAllocation_r16 != nil, ie.Format_v1610 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.InterlaceAllocation_r16 != nil {
		if err = ie.InterlaceAllocation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode InterlaceAllocation_r16", err)
		}
	}
	if ie.Format_v1610 != nil {
		if err = ie.Format_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Format_v1610", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.FormatExt_v1700 != nil || ie.Pucch_RepetitionNrofSlots_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUCCH_ResourceExt_v1610", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.FormatExt_v1700 != nil, ie.Pucch_RepetitionNrofSlots_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FormatExt_v1700 optional
			if ie.FormatExt_v1700 != nil {
				if err = ie.FormatExt_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FormatExt_v1700", err)
				}
			}
			// encode Pucch_RepetitionNrofSlots_r17 optional
			if ie.Pucch_RepetitionNrofSlots_r17 != nil {
				if err = ie.Pucch_RepetitionNrofSlots_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_RepetitionNrofSlots_r17", err)
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

func (ie *PUCCH_ResourceExt_v1610) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var InterlaceAllocation_r16Present bool
	if InterlaceAllocation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Format_v1610Present bool
	if Format_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	if InterlaceAllocation_r16Present {
		ie.InterlaceAllocation_r16 = new(PUCCH_ResourceExt_v1610_interlaceAllocation_r16)
		if err = ie.InterlaceAllocation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode InterlaceAllocation_r16", err)
		}
	}
	if Format_v1610Present {
		ie.Format_v1610 = new(PUCCH_ResourceExt_v1610_format_v1610)
		if err = ie.Format_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Format_v1610", err)
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

			FormatExt_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_RepetitionNrofSlots_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FormatExt_v1700 optional
			if FormatExt_v1700Present {
				ie.FormatExt_v1700 = new(PUCCH_ResourceExt_v1610_formatExt_v1700)
				if err = ie.FormatExt_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode FormatExt_v1700", err)
				}
			}
			// decode Pucch_RepetitionNrofSlots_r17 optional
			if Pucch_RepetitionNrofSlots_r17Present {
				ie.Pucch_RepetitionNrofSlots_r17 = new(PUCCH_ResourceExt_v1610_pucch_RepetitionNrofSlots_r17)
				if err = ie.Pucch_RepetitionNrofSlots_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pucch_RepetitionNrofSlots_r17", err)
				}
			}
		}
	}
	return nil
}
