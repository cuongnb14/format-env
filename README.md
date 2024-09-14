# format env
format-env is a Go command-line utility to generate and format environment files (.env) using a template. This tool allows you to specify the environment (e.g., dev, testing, staging, etc.) and a template file to generate the corresponding environment configuration file.

# Features
- Reads key-value pairs from an existing .env file.
- Applies the values to a template for formatting.
- Supports multiple stages (dev, testing, unstable, staging).
- Allows specifying the template path and stage dynamically through command-line arguments.

# Installation
To build and install the utility, make sure you have Go installed, and then run:
```
go build -o fenv fenv.go
```

Move the binary to your $PATH for easy access:
```
sudo mv fenv /usr/local/bin/
sudo chmod +x /usr/local/bin/fenv
```
# Usage

To use the utility, run it from the command line by specifying the template file and the environment stage you want to generate.
```
fenv <env_dir> <stages>
```
- env_dir: Path to the env dir (e.g., `env/`). The template file name `_template.env` must have in this folder
- stages: The environment stage, separate by comma (e.g. `dev,testing,staging`).

# Template Syntax
The template file should use Go’s text/template syntax. For example:
```
REDIS_ENABLE_SSL={{ df .REDIS_ENABLE_SSL "false" }}
DATABASE_URL={{ .DATABASE_URL }}
```
This template will render the environment variable REDIS_ENABLE_SSL and use a default value "false" if it is not provided.
