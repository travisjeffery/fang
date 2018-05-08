# Fang

Before:


```
cli.Flags().String("db", "", "Database connection string")
cli.Flags().String("env", "", "Environment")

viper.BindPFlag("db", cli.Flags().Lookup("db"))
viper.BindPFlag("env", cli.Flags().Lookup("env"))
viper.BindEnv("env", "ENV")
```

After:

```
fang.F(cli.Flags()).
    Flag("db", "", "Database connection string").
    Env("ENV", env", "", "Environment")
```
