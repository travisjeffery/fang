package fang

import (
	"fmt"
	"reflect"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func F(fs *pflag.FlagSet) *Fang {
	return &Fang{fs: fs}
}

type Fang struct {
	fs *pflag.FlagSet
}

// Flag takes in a creator func for the flag and its args. For example Flag(cli.Flags().String,
// "name", "value", "usage").
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

// Env takes in a creator func for the flag/env, its env var name, and the flag's args. For example Env(cli.Flags().String,
// "NAME", "name", "value", "usage").
func (f *Fang) Env(fn interface{}, args ...interface{}) *Fang {
	envvar := args[0].(string)
	name := args[1].(string)
	f.Flag(fn, args[1:]...)
	if err := viper.BindEnv(name, envvar); err != nil {
		panic(fmt.Sprintf("fang: failed to lookup flag: %s", err.Error()))
	}
	return f
}
