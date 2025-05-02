# :bookmark: difson
Compare JSON files with ease

![MIT License](https://img.shields.io/badge/license-MIT-blue "MIT License")
[![Go Report Card](https://goreportcard.com/badge/github.com/sorahashiroi/difson)](https://goreportcard.com/report/github.com/sorahashiroi/difson)
[![Coverage Status](https://coveralls.io/repos/github/sorahashiroi/difson/badge.svg?branch=main)](https://coveralls.io/github/sorahashiroi/difson?branch=main)

![Version](https://img.shields.io/badge/Version-0.1.3-blue)

## :pushpin: Overview
**difson** is a CLI tool to compare two JSON files and show their differences clearly.

## :hammer_and_pick: Usage

```sh
difson [OPTIONS] <FILE1> <FILE2>
OPTIONS:
    -c, --color          Show colorized output
    -b, --brief          Print a simple message: "Differences detected." or "No differences found."
    -p, --pretty         Print diff with indentation and formatting
    -h, --help           print this message.
```

## :file_folder: Installation

## :label: About
`difson` was built to make it easy to compare JSON files in a human-readable way.
It's especially useful for checking differences between configuration files, API responses, or exported data.

This tool is:

- :mag: Focused — Only does JSON diff, but does it well

- :cloud: Lightweight — No heavy dependencies or setup

- :sparkles: Script-friendly — Works well in CI or shell pipelines

Contributions, issues, and feature suggestions are welcome!

## Autors
- [Mizuki Yamano](https://github.com/sorahashiroi)