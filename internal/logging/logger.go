package logging

import (
	"log"
	"os"
)

type Logger struct {
	Stdout   *os.File
	Stderr   *os.File
	LogLevel LogLevel
	out      *log.Logger
	err      *log.Logger
}

func NewLogger(prefix string) Logger {
	out := log.New(os.Stdout, prefix, 0)
	err := log.New(os.Stderr, prefix, log.LstdFlags|log.Lshortfile)

	return Logger{
		Stdout:   os.Stdout,
		Stderr:   os.Stderr,
		LogLevel: globalLogLevel,
		out:      out,
		err:      err,
	}
}

func (l Logger) WithOutfile(f *os.File) Logger {
	out := log.New(f, l.out.Prefix(), 0)
	l.out = out
	return l
}

func (l Logger) WithErrfile(f *os.File) Logger {
	err := log.New(f, l.err.Prefix(), 0)
	l.err = err
	return l
}

func (l Logger) WithLogLevel(level LogLevel) Logger {
	l.LogLevel = level
	return l
}

func (l *Logger) Print(v ...interface{}) {
	l.out.Print(v...)
}

func (l *Logger) Printf(fmt string, v ...interface{}) {
	l.out.Printf(fmt, v...)
}

func (l *Logger) Println(v ...interface{}) {
	l.out.Println(v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.out.Fatal(v...)
}

func (l *Logger) Fatalf(fmt string, v ...interface{}) {
	l.out.Fatalf(fmt, v...)
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.out.Fatalln(v...)
}

func (l *Logger) Debug(v ...interface{}) {
	if l.LogLevel >= DebugLevel {
		l.err.Print(v...)
	}
}

func (l *Logger) Debugf(fmt string, v ...interface{}) {
	if l.LogLevel >= DebugLevel {
		l.err.Printf(fmt, v...)
	}
}

func (l *Logger) Debugln(v ...interface{}) {
	if l.LogLevel >= DebugLevel {
		l.err.Println(v...)
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.LogLevel >= InfoLevel {
		l.err.Print(v...)
	}
}

func (l *Logger) Infof(fmt string, v ...interface{}) {
	if l.LogLevel >= InfoLevel {
		l.err.Printf(fmt, v...)
	}
}

func (l *Logger) Infoln(v ...interface{}) {
	if l.LogLevel >= InfoLevel {
		l.err.Println(v...)
	}
}

func (l *Logger) Warn(v ...interface{}) {
	if l.LogLevel >= WarnLevel {
		l.err.Print(v...)
	}
}

func (l *Logger) Warnf(fmt string, v ...interface{}) {
	if l.LogLevel >= WarnLevel {
		l.err.Printf(fmt, v...)
	}
}

func (l *Logger) Warnln(v ...interface{}) {
	if l.LogLevel >= WarnLevel {
		l.err.Println(v...)
	}
}

func (l *Logger) Crit(v ...interface{}) {
	if l.LogLevel >= CritLevel {
		l.err.Print(v...)
	}
}

func (l *Logger) Critf(fmt string, v ...interface{}) {
	if l.LogLevel >= CritLevel {
		l.err.Printf(fmt, v...)
	}
}

func (l *Logger) Critln(v ...interface{}) {
	if l.LogLevel >= CritLevel {
		l.err.Println(v...)
	}
}
