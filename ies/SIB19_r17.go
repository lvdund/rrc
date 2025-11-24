package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB19_r17 struct {
	Ntn_Config_r17                   *NTN_Config_r17              `optional`
	T_Service_r17                    *int64                       `lb:0,ub:549755813887,optional`
	ReferenceLocation_r17            *ReferenceLocation_r17       `optional`
	DistanceThresh_r17               *int64                       `lb:0,ub:65525,optional`
	Ntn_NeighCellConfigList_r17      *NTN_NeighCellConfigList_r17 `optional`
	LateNonCriticalExtension         *[]byte                      `optional`
	Ntn_NeighCellConfigListExt_v1720 *NTN_NeighCellConfigList_r17 `optional,ext-1`
}

func (ie *SIB19_r17) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Ntn_NeighCellConfigListExt_v1720 != nil
	preambleBits := []bool{hasExtensions, ie.Ntn_Config_r17 != nil, ie.T_Service_r17 != nil, ie.ReferenceLocation_r17 != nil, ie.DistanceThresh_r17 != nil, ie.Ntn_NeighCellConfigList_r17 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ntn_Config_r17 != nil {
		if err = ie.Ntn_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ntn_Config_r17", err)
		}
	}
	if ie.T_Service_r17 != nil {
		if err = w.WriteInteger(*ie.T_Service_r17, &uper.Constraint{Lb: 0, Ub: 549755813887}, false); err != nil {
			return utils.WrapError("Encode T_Service_r17", err)
		}
	}
	if ie.ReferenceLocation_r17 != nil {
		if err = ie.ReferenceLocation_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ReferenceLocation_r17", err)
		}
	}
	if ie.DistanceThresh_r17 != nil {
		if err = w.WriteInteger(*ie.DistanceThresh_r17, &uper.Constraint{Lb: 0, Ub: 65525}, false); err != nil {
			return utils.WrapError("Encode DistanceThresh_r17", err)
		}
	}
	if ie.Ntn_NeighCellConfigList_r17 != nil {
		if err = ie.Ntn_NeighCellConfigList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ntn_NeighCellConfigList_r17", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Ntn_NeighCellConfigListExt_v1720 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB19_r17", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Ntn_NeighCellConfigListExt_v1720 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ntn_NeighCellConfigListExt_v1720 optional
			if ie.Ntn_NeighCellConfigListExt_v1720 != nil {
				if err = ie.Ntn_NeighCellConfigListExt_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ntn_NeighCellConfigListExt_v1720", err)
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

func (ie *SIB19_r17) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Ntn_Config_r17Present bool
	if Ntn_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var T_Service_r17Present bool
	if T_Service_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReferenceLocation_r17Present bool
	if ReferenceLocation_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DistanceThresh_r17Present bool
	if DistanceThresh_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ntn_NeighCellConfigList_r17Present bool
	if Ntn_NeighCellConfigList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ntn_Config_r17Present {
		ie.Ntn_Config_r17 = new(NTN_Config_r17)
		if err = ie.Ntn_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ntn_Config_r17", err)
		}
	}
	if T_Service_r17Present {
		var tmp_int_T_Service_r17 int64
		if tmp_int_T_Service_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 549755813887}, false); err != nil {
			return utils.WrapError("Decode T_Service_r17", err)
		}
		ie.T_Service_r17 = &tmp_int_T_Service_r17
	}
	if ReferenceLocation_r17Present {
		ie.ReferenceLocation_r17 = new(ReferenceLocation_r17)
		if err = ie.ReferenceLocation_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ReferenceLocation_r17", err)
		}
	}
	if DistanceThresh_r17Present {
		var tmp_int_DistanceThresh_r17 int64
		if tmp_int_DistanceThresh_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65525}, false); err != nil {
			return utils.WrapError("Decode DistanceThresh_r17", err)
		}
		ie.DistanceThresh_r17 = &tmp_int_DistanceThresh_r17
	}
	if Ntn_NeighCellConfigList_r17Present {
		ie.Ntn_NeighCellConfigList_r17 = new(NTN_NeighCellConfigList_r17)
		if err = ie.Ntn_NeighCellConfigList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ntn_NeighCellConfigList_r17", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
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

			Ntn_NeighCellConfigListExt_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ntn_NeighCellConfigListExt_v1720 optional
			if Ntn_NeighCellConfigListExt_v1720Present {
				ie.Ntn_NeighCellConfigListExt_v1720 = new(NTN_NeighCellConfigList_r17)
				if err = ie.Ntn_NeighCellConfigListExt_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ntn_NeighCellConfigListExt_v1720", err)
				}
			}
		}
	}
	return nil
}
