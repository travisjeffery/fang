# Fang

Fangs removes boilerplate using github.com/spf13/cobra with github.com/spf13/viper.

## Example

Before:


``` go
cli.Flags().String("db", "", "Database connection string")
cli.Flags().String("env", "", "Environment")

if err := viper.BindPFlag("db", cli.Flags().Lookup("db")); err != nil {
    panic(err)
}
if err := viper.BindPFlag("env", cli.Flags().Lookup("env")); err != nil {
    panic(err)
}
if err := viper.BindEnv("env", "ENV"); err != nil {
    panic(err)
}
```

After:

``` go
f := cli.Flags()
fang.F(f).
    Flag(f.String, "db", "", "Database connection string").
    Env(f.String, "ENV", env", "", "Environment")
```

## License

MIT

---

- [travisjeffery.com](http://travisjeffery.com)
- GitHub [@travisjeffery](https://github.com/travisjeffery)
- Twitter [@travisjeffery](https://twitter.com/travisjeffery)
- Medium [@travisjeffery](https://medium.com/@travisjeffery)
