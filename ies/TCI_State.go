package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TCI_State struct {
	Tci_StateId                TCI_StateId                 `madatory`
	Qcl_Type1                  QCL_Info                    `madatory`
	Qcl_Type2                  *QCL_Info                   `optional`
	AdditionalPCI_r17          *AdditionalPCIIndex_r17     `optional,ext-1`
	PathlossReferenceRS_Id_r17 *PathlossReferenceRS_Id_r17 `optional,ext-1`
	Ul_powerControl_r17        *Uplink_powerControlId_r17  `optional,ext-1`
}

func (ie *TCI_State) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.AdditionalPCI_r17 != nil || ie.PathlossReferenceRS_Id_r17 != nil || ie.Ul_powerControl_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Qcl_Type2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Tci_StateId.Encode(w); err != nil {
		return utils.WrapError("Encode Tci_StateId", err)
	}
	if err = ie.Qcl_Type1.Encode(w); err != nil {
		return utils.WrapError("Encode Qcl_Type1", err)
	}
	if ie.Qcl_Type2 != nil {
		if err = ie.Qcl_Type2.Encode(w); err != nil {
			return utils.WrapError("Encode Qcl_Type2", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.AdditionalPCI_r17 != nil || ie.PathlossReferenceRS_Id_r17 != nil || ie.Ul_powerControl_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap TCI_State", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.AdditionalPCI_r17 != nil, ie.PathlossReferenceRS_Id_r17 != nil, ie.Ul_powerControl_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode AdditionalPCI_r17 optional
			if ie.AdditionalPCI_r17 != nil {
				if err = ie.AdditionalPCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AdditionalPCI_r17", err)
				}
			}
			// encode PathlossReferenceRS_Id_r17 optional
			if ie.PathlossReferenceRS_Id_r17 != nil {
				if err = ie.PathlossReferenceRS_Id_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PathlossReferenceRS_Id_r17", err)
				}
			}
			// encode Ul_powerControl_r17 optional
			if ie.Ul_powerControl_r17 != nil {
				if err = ie.Ul_powerControl_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_powerControl_r17", err)
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

func (ie *TCI_State) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Qcl_Type2Present bool
	if Qcl_Type2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Tci_StateId.Decode(r); err != nil {
		return utils.WrapError("Decode Tci_StateId", err)
	}
	if err = ie.Qcl_Type1.Decode(r); err != nil {
		return utils.WrapError("Decode Qcl_Type1", err)
	}
	if Qcl_Type2Present {
		ie.Qcl_Type2 = new(QCL_Info)
		if err = ie.Qcl_Type2.Decode(r); err != nil {
			return utils.WrapError("Decode Qcl_Type2", err)
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

			AdditionalPCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PathlossReferenceRS_Id_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_powerControl_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode AdditionalPCI_r17 optional
			if AdditionalPCI_r17Present {
				ie.AdditionalPCI_r17 = new(AdditionalPCIIndex_r17)
				if err = ie.AdditionalPCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode AdditionalPCI_r17", err)
				}
			}
			// decode PathlossReferenceRS_Id_r17 optional
			if PathlossReferenceRS_Id_r17Present {
				ie.PathlossReferenceRS_Id_r17 = new(PathlossReferenceRS_Id_r17)
				if err = ie.PathlossReferenceRS_Id_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PathlossReferenceRS_Id_r17", err)
				}
			}
			// decode Ul_powerControl_r17 optional
			if Ul_powerControl_r17Present {
				ie.Ul_powerControl_r17 = new(Uplink_powerControlId_r17)
				if err = ie.Ul_powerControl_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_powerControl_r17", err)
				}
			}
		}
	}
	return nil
}
