# go-expert-stress-test
## 🧾 Descrição do CLI – go-expert-stress-test

go-expert-stress-test é uma ferramenta de linha de comando desenvolvida em Go para realizar testes de carga em qualquer URL. O programa executa múltiplas requisições HTTP simultâneas, de acordo com os parâmetros informados, e fornece um relatório detalhado com os resultados.
## 📌 Uso:

go-expert-stress-test --url <URL> --requests <total> --concurrency <concurrency>

### Parâmetros:
- --url (obrigatório): Endereço da URL que será testada.
- --requests (obrigatório): Número total de requisições que serão feitas.
- --concurrency (obrigatório): Número de requisições simultâneas (nível de concorrência).

## 📊 Funcionalidades:

Execução paralela de requisições utilizando goroutines.

Cancelamento imediato via SIGINT ou SIGTERM (ex: Ctrl+C).

Relatório final (ou parcial, em caso de interrupção) com:

- Tempo total de execução
- Total de requisições realizadas
- Quantidade de respostas com status 200
- Distribuição de outros códigos de status HTTP
