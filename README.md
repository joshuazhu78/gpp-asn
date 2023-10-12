# gpp-asn

gpp-asn is a 3GPP spec asn1 extractor. It reads UTF-8 format 3GPP spec and extract the asn1 syntax. It also can use the asn syntax to generate the protobuf syntax and further generate Go structure. It also automates the customerized json marshal and unmarshal functions for all generated Go interfaces. So the generated Go structure can be serialized to bytes using json or asn.1 schemas.

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

To generate proto buffer syntax from asn.1 syntax would require a [asn1c compiler](https://github.com/joshuazhu78/asn1c). With the compiler installed, simply:

```console
make protos-gen
```

From the generated proto buffer syntax to Go structure, simply:

```console
make protos-go
```

From the generated Go structure to automate the json marshal/unmarshal functions for interfaces:
```console
make code-gen
```
