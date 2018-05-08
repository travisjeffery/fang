# Fang

Before:


``` go
cli.Flags().String("db", "", "Database connection string")
cli.Flags().String("env", "", "Environment")

viper.BindPFlag("db", cli.Flags().Lookup("db"))
viper.BindPFlag("env", cli.Flags().Lookup("env"))
viper.BindEnv("env", "ENV")
```

After:

``` go
f := cli.Flags()
fang.F(fs).
    Flag(f.String, "db", "", "Database connection string").
    Env(f.String, "ENV", env", "", "Environment")
```
