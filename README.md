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
fs := cli.Flags()
fang.F(fs).
    Flag(fs.String, "db", "", "Database connection string").
    Env(fs.String, "ENV", env", "", "Environment")
```
