# SSH Manager CLI

Este é um gerenciador de conexões SSH via linha de comando. Ele permite adicionar, editar, excluir e listar conexões SSH, além de definir uma senha global que pode ser usada para conexões sem senha.

## 📦 Instalação

Para rodar o projeto, você precisa ter o [Go](https://golang.org/) instalado em sua máquina.

### Passos

1. Clone o repositório:

```bash
git clone https://github.com/seuusuario/ssh-manager.git
cd ssh-manager
go mod tidy
```

## 🚀 Comandos

### Adiciona uma nova conexão SSH.

```bash
sshm add --alias <alias> --host <host> --user <user> --port <port> --tags <tags> --group <group> --password <password>
```

- --alias: Nome da conexão.
- --host: IP ou domínio do servidor SSH.
- --user: Usuário SSH.
- --port: Porta do servidor SSH.
- --tags: Tags associadas à conexão.
- --group: Grupo ao qual a conexão pertence.
- --password: Senha da conexão.

### Edita uma conexão SSH existente.

```bash
sshm edit --alias <alias> [--host <host>] [--user <user>] [--port <port>] [--tags <tags>] [--group <group>] [--password <password>]
```

- --alias: Nome da conexão.
- --host, --user, --port, --tags, --group, --password: Campos opcionais

### Exclui uma conexão SSH pelo alias.

```bash
sshm delete --alias <alias>
```

- --alias: Nome da conexão.

### Define uma senha global para conexões SSH.

```bash
sshm set-password
```

### Lista todas as conexões SSH.

```bash
sshm list
```
