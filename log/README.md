# Log

## Requirements

go 1.20 or later, because of `errors.Join` usage.

## How to use it within a library?

Logger should passed using dependency injection.

## Hooks?

A Hooks is nothing but a custom handler. If you want a hook, simply implement log.Handler interface
and include your hook's logic in it. For simple hooks, you can simple wrap an existing handler in your custom handler implementation along with code

## Dependencies

Don't be turned away by dependencies like aws/sdk, google cloud sdk, journald etc. in the go.mod file, they are only pulled if you use their respective handlers or use `config.Automatic`.
`log` package only depends on standard library.

## Nested fields

> Use Field/Logger Namespaces!

Most logging solutions are columnar stores and nested fields are bad for most of them.
In case of ELK nested json documents are flattened anyway.


[wire]: https://github.com/google/wire
