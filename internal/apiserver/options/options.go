package options

import (
	genericoptions "github.com/TroyXia/iam/internal/pkg/options"
	"github.com/TroyXia/iam/pkg/log"
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
)

// Options runs an iam api server.
type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions       `json:"server"   mapstructure:"server"`
	GRPCOptions             *genericoptions.GRPCOptions            `json:"grpc"     mapstructure:"grpc"`
	InsecureServing         *genericoptions.InsecureServingOptions `json:"insecure" mapstructure:"insecure"`
	SecureServing           *genericoptions.SecureServingOptions   `json:"secure"   mapstructure:"secure"`
	MySQLOptions            *genericoptions.MySQLOptions           `json:"mysql"    mapstructure:"mysql"`
	RedisOptions            *genericoptions.RedisOptions           `json:"redis"    mapstructure:"redis"`
	JwtOptions              *genericoptions.JwtOptions             `json:"jwt"      mapstructure:"jwt"`
	Log                     *log.Options                           `json:"log"      mapstructure:"log"`
	FeatureOptions          *genericoptions.FeatureOptions         `json:"feature"  mapstructure:"feature"`
}

// NewOptions creates a new Options object with default parameters.
func NewOptions() *Options {
	o := Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		//GRPCOptions:             genericoptions.NewGRPCOptions(),
		InsecureServing: genericoptions.NewInsecureServingOptions(),
		SecureServing:   genericoptions.NewSecureServingOptions(),
		//MySQLOptions:            genericoptions.NewMySQLOptions(),
		//RedisOptions:            genericoptions.NewRedisOptions(),
		//JwtOptions:              genericoptions.NewJwtOptions(),
		Log: log.NewOptions(),
		//FeatureOptions:          genericoptions.NewFeatureOptions(),
	}

	return &o
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.GenericServerRunOptions.AddFlags(fss.FlagSet("generic"))

	return fss
}
