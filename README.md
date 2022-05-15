# Bhojpur SDR - Software Defined Radio

The `Bhojpur SDR` is a high performance software defined radio framework applied within the
[Bhojpur.NET Platform](https://github.com/bhojpur/platform/) ecosystem for delivery of new
wireless `applications` or `services`.

## Build Source Code

Firstly, you need [SoapySDR](https://github.com/pothosware/homebrew-pothos/wiki) library to
be able to compile the source code on a `macOS`.

```bash
brew tap pothosware/homebrew-pothos
brew update
brew install pothossoapy
```

### Generate Source

```bash
go generate ./...
```
