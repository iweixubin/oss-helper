package commons

import (
	"os"
	"os/signal"
	"syscall"
)

//var wgOnExit sync.WaitGroup

// OnExit 在程序退出的时候做一些处理工作~
func OnExit(do func()) {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal)

	go func() {
		for {
			sig := <-chSignal
			// 只选这几个的原因是因为 kingshard 也只是选这几个~
			// https://github.com/flike/kingshard/blob/master/cmd/kingshard/main.go
			if sig == syscall.SIGINT || sig == syscall.SIGTERM || sig == syscall.SIGQUIT {
				//wgOnExit.Add(1)
				do()
				//wgOnExit.Done()
			}
		}
	}()

	//go func(){
	//	wgOnExit.Wait()
	//	os.Exit()
	//}

	// https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter16/16.03.html
	//SIGBUS（总线错误）, SIGFPE（算术错误）和 SIGSEGV（段错误）称为同步信号，
	//它们在程序执行错误时触发，而不是通过 os.Process.Kill 之类的触发。通常，Go 程序会将这类信号转为 run-time panic。

	//SIGHUP（挂起）, SIGINT（中断）或 SIGTERM（终止）默认会使得程序退出。

	//SIGQUIT, SIGILL, SIGTRAP, SIGABRT, SIGSTKFLT, SIGEMT 或 SIGSYS 默认会使得程序退出，同时生成 stack dump。

	//SIGTSTP, SIGTTIN 或 SIGTTOU，这是 shell 使用的，作业控制的信号，执行系统默认的行为。

	//SIGPROF（性能分析定时器，记录 CPU 时间，包括用户态和内核态）， Go 运行时使用该信号实现 runtime.CPUProfile。

	//其他信号，Go 捕获了，但没有做任何处理。
}
