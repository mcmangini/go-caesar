# go-caesar

Caesar cipher encryption/decryption, written in Go.

***Disclaimer:*** *Please do not actually use this to encrypt anything important! The Caesar cipher is not a secure form of encryption.*

## Building and Installation

```sh
git clone https://github.com/mcmangini/go-caesar.git
cd go-caesar/src
make
sudo make install

# To uninstall:
sudo make uninstall
```

## Usage

```
Usage: caesar [-bfh] [-s <SIZE>] <TEXT>
Encrypt/decrypt text using a Caesar cipher.

Options:
  -b, --brute-force          attempt to decrypt using English-language
                             character frequency analysis;
                             unreliable for small inputs;
                             overrides -s, --shift
  -f, --file                 read text from file (interpret TEXT as file path)
  -h, --help                 display usage message
  -s, --shift <SIZE>         shift text by SIZE characters

Exit status:
 0  if OK,
 1  if error.
```

### Examples

```sh
caesar -s 5 "Hello World\! This is an example of the Caesar cipher in action." > caesar_example.txt
cat caesar_example.txt
#   Mjqqt Btwqi! Ymnx nx fs jcfruqj tk ymj Hfjxfw hnumjw ns fhynts.
caesar -b -f caesar_example.txt
#   Hello World! This is an example of the Caesar cipher in action.
```
