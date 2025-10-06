# BANK

Esse projeto realiza algumas operações de um sistema bancario.

## Principais dependências do projeto

- Golang 1.25
- Gin-gonic
- Goroutines 
- Postgres
- Docker
- Testcontainers

## Arquitetura do projeto
Para organização lógica e de pastas me inspirei nos projetos que já trabalho e que usam arquitetura hexagonal.

### Estrutura do projeto
- `internal/`: Código da aplicação
  - `infra/`: Diretório onde ficam os pacotes de infraestrutura, por exemplo configuração
  - `core/`: Diretório onde fica o core da aplicação como dominio, interfaces(ports), funções de utilidades e etc. 
    - `domain/`: Diretório onde fica os domínios da aplicacao.
    - `services/`: Diretório onde fica a camada de serviço.
    - `ports/`: Diretório onde fica interfaces do projeto.
  - `adapter/`: Comunicação com o mundo externo, seja expondo interface(drivers) para ser consumido ou consumindo recursos externos(drivens)  
- `docs/`: Diretório onde ficam os arquivos do swagger.
- `migrations/`: Arquivos com migrações de banco de dados.
- `cmd/`: Entrypoint da aplicação.
- `Dockerfile`: Arquivos de migração do banco de dados
- `Makefile`: Comandos do projeto
- `docker-compose.yml`: Arquivo de configuração usado pelo ```docker compose```
- `config/`: Diretório onde ficam gestão de variáveis de ambiente e outras configurações.


## Rodando a aplicação

Para rodar a aplicação você precisar ter o ```Docker``` e o ```Docker Compose``` instalados.
O projeto usa o ```make``` também para consolidar os principais comandos.

Crie um arquivo ```.env``` copiando o ```.env-example``` na raiz do projeto pois é nele onde ficam as variáveis de ambientes que sao usadas no projeto.

### Rodando o projeto

```
$ make run
```

Isso vai rodar os containers para o banco de dados e aplicação.

### Rodando a interface de consulta HTTP REST.
Se quiser rodar só a api, caso já tenha um container de banco rodando.

obs: Não esqueça de configurar o ```.env``` com essas informações do banco de dados. 

```
$ make run-api
```

Pronto, agora basta acessar a rota ```http://localhost:8080/bank/docs/index.html``` para visualizar a documentação interativa e poder
fazer uso da API desenvolvida.

Alguns exemplos de consultas a API desenvolvida:


```
#Criando uma account
curl -X 'POST' \
  'http://localhost:8080/bank/api/v1/accounts' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "document": "123458"
}'
```

```
# Consultando uma account
curl -X 'GET' \
  'http://localhost:8080/bank/api/v1/accounts/1' \
  -H 'accept: application/json'
```

```
# Criando uma transaction
curl -X 'POST' \
  'http://localhost:8080/bank/api/v1/transactions' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "account_id": 1,
  "amount": 300.7,
  "operation_type_id": 1
}'
```

