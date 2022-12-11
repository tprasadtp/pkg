# Log

## Requirements

- go 1.20 or later, because of `errors.Join` usage.

## How to use it within a library?

Logger should passed using dependency injection.
With tools like [wire][], you can have build time safety.

## Why No Support for Hooks?

Hooks are nothing but a custom handlers, if you want a custom hook, simply implement log.Handler interface. For simple hooks, you can simple wrap an existing handler in your custom handler implementation.

## Dependencies

Don't be turned away by dependencies like aws/sdk, google cloud sdk, journald etc. in the go.mod file, they are only pulled if you use their respective handlers or use `config.Automatic`.

## Nested fields

> Use Field/Logger Namespaces!

Most logging solutions are columnar stores and nested fields are bad for most of them.
In case of ELK nested json documents are flattened anyway.


[wire]: https://github.com/google/wire
