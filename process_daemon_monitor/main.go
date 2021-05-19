package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/things-go/x/extstr"
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
var daemon bool
var forever bool

func init() {
	rootCmd.PersistentFlags().BoolVar(&daemon, "daemon", false, "run in background")
	rootCmd.PersistentFlags().BoolVar(&forever, "forever", false, "run in forever, fail and retry")
	rootCmd.PersistentPreRun = preRun
	rootCmd.PersistentPostRun = postRun
}

func preRun(*cobra.Command, []string) {
	fmt.Println("current pid: ", os.Getpid())

	execName, args := os.Args[0], os.Args[1:]
	// daemon运行
	if daemon {
		args = extstr.DeleteAll(args, "--daemon")
		execCmd := exec.Command(execName, args...)
		execCmd.Start()

		format := "%s PID[ %d ] running...\n"
		if forever {
			format = "%s<forever> PID[ %d ] running...\n"
		}
		log.Printf(format, execName, execCmd.Process.Pid)

		os.Exit(0)
	}
}

func postRun(*cobra.Command, []string) {
	var execCmd *exec.Cmd

	execName, args := os.Args[0], os.Args[1:]
	if forever {
		go func() {
			args = extstr.DeleteAll(args, "--forever")
			for {
				if execCmd != nil {
					execCmd.Process.Kill()
					time.Sleep(time.Second * 5)
				}
				execCmd = exec.Command(execName, args...)
				// MultiPipe := func(execCmd *exec.Cmd) (io.Reader, error) {
				// 	stderrReader, err := execCmd.StderrPipe()
				// 	if err != nil {
				// 		return nil, err
				// 	}
				// 	stdoutReader, err := execCmd.StdoutPipe()
				// 	if err != nil {
				// 		return nil, err
				// 	}
				// 	return io.MultiReader(stderrReader, stdoutReader), nil
				// }

				// reader, err := MultiPipe(execCmd)
				// if err != nil {
				// 	log.Printf("-----> std multi pipe, %v, restarting...\n", err)
				// 	continue
				// }

				// go func() {
				// 	defer func() {
				// 		if err := recover(); err != nil {
				// 			log.Printf("-----> std multi pipe crashed, %v\nstack:%s", err, string(debug.Stack()))
				// 		}
				// 	}()
				// 	for scanner := bufio.NewScanner(reader); scanner.Scan(); {
				// 		log.Printf("-----> [std] %s", scanner.Text())
				// 	}
				// }()

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
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
