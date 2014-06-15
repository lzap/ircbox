ircbox
------

You have a nice reproducer but your testing box is somewhere in a lab behind
NAT. If you had a chance to give shell to your colleague... oh wait, you can!

Having a private IRC server, you can use this tiny little tool that can give
shells via IRC session. It's easy:

    ircbox -server irc.xxx.redhat.com:6667 -nick lzap

This is how session looks like:

    ircbox-680dd9c5 | Hey, you have been given an interactive shell on xyz.lab.redhat.com
    ircbox-680dd9c5 | This is not a tty. Do not run vi, emacs or mc and use grep wisely.
    ircbox-680dd9c5 | Don't be evil, you are being watched. Stdout and stderr are combined.
    ircbox-680dd9c5 | Type your commands now, use 'exit' to close your session.
                 me | ls /
    ircbox-680dd9c5 | bin
    ircbox-680dd9c5 | etc
    ircbox-680dd9c5 | blah
                 me | exit
    ircbox-680dd9c5 | Bye.

Warning: Never use public IRC servers and never expose boxes with sensitive
data. In other words, this is good for access to test labs within your company
or at home.

Download
--------

Linux 64bit binary with no dependencies (other than libc and libpthread) can
be downloaded *from this git repository*. Yes, I do that, because when I need
it, I don't want to install Go language to compile it from sources.

    $ wget --no-check-certificate -O /usr/local/bin/ircbox https://github.com/lzap/ircbox/raw/master/ircbox-linux64
    $ chmod +x /usr/local/bin/ircbox

Compilation
-----------

Can be done in three simple steps:

    $ go get github.com/thoj/go-ircevent
    $ go get github.com/lhcb-org/shell
    $ go build

Since Go compiles to static binaries by default, you can copy the executable
to your server and use it directly.

