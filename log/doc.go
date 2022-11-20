// Package log implements a simple logger which can handle
// structured logging, cli focused terminal logging and
// platform specific log backend like journald, Windows event logs,
// Google StackDriver along with helpers for http and grpc loggers.
//
// This package is inspired by [apex/log](which in turn is inspired by [sirupsen/logrus]) and
// [uber-go/zap].
//
// [apex/log]: https://github.com/apex/log
// [sirupsen/logrus]: https://github.com/sirupsen/logrus
// [uber-go/zap]: https://github.com/uber-go/zap
package log
