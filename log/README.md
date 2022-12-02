# Log


- Logger should passed using dependency injection.
With tools like [wire][], you can have build time safety.
- DOES NOT provide global logger. Having a global logger leads to library developers using it as a kitchen sink. Pass logger via dependency injection.
- DOES NOT support passing logger via context. This is because of above design decision to not have a
global logger. Global logger becomes necessary as `context.Context` might be missing Logger.
- Even though passing logger via context.Context is not supported, `log.Event` does include
the context to allow handlers to populate contextual fields like Span and Trace.
- Uses `map[string]any` to specify fields. This is better compared to passing variadic slice
`...any` as it avoids depending on a vet to check for errors. For allocation optimization reasons,
this is translated to list of attributes in `Event`.
- Hooks? hooks are nothing but a custom handlers, if you want a custom hook, simply implement log.Handler interface. For simple hooks, you can simple wrap an existing handler in your custom handler implementation.
- Don't be afraid of dependencies like aws/sdk, google cloud sdk, journald etc. in the mod file, they are only pulled if you use their respective handlers or use `handlers.AutoConfigure` (Requires g 1.17 or later)

[wire]: https://github.com/google/wire
