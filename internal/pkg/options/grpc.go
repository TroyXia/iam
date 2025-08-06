package options

// GRPCOptions are for creating an unauthenticated, unauthorized, insecure port.
// No one should be using these anymore.
type GRPCOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	BindPort    int    `json:"bind-port"    mapstructure:"bind-port"`
	MaxMsgSize  int    `json:"max-msg-size" mapstructure:"max-msg-size"`
}
