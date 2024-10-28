# 📖 O que é o padrão Decorator?

O **Decorator** é um padrão de projeto **estrutural** que permite adicionar funcionalidades a objetos dinamicamente. Em vez de modificar diretamente uma classe, ele cria "camadas" adicionais que envolvem o objeto original, permitindo uma extensão flexível e reutilizável de funcionalidades.

### 🤔 Qual problema o Decorator resolve?

O Decorator resolve o problema de adicionar funcionalidades a objetos sem precisar criar subclasses para cada combinação de funcionalidades. Em vez de gerar uma árvore complexa de classes para cobrir todas as variações, o Decorator permite que você adicione funcionalidades "embrulhando" um objeto original em decoradores.

---

### 📦 Exemplo prático

Imagine que você está criando uma cafeteria onde os clientes podem escolher diferentes tipos de café e adicionar complementos, como leite, chocolate ou caramelo. Em vez de criar classes separadas para cada combinação de café com e sem os complementos, o Decorator permite que você adicione funcionalidades adicionais a cada pedido, como uma camada extra de chocolate ou leite, conforme solicitado.

---

### 🚗 Analogia: Decorações de um Presente 🎁

Imagine que você quer dar um presente, e decide envolvê-lo em várias camadas de embrulho. Primeiro, usa um papel de presente comum, depois adiciona uma fita e, em seguida, uma etiqueta decorativa. Cada camada adiciona um "decorador" ao presente original, sem precisar modificar o conteúdo.

---

### 🧱 Estrutura do Decorator

A estrutura do padrão Decorator inclui:

- **Component**: Define a interface do objeto original que será "decorado".

- **ConcreteComponent**: Implementação do `Component`, representando o objeto principal.

- **Decorator**: Classe base para todos os decoradores, que segue a mesma interface do `Component`.

- **ConcreteDecorator**: Implementações concretas do `Decorator`, que adicionam funcionalidades específicas ao objeto.

---

### 👨‍💻 Como implementar o Decorator em Go?
Vamos ver como o padrão Decorator pode ser implementado em Go com um exemplo de cafeteria.

#### Exemplo em Go:
```go
package main

import "fmt"

// Component define a interface básica de um café.
type Coffee interface {
    Cost() float64
    Description() string
}

// ConcreteComponent: Implementação base de um tipo simples de café.
type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() float64 {
    return 5.0
}

func (s *SimpleCoffee) Description() string {
    return "Café simples"
}

// Decorator: Base para decoradores de café, implementando a interface Coffee.
type CoffeeDecorator struct {
    coffee Coffee
}

func (d *CoffeeDecorator) Cost() float64 {
    return d.coffee.Cost()
}

func (d *CoffeeDecorator) Description() string {
    return d.coffee.Description()
}

// ConcreteDecorator: Adiciona leite ao café.
type MilkDecorator struct {
    *CoffeeDecorator
}

func NewMilkDecorator(c Coffee) *MilkDecorator {
    return &MilkDecorator{
        &CoffeeDecorator{coffee: c},
    }
}

func (m *MilkDecorator) Cost() float64 {
    return m.coffee.Cost() + 1.0
}

func (m *MilkDecorator) Description() string {
    return m.coffee.Description() + ", com leite"
}

// ConcreteDecorator: Adiciona caramelo ao café.
type CaramelDecorator struct {
    *CoffeeDecorator
}

func NewCaramelDecorator(c Coffee) *CaramelDecorator {
    return &CaramelDecorator{
        &CoffeeDecorator{coffee: c},
    }
}

func (c *CaramelDecorator) Cost() float64 {
    return c.coffee.Cost() + 1.5
}

func (c *CaramelDecorator) Description() string {
    return c.coffee.Description() + ", com caramelo"
}

func main() {
    // Pedido básico de café.
    coffee := &SimpleCoffee{}
    fmt.Println(coffee.Description(), "- Custo:", coffee.Cost())

    // Adicionando leite ao café.
    coffeeWithMilk := NewMilkDecorator(coffee)
    fmt.Println(coffeeWithMilk.Description(), "- Custo:", coffeeWithMilk.Cost())

    // Adicionando caramelo ao café com leite.
    coffeeWithMilkAndCaramel := NewCaramelDecorator(coffeeWithMilk)
    fmt.Println(coffeeWithMilkAndCaramel.Description(), "- Custo:", coffeeWithMilkAndCaramel.Cost())
}
```

### 🧠 O que acontece aqui?

- **Component**: A interface `Coffee` define os métodos `Cost()` e `Description()`, para calcular o custo e a descrição de um café.

- **ConcreteComponent**: `SimpleCoffee` é o café básico sem nenhum complemento.

- **Decorator**: `CoffeeDecorator` é a base para os decoradores, que aceita um `Coffee` como parâmetro e implementa os métodos `Cost()` e `Description()`.

- **ConcreteDecorator**: `MilkDecorator` e `CaramelDecorator` adicionam os respectivos complementos ao café.

Neste exemplo, você pode facilmente adicionar mais complementos ao café, apenas criando novos decoradores.

---

### 💡 Quando usar o padrão Decorator?
- ☕ Use o Decorator quando precisar adicionar responsabilidades ou funcionalidades a objetos de maneira flexível e em tempo de execução.

- 🧩 Quando quiser evitar subclasses para cada combinação possível de funcionalidades.

- 🚀 Ideal para criar uma estrutura de funcionalidades adicionais em que você pode compor diferentes combinações dinamicamente.

---

### ⚖️ Prós e Contras

| 👍 Vantagens | 👎 Desvantagens |
| ---- | ---- |
| ✅ Flexibilidade para adicionar funcionalidades em tempo de execução | ❌ Pode aumentar a complexidade do código |
| ✅ Evita a criação de várias subclasses | ❌ Pode dificultar a compreensão para novas pessoas no projeto |
| ✅ Cada decorador é reutilizável e independente | |

---

### 🎯 Conclusão

O **padrão Decorator** permite adicionar funcionalidades a objetos de maneira dinâmica e flexível, sem necessidade de modificar suas classes ou criar muitas subclasses. Em Go, a estrutura do Decorator é eficiente e permite a composição dinâmica de funcionalidades usando composição e interfaces.