// Fangs removes boilerplate using github.com/spf13/cobra with github.com/spf13/viper.
package fang

import (
	"fmt"
	"reflect"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// F takes a flag set and creates a fang instance to create flags/envs with.
func F(fs *pflag.FlagSet) *Fang {
	return &Fang{fs: fs}
}

// Fang is used to chain Flag and Env calls.
type Fang struct {
	fs *pflag.FlagSet
}

// Flag takes a creator func for the flag and its args.
func (f *Fang) Flag(fn interface{}, args ...interface{}) *Fang {
	name := args[0].(string)
	argsv := make([]reflect.Value, len(args))
	for i, v := range args {
		argsv[i] = reflect.ValueOf(v)
	}
	fv := reflect.ValueOf(fn)
	fv.Call(argsv)
	if err := viper.BindPFlag(name, f.fs.Lookup(name)); err != nil {
		panic(fmt.Sprintf("fang: failed to lookup flag: %s", err.Error()))
	}
	return f
}

// Env takes a creator func for the flag/env, its env var name, and the flag's args.
func (f *Fang) Env(fn interface{}, args ...interface{}) *Fang {
	envvar := args[0].(string)
	name := args[1].(string)
	f.Flag(fn, args[1:]...)
	if err := viper.BindEnv(name, envvar); err != nil {
		panic(fmt.Sprintf("fang: failed to lookup flag: %s", err.Error()))
	}
	return f
}
