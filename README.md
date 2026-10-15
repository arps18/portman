<h1><img src="assets/portman-header.svg" alt="Portman" width="220" height="56"></h1>

Your friendly port manager.

See what's listening on a local TCP port.

## Why?

When a development server says “port already in use,” Portman shows which process is listening there with one command.

This is a hobby project I'm building to learn Go, testing, CLI design, and operating-system fundamentals.

## Install

Requires macOS and `lsof`.

Install through the Portman Homebrew tap:

```sh
brew install arps18/tap/portman
```

## Usage

Inspect a TCP port:

```sh
portman 3000
```

Example output, with identifying details replaced:

```text
COMMAND   PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
Python  12345 dev     3u  IPv4 0x0000      0t0  TCP 127.0.0.1:3000 (LISTEN)
```

If no listeners are found:

```text
No TCP listeners found on port 3000
```

Show help:

```sh
portman --help
```

`-h` also works.

Show the installed version:

```sh
portman --version
```

Ports must be between `1` and `65535`. Invalid input or an inspection failure exits with status `1`. A completed inspection exits with status `0`, including when no listeners are found.

## v0.2.0 changes

- Added `--help` and `-h`.
- Added `--version`, with the release version embedded during builds.
- Improved error messages for invalid port input.
- Added a clear error when `lsof` cannot be found in `PATH`.

```sh
portman 3000
```

Example output, with identifying details replaced:

```text
COMMAND   PID USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
Python  12345 dev     3u  IPv4 0x0000      0t0  TCP 127.0.0.1:3000 (LISTEN)
```

If no listeners are found:

```text
No TCP listeners found on port 3000
```

Ports must be between `1` and `65535`. Invalid input or an inspection failure exits with status `1`. A completed inspection exits with status `0`, including when no listeners are found.

## Use with AI coding agents

Portman works as a standalone CLI. Coding agents with shell access can also run it when it's installed and available on their `PATH`. Add this instruction to `CLAUDE.md`, `AGENTS.md`, or your agent's project instructions:

> When a development server reports a port conflict, run `portman <port>` to inspect its TCP listeners.

> [!NOTE]
> Portman provides terminal output; no dedicated agent integration is required.

## Development

Requires Go 1.25.1 or later.

```sh
git clone https://github.com/arps18/portman.git
cd portman
go run . 3000
```

Run tests:

```sh
go test ./...
```

Set up Git hooks:

```sh
brew install lefthook
lefthook install
```

The hooks format staged Go files before commit and run tests before push.

## Limitations

- macOS only.
- TCP listeners only.
- Results depend on the processes `lsof` can see with your permissions.
- Inspection only; Portman does not stop processes.

## Roadmap

Planned iterations after v0.1.0. Scope may change as the project develops.

- **v0.2.0:** Help and version commands, clearer error messages.
- **v0.3.0:** List all visible TCP listeners.
- **v0.4.0:** JSON output for scripts and other tools.
- **v0.5.0:** Stop a process listening on a port, with confirmation and graceful termination.
- **v0.6.0:** Linux support.
- **v0.7.0:** Watch a port and report when its listeners change.

## License

Portman's code is licensed under the [MIT License](LICENSE).
The Google Cloud icon artwork retains its separate Apache 2.0 license.

## Credits

Icon sourced from [OpenSVG](https://opensvg.dev/). Original artwork: Google Cloud's “Developer Portal” icon from [Google Cloud Icons](https://github.com/iconify/icon-sets/blob/master/json/gcp.json), listed under [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0).
