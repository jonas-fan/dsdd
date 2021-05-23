# Deep Security Diagnostic Debugger

`dsdd` is a command line tool providing simplied commands to describe Deep Security diagnostic packages, enabling interactions with UNIX tools easier and faster.

## Build

Building the source on Linux with the `build` script, or specify a target platform.

```bash
Usage: ./build [option]

Option:
    linux/amd64
    windows/amd64
```

## Usage

A couple of commands are supported, such as `ps`, `events`, etc. You can also list all of them for more details.

```bash
Usage:
  dsdd [command]

Available Commands:
  agent       Display host-specific data
  events      Display various events
  help        Help about any command
  logs        Show logs
  ps          List running processes

Flags:
  -h, --help   help for dsdd
```

Here is an example of using `dsdd` to describe Deep Security diagnostic packages.

```bash
$ cd diagnostic1163886380    # change workding directory
$ dsdd events                # list events
TIME                 ORIGIN   LEVEL  EVENT ID  EVENT
2021-03-16 08:51:44  Agent    Info   2204      Security Update: Pattern Update on Agents/Appliances Successful
2021-03-16 08:51:01  Manager  Info   273       Security Update: Security Update Check and Download Requested
```
