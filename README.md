
# stopwatch - a time tracking tool

stopwatch is a simple time tracking tool that allows you to track time spent on
different tasks. It is a command line tool that is build with a minimalistic
interface and the option to edit entries manually that are stored in a human
readable and friendly format.

## Usage

The stopwatch command line tool `stpw` offers a number of commands to manage
time tracking. Information is stored in a plain text file placed in your home
directory that you can access and edit. The file is located at `~/.stpw.txt`.

```bash
$ stpw # show help
$ stpw start <task>
$ stpw break # mark a break or end of day
$ stpw status # daily report
$ stpw list # a full report of last 5 days

# access duration summaries with different timeframes and filters
$ stpw list --filter meeting # filter by tag #meeting
$ stpw weekly # (wip)
$ stpw monthly # (wip)
```

The time entries are stored in a plain text file and entries can be edited
there on demand without breaking the tool.

```bash
$ cat ~/.stpw.txt
```

You can use the `--file` flag to specify a different file location. In case
this should be permanent use the `STPW_FILE` environment variable (WIP) or
create an alias.

```bash
stpw --file ~/.stopwatch-file.txt list
STPW_FILE=~/.stopwatch-file.txt stpw list
alias stpw="stpw --file ~/.stopwatch-file.txt"
```

## Development

```bash
make # build binary
make run ARGS="help" # run on-the-fly with arguments
make test # run testsuite
make clean # clean binary
```

