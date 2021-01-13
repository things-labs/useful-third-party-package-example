package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/thinkgos/go-core-package/extstr"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "process",
	Short: "process tool with command",
	Long:  `process tool with command`,
	Run: func(cmd *cobra.Command, args []string) {
		if forever {
			return
		}
		for {
			time.Sleep(time.Second * 3)
			log.Println("it is running")
		}
	},
}
var (
	hasDebug bool
	daemon   bool
	forever  bool
)

func init() {
	rootCmd.PersistentFlags().BoolVar(&hasDebug, "debug", false, "debug log output")
	rootCmd.PersistentFlags().BoolVar(&daemon, "daemon", false, "run in background")
	rootCmd.PersistentFlags().BoolVar(&forever, "forever", false, "run in forever, fail and retry")
	rootCmd.PersistentPreRun = preRun
	rootCmd.PersistentPostRun = postRun
}

func preRun(cmd *cobra.Command, args []string) {
	execName := os.Args[0]

	// daemon运行
	if daemon {
		args := extstr.DeleteAll(os.Args[1:], "--daemon")
		execCmd := exec.Command(execName, args...)
		execCmd.Start()
		// TODO: 检查相关进程是否已启动??
		format := "%s PID[ %d ] running...\n"
		if forever {
			format = "%s<forever> PID[ %d ] running...\n"
		}
		log.Println(format, execName, execCmd.Process.Pid)

		os.Exit(0)
	}

	if hasDebug {
		// cpuProfilingFile, _ = os.Create("cpu.prof")
		// memProfilingFile, _ = os.Create("memory.prof")
		// blockProfilingFile, _ = os.Create("block.prof")
		// goroutineProfilingFile, _ = os.Create("goroutine.prof")
		// threadcreateProfilingFile, _ = os.Create("threadcreate.prof")
		// pprof.StartCPUProfile(cpuProfilingFile)
	}
}

func postRun(cmd *cobra.Command, args []string) {
	var execCmd *exec.Cmd

	execName := os.Args[0]
	if forever {
		go func() {
			args := extstr.DeleteAll(os.Args[1:], "--forever")
			for {
				if execCmd != nil {
					execCmd.Process.Kill()
					time.Sleep(time.Second * 5)
				}
				execCmd = exec.Command(execName, args...)

				MultiPipe := func(execCmd *exec.Cmd) (io.Reader, error) {
					stderrReader, err := execCmd.StderrPipe()
					if err != nil {
						return nil, err
					}
					stdoutReader, err := execCmd.StdoutPipe()
					if err != nil {
						return nil, err
					}
					return io.MultiReader(stderrReader, stdoutReader), nil
				}

				reader, err := MultiPipe(execCmd)
				if err != nil {
					log.Printf("-----> std multi pipe, %v, restarting...\n", err)
					continue
				}

				go func() {
					defer func() {
						if err := recover(); err != nil {
							log.Printf("-----> std multi pipe crashed, %v\nstack:%s", err, string(debug.Stack()))
						}
					}()
					for scanner := bufio.NewScanner(reader); scanner.Scan(); {
						log.Printf("-----> [std] %s", scanner.Text())
					}
				}()

				if err := execCmd.Start(); err != nil {
					log.Printf("-----> child process start failed, %v, restarting...\n", err)
					continue
				}
				pid := execCmd.Process.Pid

				log.Printf("-----> worker name: %s PID[ %d ] running...\n", execName, pid)
				if err := execCmd.Wait(); err != nil {
					log.Printf("-----> parent process wait, %v, restarting...", err)
					continue
				}
				log.Printf("-----> worker name: %s PID[ %d ] unexpected exited, restarting...\n", execName, pid)
			}
		}()
	}

	sysSignal := make(chan os.Signal, 1)
	signal.Notify(sysSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	log.Println(<-sysSignal)
	if execCmd != nil {
		log.Printf("--> kill process PID[ %d ]", execCmd.Process.Pid)
		execCmd.Process.Kill()
	}
	if hasDebug {
		SaveProfiling()
	}
}

func SaveProfiling() {
	// pprof.Lookup("goroutine").WriteTo(goroutineProfilingFile, 1)
	// pprof.Lookup("heap").WriteTo(memProfilingFile, 1)
	// pprof.Lookup("block").WriteTo(blockProfilingFile, 1)
	// pprof.Lookup("threadcreate").WriteTo(threadcreateProfilingFile, 1)
	// pprof.StopCPUProfile()
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
