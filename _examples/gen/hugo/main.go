// Command hugo is a xo/ox version of `+hugo`.
//
// Generated from _examples/gen.go.
package main

import (
	"context"

	"github.com/xo/ox"
)

func main() {
	ox.RunContext(
		context.Background(),
		ox.Defaults(),
		ox.Banner("hugo is the main command, used to build your Hugo site.\n\nHugo is a Fast and Flexible Static Site Generator\nbuilt with love by spf13 and friends in Go.\n\nComplete documentation is available at https://gohugo.io/."),
		ox.Usage("hugo", ""),
		ox.Spec("[flags]"),
		ox.Footer("Use \"hugo [command] --help\" for more information about a command."),
		ox.Sub(
			ox.Banner("build is the main command, used to build your Hugo site.\n\nHugo is a Fast and Flexible Static Site Generator\nbuilt with love by spf13 and friends in Go.\n\nComplete documentation is available at https://gohugo.io/."),
			ox.Usage("build", "Build your site"),
			ox.Spec("[flags]"),
			ox.Flags().
				String("baseURL", "hostname (and path) to the root, e.g. https://spf13.com/", ox.Short("b")).
				Bool("buildDrafts", "include content marked as draft", ox.Short("D")).
				Bool("buildExpired", "include expired content", ox.Short("E")).
				Bool("buildFuture", "include content with publishdate in the future", ox.Short("F")).
				String("cacheDir", "filesystem path to cache directory").
				Bool("cleanDestinationDir", "remove files from destination not found in static directories").
				String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
				String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
				String("configDir", "config dir", ox.Default("config")).
				String("contentDir", "filesystem path to content directory", ox.Short("c")).
				String("destination", "filesystem path to write files to", ox.Short("d")).
				Slice("disableKinds", "disable different kind of pages (home, RSS etc.)", ox.Elem(ox.StringT)).
				Bool("enableGitInfo", "add Git revision, date, author, and CODEOWNERS info to the pages").
				String("environment", "build environment", ox.Short("e")).
				Bool("forceSyncStatic", "copy all files when static is changed.").
				Bool("gc", "enable to run some cleanup tasks (remove unused cache files) after the build").
				Bool("ignoreCache", "ignores the cache directory").
				String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
				String("layoutDir", "filesystem path to layout directory", ox.Short("l")).
				String("logLevel", "log level (debug|info|warn|error)").
				Bool("minify", "minify any supported output format (HTML, XML etc.)").
				Bool("noBuildLock", "don't create .hugo_build.lock file").
				Bool("noChmod", "don't sync permission mode of files").
				Bool("noTimes", "don't sync modification time of files").
				Bool("panicOnWarning", "panic on first WARNING log").
				String("poll", "set this to a poll interval, e.g --poll 700ms, to use a poll based approach to watch for file system changes").
				Bool("printI18nWarnings", "print missing translations").
				Bool("printMemoryUsage", "print memory usage to screen at intervals").
				Bool("printPathWarnings", "print warnings on duplicate target paths etc.").
				Bool("printUnusedTemplates", "print warnings on unused templates.").
				Bool("quiet", "build in quiet mode").
				Slice("renderSegments", "named segments to render (configured in the segments config)", ox.Elem(ox.StringT)).
				Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
				String("source", "filesystem path to read files relative from", ox.Short("s")).
				Bool("templateMetrics", "display metrics about template executions").
				Bool("templateMetricsHints", "calculate some improvement hints when combined with --templateMetrics").
				Slice("theme", "themes to use (located in /themes/THEMENAME/)", ox.Elem(ox.StringT), ox.Short("t")).
				String("themesDir", "filesystem path to themes directory").
				String("trace", "write trace to file (not useful in general)", ox.Spec("file")).
				Bool("watch", "watch filesystem for changes and recreate as needed", ox.Short("w")),
		), ox.Sub(
			ox.Banner("Display site configuration, both default and custom settings."),
			ox.Usage("config", "Display site configuration"),
			ox.Spec("[command] [flags]"),
			ox.Footer("Use \"hugo config [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("Print the configured file mounts"),
				ox.Usage("mounts", "Print the configured file mounts"),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("baseURL", "hostname (and path) to the root, e.g. https://spf13.com/", ox.Short("b")).
					String("cacheDir", "filesystem path to cache directory").
					String("contentDir", "filesystem path to content directory", ox.Short("c")).
					Slice("renderSegments", "named segments to render (configured in the segments config)", ox.Elem(ox.StringT)).
					Slice("theme", "themes to use (located in /themes/THEMENAME/)", ox.Elem(ox.StringT), ox.Short("t")).
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Flags().
				String("baseURL", "hostname (and path) to the root, e.g. https://spf13.com/", ox.Short("b")).
				String("cacheDir", "filesystem path to cache directory").
				String("contentDir", "filesystem path to content directory", ox.Short("c")).
				String("format", "preferred file format (toml, yaml or json)", ox.Default("toml")).
				String("lang", "the language to display config for. Defaults to the first language defined.").
				Slice("renderSegments", "named segments to render (configured in the segments config)", ox.Elem(ox.StringT)).
				Slice("theme", "themes to use (located in /themes/THEMENAME/)", ox.Elem(ox.StringT), ox.Short("t")).
				String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
				String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
				String("configDir", "config dir", ox.Default("config")).
				String("destination", "filesystem path to write files to", ox.Short("d")).
				String("environment", "build environment", ox.Short("e")).
				String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
				String("logLevel", "log level (debug|info|warn|error)").
				Bool("quiet", "build in quiet mode").
				Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
				String("source", "filesystem path to read files relative from", ox.Short("s")).
				String("themesDir", "filesystem path to themes directory"),
		), ox.Sub(
			ox.Banner("Convert front matter to another format.\n\nSee convert's subcommands toJSON, toTOML and toYAML for more information."),
			ox.Usage("convert", "Convert front matter to another format"),
			ox.Spec("[command]"),
			ox.Footer("Use \"hugo convert [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("toJSON converts all front matter in the content directory\nto use JSON for the front matter."),
				ox.Usage("toJSON", "Convert front matter to JSON"),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					String("output", "filesystem path to write files to", ox.Short("o")).
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory").
					Bool("unsafe", "enable less safe operations, please backup first"),
			), ox.Sub(
				ox.Banner("toTOML converts all front matter in the content directory\nto use TOML for the front matter."),
				ox.Usage("toTOML", "Convert front matter to TOML"),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					String("output", "filesystem path to write files to", ox.Short("o")).
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory").
					Bool("unsafe", "enable less safe operations, please backup first"),
			), ox.Sub(
				ox.Banner("toYAML converts all front matter in the content directory\nto use YAML for the front matter."),
				ox.Usage("toYAML", "Convert front matter to YAML"),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					String("output", "filesystem path to write files to", ox.Short("o")).
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory").
					Bool("unsafe", "enable less safe operations, please backup first"),
			), ox.Flags().
				String("output", "filesystem path to write files to", ox.Short("o")).
				Bool("unsafe", "enable less safe operations, please backup first").
				String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
				String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
				String("configDir", "config dir", ox.Default("config")).
				String("destination", "filesystem path to write files to", ox.Short("d")).
				String("environment", "build environment", ox.Short("e")).
				String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
				String("logLevel", "log level (debug|info|warn|error)").
				Bool("quiet", "build in quiet mode").
				Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
				String("source", "filesystem path to read files relative from", ox.Short("s")).
				String("themesDir", "filesystem path to themes directory"),
		), ox.Sub(
			ox.Banner("Display version and environment info. This is useful in Hugo bug reports"),
			ox.Usage("env", "Display version and environment info"),
			ox.Spec("[flags] [args]"),
			ox.Flags().
				String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
				String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
				String("configDir", "config dir", ox.Default("config")).
				String("destination", "filesystem path to write files to", ox.Short("d")).
				String("environment", "build environment", ox.Short("e")).
				String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
				String("logLevel", "log level (debug|info|warn|error)").
				Bool("quiet", "build in quiet mode").
				Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
				String("source", "filesystem path to read files relative from", ox.Short("s")).
				String("themesDir", "filesystem path to themes directory"),
		), ox.Sub(
			ox.Banner("Generate documentation for your project using Hugo's documentation engine, including syntax highlighting for various programming languages."),
			ox.Usage("gen", "Generate documentation and syntax highlighting styles"),
			ox.Spec("[command]"),
			ox.Footer("Use \"hugo gen [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("Generate CSS stylesheet for the Chroma code highlighter for a given style. This stylesheet is needed if markup.highlight.noClasses is disabled in config.\n\nSee https://xyproto.github.io/splash/docs/all.html for a preview of the available styles"),
				ox.Usage("chromastyles", "Generate CSS stylesheet for the Chroma code highlighter"),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("highlightStyle", "foreground and background colors for highlighted lines, e.g. --highlightStyle \"#fff000 bg:#000fff\"").
					String("lineNumbersInlineStyle", "foreground and background colors for inline line numbers, e.g. --lineNumbersInlineStyle \"#fff000 bg:#000fff\"").
					String("lineNumbersTableStyle", "foreground and background colors for table line numbers, e.g. --lineNumbersTableStyle \"#fff000 bg:#000fff\"").
					String("style", "highlighter style (see https://xyproto.github.io/splash/docs/)", ox.Default("friendly")).
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("Generate Markdown documentation for the Hugo CLI.\n\t\t\tThis command is, mostly, used to create up-to-date documentation\n\tof Hugo's command-line interface for https://gohugo.io/.\n\n\tIt creates one Markdown file per command with front matter suitable\n\tfor rendering in Hugo."),
				ox.Usage("doc", "Generate Markdown documentation for the Hugo CLI."),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("dir", "the directory to write the doc.", ox.Default("/tmp/$APPNAMEdoc/")).
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("This command automatically generates up-to-date man pages of Hugo's\n\tcommand-line interface.  By default, it creates the man page files\n\tin the \"man\" directory under the current directory."),
				ox.Usage("man", "Generate man pages for the Hugo CLI"),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("dir", "the directory to write the man pages.", ox.Default("man/")).
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Flags().
				String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
				String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
				String("configDir", "config dir", ox.Default("config")).
				String("destination", "filesystem path to write files to", ox.Short("d")).
				String("environment", "build environment", ox.Short("e")).
				String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
				String("logLevel", "log level (debug|info|warn|error)").
				Bool("quiet", "build in quiet mode").
				Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
				String("source", "filesystem path to read files relative from", ox.Short("s")).
				String("themesDir", "filesystem path to themes directory"),
		), ox.Sub(
			ox.Banner("Import a site from another system.\n\nImport requires a subcommand, e.g. `hugo import jekyll jekyll_root_path target_path`."),
			ox.Usage("import", "Import a site from another system"),
			ox.Spec("[command]"),
			ox.Footer("Use \"hugo import [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("hugo import from Jekyll.\n\t\t\nImport from Jekyll requires two paths, e.g. `hugo import jekyll jekyll_root_path target_path`."),
				ox.Usage("jekyll", "hugo import from Jekyll"),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					Bool("force", "allow import into non-empty target directory").
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Flags().
				String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
				String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
				String("configDir", "config dir", ox.Default("config")).
				String("destination", "filesystem path to write files to", ox.Short("d")).
				String("environment", "build environment", ox.Short("e")).
				String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
				String("logLevel", "log level (debug|info|warn|error)").
				Bool("quiet", "build in quiet mode").
				Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
				String("source", "filesystem path to read files relative from", ox.Short("s")).
				String("themesDir", "filesystem path to themes directory"),
		), ox.Sub(
			ox.Banner("List content.\n\nList requires a subcommand, e.g. hugo list drafts"),
			ox.Usage("list", "List content"),
			ox.Spec("[command]"),
			ox.Footer("Use \"hugo list [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("List all content including draft, future, and expired."),
				ox.Usage("all", "List all content"),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("List draft content."),
				ox.Usage("drafts", "List draft content"),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("List content with a past expiration date."),
				ox.Usage("expired", "List expired content"),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("List content with a future publication date."),
				ox.Usage("future", "List future content"),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("List content that is not draft, future, or expired."),
				ox.Usage("published", "List published content"),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Flags().
				String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
				String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
				String("configDir", "config dir", ox.Default("config")).
				String("destination", "filesystem path to write files to", ox.Short("d")).
				String("environment", "build environment", ox.Short("e")).
				String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
				String("logLevel", "log level (debug|info|warn|error)").
				Bool("quiet", "build in quiet mode").
				Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
				String("source", "filesystem path to read files relative from", ox.Short("s")).
				String("themesDir", "filesystem path to themes directory"),
		), ox.Sub(
			ox.Banner("Various helpers to help manage the modules in your project's dependency graph.\nMost operations here requires a Go version installed on your system (>= Go 1.12) and the relevant VCS client (typically Git).\nThis is not needed if you only operate on modules inside /themes or if you have vendored them via \"hugo mod vendor\".\n\n\nNote that Hugo will always start out by resolving the components defined in the site\nconfiguration, provided by a _vendor directory (if no --ignoreVendorPaths flag provided),\nGo Modules, or a folder inside the themes directory, in that order.\n\nSee https://gohugo.io/hugo-modules/ for more information."),
			ox.Usage("mod", "Manage modules"),
			ox.Spec("[command]"),
			ox.Footer("Use \"hugo mod [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("Delete the Hugo Module cache for the current project."),
				ox.Usage("clean", "Delete the Hugo Module cache for the current project."),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					Bool("all", "clean entire module cache").
					String("baseURL", "hostname (and path) to the root, e.g. https://spf13.com/", ox.Short("b")).
					String("cacheDir", "filesystem path to cache directory").
					String("contentDir", "filesystem path to content directory", ox.Short("c")).
					String("pattern", "pattern matching module paths to clean (all if not set), e.g. \"**hugo*\"").
					Slice("renderSegments", "named segments to render (configured in the segments config)", ox.Elem(ox.StringT)).
					Slice("theme", "themes to use (located in /themes/THEMENAME/)", ox.Elem(ox.StringT), ox.Short("t")).
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("\nResolves dependencies in your current Hugo Project.\n\nSome examples:\n\nInstall the latest version possible for a given module:\n\n    hugo mod get github.com/gohugoio/testshortcodes\n    \nInstall a specific version:\n\n    hugo mod get github.com/gohugoio/testshortcodes@v0.3.0\n\nInstall the latest versions of all direct module dependencies:\n\n    hugo mod get\n    hugo mod get ./... (recursive)\n\nInstall the latest versions of all module dependencies (direct and indirect):\n\n    hugo mod get -u\n    hugo mod get -u ./... (recursive)\n\nRun \"go help get\" for more information. All flags available for \"go get\" is also relevant here.\n\nNote that Hugo will always start out by resolving the components defined in the site\nconfiguration, provided by a _vendor directory (if no --ignoreVendorPaths flag provided),\nGo Modules, or a folder inside the themes directory, in that order.\n\nSee https://gohugo.io/hugo-modules/ for more information."),
				ox.Usage("get", "Resolves dependencies in your current Hugo Project."),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("Print a module dependency graph with information about module status (disabled, vendored).\nNote that for vendored modules, that is the version listed and not the one from go.mod."),
				ox.Usage("graph", "Print a module dependency graph."),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("baseURL", "hostname (and path) to the root, e.g. https://spf13.com/", ox.Short("b")).
					String("cacheDir", "filesystem path to cache directory").
					Bool("clean", "delete module cache for dependencies that fail verification").
					String("contentDir", "filesystem path to content directory", ox.Short("c")).
					Slice("renderSegments", "named segments to render (configured in the segments config)", ox.Elem(ox.StringT)).
					Slice("theme", "themes to use (located in /themes/THEMENAME/)", ox.Elem(ox.StringT), ox.Short("t")).
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("Initialize this project as a Hugo Module.\n\tIt will try to guess the module path, but you may help by passing it as an argument, e.g:\n\t\n\t\thugo mod init github.com/gohugoio/testshortcodes\n\t\n\tNote that Hugo Modules supports multi-module projects, so you can initialize a Hugo Module\n\tinside a subfolder on GitHub, as one example."),
				ox.Usage("init", "Initialize this project as a Hugo Module."),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("baseURL", "hostname (and path) to the root, e.g. https://spf13.com/", ox.Short("b")).
					String("cacheDir", "filesystem path to cache directory").
					String("contentDir", "filesystem path to content directory", ox.Short("c")).
					Slice("renderSegments", "named segments to render (configured in the segments config)", ox.Elem(ox.StringT)).
					Slice("theme", "themes to use (located in /themes/THEMENAME/)", ox.Elem(ox.StringT), ox.Short("t")).
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("Various npm (Node package manager) helpers."),
				ox.Usage("npm", "Various npm helpers."),
				ox.Spec("[command] [flags]"),
				ox.Footer("Use \"hugo mod npm [command] --help\" for more information about a command."),
				ox.Sub(
					ox.Banner("Prepares and writes a composite package.json file for your project.\n\nOn first run it creates a \"package.hugo.json\" in the project root if not already there. This file will be used as a template file\nwith the base dependency set. \n\nThis set will be merged with all \"package.hugo.json\" files found in the dependency tree, picking the version closest to the project.\n\nThis command is marked as 'Experimental'. We think it's a great idea, so it's not likely to be\nremoved from Hugo, but we need to test this out in \"real life\" to get a feel of it,\nso this may/will change in future versions of Hugo."),
					ox.Usage("pack", "Experimental: Prepares and writes a composite package.json file for your project."),
					ox.Spec("[flags] [args]"),
					ox.Flags().
						String("baseURL", "hostname (and path) to the root, e.g. https://spf13.com/", ox.Short("b")).
						String("cacheDir", "filesystem path to cache directory").
						String("contentDir", "filesystem path to content directory", ox.Short("c")).
						Slice("renderSegments", "named segments to render (configured in the segments config)", ox.Elem(ox.StringT)).
						Slice("theme", "themes to use (located in /themes/THEMENAME/)", ox.Elem(ox.StringT), ox.Short("t")).
						String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
						String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
						String("configDir", "config dir", ox.Default("config")).
						String("destination", "filesystem path to write files to", ox.Short("d")).
						String("environment", "build environment", ox.Short("e")).
						String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
						String("logLevel", "log level (debug|info|warn|error)").
						Bool("quiet", "build in quiet mode").
						Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
						String("source", "filesystem path to read files relative from", ox.Short("s")).
						String("themesDir", "filesystem path to themes directory"),
				), ox.Flags().
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("Remove unused entries in go.mod and go.sum."),
				ox.Usage("tidy", "Remove unused entries in go.mod and go.sum."),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("baseURL", "hostname (and path) to the root, e.g. https://spf13.com/", ox.Short("b")).
					String("cacheDir", "filesystem path to cache directory").
					String("contentDir", "filesystem path to content directory", ox.Short("c")).
					Slice("renderSegments", "named segments to render (configured in the segments config)", ox.Elem(ox.StringT)).
					Slice("theme", "themes to use (located in /themes/THEMENAME/)", ox.Elem(ox.StringT), ox.Short("t")).
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("Vendor all module dependencies into the _vendor directory.\n\tIf a module is vendored, that is where Hugo will look for it's dependencies."),
				ox.Usage("vendor", "Vendor all module dependencies into the _vendor directory."),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("baseURL", "hostname (and path) to the root, e.g. https://spf13.com/", ox.Short("b")).
					String("cacheDir", "filesystem path to cache directory").
					String("contentDir", "filesystem path to content directory", ox.Short("c")).
					Slice("renderSegments", "named segments to render (configured in the segments config)", ox.Elem(ox.StringT)).
					Slice("theme", "themes to use (located in /themes/THEMENAME/)", ox.Elem(ox.StringT), ox.Short("t")).
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("Verify checks that the dependencies of the current module, which are stored in a local downloaded source cache, have not been modified since being downloaded."),
				ox.Usage("verify", "Verify dependencies."),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					String("baseURL", "hostname (and path) to the root, e.g. https://spf13.com/", ox.Short("b")).
					String("cacheDir", "filesystem path to cache directory").
					Bool("clean", "delete module cache for dependencies that fail verification").
					String("contentDir", "filesystem path to content directory", ox.Short("c")).
					Slice("renderSegments", "named segments to render (configured in the segments config)", ox.Elem(ox.StringT)).
					Slice("theme", "themes to use (located in /themes/THEMENAME/)", ox.Elem(ox.StringT), ox.Short("t")).
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Flags().
				String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
				String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
				String("configDir", "config dir", ox.Default("config")).
				String("destination", "filesystem path to write files to", ox.Short("d")).
				String("environment", "build environment", ox.Short("e")).
				String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
				String("logLevel", "log level (debug|info|warn|error)").
				Bool("quiet", "build in quiet mode").
				Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
				String("source", "filesystem path to read files relative from", ox.Short("s")).
				String("themesDir", "filesystem path to themes directory"),
		), ox.Sub(
			ox.Banner("Create a new content file and automatically set the date and title.\nIt will guess which kind of file to create based on the path provided.\n\nYou can also specify the kind with `-k KIND`.\n\nIf archetypes are provided in your theme or site, they will be used.\n\nEnsure you run this within the root directory of your site."),
			ox.Usage("new", "Create new content"),
			ox.Spec("[command]"),
			ox.Footer("Use \"hugo new [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("Create a new content file and automatically set the date and title.\nIt will guess which kind of file to create based on the path provided.\n\nYou can also specify the kind with `-k KIND`.\n\nIf archetypes are provided in your theme or site, they will be used.\n\nEnsure you run this within the root directory of your site."),
				ox.Usage("content", "Create new content"),
				ox.Spec("[path] [flags]"),
				ox.Flags().
					String("baseURL", "hostname (and path) to the root, e.g. https://spf13.com/", ox.Short("b")).
					String("cacheDir", "filesystem path to cache directory").
					String("contentDir", "filesystem path to content directory", ox.Short("c")).
					String("editor", "edit new content with this editor, if provided").
					Bool("force", "overwrite file if it already exists", ox.Short("f")).
					String("kind", "content type to create", ox.Short("k")).
					Slice("renderSegments", "named segments to render (configured in the segments config)", ox.Elem(ox.StringT)).
					Slice("theme", "themes to use (located in /themes/THEMENAME/)", ox.Elem(ox.StringT), ox.Short("t")).
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("Create a new site in the provided directory.\nThe new site will have the correct structure, but no content or theme yet.\nUse `hugo new [contentPath]` to create new content."),
				ox.Usage("site", "Create a new site (skeleton)"),
				ox.Spec("[path] [flags]"),
				ox.Flags().
					Bool("force", "init inside non-empty directory", ox.Short("f")).
					String("format", "preferred file format (toml, yaml or json)", ox.Default("toml")).
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Sub(
				ox.Banner("Create a new theme (skeleton) called [name] in ./themes.\nNew theme is a skeleton. Please add content to the touched files. Add your\nname to the copyright line in the license and adjust the theme.toml file\naccording to your needs."),
				ox.Usage("theme", "Create a new theme (skeleton)"),
				ox.Spec("[name] [flags]"),
				ox.Flags().
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Flags().
				String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
				String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
				String("configDir", "config dir", ox.Default("config")).
				String("destination", "filesystem path to write files to", ox.Short("d")).
				String("environment", "build environment", ox.Short("e")).
				String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
				String("logLevel", "log level (debug|info|warn|error)").
				Bool("quiet", "build in quiet mode").
				Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
				String("source", "filesystem path to read files relative from", ox.Short("s")).
				String("themesDir", "filesystem path to themes directory"),
		), ox.Sub(
			ox.Banner("Hugo provides its own webserver which builds and serves the site.\nWhile hugo server is high performance, it is a webserver with limited options.\n\nThe `hugo server` command will by default write and serve files from disk, but\nyou can render to memory by using the `--renderToMemory` flag. This can be\nfaster in some cases, but it will consume more memory.\n\nBy default hugo will also watch your files for any changes you make and\nautomatically rebuild the site. It will then live reload any open browser pages\nand push the latest content to them. As most Hugo sites are built in a fraction\nof a second, you will be able to save and see your changes nearly instantly."),
			ox.Usage("server", "Start the embedded web server"),
			ox.Spec("[command] [flags]"),
			ox.Aliases("server", "serve"),
			ox.Footer("Use \"hugo server [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("Install the local CA in the system trust store."),
				ox.Usage("trust", "Install the local CA in the system trust store."),
				ox.Spec("[flags] [args]"),
				ox.Flags().
					Bool("uninstall", "Uninstall the local CA (but do not delete it).").
					String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
					String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
					String("configDir", "config dir", ox.Default("config")).
					String("destination", "filesystem path to write files to", ox.Short("d")).
					String("environment", "build environment", ox.Short("e")).
					String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
					String("logLevel", "log level (debug|info|warn|error)").
					Bool("quiet", "build in quiet mode").
					Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
					String("source", "filesystem path to read files relative from", ox.Short("s")).
					String("themesDir", "filesystem path to themes directory"),
			), ox.Flags().
				Bool("appendPort", "append port to baseURL", ox.Default("true")).
				String("baseURL", "hostname (and path) to the root, e.g. https://spf13.com/", ox.Short("b")).
				String("bind", "interface to which the server will bind", ox.Default("127.0.0.1")).
				Bool("buildDrafts", "include content marked as draft", ox.Short("D")).
				Bool("buildExpired", "include expired content", ox.Short("E")).
				Bool("buildFuture", "include content with publishdate in the future", ox.Short("F")).
				String("cacheDir", "filesystem path to cache directory").
				Bool("cleanDestinationDir", "remove files from destination not found in static directories").
				String("contentDir", "filesystem path to content directory", ox.Short("c")).
				Bool("disableBrowserError", "do not show build errors in the browser").
				Bool("disableFastRender", "enables full re-renders on changes").
				Slice("disableKinds", "disable different kind of pages (home, RSS etc.)", ox.Elem(ox.StringT)).
				Bool("disableLiveReload", "watch without enabling live browser reload on rebuild").
				Bool("enableGitInfo", "add Git revision, date, author, and CODEOWNERS info to the pages").
				Bool("forceSyncStatic", "copy all files when static is changed.").
				Bool("gc", "enable to run some cleanup tasks (remove unused cache files) after the build").
				Bool("ignoreCache", "ignores the cache directory").
				String("layoutDir", "filesystem path to layout directory", ox.Short("l")).
				Int("liveReloadPort", "port for live reloading (i.e. 443 in HTTPS proxy situations)", ox.Default("-1")).
				Bool("minify", "minify any supported output format (HTML, XML etc.)").
				Bool("navigateToChanged", "navigate to changed content file on live browser reload", ox.Short("N")).
				Bool("noBuildLock", "don't create .hugo_build.lock file").
				Bool("noChmod", "don't sync permission mode of files").
				Bool("noHTTPCache", "prevent HTTP caching").
				Bool("noTimes", "don't sync modification time of files").
				Bool("openBrowser", "open the site in a browser after server startup", ox.Short("O")).
				Bool("panicOnWarning", "panic on first WARNING log").
				String("poll", "set this to a poll interval, e.g --poll 700ms, to use a poll based approach to watch for file system changes").
				Int("port", "port on which the server will listen", ox.Default("1313"), ox.Short("p")).
				Bool("pprof", "enable the pprof server (port 8080)").
				Bool("printI18nWarnings", "print missing translations").
				Bool("printMemoryUsage", "print memory usage to screen at intervals").
				Bool("printPathWarnings", "print warnings on duplicate target paths etc.").
				Bool("printUnusedTemplates", "print warnings on unused templates.").
				Slice("renderSegments", "named segments to render (configured in the segments config)", ox.Elem(ox.StringT)).
				Bool("renderStaticToDisk", "serve static files from disk and dynamic files from memory").
				Bool("templateMetrics", "display metrics about template executions").
				Bool("templateMetricsHints", "calculate some improvement hints when combined with --templateMetrics").
				Slice("theme", "themes to use (located in /themes/THEMENAME/)", ox.Elem(ox.StringT), ox.Short("t")).
				Bool("tlsAuto", "generate and use locally-trusted certificates.").
				String("tlsCertFile", "path to TLS certificate file").
				String("tlsKeyFile", "path to TLS key file").
				String("trace", "write trace to file (not useful in general)", ox.Spec("file")).
				Bool("watch", "watch filesystem for changes and recreate as needed", ox.Default("true"), ox.Short("w")).
				String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
				String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
				String("configDir", "config dir", ox.Default("config")).
				String("destination", "filesystem path to write files to", ox.Short("d")).
				String("environment", "build environment", ox.Short("e")).
				String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
				String("logLevel", "log level (debug|info|warn|error)").
				Bool("quiet", "build in quiet mode").
				Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
				String("source", "filesystem path to read files relative from", ox.Short("s")).
				String("themesDir", "filesystem path to themes directory"),
		), ox.Flags().
			String("baseURL", "hostname (and path) to the root, e.g. https://spf13.com/", ox.Short("b")).
			Bool("buildDrafts", "include content marked as draft", ox.Short("D")).
			Bool("buildExpired", "include expired content", ox.Short("E")).
			Bool("buildFuture", "include content with publishdate in the future", ox.Short("F")).
			String("cacheDir", "filesystem path to cache directory").
			Bool("cleanDestinationDir", "remove files from destination not found in static directories").
			String("clock", "set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00").
			String("config", "config file", ox.Default("is $APPNAME.yaml|json|toml")).
			String("configDir", "config dir", ox.Default("config")).
			String("contentDir", "filesystem path to content directory", ox.Short("c")).
			String("destination", "filesystem path to write files to", ox.Short("d")).
			Slice("disableKinds", "disable different kind of pages (home, RSS etc.)", ox.Elem(ox.StringT)).
			Bool("enableGitInfo", "add Git revision, date, author, and CODEOWNERS info to the pages").
			String("environment", "build environment", ox.Short("e")).
			Bool("forceSyncStatic", "copy all files when static is changed.").
			Bool("gc", "enable to run some cleanup tasks (remove unused cache files) after the build").
			Bool("ignoreCache", "ignores the cache directory").
			String("ignoreVendorPaths", "ignores any _vendor for module paths matching the given Glob pattern").
			String("layoutDir", "filesystem path to layout directory", ox.Short("l")).
			String("logLevel", "log level (debug|info|warn|error)").
			Bool("minify", "minify any supported output format (HTML, XML etc.)").
			Bool("noBuildLock", "don't create .hugo_build.lock file").
			Bool("noChmod", "don't sync permission mode of files").
			Bool("noTimes", "don't sync modification time of files").
			Bool("panicOnWarning", "panic on first WARNING log").
			String("poll", "set this to a poll interval, e.g --poll 700ms, to use a poll based approach to watch for file system changes").
			Bool("printI18nWarnings", "print missing translations").
			Bool("printMemoryUsage", "print memory usage to screen at intervals").
			Bool("printPathWarnings", "print warnings on duplicate target paths etc.").
			Bool("printUnusedTemplates", "print warnings on unused templates.").
			Bool("quiet", "build in quiet mode").
			Slice("renderSegments", "named segments to render (configured in the segments config)", ox.Elem(ox.StringT)).
			Bool("renderToMemory", "render to memory (mostly useful when running the server)", ox.Short("M")).
			String("source", "filesystem path to read files relative from", ox.Short("s")).
			Bool("templateMetrics", "display metrics about template executions").
			Bool("templateMetricsHints", "calculate some improvement hints when combined with --templateMetrics").
			Slice("theme", "themes to use (located in /themes/THEMENAME/)", ox.Elem(ox.StringT), ox.Short("t")).
			String("themesDir", "filesystem path to themes directory").
			String("trace", "write trace to file (not useful in general)", ox.Spec("file")).
			Bool("watch", "watch filesystem for changes and recreate as needed", ox.Short("w")),
	)
}
