package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_Allocation_r16 struct {
	MappingType_r16            *PUSCH_Allocation_r16_mappingType_r16            `optional`
	StartSymbolAndLength_r16   *int64                                           `lb:0,ub:127,optional`
	StartSymbol_r16            *int64                                           `lb:0,ub:13,optional`
	Length_r16                 *int64                                           `lb:1,ub:14,optional`
	NumberOfRepetitions_r16    *PUSCH_Allocation_r16_numberOfRepetitions_r16    `optional`
	NumberOfRepetitionsExt_r17 *PUSCH_Allocation_r16_numberOfRepetitionsExt_r17 `optional,ext-1`
	NumberOfSlotsTBoMS_r17     *PUSCH_Allocation_r16_numberOfSlotsTBoMS_r17     `optional,ext-1`
	ExtendedK2_r17             *int64                                           `lb:0,ub:128,optional,ext-1`
}

func (ie *PUSCH_Allocation_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.NumberOfRepetitionsExt_r17 != nil || ie.NumberOfSlotsTBoMS_r17 != nil || ie.ExtendedK2_r17 != nil
	preambleBits := []bool{hasExtensions, ie.MappingType_r16 != nil, ie.StartSymbolAndLength_r16 != nil, ie.StartSymbol_r16 != nil, ie.Length_r16 != nil, ie.NumberOfRepetitions_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MappingType_r16 != nil {
		if err = ie.MappingType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MappingType_r16", err)
		}
	}
	if ie.StartSymbolAndLength_r16 != nil {
		if err = w.WriteInteger(*ie.StartSymbolAndLength_r16, &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Encode StartSymbolAndLength_r16", err)
		}
	}
	if ie.StartSymbol_r16 != nil {
		if err = w.WriteInteger(*ie.StartSymbol_r16, &uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
			return utils.WrapError("Encode StartSymbol_r16", err)
		}
	}
	if ie.Length_r16 != nil {
		if err = w.WriteInteger(*ie.Length_r16, &uper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
			return utils.WrapError("Encode Length_r16", err)
		}
	}
	if ie.NumberOfRepetitions_r16 != nil {
		if err = ie.NumberOfRepetitions_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NumberOfRepetitions_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.NumberOfRepetitionsExt_r17 != nil || ie.NumberOfSlotsTBoMS_r17 != nil || ie.ExtendedK2_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUSCH_Allocation_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.NumberOfRepetitionsExt_r17 != nil, ie.NumberOfSlotsTBoMS_r17 != nil, ie.ExtendedK2_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode NumberOfRepetitionsExt_r17 optional
			if ie.NumberOfRepetitionsExt_r17 != nil {
				if err = ie.NumberOfRepetitionsExt_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NumberOfRepetitionsExt_r17", err)
				}
			}
			// encode NumberOfSlotsTBoMS_r17 optional
			if ie.NumberOfSlotsTBoMS_r17 != nil {
				if err = ie.NumberOfSlotsTBoMS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NumberOfSlotsTBoMS_r17", err)
				}
			}
			// encode ExtendedK2_r17 optional
			if ie.ExtendedK2_r17 != nil {
				if err = extWriter.WriteInteger(*ie.ExtendedK2_r17, &uper.Constraint{Lb: 0, Ub: 128}, false); err != nil {
					return utils.WrapError("Encode ExtendedK2_r17", err)
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

func (ie *PUSCH_Allocation_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var MappingType_r16Present bool
	if MappingType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var StartSymbolAndLength_r16Present bool
	if StartSymbolAndLength_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var StartSymbol_r16Present bool
	if StartSymbol_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Length_r16Present bool
	if Length_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NumberOfRepetitions_r16Present bool
	if NumberOfRepetitions_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MappingType_r16Present {
		ie.MappingType_r16 = new(PUSCH_Allocation_r16_mappingType_r16)
		if err = ie.MappingType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MappingType_r16", err)
		}
	}
	if StartSymbolAndLength_r16Present {
		var tmp_int_StartSymbolAndLength_r16 int64
		if tmp_int_StartSymbolAndLength_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode StartSymbolAndLength_r16", err)
		}
		ie.StartSymbolAndLength_r16 = &tmp_int_StartSymbolAndLength_r16
	}
	if StartSymbol_r16Present {
		var tmp_int_StartSymbol_r16 int64
		if tmp_int_StartSymbol_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
			return utils.WrapError("Decode StartSymbol_r16", err)
		}
		ie.StartSymbol_r16 = &tmp_int_StartSymbol_r16
	}
	if Length_r16Present {
		var tmp_int_Length_r16 int64
		if tmp_int_Length_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
			return utils.WrapError("Decode Length_r16", err)
		}
		ie.Length_r16 = &tmp_int_Length_r16
	}
	if NumberOfRepetitions_r16Present {
		ie.NumberOfRepetitions_r16 = new(PUSCH_Allocation_r16_numberOfRepetitions_r16)
		if err = ie.NumberOfRepetitions_r16.Decode(r); err != nil {
			return utils.WrapError("Decode NumberOfRepetitions_r16", err)
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

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			NumberOfRepetitionsExt_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NumberOfSlotsTBoMS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ExtendedK2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode NumberOfRepetitionsExt_r17 optional
			if NumberOfRepetitionsExt_r17Present {
				ie.NumberOfRepetitionsExt_r17 = new(PUSCH_Allocation_r16_numberOfRepetitionsExt_r17)
				if err = ie.NumberOfRepetitionsExt_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode NumberOfRepetitionsExt_r17", err)
				}
			}
			// decode NumberOfSlotsTBoMS_r17 optional
			if NumberOfSlotsTBoMS_r17Present {
				ie.NumberOfSlotsTBoMS_r17 = new(PUSCH_Allocation_r16_numberOfSlotsTBoMS_r17)
				if err = ie.NumberOfSlotsTBoMS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode NumberOfSlotsTBoMS_r17", err)
				}
			}
			// decode ExtendedK2_r17 optional
			if ExtendedK2_r17Present {
				var tmp_int_ExtendedK2_r17 int64
				if tmp_int_ExtendedK2_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 128}, false); err != nil {
					return utils.WrapError("Decode ExtendedK2_r17", err)
				}
				ie.ExtendedK2_r17 = &tmp_int_ExtendedK2_r17
			}
		}
	}
	return nil
}
