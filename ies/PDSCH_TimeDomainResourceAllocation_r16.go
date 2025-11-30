package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_TimeDomainResourceAllocation_r16 struct {
	K0_r16                   *int64                                                         `lb:0,ub:32,optional`
	MappingType_r16          PDSCH_TimeDomainResourceAllocation_r16_mappingType_r16         `madatory`
	StartSymbolAndLength_r16 int64                                                          `lb:0,ub:127,madatory`
	RepetitionNumber_r16     *PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16   `optional`
	K0_v1710                 *int64                                                         `lb:33,ub:128,optional,ext-1`
	RepetitionNumber_v1730   *PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730 `optional,ext-2`
}

func (ie *PDSCH_TimeDomainResourceAllocation_r16) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.K0_v1710 != nil || ie.RepetitionNumber_v1730 != nil
	preambleBits := []bool{hasExtensions, ie.K0_r16 != nil, ie.RepetitionNumber_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.K0_r16 != nil {
		if err = w.WriteInteger(*ie.K0_r16, &aper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode K0_r16", err)
		}
	}
	if err = ie.MappingType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MappingType_r16", err)
	}
	if err = w.WriteInteger(ie.StartSymbolAndLength_r16, &aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("WriteInteger StartSymbolAndLength_r16", err)
	}
	if ie.RepetitionNumber_r16 != nil {
		if err = ie.RepetitionNumber_r16.Encode(w); err != nil {
			return utils.WrapError("Encode RepetitionNumber_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.K0_v1710 != nil, ie.RepetitionNumber_v1730 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDSCH_TimeDomainResourceAllocation_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.K0_v1710 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode K0_v1710 optional
			if ie.K0_v1710 != nil {
				if err = extWriter.WriteInteger(*ie.K0_v1710, &aper.Constraint{Lb: 33, Ub: 128}, false); err != nil {
					return utils.WrapError("Encode K0_v1710", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.RepetitionNumber_v1730 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RepetitionNumber_v1730 optional
			if ie.RepetitionNumber_v1730 != nil {
				if err = ie.RepetitionNumber_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RepetitionNumber_v1730", err)
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

func (ie *PDSCH_TimeDomainResourceAllocation_r16) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var K0_r16Present bool
	if K0_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RepetitionNumber_r16Present bool
	if RepetitionNumber_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if K0_r16Present {
		var tmp_int_K0_r16 int64
		if tmp_int_K0_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode K0_r16", err)
		}
		ie.K0_r16 = &tmp_int_K0_r16
	}
	if err = ie.MappingType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MappingType_r16", err)
	}
	var tmp_int_StartSymbolAndLength_r16 int64
	if tmp_int_StartSymbolAndLength_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("ReadInteger StartSymbolAndLength_r16", err)
	}
	ie.StartSymbolAndLength_r16 = tmp_int_StartSymbolAndLength_r16
	if RepetitionNumber_r16Present {
		ie.RepetitionNumber_r16 = new(PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16)
		if err = ie.RepetitionNumber_r16.Decode(r); err != nil {
			return utils.WrapError("Decode RepetitionNumber_r16", err)
		}
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			K0_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode K0_v1710 optional
			if K0_v1710Present {
				var tmp_int_K0_v1710 int64
				if tmp_int_K0_v1710, err = extReader.ReadInteger(&aper.Constraint{Lb: 33, Ub: 128}, false); err != nil {
					return utils.WrapError("Decode K0_v1710", err)
				}
				ie.K0_v1710 = &tmp_int_K0_v1710
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			RepetitionNumber_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RepetitionNumber_v1730 optional
			if RepetitionNumber_v1730Present {
				ie.RepetitionNumber_v1730 = new(PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730)
				if err = ie.RepetitionNumber_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode RepetitionNumber_v1730", err)
				}
			}
		}
	}
	return nil
}
