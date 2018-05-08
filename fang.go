package fang

import (
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

func (f *Fang) Flag(fn interface{}, args ...interface{}) {
	name := args[0].(string)
	argsv := make([]reflect.Value, len(args))
	for i, v := range args {
		argsv[i] = reflect.ValueOf(v)
	}
	fv := reflect.ValueOf(fn)
	fv.Call(argsv)
	viper.BindPFlag(name, f.fs.Lookup(name))
}

func (f *Fang) Env(fn interface{}, args ...interface{}) {
	envvar := args[0].(string)
	name := args[1].(string)
	f.Flag(fn, args[1:]...)
	viper.BindEnv(name, envvar)
}
