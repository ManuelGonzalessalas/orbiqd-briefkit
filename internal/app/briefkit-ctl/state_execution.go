package briefkitctl

type StateExecutionCmd struct {
	Create StateExecutionCreateCmd `cmd:"" help:"Create a new execution"`
	List   StateExecutionListCmd   `cmd:"" help:"List executions"`
	Show   StateExecutionShowCmd   `cmd:"" help:"Show execution details"`
}
