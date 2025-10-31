package ies

// MUST-Parameters-r14 ::= SEQUENCE
type MustParametersR14 struct {
	MustTm234Upto2txR14                   *MustParametersR14MustTm234Upto2txR14
	MustTm89UptooneinterferinglayerR14    *MustParametersR14MustTm89UptooneinterferinglayerR14
	MustTm10UptooneinterferinglayerR14    *MustParametersR14MustTm10UptooneinterferinglayerR14
	MustTm89UptothreeinterferinglayersR14 *MustParametersR14MustTm89UptothreeinterferinglayersR14
	MustTm10UptothreeinterferinglayersR14 *MustParametersR14MustTm10UptothreeinterferinglayersR14
}
