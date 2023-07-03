module github.com/tprasadtp/pkg

go 1.21rc2

require (
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	golang.org/x/sys v0.9.0
)

require github.com/inconshreveable/mousetrap v1.1.0 // indirect

retract	[v1.0.0, v1.2.4] // this module provides no compatibility gurantees.
retract [v0.1.0, v0.2.1] // this module provides no compatibility gurantees.
