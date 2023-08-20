## Dêe uma estreal! :star:
Se vc gostou do projeto fc-client-server-api, por favor dêe uma estrela

## Como executar:
Execute do docker-compose na pasta 'server' para criar o banco de dados sql server:
```
docker-compose up -d
```

Execute o comando abaixo na pasta 'server':
```
go run .\cmd\server\main.go
```

Execute o comando abaixo na pasta 'client':
```
go run .\cmd\client\main.go
```

O arquivo 'cotacao.txt' gerado no 'client' fica dentro da pasta 'txt'.

Por isso estou utilizando o banco de dados sql server.

## Tecnologias implementadas:

- go 1.20
 - Router chi
 - gorm (sql server)
 - context
 - net/http
 - os
 