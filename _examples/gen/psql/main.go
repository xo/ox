// Command psql is a xo/ox version of `psql`.
//
// Generated from _examples/gen.go.
package main

import (
	"context"

	"github.com/xo/ox"
	"github.com/xo/ox/text"
)

func init() {
	text.FlagSpecSpacer = "="
}

func main() {
	ox.RunContext(
		context.Background(),
		ox.Defaults(),
		ox.Usage(`psql`, ``),
		ox.Banner(`psql is the PostgreSQL interactive terminal.`),
		ox.Spec(`[OPTION]... [DBNAME [USERNAME]]`),
		ox.Help(ox.Sections(
			"General options",
			"Input and output options",
			"Output format options",
			"Connection options",
		)),
		ox.Footer(`For more information, type "\?" (for internal commands) or "\help" (for SQL
commands) from within psql, or consult the psql section in the PostgreSQL
documentation.

Report bugs to <pgsql-bugs@lists.postgresql.org>.
PostgreSQL home page - <https://www.postgresql.org/>`),
		ox.Flags().
			String(`command`, `run only single command (SQL or internal) and exit`, ox.Spec(`COMMAND`), ox.Short("c"), ox.Section(0)).
			String(`dbname`, `database name to connect to`, ox.Spec(`DBNAME`), ox.Short("d"), ox.Section(0)).
			String(`file`, `execute commands from file, then exit`, ox.Spec(`FILENAME`), ox.Short("f"), ox.Section(0)).
			Bool(`list`, `list available databases, then exit`, ox.Short("l"), ox.Section(0)).
			String(`set`, `--variable=NAME=VALUE`, ox.Spec(`,`), ox.Short("v"), ox.Section(0)).
			Bool(`no-psqlrc`, `do not read startup file (~/.psqlrc)`, ox.Short("X"), ox.Section(0)).
			String(`help[`, `show this help, then exit`, ox.Spec(`options]`), ox.Short("?"), ox.Section(0)).
			Bool(`echo-all`, `echo all input from script`, ox.Short("a"), ox.Section(1)).
			Bool(`echo-errors`, `echo failed commands`, ox.Short("b"), ox.Section(1)).
			Bool(`echo-queries`, `echo commands sent to server`, ox.Short("e"), ox.Section(1)).
			Bool(`echo-hidden`, `display queries that internal commands generate`, ox.Short("E"), ox.Section(1)).
			String(`log-file`, `send session log to file`, ox.Spec(`FILENAME`), ox.Short("L"), ox.Section(1)).
			Bool(`no-readline`, `disable enhanced command line editing (readline)`, ox.Short("n"), ox.Section(1)).
			String(`output`, `send query results to file (or |pipe)`, ox.Spec(`FILENAME`), ox.Short("o"), ox.Section(1)).
			Bool(`quiet`, `run quietly (no messages, only query output)`, ox.Short("q"), ox.Section(1)).
			Bool(`single-step`, `single-step mode (confirm each query)`, ox.Short("s"), ox.Section(1)).
			Bool(`single-line`, `single-line mode (end of line terminates SQL command)`, ox.Short("S"), ox.Section(1)).
			Bool(`no-align`, `unaligned table output mode`, ox.Short("A"), ox.Section(2)).
			Bool(`csv`, `CSV (Comma-Separated Values) table output mode`, ox.Section(2)).
			String(`field-separator`, `field separator for unaligned output (default: "|")`, ox.Spec(`STRING`), ox.Short("F"), ox.Section(2)).
			Bool(`html`, `HTML table output mode`, ox.Short("H"), ox.Section(2)).
			Map(`pset`, `set printing option VAR to ARG (see \pset command)`, ox.Spec(`VAR[=ARG]`), ox.Short("P"), ox.Section(2)).
			String(`record-separator`, `record separator for unaligned output (default: newline)`, ox.Spec(`STRING`), ox.Short("R"), ox.Section(2)).
			Bool(`tuples-only`, `print rows only`, ox.Short("t"), ox.Section(2)).
			String(`table-attr`, `set HTML table tag attributes (e.g., width, border)`, ox.Spec(`TEXT`), ox.Short("T"), ox.Section(2)).
			Bool(`expanded`, `turn on expanded table output`, ox.Short("x"), ox.Section(2)).
			String(`host`, `database server host or socket directory`, ox.Spec(`HOSTNAME`), ox.Short("h"), ox.Section(3)).
			String(`port`, `database server port`, ox.Spec(`PORT`), ox.Short("p"), ox.Section(3)).
			String(`username`, `database user name`, ox.Spec(`USERNAME`), ox.Short("U"), ox.Section(3)).
			Bool(`no-password`, `never prompt for password`, ox.Short("w"), ox.Section(3)).
			Bool(`password`, `force password prompt (should happen automatically)`, ox.Short("W"), ox.Section(3)),
	)
}
