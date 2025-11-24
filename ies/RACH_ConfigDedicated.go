package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigDedicated struct {
	Cfra                         *CFRA              `optional`
	Ra_Prioritization            *RA_Prioritization `optional`
	Ra_PrioritizationTwoStep_r16 *RA_Prioritization `optional,ext-1`
	Cfra_TwoStep_r16             *CFRA_TwoStep_r16  `optional,ext-1`
}

func (ie *RACH_ConfigDedicated) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Ra_PrioritizationTwoStep_r16 != nil || ie.Cfra_TwoStep_r16 != nil
	preambleBits := []bool{hasExtensions, ie.Cfra != nil, ie.Ra_Prioritization != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Cfra != nil {
		if err = ie.Cfra.Encode(w); err != nil {
			return utils.WrapError("Encode Cfra", err)
		}
	}
	if ie.Ra_Prioritization != nil {
		if err = ie.Ra_Prioritization.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_Prioritization", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Ra_PrioritizationTwoStep_r16 != nil || ie.Cfra_TwoStep_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RACH_ConfigDedicated", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Ra_PrioritizationTwoStep_r16 != nil, ie.Cfra_TwoStep_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ra_PrioritizationTwoStep_r16 optional
			if ie.Ra_PrioritizationTwoStep_r16 != nil {
				if err = ie.Ra_PrioritizationTwoStep_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ra_PrioritizationTwoStep_r16", err)
				}
			}
			// encode Cfra_TwoStep_r16 optional
			if ie.Cfra_TwoStep_r16 != nil {
				if err = ie.Cfra_TwoStep_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cfra_TwoStep_r16", err)
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

func (ie *RACH_ConfigDedicated) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var CfraPresent bool
	if CfraPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_PrioritizationPresent bool
	if Ra_PrioritizationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CfraPresent {
		ie.Cfra = new(CFRA)
		if err = ie.Cfra.Decode(r); err != nil {
			return utils.WrapError("Decode Cfra", err)
		}
	}
	if Ra_PrioritizationPresent {
		ie.Ra_Prioritization = new(RA_Prioritization)
		if err = ie.Ra_Prioritization.Decode(r); err != nil {
			return utils.WrapError("Decode Ra_Prioritization", err)
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

			Ra_PrioritizationTwoStep_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cfra_TwoStep_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ra_PrioritizationTwoStep_r16 optional
			if Ra_PrioritizationTwoStep_r16Present {
				ie.Ra_PrioritizationTwoStep_r16 = new(RA_Prioritization)
				if err = ie.Ra_PrioritizationTwoStep_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ra_PrioritizationTwoStep_r16", err)
				}
			}
			// decode Cfra_TwoStep_r16 optional
			if Cfra_TwoStep_r16Present {
				ie.Cfra_TwoStep_r16 = new(CFRA_TwoStep_r16)
				if err = ie.Cfra_TwoStep_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cfra_TwoStep_r16", err)
				}
			}
		}
	}
	return nil
}
