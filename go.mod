module github.com/tprasadtp/pkg

go 1.20

require (
	github.com/spf13/cobra v1.6.1
	go.opentelemetry.io/otel/trace v1.11.2
	golang.org/x/sys v0.4.0
	golang.org/x/term v0.4.0
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.opentelemetry.io/otel v1.11.2 // indirect
)

retract [v0.1.0, v0.2.1]

retract [v1.0.0, v1.2.4]
