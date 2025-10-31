package ies

// SystemInformationBlockType18-r12-commConfig-r12 ::= SEQUENCE
type Systeminformationblocktype18R12CommconfigR12 struct {
	CommrxpoolR12             SlCommrxpoollistR12
	CommtxpoolnormalcommonR12 *SlCommtxpoollistR12
	CommtxpoolexceptionalR12  *SlCommtxpoollistR12
	CommsyncconfigR12         *SlSyncconfiglistR12
}
