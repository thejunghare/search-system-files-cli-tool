# File Search CLI Tool

This is a command-line interface (CLI) tool that allows you to search for files in your system.

## Usage

To use this tool, run the `search` command followed by the `--path` or `-p` flag and the name of the file you want to search for. For example:

```
Search --path myfile.txt
```

or

```
Search -p myfile.txt
```

If the file exists, the tool will return the absolute path of the file. If the file does not exist, an error message will be printed.

## Installation

To install this tool, follow these steps:

1. Clone this repository to your local machine.
2. Navigate to the root directory of the repository.
3. Run `go build` to build the binary.
4. Move the binary to a directory in your `PATH`.

## Contributing

Contributions are welcome! If you would like to contribute to this project, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
