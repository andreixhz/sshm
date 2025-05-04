# SSH Manager CLI

This is an SSH connection manager via command line. It allows you to add, edit, delete, and list SSH connections, as well as set a global password that can be used for connections without a password.

## 📦 Installation

### Download via git

```bash
go install github.com/andreixhz/sshm@latest
```

### Using the repository

To run the project, you need to have [Go](https://golang.org/) installed on your machine.

### Steps

1. Clone the repository:

```bash
git clone https://github.com/andreixhz/sshm.git
cd ssh-manager
go mod tidy
```

## 🚀 Commands

### `add` - Add a new SSH connection

```bash
sshm add --alias <alias> --host <host> --user <user> --port <port> --tags <tags> --group <group> --password <password>
```

Parameters:
- `--alias`: Connection name
- `--host`: Server IP or domain
- `--user`: SSH user
- `--port`: SSH server port
- `--tags`: Tags associated with the connection
- `--group`: Group to which the connection belongs
- `--password`: Connection password

### `edit` - Edit an existing SSH connection

```bash
sshm edit --alias <alias> [--host <host>] [--user <user>] [--port <port>] [--tags <tags>] [--group <group>] [--password <password>]
```

Parameters:
- `--alias`: Connection name
- `--host`, `--user`, `--port`, `--tags`, `--group`, `--password`: Optional fields to update

### `delete` - Delete an SSH connection

```bash
sshm delete --alias <alias>
```

Parameters:
- `--alias`: Name of the connection to be deleted

### `set-password` - Set global password

```bash
sshm set-password
```

Sets a global password for SSH connections that don't have their own password.

### `list` - List all SSH connections

```bash
sshm list
```

Displays all SSH connections registered in the system.

## 📄 License

This project is licensed under the MIT license.
