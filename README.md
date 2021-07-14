# Hashmatch

<p>
<a href="https://github.com/agrimrules/hashmatch/releases/"><img alt="Version" src="https://img.shields.io/github/release-pre/agrimrules/hashmatch.svg"></a>
<img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/agrimrules/hashmatch">
<a href="https://goreportcard.com/report/github.com/agrimrules/hashmatch"><img src="https://goreportcard.com/badge/github.com/agrimrules/hashmatch" alt="Go Report Card" /a>
<a href="https://github.com/agrimrules/hashmatch/blob/main/LICENSE"><img src="https://img.shields.io/github/license/agrimrules/hashmatch" alt="License"></a>
<a href="https://github.com/agrimrules/hashmatch/actions/workflows/release.yml"><img srv="https://github.com/agrimrules/hashmatch/actions/workflows/release.yml/badge.svg"></a>
</p>

A simple CLI tool written to verify files based on various hashing algorithms.

<p align="center"><img src=".github/hashmatch.gif?raw=true"/></p>

## Installation
Hashmatch is cross platform and available on Linux, macOS and Windows.

* Install via the golang toolchain
```shell
go get -u github.com/agrimrules/hashmatch
```

* Binaries are available at the [Releases](https://github.com/agrimrules/hashmatch/releases) page, download the necessary binary for your platform and add it to your `$PATH`

## Usage
```shell
hashmatch file1 file2
```
Can be used to see if both files are the same via matching md5 sums

```shell
hashmatch /path/to/directory1 /path/to/directory2
```
Will traverse both directories and indicate if all files within them match or not.

## License

The Hashmatch cli tool is open-sourced software licensed under the [Apache-2.0 License](./LICENSE).

## Acknowledgments
  
The following projects had particular influence on the hashmatch cli.
  
- [karrick/godirwalk](https://github.com/karrick/godirwalk) helped provide quick directory traversal using a simpler API.
- [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter) Provides a simple table TUI for displaying the results of the comparison.
- [spf13/cobra](https://github.com/spf13/cobra) A Go framework for building CLI applications.
- [thoas/go-funk](https://github.com/thoas/go-funk) A Go library providing functional utilities similar to lodash.