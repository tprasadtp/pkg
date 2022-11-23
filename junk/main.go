//go:build ignore

package main

import (
	"fmt"
	"log"
	"sync"
)

type core struct {
	handlers []string
}

type Logger struct {
	mu   sync.Mutex
	name string
	core *core
}

func (l *Logger) init() {
	if l.core == nil {
		log.Printf("[LOGGER-CORE] Initialize")
		l.core = &core{}
	}
}

func (l *Logger) AddHandler(handler string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.init()
	l.core.handlers = append(l.core.handlers, handler)
}

func (l *Logger) WithName(name string) Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.init()
	return Logger{
		name: name,
		core: l.core,
	}
}

func main() {
	l1 := &Logger{
		name: "logger-01",
	}

	var wg sync.WaitGroup
	log.Printf("---> Basic Logger")
	wg.Add(1)
	go func() {
		defer wg.Done()
		l1.AddHandler("handler-01")
		log.Printf("name=%s, handler=%s", l1.name, l1.core.handlers[0])
	}()
	wg.Wait()

	log.Printf("---> Logger with multiple handlers")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			handler_name := fmt.Sprintf("handler-%d", j+1)
			l1.AddHandler(handler_name)
			log.Printf("name=%s, handler=%d(%s)", l1.name, len(l1.core.handlers), l1.core.handlers[len(l1.core.handlers)-1])
		}(i)
	}
	wg.Wait()

	log.Printf("---> Multiple Loggers with multiple handlers")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			logger_name := fmt.Sprintf("logger-%d", j+1)
			handler_name := fmt.Sprintf("handler-%d", j+1)
			looper := l1.WithName(logger_name)
			looper.AddHandler(handler_name)
			log.Printf("name=%s, handler=%d(%s)", looper.name, len(looper.core.handlers), looper.core.handlers[len(looper.core.handlers)-1])
		}(i)
	}
	wg.Wait()

	log.Printf("[MAIN] Exited all GoRoutines")
}
