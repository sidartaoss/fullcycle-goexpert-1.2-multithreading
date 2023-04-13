# Go Expert - Multithreading

Este desafio compreende juntar conhecimentos de _Multithreading e APIs_ para buscar o resultado mais rápido entre 2 _endpoints_ distintos.

Antes, uma nota sobre estrutura de diretórios.

Geralmente, ao iniciar um projeto, é comum ficar em dúvida com relação à organização dos diretórios. Em vista disso, foi criado um projeto no _GitHub_, o _Standard Go Project Layout_ (https://github.com/golang-standards/project-layout), que se propõe servir como um _layout_ básico de aplicações _Go_.

Não é obrigatório e nem todos os projetos seguem esse _layout_, mas, como um ponto de partida, é de grande ajuda.

O _Go Project Layout_, então, guarda uma convenção de como nomear os diretórios conforme o papel de cada um:

- _api_: guarda arquivos relacionados à documentação dos _endpoints_ das _APIs_. (Não é o local onde fica o código da _API_);

- _internal_: diz respeito aos pacotes privados da aplicação, sendo que o código não deve ser reaproveitado/importado em outras aplicações, porque ele resolve problemas específicos do negócio da aplicação;

- _pkg_: é onde ficam as bibliotecas que podem ser compartilhadas/importadas em outras aplicações. Por exemplo, uma biblioteca de autenticação pode ser usada por diversas aplicações;

- _cmd_: é o local aonde é gerado o executável. Corresponde ao diretório a partir da onde se roda a aplicação.

- _configs_: é o lugar onde podem ser guardados _templates_ de configuração ou arquivos _Go_ de configuração para fazer o _boot_ da aplicação; inclui configurações padrão de como a aplicação vai iniciar, aonde se define o padrão das variáveis de ambiente, por exemplo;

- _test_: contempla arquivos adicionais que vão ajudar no processo de teste, como documentação de teste, exemplos de teste, _stubs_ de teste, arquivos _HTTP_ para testes _end-to-end_; portanto, não necessariamente inclui arquivos ._go_.

Neste desafio, 2 requisições são feitas simultaneamente para os _endpoints_: 1. https://viacep.com.br/ws/%s/json; 2. https://example.api.findcep.com/v1/cep/%s.json, onde %s refere-se ao valor do _CEP_.

Assim, os requisitos para este desafio incluem:

- Acatar o _endpoint_ que entregar a resposta mais rápida e descartar a resposta mais lenta;
- Exibir no _command-line_ (_CLI_) o resultado do _request_, assim como qual _endpoint_ que o enviou;
- Limitar o tempo de resposta em 1 segundo. Caso contrário, exibir o erro de _timeout_.

### Execução

```sh
curl localhost:8000?cep=01234000
```
