# Fang

Fangs removes boilerplate using github.com/spf13/cobra with github.com/spf13/viper.

## Example

Before:


``` go
cli.Flags().String("db-host", "", "Database host")
cli.Flags().String("db-user", "", "Database user")
cli.Flags().String("db-pass", "", "Database password.")

if err := viper.BindPFlag("db-host", cli.Flags().Lookup("db-host")); err != nil {
    panic(err)
}
if err := viper.BindEnv("db-user", "DB_USER"); err != nil {
    panic(err)
}
if err := viper.BindPFlag("db-pass", cli.Flags().Lookup("db-pass")); err != nil {
    panic(err)
}
if err := viper.BindEnv("db-pass", "DB_PASS"); err != nil {
    panic(err)
}
```

After:

``` go
f := cli.Flags()
fang.F(f).
    Flag(f.String, "db-host", "", "Database host").
    Env(f.String, "DB_USER", "db-user", "", "Database user").
    Env(f.String, "DB_PASS", db-pass", "", "Database password")
```

## License

MIT

---

- [travisjeffery.com](http://travisjeffery.com)
- GitHub [@travisjeffery](https://github.com/travisjeffery)
- Twitter [@travisjeffery](https://twitter.com/travisjeffery)
- Medium [@travisjeffery](https://medium.com/@travisjeffery)
