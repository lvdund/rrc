package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LogicalChannelConfig struct {
	Ul_SpecificParameters     *LogicalChannelConfig_ul_SpecificParameters `lb:1,ub:maxNrofServingCells_1,optional`
	ChannelAccessPriority_r16 *int64                                      `lb:1,ub:4,optional,ext-3`
	BitRateMultiplier_r16     *LogicalChannelConfig_bitRateMultiplier_r16 `optional,ext-3`
}

func (ie *LogicalChannelConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ChannelAccessPriority_r16 != nil || ie.BitRateMultiplier_r16 != nil
	preambleBits := []bool{hasExtensions, ie.Ul_SpecificParameters != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ul_SpecificParameters != nil {
		if err = ie.Ul_SpecificParameters.Encode(w); err != nil {
			return utils.WrapError("Encode Ul_SpecificParameters", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.ChannelAccessPriority_r16 != nil || ie.BitRateMultiplier_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap LogicalChannelConfig", err)
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.ChannelAccessPriority_r16 != nil, ie.BitRateMultiplier_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ChannelAccessPriority_r16 optional
			if ie.ChannelAccessPriority_r16 != nil {
				if err = extWriter.WriteInteger(*ie.ChannelAccessPriority_r16, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode ChannelAccessPriority_r16", err)
				}
			}
			// encode BitRateMultiplier_r16 optional
			if ie.BitRateMultiplier_r16 != nil {
				if err = ie.BitRateMultiplier_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BitRateMultiplier_r16", err)
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

func (ie *LogicalChannelConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Ul_SpecificParametersPresent bool
	if Ul_SpecificParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ul_SpecificParametersPresent {
		ie.Ul_SpecificParameters = new(LogicalChannelConfig_ul_SpecificParameters)
		if err = ie.Ul_SpecificParameters.Decode(r); err != nil {
			return utils.WrapError("Decode Ul_SpecificParameters", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ChannelAccessPriority_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BitRateMultiplier_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ChannelAccessPriority_r16 optional
			if ChannelAccessPriority_r16Present {
				var tmp_int_ChannelAccessPriority_r16 int64
				if tmp_int_ChannelAccessPriority_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode ChannelAccessPriority_r16", err)
				}
				ie.ChannelAccessPriority_r16 = &tmp_int_ChannelAccessPriority_r16
			}
			// decode BitRateMultiplier_r16 optional
			if BitRateMultiplier_r16Present {
				ie.BitRateMultiplier_r16 = new(LogicalChannelConfig_bitRateMultiplier_r16)
				if err = ie.BitRateMultiplier_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode BitRateMultiplier_r16", err)
				}
			}
		}
	}
	return nil
}
