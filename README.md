# cloudkit

CLI em Go que interage com serviços AWS reais via `aws-sdk-go-v2`, rodando 100% local com **MiniStack** — um emulador AWS open-source, sem conta, sem login e sem custo. O objetivo é exercitar os serviços AWS mais usados no dia a dia (S3, SQS, Secrets Manager, IAM e EKS) num ambiente que qualquer pessoa pode subir com um `docker run`.

## Por que MiniStack

Para usar a AWS de verdade você precisa de conta, cartão cadastrado e paga por uso. O MiniStack finge ser a AWS na sua máquina: é open-source (MIT), **não pede conta nem token** e já suporta 40+ serviços — incluindo EKS de graça (que roda um **k3s real** por trás, não um mock).

Vantagem prática: clonou o repositório, subiu o MiniStack, está funcionando. Sem cadastro em lugar nenhum.

## Serviços cobertos

| Serviço | O que faz |
|---|---|
| **S3** | Armazenamento de arquivos (buckets e objetos) |
| **SQS** | Fila de mensagens (produtor/consumidor) |
| **Secrets Manager** | Cofre de senhas e tokens buscados em runtime |
| **IAM** | Controle de acesso (roles e policies) |
| **EKS** | Kubernetes gerenciado (k3s real por baixo) |

## Pré-requisitos

- Go 1.22+
- Docker

## Como rodar

### 1. Suba o MiniStack

```bash
docker run -d --name ministack -p 4566:4566 nahuelnucera/ministack
```

Ou via Docker Compose:

```bash
docker-compose up -d
```

Verifique se subiu:

```bash
curl http://localhost:4566/_ministack/health
```

> A flag `privileged: true` no `docker-compose.yml` só é necessária para o EKS, porque ele roda um k3s real internamente.

### 2. Rode os comandos

```bash
# S3
cloudkit s3 upload arquivo.txt meu-bucket
cloudkit s3 list meu-bucket
cloudkit s3 download meu-bucket arquivo.txt

# SQS
cloudkit sqs send minha-fila "mensagem de teste"
cloudkit sqs receive minha-fila

# Secrets Manager
cloudkit secrets get db-password

# EKS
cloudkit eks create-cluster meu-cluster
cloudkit eks list-clusters
```

## Exemplo de output

```
$ cloudkit s3 list meu-bucket
Nome: teste.txt | Tamanho: 42 bytes
```

## Como o Go se conecta na AWS

Toda chamada segue o mesmo padrão, que se repete para qualquer um dos +200 serviços da AWS:

1. Cria uma configuração (`aws.Config`) — diz onde a AWS (ou o MiniStack) está
2. Cria um client do serviço (`s3`, `sqs`, etc) a partir dessa config
3. Chama os métodos do client (`PutObject`, `SendMessage`, ...)

O client aponta para o MiniStack via `BaseEndpoint` (`http://localhost:4566`) e usa credenciais fake — em produção, bastaria remover o endpoint customizado e usar credenciais reais.

## Estrutura

```
cloudkit/
├── cmd/
│   └── main.go        # CLI com subcomandos
├── internal/
│   ├── s3client/
│   ├── sqsclient/
│   ├── secrets/
│   └── eksclient/
├── docker-compose.yml
└── README.md
```

## Licença

MIT
