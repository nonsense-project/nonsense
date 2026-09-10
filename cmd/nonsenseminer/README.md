# nonsenseminer

`nonsenseminer` is a CPU-based miner for `nonsensed`.

## Requirements

Go 1.24 or later.

## Build from Source

* Install Go according to the installation instructions here:
  http://golang.org/doc/install

* Ensure Go was installed properly and is a supported version:

```bash
go version
```

* Run the following commands to obtain and install `nonsensed`
  including all dependencies:

```bash
git clone https://github.com/nonsense-project/nonsense
cd nonsensed/cmd/nonsenseminer
go install .
```

* `nonsenseminer` should now be installed in `$(go env GOPATH)/bin`.
  If you did not already add the bin directory to your system path
  during Go installation, you are encouraged to do so now.

## Usage

The full `nonsenseminer` configuration options can be seen with:

```bash
nonsenseminer --help
```

But the minimum configuration needed to run it is:

```bash
nonsenseminer --miningaddr=<YOUR_MINING_ADDRESS>
```
