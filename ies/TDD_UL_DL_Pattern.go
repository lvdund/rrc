package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_Pattern struct {
	Dl_UL_TransmissionPeriodicity       TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity        `madatory`
	NrofDownlinkSlots                   int64                                                  `lb:0,ub:maxNrofSlots,madatory`
	NrofDownlinkSymbols                 int64                                                  `lb:0,ub:maxNrofSymbols_1,madatory`
	NrofUplinkSlots                     int64                                                  `lb:0,ub:maxNrofSlots,madatory`
	NrofUplinkSymbols                   int64                                                  `lb:0,ub:maxNrofSymbols_1,madatory`
	Dl_UL_TransmissionPeriodicity_v1530 *TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530 `optional,ext-1`
}

func (ie *TDD_UL_DL_Pattern) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Dl_UL_TransmissionPeriodicity_v1530 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Dl_UL_TransmissionPeriodicity.Encode(w); err != nil {
		return utils.WrapError("Encode Dl_UL_TransmissionPeriodicity", err)
	}
	if err = w.WriteInteger(ie.NrofDownlinkSlots, &uper.Constraint{Lb: 0, Ub: maxNrofSlots}, false); err != nil {
		return utils.WrapError("WriteInteger NrofDownlinkSlots", err)
	}
	if err = w.WriteInteger(ie.NrofDownlinkSymbols, &uper.Constraint{Lb: 0, Ub: maxNrofSymbols_1}, false); err != nil {
		return utils.WrapError("WriteInteger NrofDownlinkSymbols", err)
	}
	if err = w.WriteInteger(ie.NrofUplinkSlots, &uper.Constraint{Lb: 0, Ub: maxNrofSlots}, false); err != nil {
		return utils.WrapError("WriteInteger NrofUplinkSlots", err)
	}
	if err = w.WriteInteger(ie.NrofUplinkSymbols, &uper.Constraint{Lb: 0, Ub: maxNrofSymbols_1}, false); err != nil {
		return utils.WrapError("WriteInteger NrofUplinkSymbols", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Dl_UL_TransmissionPeriodicity_v1530 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap TDD_UL_DL_Pattern", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Dl_UL_TransmissionPeriodicity_v1530 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Dl_UL_TransmissionPeriodicity_v1530 optional
			if ie.Dl_UL_TransmissionPeriodicity_v1530 != nil {
				if err = ie.Dl_UL_TransmissionPeriodicity_v1530.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_UL_TransmissionPeriodicity_v1530", err)
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

func (ie *TDD_UL_DL_Pattern) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Dl_UL_TransmissionPeriodicity.Decode(r); err != nil {
		return utils.WrapError("Decode Dl_UL_TransmissionPeriodicity", err)
	}
	var tmp_int_NrofDownlinkSlots int64
	if tmp_int_NrofDownlinkSlots, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSlots}, false); err != nil {
		return utils.WrapError("ReadInteger NrofDownlinkSlots", err)
	}
	ie.NrofDownlinkSlots = tmp_int_NrofDownlinkSlots
	var tmp_int_NrofDownlinkSymbols int64
	if tmp_int_NrofDownlinkSymbols, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSymbols_1}, false); err != nil {
		return utils.WrapError("ReadInteger NrofDownlinkSymbols", err)
	}
	ie.NrofDownlinkSymbols = tmp_int_NrofDownlinkSymbols
	var tmp_int_NrofUplinkSlots int64
	if tmp_int_NrofUplinkSlots, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSlots}, false); err != nil {
		return utils.WrapError("ReadInteger NrofUplinkSlots", err)
	}
	ie.NrofUplinkSlots = tmp_int_NrofUplinkSlots
	var tmp_int_NrofUplinkSymbols int64
	if tmp_int_NrofUplinkSymbols, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofSymbols_1}, false); err != nil {
		return utils.WrapError("ReadInteger NrofUplinkSymbols", err)
	}
	ie.NrofUplinkSymbols = tmp_int_NrofUplinkSymbols

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

			Dl_UL_TransmissionPeriodicity_v1530Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Dl_UL_TransmissionPeriodicity_v1530 optional
			if Dl_UL_TransmissionPeriodicity_v1530Present {
				ie.Dl_UL_TransmissionPeriodicity_v1530 = new(TDD_UL_DL_Pattern_dl_UL_TransmissionPeriodicity_v1530)
				if err = ie.Dl_UL_TransmissionPeriodicity_v1530.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_UL_TransmissionPeriodicity_v1530", err)
				}
			}
		}
	}
	return nil
}
