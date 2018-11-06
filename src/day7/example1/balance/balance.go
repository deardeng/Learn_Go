package balance

type Balancer interface {
	DoBalance(insts []*Instance, key ...string) (inst *Instance, err error)
}
