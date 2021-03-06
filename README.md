# templr
[![Travis CI](https://img.shields.io/travis/gesquive/templr/master.svg?style=flat-square)](https://travis-ci.org/gesquive/templr)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/gesquive/templr/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/gesquive/templr)

An iptables firewall manager.

This program was created to help account for some of the shortcomings of iptable rules. It allows you to define rules based on url's instead of just IP addresses.

## Installing

### Compile
This project has been tested with go1.9+ on Ubuntu 16.04. Just run `go get -u github.com/gesquive/templr` and the executable should be built for you automatically in your `$GOPATH`.

Optionally you can clone the repo and run `make install` to build and copy the executable to `/usr/local/bin/` with correct permissions.

### Download
Alternately, you can download the latest release for your platform from [github](https://github.com/gesquive/templr/releases/latest).

Once you have an executable, make sure to copy it somewhere on your path like `/usr/local/bin`.
If on a \*nix system, make sure to run `chmod +x /path/to/templr`.

## Configuration

### Precedence Order
The application looks for variables in the following order:
 - command line flag
 - environment variable
 - config file variable
 - default

So any variable specified on the command line would override values set in the environment or config file.

### Config File
The application looks for a configuration file at the following locations in order:
 - `config.yml`
 - `~/.config/templr/config.yml`
 - `/etc/templr/config.yml`

If you are planning to run this app as a cron job, it is recommended that you place the config in `/etc/templr/config.yml`.

### Environment Variables
Optionally, instead of using a config file you can specify config entries as environment variables. Use the prefix "TEMPLR_" in front of the uppercased variable name. For example, the config variable `ipv4-only` would be the environment variable `TEMPLR_IPV4_ONLY`.

### Firewall Rules
`templr` uses the golang [text template engine](https://golang.org/pkg/text/template/) to generate the final ruleset. In addition to the standard [functions](https://golang.org/pkg/text/template/#hdr-Functions), `templr` has a number of helper functions designed to ease the creation of iptable rules. Please refer to the [helper documentation](https://gesquive.github.io/templr/) for a list of helper functions available.

## Imports
Other rulesets can be imported by using the `{@ glob @}` brackets, where the `glob` can be:

 - relative path to file (`../../path/to/file`)
 - absolute path to file - (`/path/to/file`)
 - a glob of files - (`*.tr`)
 - a glob of files in a directory - (`path/to/dir/*`)
 - a glob of files in multiple directories - (`path/to/dir/*/.tr`)

The specified glob is checked to see if there are file matches under a relative path first, if no relative matches are found the absolute path is checked. When multiple matches are found, they are imported in alphabetical order and delimited by a newline.

A single `glob` can be imported per set of brackets. For example:
```yaml
{@ /path/to/file @}
```

## Variables
In addition, yaml variables can be defined from within a template by using the `{$ $}` brackets. Anything within the brackets will be parsed as yaml and passed to the template as variables. For example:
```yaml
{$ dnsServers: ["google-public-dns-a.google.com", "google-public-dns-b.google.com"] $}
```
The above code creates a `.dnsServers` variable that can be referenced from elsewhere in the document like so:
```yaml
# Allow DNS lookups from {{ list .dnsServers }}
```
will print out:
```yaml
# Allow DNS lookups from google-public-dns-a.google.com, google-public-dns-b.google.com
```

An example rule template can be found at [`pkg/rules.example.yml`](https://github.com/gesquive/templr/blob/master/pkg/rules.example.yml).

### Cron Job
This application was developed to run from a scheduler such as cron.

You can use any scheduler that can run the `templr` with sufficient privledges. An example cron script can be found in the `pkg/services` directory. A logrotate script can also be found in the `pkg/services` directory. All of the configs assume the user to run as is named `templr`, make sure to change this if needed.

## Usage

```console
Manage and update your iptables firewall rules

Usage:
  templr [command]

Available Commands:
  help        Help about any command
  reload      Reload the firewall rules
  save        Output the generated firewall rules
  status      Report the firewall status
  unload      Clear the firewall, accept all traffic
  up          Bring up the firewall(s)

Flags:
  -c, --config string     config file (default is $HOME/.config/templr.yml)
  -h, --help              help for templr
  -4, --ipv4-only         Apply command to IPv4 rules only.
  -6, --ipv6-only         Apply command to IPv6 rules only.
  -l, --log-file string   Path to log file
  -p, --persist           Save the firewall configuration to netfilter-persistent
  -r, --rules string      The templated firewall rules
  -V, --version           Show the version and exit
```

Optionally, a hidden debug flag is available in case you need additional output.
```console
Hidden Flags:
  -D, --debug                  Include debug statements in log output
```

## Documentation

This documentation can be found at github.com/gesquive/templr

## License

This package is made available under an MIT-style license. See LICENSE.

## Contributing

PRs are always welcome!
