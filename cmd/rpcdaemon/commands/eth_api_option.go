package commands

// EthAPIOption is an option that can be set with NewEthAPI
type EthAPIOption interface {
	Apply(*APIImpl)
}

// EthAPIOptionFunc is a function defined for *APIImpl
type EthAPIOptionFunc func(*APIImpl)

// Apply is a wrapper function used to call the actual function
func (f EthAPIOptionFunc) Apply(sd *APIImpl) {
	f(sd)
}

// SetAllowUnprotectedTxs is an option that can set allowUnprotectedTxs for *APIImpl
func SetAllowUnprotectedTxs(allowUnprotectedTxs bool) EthAPIOption {
	return EthAPIOptionFunc(func(p *APIImpl) {
		p.allowUnprotectedTxs = allowUnprotectedTxs
	})
}
