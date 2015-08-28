/*
@author: xuchengxuan(bigpyer@126.com)
@brief: logger module
*/
package ymtlog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

const (
	DATEFORMAT        = "2006-01-02"
	DEFAULT_LOG_SCAN  = 300
	DEFAULT_LOG_LEVEL = INFO
)

type LEVEL byte

const (
	INFO LEVEL = iota
	DEBUG
	WARN
	ERROR
)

type logger struct {
	mu       *sync.RWMutex
	fileDir  string
	fileName string

	date *time.Time

	logFile  *os.File
	lger     *log.Logger
	timeScan int64

	logChan  chan string
	logLevel LEVEL
}

// logger handler constructor
func NewLogger(dir string, name string) *logger {
	dailyLogger := &logger{
		mu:       new(sync.RWMutex),
		fileDir:  dir,
		fileName: name,
		logChan:  make(chan string, 500),
		logLevel: DEFAULT_LOG_LEVEL,
	}

	dailyLogger.initDailyLogger()

	return dailyLogger
}

func (f *logger) initDailyLogger() {

	t, _ := time.Parse(DATEFORMAT, time.Now().Format(DATEFORMAT))

	f.date = &t
	f.mu.Lock()
	defer f.mu.Unlock()

	logFile := joinFilePath(f.fileDir, f.fileName)
	f.logFile, _ = os.OpenFile(logFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	f.lger = log.New(f.logFile, "", log.LstdFlags|log.Lmicroseconds)

	go f.writeLog()
	go f.monitorFile()
}

func (f *logger) isNeedRotate() bool {
	t, _ := time.Parse(DATEFORMAT, time.Now().Format(DATEFORMAT))
	if t.After(*f.date) {
		return true
	}
	return true
}

func (f *logger) rotate() {
	logFile := joinFilePath(f.fileDir, f.fileName)
	logFileBak := logFile + "." + f.date.Format(DATEFORMAT)

	if !isExist(logFileBak) && f.isNeedRotate() {
		if f.logFile != nil {
			f.logFile.Close()
		}

		err := os.Rename(logFile, logFileBak)
		if err != nil {
			f.lger.Printf("logger rename error: %v", err.Error())
		}

		t, _ := time.Parse(DATEFORMAT, time.Now().Format(DATEFORMAT))
		f.date = &t
		f.logFile, _ = os.Create(logFile)
		f.lger = log.New(f.logFile, "", log.LstdFlags|log.Lmicroseconds)
	}
}

func (f *logger) monitorFile() {
	defer func() {
		if err := recover(); err != nil {
			f.lger.Panic("logger's FileMonitor() catch panic: %v\n", err)
		}
	}()

	// check frequency
	logScan := DEFAULT_LOG_SCAN

	timer := time.NewTicker(time.Duration(logScan) * time.Second)
	for {
		select {
		case <-timer.C:
			f.checkFile()
		}
	}
}

func (f *logger) checkFile() {
	defer func() {
		if err := recover(); err != nil {
			f.lger.Printf("logger's FileCheck() catch panic: %v\n", err)
		}
	}()
	if f.isNeedRotate() {
		f.mu.Lock()
		defer f.mu.Unlock()

		f.rotate()
	}
}

// passive to close filelogger
func (f *logger) Close() error {

	close(f.logChan)
	f.lger = nil

	return f.logFile.Close()
}

// Receive logStr from f's logChan and print logstr to file
func (f *logger) writeLog() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf(" writeLog catch panic: %v\n", err)
		}
	}()

	for {
		select {
		case str := <-f.logChan:
			f.outPut(str)
		}
	}
}

// print log
func (f *logger) outPut(str string) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	f.lger.Output(2, str)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func joinFilePath(path, file string) string {
	return filepath.Join(path, file)
}

func shortFileName(file string) string {
	return filepath.Base(file)
}

// info log
func (f *logger) Info(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if f.logLevel <= INFO {
		f.logChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[INFO] "+format, v...)
	}
}

// debug log
func (f *logger) Debug(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if f.logLevel <= DEBUG {
		f.logChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[DEBUG] "+format, v...)
	}
}

// warn log
func (f *logger) Warn(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if f.logLevel <= WARN {
		f.logChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[WARN] "+format, v...)
	}
}

// error log
func (f *logger) Error(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if f.logLevel <= ERROR {
		f.logChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[ERROR] "+format, v...)
	}
}
