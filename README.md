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

### `add` - Adicionar uma nova conexão SSH

```bash
sshm add --alias <alias> --host <host> --user <user> --port <port> --tags <tags> --group <group> --password <password>
```

Parâmetros:
- `--alias`: Nome da conexão
- `--host`: IP ou domínio do servidor SSH
- `--user`: Usuário SSH
- `--port`: Porta do servidor SSH
- `--tags`: Tags associadas à conexão
- `--group`: Grupo ao qual a conexão pertence
- `--password`: Senha da conexão

### `edit` - Editar uma conexão SSH existente

```bash
sshm edit --alias <alias> [--host <host>] [--user <user>] [--port <port>] [--tags <tags>] [--group <group>] [--password <password>]
```

Parâmetros:
- `--alias`: Nome da conexão
- `--host`, `--user`, `--port`, `--tags`, `--group`, `--password`: Campos opcionais para atualização

### `delete` - Excluir uma conexão SSH

```bash
sshm delete --alias <alias>
```

Parâmetros:
- `--alias`: Nome da conexão a ser excluída

### `set-password` - Definir senha global

```bash
sshm set-password
```

Define uma senha global para conexões SSH que não possuem senha própria.

### `list` - Listar todas as conexões SSH

```bash
sshm list
```

Exibe todas as conexões SSH cadastradas no sistema.

## 📄 Licença

Este projeto está licenciado sob a licença MIT.
