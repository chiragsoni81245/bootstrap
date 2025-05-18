# How to run

### Install Golang

### Clone this repository

### Install packages 
```bash
go mod tidy
```

### Make a config file for the project you want to initiate
Example config for `go_with_vanilla_auth` template
```yaml
project:
    folder_name: testing
```

### Run this command to initialize this project
```bash
make run ARGS="--template-name go_with_vanilla_auth --config config.yaml"
```
