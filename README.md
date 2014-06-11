ircbox
------

You have a nice reproducer but your testing box is somewhere in a lab behind
NAT. If you had a chance to give shell to your colleague. But wait, you can!

If you have a private IRC server, you can use this tiny little tool that can
give shells via IRC session. It's easy:

    ircbox -server irc.xxx.redhat.com:6667 -nick lzap

This is how session looks like:

    ircbox-680dd9c5 | Hey, you have been given an interactive shell on lzapx.brq.redhat.com
    ircbox-680dd9c5 | This is not a tty. Do not run vi, emacs or mc and use grep wisely.
    ircbox-680dd9c5 | Don't be evil, you are being watched. Stdout and stderr are combined.
    ircbox-680dd9c5 | Type your commands now, use 'exit' to close your session.
    > ls /
    bin
    etc
    blah
    > exit
    Bye.

Warning: Never use public IRC servers and never expose boxes with sensitive data.
In other words, this is good for access to test labs within your company or at
home.

Compilation
-----------

    go get github.com/thoj/go-ircevent
    go get github.com/lhcb-org/shell
    go build

Since Go compiles to static binaries, you can copy the executable to your server
(if you build on linux 64bit it should work on all modern distros) and use it
directly.
