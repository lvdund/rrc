package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellReselectionPriorities struct {
	FreqPriorityListEUTRA                *FreqPriorityListEUTRA                `optional`
	FreqPriorityListNR                   *FreqPriorityListNR                   `optional`
	T320                                 *CellReselectionPriorities_t320       `optional`
	FreqPriorityListDedicatedSlicing_r17 *FreqPriorityListDedicatedSlicing_r17 `optional,ext-1`
}

func (ie *CellReselectionPriorities) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.FreqPriorityListDedicatedSlicing_r17 != nil
	preambleBits := []bool{hasExtensions, ie.FreqPriorityListEUTRA != nil, ie.FreqPriorityListNR != nil, ie.T320 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FreqPriorityListEUTRA != nil {
		if err = ie.FreqPriorityListEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode FreqPriorityListEUTRA", err)
		}
	}
	if ie.FreqPriorityListNR != nil {
		if err = ie.FreqPriorityListNR.Encode(w); err != nil {
			return utils.WrapError("Encode FreqPriorityListNR", err)
		}
	}
	if ie.T320 != nil {
		if err = ie.T320.Encode(w); err != nil {
			return utils.WrapError("Encode T320", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.FreqPriorityListDedicatedSlicing_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CellReselectionPriorities", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.FreqPriorityListDedicatedSlicing_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FreqPriorityListDedicatedSlicing_r17 optional
			if ie.FreqPriorityListDedicatedSlicing_r17 != nil {
				if err = ie.FreqPriorityListDedicatedSlicing_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FreqPriorityListDedicatedSlicing_r17", err)
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

func (ie *CellReselectionPriorities) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var FreqPriorityListEUTRAPresent bool
	if FreqPriorityListEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FreqPriorityListNRPresent bool
	if FreqPriorityListNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var T320Present bool
	if T320Present, err = r.ReadBool(); err != nil {
		return err
	}
	if FreqPriorityListEUTRAPresent {
		ie.FreqPriorityListEUTRA = new(FreqPriorityListEUTRA)
		if err = ie.FreqPriorityListEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode FreqPriorityListEUTRA", err)
		}
	}
	if FreqPriorityListNRPresent {
		ie.FreqPriorityListNR = new(FreqPriorityListNR)
		if err = ie.FreqPriorityListNR.Decode(r); err != nil {
			return utils.WrapError("Decode FreqPriorityListNR", err)
		}
	}
	if T320Present {
		ie.T320 = new(CellReselectionPriorities_t320)
		if err = ie.T320.Decode(r); err != nil {
			return utils.WrapError("Decode T320", err)
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

			FreqPriorityListDedicatedSlicing_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FreqPriorityListDedicatedSlicing_r17 optional
			if FreqPriorityListDedicatedSlicing_r17Present {
				ie.FreqPriorityListDedicatedSlicing_r17 = new(FreqPriorityListDedicatedSlicing_r17)
				if err = ie.FreqPriorityListDedicatedSlicing_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode FreqPriorityListDedicatedSlicing_r17", err)
				}
			}
		}
	}
	return nil
}
