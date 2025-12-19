# Contribuindo para To-Do List API

Obrigado por considerar contribuir com este projeto! 🎉

## 📋 Índice

- [Código de Conduta](#código-de-conduta)
- [Como Posso Contribuir?](#como-posso-contribuir)
- [Diretrizes de Desenvolvimento](#diretrizes-de-desenvolvimento)
- [Processo de Pull Request](#processo-de-pull-request)
- [Convenções de Código](#convenções-de-código)

## 📜 Código de Conduta

Este projeto segue um código de conduta. Ao participar, você concorda em manter um ambiente respeitoso e acolhedor para todos.

## 🤝 Como Posso Contribuir?

### Reportando Bugs

Antes de reportar um bug, verifique se ele já não foi reportado. Se não, crie uma issue incluindo:

- **Descrição clara** do problema
- **Passos para reproduzir** o bug
- **Comportamento esperado** vs **comportamento atual**
- **Versão do Go** e sistema operacional
- **Logs** relevantes, se houver

### Sugerindo Melhorias

Sugestões de melhorias são sempre bem-vindas! Ao criar uma issue de feature:

- Use um **título claro e descritivo**
- Forneça uma **descrição detalhada** da melhoria sugerida
- Explique **por que** essa melhoria seria útil
- Se possível, forneça **exemplos** de como seria usado

### Contribuindo com Código

1. **Fork** o repositório
2. **Clone** seu fork localmente
3. **Crie uma branch** para sua feature (`git checkout -b feature/MinhaFeature`)
4. **Faça suas alterações** seguindo as [convenções de código](#convenções-de-código)
5. **Commit** suas mudanças (`git commit -m 'Add: MinhaFeature'`)
6. **Push** para a branch (`git push origin feature/MinhaFeature`)
7. Abra um **Pull Request**

## 🛠️ Diretrizes de Desenvolvimento

### Configuração do Ambiente

```bash
# Clone seu fork
git clone https://github.com/seu-usuario/to-do-list.git
cd to-do-list

# Instale as dependências
make install

# Configure o banco de dados
docker-compose up -d db
make db-migrate

# Execute os testes
make test

# Execute a aplicação
make run
```

### Executando Testes

Sempre execute os testes antes de submeter um PR:

```bash
# Testes unitários
make test

# Testes com cobertura
make test-coverage

# Linter
make lint

# Formatação
make fmt
```

### Escrevendo Testes

- Todo novo código deve ter testes correspondentes
- Mantenha a cobertura de código acima de 80%
- Use table-driven tests quando apropriado
- Nomeie os testes de forma descritiva: `TestFunctionName_Scenario_ExpectedResult`

Exemplo:

```go
func TestCreateTask_EmptyTitle_ReturnsError(t *testing.T) {
    // Arrange
    service := NewTaskService(mockRepo)
    
    // Act
    err := service.CreateTask("", "description")
    
    // Assert
    if err == nil {
        t.Error("Expected error for empty title")
    }
}
```

## 🔄 Processo de Pull Request

1. **Atualize** o README.md se necessário
2. **Atualize** a documentação da API se adicionar/modificar endpoints
3. **Garanta** que todos os testes passam
4. **Formate** o código com `make fmt`
5. **Execute** o linter com `make lint`
6. **Mantenha** commits organizados e com mensagens claras
7. **Referencie** issues relacionadas no PR (ex: "Closes #123")

### Revisão

- Pelo menos um mantenedor deve revisar e aprovar o PR
- Mudanças podem ser solicitadas antes da aprovação
- Uma vez aprovado, um mantenedor fará o merge

## 📝 Convenções de Código

### Estilo de Código

- Siga as [Effective Go guidelines](https://golang.org/doc/effective_go)
- Use `gofmt` para formatação
- Use `golangci-lint` para verificar problemas de código
- Mantenha funções pequenas e focadas (idealmente < 50 linhas)
- Documente funções exportadas

### Convenções de Nomenclatura

- **Variáveis**: camelCase (`userName`, `taskList`)
- **Constantes**: UPPER_CASE (`MAX_RETRIES`, `DB_TIMEOUT`)
- **Funções/Métodos**: PascalCase para exportados, camelCase para privados
- **Interfaces**: Use sufixo `-er` quando apropriado (`TaskRepository`)

### Commits

Use mensagens de commit semânticas:

- `Add:` para novas features
- `Fix:` para correções de bugs
- `Update:` para atualizações de código existente
- `Refactor:` para refatoração sem mudança de funcionalidade
- `Docs:` para mudanças na documentação
- `Test:` para adicionar/modificar testes
- `Chore:` para tarefas de manutenção

Exemplos:

```
Add: endpoint para atualizar tarefa
Fix: erro ao buscar tarefa inexistente
Update: validação de título da tarefa
Refactor: extrai lógica de validação para função separada
Docs: adiciona exemplos de uso da API
Test: adiciona testes para TaskService
Chore: atualiza dependências
```

### Estrutura de Branches

- `main`: branch principal (protegida)
- `develop`: branch de desenvolvimento
- `feature/*`: novas features
- `fix/*`: correções de bugs
- `refactor/*`: refatorações
- `docs/*`: documentação

## 🔍 Checklist Antes de Submeter

- [ ] Código formatado com `make fmt`
- [ ] Linter passou sem erros (`make lint`)
- [ ] Todos os testes passam (`make test`)
- [ ] Novos testes adicionados para novo código
- [ ] Documentação atualizada se necessário
- [ ] Commit messages seguem as convenções
- [ ] Branch atualizada com a main/develop

## 🆘 Precisa de Ajuda?

- Abra uma **issue** com sua dúvida
- Entre em contato através dos canais de comunicação do projeto
- Consulte a **documentação** no README.md

## 🙏 Agradecimentos

Obrigado por contribuir! Sua ajuda é muito valiosa para o projeto. ❤️

