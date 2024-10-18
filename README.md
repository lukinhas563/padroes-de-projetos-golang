# 📚 Introdução: Padrões de Projeto em Go

Este artigo vai te mostrar como implementar **padrões de projeto** (_design patterns_) na linguagem **Go** de forma prática e eficiente! Padrões de projeto são soluções reutilizáveis para problemas comuns no desenvolvimento de software, e ajudam a deixar seu código mais flexível e organizado.

### 🤔 Por que usar padrões de projeto?

Os padrões de projeto são como "receitas de bolo" que te ajudam a resolver desafios frequentes de design. Eles oferecem:

- 🧹 **Código mais organizado**: Ajuda a manter tudo claro e fácil de entender.

- 📝 **Reutilização de soluções**: Usar abordagens comprovadas para diferentes problemas.

- 📢 **Comunicação mais fácil**: Falar a "mesma língua" com outros desenvolvedores.

- 🐞 **Menos bugs**: Soluções já testadas e otimizadas.

- 🪜 **Flexibilidade**: Adaptar o código a mudanças futuras sem grandes dores de cabeça.

### Quando e quem criou os padrões?

Essas ideias vieram da arquitetura, mas no mundo do software foram popularizadas em 1994 por um grupo chamado **_Gang of Four_** (**GoF**) no famoso livro "**_Design Patterns: Elements of Reusable Object-Oriented Software_**". Eles documentaram 23 padrões que hoje são amplamente usados no desenvolvimento de software.

---

### 🏗️ Padrões Criacionais

Os padrões criacionais ajudam a criar objetos de forma mais flexível, sem precisar expor muita lógica para quem vai usar.

Aqui estão alguns exemplos implementados em Go:

1. **Singleton**: [singleton.md](./examples/criacional/singleton/singleton.md)
    - **O que faz**: Garante que uma classe tenha apenas uma instância.

2. **Builder**: [builder.md](./examples/criacional/builder/builder.md)
    - **O que faz**: Permite construir objetos complexos passo a passo.

3. **Prototype**: [prototype.md](./examples/criacional/prototype/prototype.md)
    - **O que faz**: Cria novos objetos copiando uma instância existente.

4. **Factory Method**: [factory.md](./examples/criacional/factory/factory.md)
    - **O que faz**: Fornece uma interface para criar objetos, mas permite que subclasses decidam quais classes instanciar.

5. **Abstract Factory**: [abstract_factory.md](./examples/criacional/abstract_factory/abstract_factory.md)
    - **O que faz**: Cria famílias de objetos relacionados sem especificar suas classes concretas.

---

### 🧱 Padrões Estruturais
Esses padrões lidam com a organização de classes e objetos para criar estruturas maiores e eficientes.

Exemplos em Go:

1. **Adapter**: Ver código
    - **O que faz**: Converte a interface de um objeto para outra, permitindo que trabalhem juntos.

2. **Composite**: [composite.md](./examples/estrutural/composite/composite.md)
    - **O que faz**: Organiza objetos em estruturas hierárquicas de árvore, facilitando o gerenciamento.

3. **Decorator**: Ver código
    - **O que faz**: Adiciona dinamicamente responsabilidades adicionais a um objeto sem alterar sua estrutura.

4. **Facade**: Ver código
    - **O que faz**: Fornece uma interface simplificada para um conjunto complexo de subsistemas.

5. **Proxy**: Ver código
    - **O que faz**: Fornece um substituto ou placeholder para outro objeto para controlar o acesso a ele.

---

### 🎯 Conclusão

Os padrões de projeto são ferramentas essenciais para qualquer desenvolvedor. Mesmo que Go não seja uma linguagem orientada a objetos pura, podemos adaptar esses padrões usando interfaces e funções. A ideia é manter o código **limpo**, **flexível** e **fácil de manter**.
