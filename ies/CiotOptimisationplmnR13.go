package ies

// CIOT-OptimisationPLMN-r13 ::= SEQUENCE
type CiotOptimisationplmnR13 struct {
	UpCiotEpsOptimisationR13        *bool
	CpCiotEpsOptimisationR13        *bool
	AttachwithoutpdnConnectivityR13 *bool
}
