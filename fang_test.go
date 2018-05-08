package fang_test

import (
	"github.com/spf13/cobra"
	"github.com/travisjeffery/fang"
)

func ExampleFang() {
	cli := &cobra.Command{
		Use: "example",
		Run: func(c *cobra.Command, args []string) {},
	}
	f := cli.Flags()
	fang.F(f).
		Flag(f.String, "db-conn", "localhost:5432", "DB connection string.").
		Env(f.String, "ENV", "env", "dev", "Environment command is running in.")
}
