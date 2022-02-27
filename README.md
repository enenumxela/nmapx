# nmapx

A [go(golang)](http://golang.org/) CLI utility to merge and/or parse nmap XML output.

## Resources

* [Installation](#installation)
	* [From Binary](#from-binary)
	* [From source](#from-source)
	* [From github](#from-github)
* [Contribution](#contribution)

## Installation

### From Binary

You can download the pre-built binary for your platform from this repository's [releases](https://github.com/enenumxela/nmapx/releases/) page, extract, then move it to your `$PATH`and you're ready to go.

### From Source

nmapx requires **go1.17+** to install successfully. Run the following command to get the repo

```bash
go install -v github.com/enenumxela/nmapx/cmd/nmapx@latest
```

### From Github

```bash
git clone https://github.com/enenumxela/nmapx.git && \
cd nmapx/cmd/nmapx/ && \
go build; mv nmapx /usr/local/bin/ && \
nmapx -h
```

## Contribution

[Issues](https://github.com/enenumxela/nmapx/issues) and [Pull Requests](https://github.com/enenumxela/nmapx/pulls) are welcome!