package proxy

import (
	"fmt"
	"log/syslog"
	"sync/atomic"
	"time"
)

// Logger logs a message into local filesytem or remote logging system
type Logger interface {
	Debug(msg string, args ...interface{})
}

// SysLogger reports log messages to remote machine which enables syslog,
// this Proxy object proxys log request to the real object syslog.Writer,
// but this Proxy object add other access to the syslog.Writer.
//
// For example, it won't acccess the real object immediately, to avoid too
// frequent requests to the remote syslog service. It queues the request
// and send it to the remote every 5s or the pending log entries is more
// than 100.
func NewSysLogger(addr, tag string) (*SysLogger, error) {
	l := &SysLogger{tag: tag, addr: addr}
	w, err := syslog.Dial("tcp", l.addr, syslog.LOG_DEBUG, l.tag)
	if err != nil {
		return nil, err
	}
	l.writer = w
	l.ch = make(chan string, 1024)

	go l.report()

	return l, nil
}

type SysLogger struct {
	tag    string
	addr   string
	writer *syslog.Writer

	ch    chan string
	count atomic.Int32
}

func (l *SysLogger) Debug(msg string, args ...interface{}) {
	select {
	case l.ch <- fmt.Sprintf(msg, args...):
		l.count.Add(1)
	default:
		fmt.Println("err: queue full, log message dropped")
	}
}

func (l *SysLogger) report() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for {
		// report every 5s or 100 log entries pending
		select {
		case <-ticker.C:
		default:
			if l.count.Load() < 100 {
				continue
			}
		}

		// report
		msg, ok := <-l.ch
		l.count.Add(-1)
		if !ok {
			fmt.Println("logger closed")
			break
		}
		// report
		err := l.writer.Debug(msg)
		if err != nil {
			fmt.Println("report msg fail:", err)
		}
	}
}
