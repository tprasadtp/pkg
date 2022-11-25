# cli

- This is fork of `github.com/spf13/cobra`.

### Changes from upstream

> This only lists public API changes. Please refer to [godoc](https://pkg.go.dev/github.com/tprasadtp/cli) for API documentation.

- Because zsh shell completion now always includes command descriptions, following functions are removed.
    - `GenZshCompletionFileNoDesc`
    - `GenZshCompletionNoDesc`
- Because fish shell completion now always includes command descriptions, following functions signatures have changed.
    - `GenZshCompletionFileNoDesc`
    - `GenZshCompletionNoDesc`
- Because bash shell completion is now always uses V2 completion format, following functions have been removed use `GenBashCompletion` and `GenBashCompletionFile` instead.
    - `GenBashCompletionV2`
    - `GenBashCompletionV2File`
- Following exported fields in `Command` struct has changed.
    - Removed `BashCompletionFunction` (because legacy bash completion was removed)
    - Removed `DisableAutoGenTag` (because of reworked documentation and manpage generation)
    - Removed `CompletionOptions` (because root command no longer includes completion command)
- `CompletionOptions` struct removed
- Documentation generation is completely reworked, and now **only** supports `manpages` and markdown file generation.
- Adds `version` package to inject build time version info and version command.
