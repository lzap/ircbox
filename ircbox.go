package main

import "github.com/thoj/go-ircevent"
import "github.com/lhcb-org/shell"
import "os"
import "os/signal"
import "syscall"
import "fmt"
import "log"
import "math/rand"
import "time"
import "strings"
import "strconv"
import "flag"
import "regexp"

func main() {
	server := flag.String("server", "", "IRC server and port (e.g. irc.xxx.redhat.com:6667")
	nick := flag.String("nick", "", "nickname to give shell to")
	max_lines := flag.Int("max-lines", 20, "maximum lines returned to prevent flooding")
	flag.Parse()

	if !regexp.MustCompile(`^.*:\d+$`).MatchString(*server) {
		fmt.Println("You must provide server option")
		os.Exit(1)
	}
	if len(*nick) == 0 {
		fmt.Println("You must provide nick option")
		os.Exit(1)
	}

	sh, err := shell.New()
	if err != nil {
		log.Fatal(err)
	}
	defer sh.Delete()

	rand.Seed(time.Now().UTC().UnixNano())
	name := "ircbox-" + strconv.FormatInt(int64(rand.Int31()), 16)
	conn := irc.IRC(name, name)
	conn.Connect(*server)
	defer conn.Quit()

	// handle ctrl-c properly
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			fmt.Println("Interrupted")
			conn.Quit()
			sh.Delete()
			os.Exit(0)
		case syscall.SIGTERM:
			fmt.Println("Terminated")
			conn.Quit()
			sh.Delete()
			os.Exit(0)
		}
	}()

	hostname, _ := os.Hostname()
	conn.Privmsg(*nick, "Hey, you have been given an interactive shell on "+hostname)
	conn.Privmsg(*nick, "This is not a tty. Do not run vi, emacs or mc and use grep wisely.")
	conn.Privmsg(*nick, "Don't be evil, you are being watched. Stdout and stderr are combined.")
	conn.Privmsg(*nick, "Type your commands now, use 'exit' to close your session.")
	conn.AddCallback("PRIVMSG", func(event *irc.Event) {
		msg := event.Message()
		if msg == "exit" {
			conn.Privmsg(*nick, "Bye. You were using https://github.com/lzap/ircbox")
			conn.Quit()
			sh.Delete()
			os.Exit(0)
		} else {
			run(&sh, conn, nick, msg, *max_lines)
		}
	})
	fmt.Println("The session log now follows (Ctrl-c to cancel)")
	conn.Loop()
	fmt.Println("Finished")
}

func run(sh *shell.Shell, conn *irc.Connection, nick *string, command string, max_lines int) {
	fmt.Println("$ " + command)
	out, err := sh.Run(command, "|", "head -n"+strconv.FormatInt(int64(max_lines), 10))
	if err != nil {
		fmt.Println(err.Error())
		conn.Privmsg(*nick, err.Error())
	}
	fmt.Println(string(out))
	for i, msg := range strings.Split(string(out), "\n") {
		if i >= max_lines {
			conn.Privmsg(*nick, "... (cut) use grep or something ...")
			break
		} else {
			conn.Privmsg(*nick, msg)
		}
	}
	ps1, _ := sh.Run("echo $(whoami)@$(hostname):$(pwd)")
	conn.Privmsg(*nick, string(ps1))
}
