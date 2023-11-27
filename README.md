# Aplicação Golang

Aplicação simples usando golang, rabbitmq e sqlite3

Deploy através do docker e kubernetes

## Executando api

### Pré requisitos

- Docker
- kubectl
- Kind

### Para executar api siga os passos:

- Criação de um novo cluster através do Kind
    ```
    kind create cluster
    ```

- Realizando o deploy conforme configuração do YAML
    ```
    kubectl apply -f k8s/deployment.yaml
    kubectl apply -f k8s/service.yaml
    ```

- Compartilhando o serviço para ser possível acessar no browser
    ```
    kubectl port-forward svc/goapp-service 8888:8888
    ```

## Executando order

Executar o serviço de backend qual fica escutando a fila do rabbitmq e no caso de um novo registro publicado será feito a leitura e consolidado no banco de dados

### Pré requisitos

- sqlite3
- tabela criada conforme arquivo table.sql

### Para executar siga os passos:

- Start da aplicação local, com isso o rabbitmq ficará escutando a fila
    ```
    go run cmd/order/main.go
    ```

- Publicar novo item na fila, no padrão JSON abaixo
    ```
    {"id": "1", "price": 10.0, "tax": 0.1}
    ```

A saída deve ser a mensagem:
```
Mensagem processada e salva no banco: &<JSON>
``````