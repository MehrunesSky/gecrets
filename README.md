# Gecrets

Gecrets is a Go application that allows you to edit sensitive secrets directly from the command line using VIM. It’s designed to simplify the management and updating of secrets while ensuring security.
Features

    Securely load and edit secrets in a controlled environment.
    Uses VIM as the default editor for seamless editing.
    Easily manage and update sensitive information.

## Installation

Clone the repository and build the application:

```shell
git clone https://github.com/MehrunesSky/gecrets.git
cd gecrets
go build -o gecrets
```

Then add the binary to your PATH:

```shell
export PATH=$PATH:$(pwd)
```

## Usage

To show secrets, use the following command:
```shell
gecrets list --ks <keyStoreName>
```
This will open the specified secret in VIM for secure editing.
Examples

Edit secrets named :
```shell
gecrets update --ks <keyStoreName>
```

## Contributing

Contributions are welcome! Feel free to submit pull requests or open issues.
License

This project is licensed under the MIT License.