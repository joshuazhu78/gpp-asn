# gpp-asn

gpp-asn1 is a 3GPP spec asn1 extractor. It reads UTF-8 format 3GPP spec and extract the asn1 syntax. It also can use the asn syntax to generate protobuf syntax and further generating Go structure.

## How to use

The asn syntax extractor is used as below:

```console
Usage of ./build/_output/extractor:
  -asnPath string
        destination path to save the asn1 syntex (default "./asn/")
  -specTxt string
        spec text file with .txt extension and UTF-8 format (default "38331-gb0.txt")
```

To generate default NR RRC asn.1 syntax, simply:

```console
make extract
```

To generate proto buffer syntax from asn.1 syntax would require a [asn1c compiler](https://github.com/onosproject/asn1c). With the compiler installed, simply:

```console
make protos-gen
```

From the generated proto buffer syntax to Go structure, simply:

```console
make protos-go
```
