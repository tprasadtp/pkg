module github.com/tprasadtp/pkg

go 1.20

require (
	github.com/spf13/cobra v1.6.1
	go.opentelemetry.io/otel/trace v1.13.0
	golang.org/x/sys v0.5.0
	golang.org/x/term v0.5.0
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.opentelemetry.io/otel v1.13.0 // indirect
)

retract [v0.1.0, v0.2.1] // Unstable release which is not used.

retract [v1.0.0, v1.2.4] // Unstable release which is not used.
