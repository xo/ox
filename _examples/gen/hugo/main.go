// Command hugo is a xo/ox version of `hugo`.
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
		ox.Usage(`hugo`, ``),
		ox.Banner(`hugo is the main command, used to build your Hugo site.

Hugo is a Fast and Flexible Static Site Generator
built with love by spf13 and friends in Go.

Complete documentation is available at https://gohugo.io/.`),
		ox.Spec(`[flags]`),
		ox.Footer(`Use "hugo [command] --help" for more information about a command.`),
		ox.Flags().
			String(`baseURL`, `hostname (and path) to the root, e.g. https://spf13.com/`, ox.Short("b")).
			Bool(`buildDrafts`, `include content marked as draft`, ox.Short("D")).
			Bool(`buildExpired`, `include expired content`, ox.Short("E")).
			Bool(`buildFuture`, `include content with publishdate in the future`, ox.Short("F")).
			String(`cacheDir`, `filesystem path to cache directory`).
			Bool(`cleanDestinationDir`, `remove files from destination not found in static directories`).
			String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`).
			String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml")).
			String(`configDir`, `config dir`, ox.Default("config")).
			String(`contentDir`, `filesystem path to content directory`, ox.Short("c")).
			String(`destination`, `filesystem path to write files to`, ox.Short("d")).
			Slice(`disableKinds`, `disable different kind of pages (home, RSS etc.)`).
			Bool(`enableGitInfo`, `add Git revision, date, author, and CODEOWNERS info to the pages`).
			String(`environment`, `build environment`, ox.Short("e")).
			Bool(`forceSyncStatic`, `copy all files when static is changed.`).
			Bool(`gc`, `enable to run some cleanup tasks (remove unused cache files) after the build`).
			Bool(`ignoreCache`, `ignores the cache directory`).
			String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`).
			String(`layoutDir`, `filesystem path to layout directory`, ox.Short("l")).
			String(`logLevel`, `log level (debug|info|warn|error)`).
			Bool(`minify`, `minify any supported output format (HTML, XML etc.)`).
			Bool(`noBuildLock`, `don't create .hugo_build.lock file`).
			Bool(`noChmod`, `don't sync permission mode of files`).
			Bool(`noTimes`, `don't sync modification time of files`).
			Bool(`panicOnWarning`, `panic on first WARNING log`).
			String(`poll`, `set this to a poll interval, e.g --poll 700ms, to use a poll based approach to watch for file system changes`).
			Bool(`printI18nWarnings`, `print missing translations`).
			Bool(`printMemoryUsage`, `print memory usage to screen at intervals`).
			Bool(`printPathWarnings`, `print warnings on duplicate target paths etc.`).
			Bool(`printUnusedTemplates`, `print warnings on unused templates.`).
			Bool(`quiet`, `build in quiet mode`).
			Slice(`renderSegments`, `named segments to render (configured in the segments config)`).
			Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M")).
			String(`source`, `filesystem path to read files relative from`, ox.Short("s")).
			Bool(`templateMetrics`, `display metrics about template executions`).
			Bool(`templateMetricsHints`, `calculate some improvement hints when combined with --templateMetrics`).
			Slice(`theme`, `themes to use (located in /themes/THEMENAME/)`, ox.Short("t")).
			String(`themesDir`, `filesystem path to themes directory`).
			String(`trace`, `write trace to file (not useful in general)`, ox.Spec(`file`)).
			Bool(`watch`, `watch filesystem for changes and recreate as needed`, ox.Short("w")),
		ox.Sub( // hugo build
			ox.Usage(`build`, `Build your site`),
			ox.Banner(`build is the main command, used to build your Hugo site.

Hugo is a Fast and Flexible Static Site Generator
built with love by spf13 and friends in Go.

Complete documentation is available at https://gohugo.io/.`),
			ox.Spec(`[flags]`),
			ox.Flags().
				String(`baseURL`, `hostname (and path) to the root, e.g. https://spf13.com/`, ox.Short("b")).
				Bool(`buildDrafts`, `include content marked as draft`, ox.Short("D")).
				Bool(`buildExpired`, `include expired content`, ox.Short("E")).
				Bool(`buildFuture`, `include content with publishdate in the future`, ox.Short("F")).
				String(`cacheDir`, `filesystem path to cache directory`).
				Bool(`cleanDestinationDir`, `remove files from destination not found in static directories`).
				String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`).
				String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml")).
				String(`configDir`, `config dir`, ox.Default("config")).
				String(`contentDir`, `filesystem path to content directory`, ox.Short("c")).
				String(`destination`, `filesystem path to write files to`, ox.Short("d")).
				Slice(`disableKinds`, `disable different kind of pages (home, RSS etc.)`).
				Bool(`enableGitInfo`, `add Git revision, date, author, and CODEOWNERS info to the pages`).
				String(`environment`, `build environment`, ox.Short("e")).
				Bool(`forceSyncStatic`, `copy all files when static is changed.`).
				Bool(`gc`, `enable to run some cleanup tasks (remove unused cache files) after the build`).
				Bool(`ignoreCache`, `ignores the cache directory`).
				String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`).
				String(`layoutDir`, `filesystem path to layout directory`, ox.Short("l")).
				String(`logLevel`, `log level (debug|info|warn|error)`).
				Bool(`minify`, `minify any supported output format (HTML, XML etc.)`).
				Bool(`noBuildLock`, `don't create .hugo_build.lock file`).
				Bool(`noChmod`, `don't sync permission mode of files`).
				Bool(`noTimes`, `don't sync modification time of files`).
				Bool(`panicOnWarning`, `panic on first WARNING log`).
				String(`poll`, `set this to a poll interval, e.g --poll 700ms, to use a poll based approach to watch for file system changes`).
				Bool(`printI18nWarnings`, `print missing translations`).
				Bool(`printMemoryUsage`, `print memory usage to screen at intervals`).
				Bool(`printPathWarnings`, `print warnings on duplicate target paths etc.`).
				Bool(`printUnusedTemplates`, `print warnings on unused templates.`).
				Bool(`quiet`, `build in quiet mode`).
				Slice(`renderSegments`, `named segments to render (configured in the segments config)`).
				Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M")).
				String(`source`, `filesystem path to read files relative from`, ox.Short("s")).
				Bool(`templateMetrics`, `display metrics about template executions`).
				Bool(`templateMetricsHints`, `calculate some improvement hints when combined with --templateMetrics`).
				Slice(`theme`, `themes to use (located in /themes/THEMENAME/)`, ox.Short("t")).
				String(`themesDir`, `filesystem path to themes directory`).
				String(`trace`, `write trace to file (not useful in general)`, ox.Spec(`file`)).
				Bool(`watch`, `watch filesystem for changes and recreate as needed`, ox.Short("w")),
		),
		ox.Sub( // hugo config
			ox.Usage(`config`, `Display site configuration`),
			ox.Banner(`Display site configuration, both default and custom settings.`),
			ox.Spec(`[command] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "hugo config [command] --help" for more information about a command.`),
			ox.Flags().
				String(`baseURL`, `hostname (and path) to the root, e.g. https://spf13.com/`, ox.Short("b")).
				String(`cacheDir`, `filesystem path to cache directory`).
				String(`contentDir`, `filesystem path to content directory`, ox.Short("c")).
				String(`format`, `preferred file format (toml, yaml or json)`, ox.Default("toml")).
				String(`lang`, `the language to display config for. Defaults to the first language defined.`).
				Slice(`renderSegments`, `named segments to render (configured in the segments config)`).
				Slice(`theme`, `themes to use (located in /themes/THEMENAME/)`, ox.Short("t")).
				String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
				String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
				String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
				String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
				String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
				String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
				String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
				Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
				Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
				String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
				String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			ox.Sub( // hugo config mounts
				ox.Usage(`mounts`, `Print the configured file mounts`),
				ox.Banner(`Print the configured file mounts`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`baseURL`, `hostname (and path) to the root, e.g. https://spf13.com/`, ox.Short("b")).
					String(`cacheDir`, `filesystem path to cache directory`).
					String(`contentDir`, `filesystem path to content directory`, ox.Short("c")).
					Slice(`renderSegments`, `named segments to render (configured in the segments config)`).
					Slice(`theme`, `themes to use (located in /themes/THEMENAME/)`, ox.Short("t")).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
		),
		ox.Sub( // hugo convert
			ox.Usage(`convert`, `Convert front matter to another format`),
			ox.Banner(`Convert front matter to another format.

See convert's subcommands toJSON, toTOML and toYAML for more information.`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "hugo convert [command] --help" for more information about a command.`),
			ox.Flags().
				String(`output`, `filesystem path to write files to`, ox.Short("o")).
				Bool(`unsafe`, `enable less safe operations, please backup first`).
				String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
				String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
				String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
				String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
				String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
				String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
				String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
				Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
				Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
				String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
				String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			ox.Sub( // hugo convert toJSON
				ox.Usage(`toJSON`, `Convert front matter to JSON`),
				ox.Banner(`toJSON converts all front matter in the content directory
to use JSON for the front matter.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					String(`output`, `filesystem path to write files to`, ox.Short("o"), ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)).
					Bool(`unsafe`, `enable less safe operations, please backup first`, ox.Section(0)),
			),
			ox.Sub( // hugo convert toTOML
				ox.Usage(`toTOML`, `Convert front matter to TOML`),
				ox.Banner(`toTOML converts all front matter in the content directory
to use TOML for the front matter.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					String(`output`, `filesystem path to write files to`, ox.Short("o"), ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)).
					Bool(`unsafe`, `enable less safe operations, please backup first`, ox.Section(0)),
			),
			ox.Sub( // hugo convert toYAML
				ox.Usage(`toYAML`, `Convert front matter to YAML`),
				ox.Banner(`toYAML converts all front matter in the content directory
to use YAML for the front matter.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					String(`output`, `filesystem path to write files to`, ox.Short("o"), ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)).
					Bool(`unsafe`, `enable less safe operations, please backup first`, ox.Section(0)),
			),
		),
		ox.Sub( // hugo env
			ox.Usage(`env`, `Display version and environment info`),
			ox.Banner(`Display version and environment info. This is useful in Hugo bug reports`),
			ox.Spec(`[flags] [args]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
				String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
				String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
				String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
				String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
				String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
				String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
				Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
				Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
				String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
				String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
		),
		ox.Sub( // hugo gen
			ox.Usage(`gen`, `Generate documentation and syntax highlighting styles`),
			ox.Banner(`Generate documentation for your project using Hugo's documentation engine, including syntax highlighting for various programming languages.`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "hugo gen [command] --help" for more information about a command.`),
			ox.Flags().
				String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
				String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
				String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
				String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
				String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
				String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
				String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
				Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
				Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
				String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
				String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			ox.Sub( // hugo gen chromastyles
				ox.Usage(`chromastyles`, `Generate CSS stylesheet for the Chroma code highlighter`),
				ox.Banner(`Generate CSS stylesheet for the Chroma code highlighter for a given style. This stylesheet is needed if markup.highlight.noClasses is disabled in config.

See https://xyproto.github.io/splash/docs/all.html for a preview of the available styles`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`highlightStyle`, `foreground and background colors for highlighted lines, e.g. --highlightStyle "#fff000 bg:#000fff"`).
					String(`lineNumbersInlineStyle`, `foreground and background colors for inline line numbers, e.g. --lineNumbersInlineStyle "#fff000 bg:#000fff"`).
					String(`lineNumbersTableStyle`, `foreground and background colors for table line numbers, e.g. --lineNumbersTableStyle "#fff000 bg:#000fff"`).
					String(`style`, `highlighter style (see https://xyproto.github.io/splash/docs/)`, ox.Default("friendly")).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo gen doc
				ox.Usage(`doc`, `Generate Markdown documentation for the Hugo CLI.`),
				ox.Banner(`Generate Markdown documentation for the Hugo CLI.
			This command is, mostly, used to create up-to-date documentation
	of Hugo's command-line interface for https://gohugo.io/.

	It creates one Markdown file per command with front matter suitable
	for rendering in Hugo.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`dir`, `the directory to write the doc.`, ox.Default("/tmp/$APPNAMEdoc/")).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo gen man
				ox.Usage(`man`, `Generate man pages for the Hugo CLI`),
				ox.Banner(`This command automatically generates up-to-date man pages of Hugo's
	command-line interface.  By default, it creates the man page files
	in the "man" directory under the current directory.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`dir`, `the directory to write the man pages.`, ox.Default("man/")).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
		),
		ox.Sub( // hugo import
			ox.Usage(`import`, `Import a site from another system`),
			ox.Banner(`Import a site from another system.

Import requires a subcommand, e.g. `+"`"+`hugo import jekyll jekyll_root_path target_path`+"`"+`.`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "hugo import [command] --help" for more information about a command.`),
			ox.Flags().
				String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
				String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
				String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
				String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
				String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
				String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
				String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
				Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
				Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
				String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
				String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			ox.Sub( // hugo import jekyll
				ox.Usage(`jekyll`, `hugo import from Jekyll`),
				ox.Banner(`hugo import from Jekyll.
		
Import from Jekyll requires two paths, e.g. `+"`"+`hugo import jekyll jekyll_root_path target_path`+"`"+`.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`force`, `allow import into non-empty target directory`).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
		),
		ox.Sub( // hugo list
			ox.Usage(`list`, `List content`),
			ox.Banner(`List content.

List requires a subcommand, e.g. hugo list drafts`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "hugo list [command] --help" for more information about a command.`),
			ox.Flags().
				String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
				String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
				String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
				String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
				String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
				String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
				String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
				Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
				Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
				String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
				String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			ox.Sub( // hugo list all
				ox.Usage(`all`, `List all content`),
				ox.Banner(`List all content including draft, future, and expired.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo list drafts
				ox.Usage(`drafts`, `List draft content`),
				ox.Banner(`List draft content.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo list expired
				ox.Usage(`expired`, `List expired content`),
				ox.Banner(`List content with a past expiration date.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo list future
				ox.Usage(`future`, `List future content`),
				ox.Banner(`List content with a future publication date.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo list published
				ox.Usage(`published`, `List published content`),
				ox.Banner(`List content that is not draft, future, or expired.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
		),
		ox.Sub( // hugo mod
			ox.Usage(`mod`, `Manage modules`),
			ox.Banner(`Various helpers to help manage the modules in your project's dependency graph.
Most operations here requires a Go version installed on your system (>= Go 1.12) and the relevant VCS client (typically Git).
This is not needed if you only operate on modules inside /themes or if you have vendored them via "hugo mod vendor".


Note that Hugo will always start out by resolving the components defined in the site
configuration, provided by a _vendor directory (if no --ignoreVendorPaths flag provided),
Go Modules, or a folder inside the themes directory, in that order.

See https://gohugo.io/hugo-modules/ for more information.`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "hugo mod [command] --help" for more information about a command.`),
			ox.Flags().
				String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
				String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
				String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
				String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
				String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
				String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
				String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
				Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
				Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
				String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
				String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			ox.Sub( // hugo mod clean
				ox.Usage(`clean`, `Delete the Hugo Module cache for the current project.`),
				ox.Banner(`Delete the Hugo Module cache for the current project.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`all`, `clean entire module cache`).
					String(`baseURL`, `hostname (and path) to the root, e.g. https://spf13.com/`, ox.Short("b")).
					String(`cacheDir`, `filesystem path to cache directory`).
					String(`contentDir`, `filesystem path to content directory`, ox.Short("c")).
					String(`pattern`, `pattern matching module paths to clean (all if not set), e.g. "**hugo*"`).
					Slice(`renderSegments`, `named segments to render (configured in the segments config)`).
					Slice(`theme`, `themes to use (located in /themes/THEMENAME/)`, ox.Short("t")).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo mod get
				ox.Usage(`get`, `Resolves dependencies in your current Hugo Project.`),
				ox.Banner(`
Resolves dependencies in your current Hugo Project.

Some examples:

Install the latest version possible for a given module:

    hugo mod get github.com/gohugoio/testshortcodes
    
Install a specific version:

    hugo mod get github.com/gohugoio/testshortcodes@v0.3.0

Install the latest versions of all direct module dependencies:

    hugo mod get
    hugo mod get ./... (recursive)

Install the latest versions of all module dependencies (direct and indirect):

    hugo mod get -u
    hugo mod get -u ./... (recursive)

Run "go help get" for more information. All flags available for "go get" is also relevant here.

Note that Hugo will always start out by resolving the components defined in the site
configuration, provided by a _vendor directory (if no --ignoreVendorPaths flag provided),
Go Modules, or a folder inside the themes directory, in that order.

See https://gohugo.io/hugo-modules/ for more information.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo mod graph
				ox.Usage(`graph`, `Print a module dependency graph.`),
				ox.Banner(`Print a module dependency graph with information about module status (disabled, vendored).
Note that for vendored modules, that is the version listed and not the one from go.mod.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`baseURL`, `hostname (and path) to the root, e.g. https://spf13.com/`, ox.Short("b")).
					String(`cacheDir`, `filesystem path to cache directory`).
					Bool(`clean`, `delete module cache for dependencies that fail verification`).
					String(`contentDir`, `filesystem path to content directory`, ox.Short("c")).
					Slice(`renderSegments`, `named segments to render (configured in the segments config)`).
					Slice(`theme`, `themes to use (located in /themes/THEMENAME/)`, ox.Short("t")).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo mod init
				ox.Usage(`init`, `Initialize this project as a Hugo Module.`),
				ox.Banner(`Initialize this project as a Hugo Module.
	It will try to guess the module path, but you may help by passing it as an argument, e.g:
	
		hugo mod init github.com/gohugoio/testshortcodes
	
	Note that Hugo Modules supports multi-module projects, so you can initialize a Hugo Module
	inside a subfolder on GitHub, as one example.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`baseURL`, `hostname (and path) to the root, e.g. https://spf13.com/`, ox.Short("b")).
					String(`cacheDir`, `filesystem path to cache directory`).
					String(`contentDir`, `filesystem path to content directory`, ox.Short("c")).
					Slice(`renderSegments`, `named segments to render (configured in the segments config)`).
					Slice(`theme`, `themes to use (located in /themes/THEMENAME/)`, ox.Short("t")).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo mod npm
				ox.Usage(`npm`, `Various npm helpers.`),
				ox.Banner(`Various npm (Node package manager) helpers.`),
				ox.Spec(`[command] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "hugo mod npm [command] --help" for more information about a command.`),
				ox.Flags().
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
				ox.Sub( // hugo mod npm pack
					ox.Usage(`pack`, `Experimental: Prepares and writes a composite package.json file for your project.`),
					ox.Banner(`Prepares and writes a composite package.json file for your project.

On first run it creates a "package.hugo.json" in the project root if not already there. This file will be used as a template file
with the base dependency set. 

This set will be merged with all "package.hugo.json" files found in the dependency tree, picking the version closest to the project.

This command is marked as 'Experimental'. We think it's a great idea, so it's not likely to be
removed from Hugo, but we need to test this out in "real life" to get a feel of it,
so this may/will change in future versions of Hugo.`),
					ox.Spec(`[flags] [args]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`baseURL`, `hostname (and path) to the root, e.g. https://spf13.com/`, ox.Short("b")).
						String(`cacheDir`, `filesystem path to cache directory`).
						String(`contentDir`, `filesystem path to content directory`, ox.Short("c")).
						Slice(`renderSegments`, `named segments to render (configured in the segments config)`).
						Slice(`theme`, `themes to use (located in /themes/THEMENAME/)`, ox.Short("t")).
						String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
						String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
						String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
						String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
						String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
						String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
						String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
						Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
						Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
						String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
						String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
				),
			),
			ox.Sub( // hugo mod tidy
				ox.Usage(`tidy`, `Remove unused entries in go.mod and go.sum.`),
				ox.Banner(`Remove unused entries in go.mod and go.sum.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`baseURL`, `hostname (and path) to the root, e.g. https://spf13.com/`, ox.Short("b")).
					String(`cacheDir`, `filesystem path to cache directory`).
					String(`contentDir`, `filesystem path to content directory`, ox.Short("c")).
					Slice(`renderSegments`, `named segments to render (configured in the segments config)`).
					Slice(`theme`, `themes to use (located in /themes/THEMENAME/)`, ox.Short("t")).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo mod vendor
				ox.Usage(`vendor`, `Vendor all module dependencies into the _vendor directory.`),
				ox.Banner(`Vendor all module dependencies into the _vendor directory.
	If a module is vendored, that is where Hugo will look for it's dependencies.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`baseURL`, `hostname (and path) to the root, e.g. https://spf13.com/`, ox.Short("b")).
					String(`cacheDir`, `filesystem path to cache directory`).
					String(`contentDir`, `filesystem path to content directory`, ox.Short("c")).
					Slice(`renderSegments`, `named segments to render (configured in the segments config)`).
					Slice(`theme`, `themes to use (located in /themes/THEMENAME/)`, ox.Short("t")).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo mod verify
				ox.Usage(`verify`, `Verify dependencies.`),
				ox.Banner(`Verify checks that the dependencies of the current module, which are stored in a local downloaded source cache, have not been modified since being downloaded.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`baseURL`, `hostname (and path) to the root, e.g. https://spf13.com/`, ox.Short("b")).
					String(`cacheDir`, `filesystem path to cache directory`).
					Bool(`clean`, `delete module cache for dependencies that fail verification`).
					String(`contentDir`, `filesystem path to content directory`, ox.Short("c")).
					Slice(`renderSegments`, `named segments to render (configured in the segments config)`).
					Slice(`theme`, `themes to use (located in /themes/THEMENAME/)`, ox.Short("t")).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
		),
		ox.Sub( // hugo new
			ox.Usage(`new`, `Create new content`),
			ox.Banner(`Create a new content file and automatically set the date and title.
It will guess which kind of file to create based on the path provided.

You can also specify the kind with `+"`"+`-k KIND`+"`"+`.

If archetypes are provided in your theme or site, they will be used.

Ensure you run this within the root directory of your site.`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "hugo new [command] --help" for more information about a command.`),
			ox.Flags().
				String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
				String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
				String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
				String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
				String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
				String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
				String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
				Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
				Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
				String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
				String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			ox.Sub( // hugo new content
				ox.Usage(`content`, `Create new content`),
				ox.Banner(`Create a new content file and automatically set the date and title.
It will guess which kind of file to create based on the path provided.

You can also specify the kind with `+"`"+`-k KIND`+"`"+`.

If archetypes are provided in your theme or site, they will be used.

Ensure you run this within the root directory of your site.`),
				ox.Spec(`[path] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`baseURL`, `hostname (and path) to the root, e.g. https://spf13.com/`, ox.Short("b")).
					String(`cacheDir`, `filesystem path to cache directory`).
					String(`contentDir`, `filesystem path to content directory`, ox.Short("c")).
					String(`editor`, `edit new content with this editor, if provided`).
					Bool(`force`, `overwrite file if it already exists`, ox.Short("f")).
					String(`kind`, `content type to create`, ox.Short("k")).
					Slice(`renderSegments`, `named segments to render (configured in the segments config)`).
					Slice(`theme`, `themes to use (located in /themes/THEMENAME/)`, ox.Short("t")).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo new site
				ox.Usage(`site`, `Create a new site (skeleton)`),
				ox.Banner(`Create a new site in the provided directory.
The new site will have the correct structure, but no content or theme yet.
Use `+"`"+`hugo new [contentPath]`+"`"+` to create new content.`),
				ox.Spec(`[path] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`force`, `init inside non-empty directory`, ox.Short("f")).
					String(`format`, `preferred file format (toml, yaml or json)`, ox.Default("toml")).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
			ox.Sub( // hugo new theme
				ox.Usage(`theme`, `Create a new theme (skeleton)`),
				ox.Banner(`Create a new theme (skeleton) called [name] in ./themes.
New theme is a skeleton. Please add content to the touched files. Add your
name to the copyright line in the license and adjust the theme.toml file
according to your needs.`),
				ox.Spec(`[name] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
		),
		ox.Sub( // hugo server
			ox.Usage(`server`, `Start the embedded web server`),
			ox.Banner(`Hugo provides its own webserver which builds and serves the site.
While hugo server is high performance, it is a webserver with limited options.

The `+"`"+`hugo server`+"`"+` command will by default write and serve files from disk, but
you can render to memory by using the `+"`"+`--renderToMemory`+"`"+` flag. This can be
faster in some cases, but it will consume more memory.

By default hugo will also watch your files for any changes you make and
automatically rebuild the site. It will then live reload any open browser pages
and push the latest content to them. As most Hugo sites are built in a fraction
of a second, you will be able to save and see your changes nearly instantly.`),
			ox.Spec(`[command] [flags]`),
			ox.Aliases("serve"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "hugo server [command] --help" for more information about a command.`),
			ox.Flags().
				Bool(`appendPort`, `append port to baseURL`, ox.Default("true")).
				String(`baseURL`, `hostname (and path) to the root, e.g. https://spf13.com/`, ox.Short("b")).
				String(`bind`, `interface to which the server will bind`, ox.Default("127.0.0.1")).
				Bool(`buildDrafts`, `include content marked as draft`, ox.Short("D")).
				Bool(`buildExpired`, `include expired content`, ox.Short("E")).
				Bool(`buildFuture`, `include content with publishdate in the future`, ox.Short("F")).
				String(`cacheDir`, `filesystem path to cache directory`).
				Bool(`cleanDestinationDir`, `remove files from destination not found in static directories`).
				String(`contentDir`, `filesystem path to content directory`, ox.Short("c")).
				Bool(`disableBrowserError`, `do not show build errors in the browser`).
				Bool(`disableFastRender`, `enables full re-renders on changes`).
				Slice(`disableKinds`, `disable different kind of pages (home, RSS etc.)`).
				Bool(`disableLiveReload`, `watch without enabling live browser reload on rebuild`).
				Bool(`enableGitInfo`, `add Git revision, date, author, and CODEOWNERS info to the pages`).
				Bool(`forceSyncStatic`, `copy all files when static is changed.`).
				Bool(`gc`, `enable to run some cleanup tasks (remove unused cache files) after the build`).
				Bool(`ignoreCache`, `ignores the cache directory`).
				String(`layoutDir`, `filesystem path to layout directory`, ox.Short("l")).
				Int(`liveReloadPort`, `port for live reloading (i.e. 443 in HTTPS proxy situations)`, ox.Default("-1")).
				Bool(`minify`, `minify any supported output format (HTML, XML etc.)`).
				Bool(`navigateToChanged`, `navigate to changed content file on live browser reload`, ox.Short("N")).
				Bool(`noBuildLock`, `don't create .hugo_build.lock file`).
				Bool(`noChmod`, `don't sync permission mode of files`).
				Bool(`noHTTPCache`, `prevent HTTP caching`).
				Bool(`noTimes`, `don't sync modification time of files`).
				Bool(`openBrowser`, `open the site in a browser after server startup`, ox.Short("O")).
				Bool(`panicOnWarning`, `panic on first WARNING log`).
				String(`poll`, `set this to a poll interval, e.g --poll 700ms, to use a poll based approach to watch for file system changes`).
				Int(`port`, `port on which the server will listen`, ox.Default("1313"), ox.Short("p")).
				Bool(`pprof`, `enable the pprof server (port 8080)`).
				Bool(`printI18nWarnings`, `print missing translations`).
				Bool(`printMemoryUsage`, `print memory usage to screen at intervals`).
				Bool(`printPathWarnings`, `print warnings on duplicate target paths etc.`).
				Bool(`printUnusedTemplates`, `print warnings on unused templates.`).
				Slice(`renderSegments`, `named segments to render (configured in the segments config)`).
				Bool(`renderStaticToDisk`, `serve static files from disk and dynamic files from memory`).
				Bool(`templateMetrics`, `display metrics about template executions`).
				Bool(`templateMetricsHints`, `calculate some improvement hints when combined with --templateMetrics`).
				Slice(`theme`, `themes to use (located in /themes/THEMENAME/)`, ox.Short("t")).
				Bool(`tlsAuto`, `generate and use locally-trusted certificates.`).
				String(`tlsCertFile`, `path to TLS certificate file`).
				String(`tlsKeyFile`, `path to TLS key file`).
				String(`trace`, `write trace to file (not useful in general)`, ox.Spec(`file`)).
				Bool(`watch`, `watch filesystem for changes and recreate as needed`, ox.Default("true"), ox.Short("w")).
				String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
				String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
				String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
				String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
				String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
				String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
				String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
				Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
				Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
				String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
				String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			ox.Sub( // hugo server trust
				ox.Usage(`trust`, `Install the local CA in the system trust store.`),
				ox.Banner(`Install the local CA in the system trust store.`),
				ox.Spec(`[flags] [args]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`uninstall`, `Uninstall the local CA (but do not delete it).`).
					String(`clock`, `set the clock used by Hugo, e.g. --clock 2021-11-06T22:30:00.00+09:00`, ox.Section(0)).
					String(`config`, `config file`, ox.Default("is $APPNAME.yaml|json|toml"), ox.Section(0)).
					String(`configDir`, `config dir`, ox.Default("config"), ox.Section(0)).
					String(`destination`, `filesystem path to write files to`, ox.Short("d"), ox.Section(0)).
					String(`environment`, `build environment`, ox.Short("e"), ox.Section(0)).
					String(`ignoreVendorPaths`, `ignores any _vendor for module paths matching the given Glob pattern`, ox.Section(0)).
					String(`logLevel`, `log level (debug|info|warn|error)`, ox.Section(0)).
					Bool(`quiet`, `build in quiet mode`, ox.Section(0)).
					Bool(`renderToMemory`, `render to memory (mostly useful when running the server)`, ox.Short("M"), ox.Section(0)).
					String(`source`, `filesystem path to read files relative from`, ox.Short("s"), ox.Section(0)).
					String(`themesDir`, `filesystem path to themes directory`, ox.Section(0)),
			),
		),
	)
}
