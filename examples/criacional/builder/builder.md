# Builder em Go
### 📖 O que é o padrão Builder?
O **Builder** é um padrão de projeto **criacional** que permite a você construir objetos complexos passo a passo. Ao invés de criar o objeto diretamente, o Builder oferece uma maneira de criar uma representação complexa de forma controlada e flexível.

### 🤔 Qual problema o Builder resolve?
O Builder é útil quando você precisa criar objetos **complexos** que possuem muitos parâmetros opcionais ou que precisam ser construídos em múltiplos passos. Ao invés de ter um construtor gigantesco com muitos parâmetros (muitos dos quais podem ser opcionais), o Builder divide a construção do objeto em passos lógicos.

### 📦 Exemplo prático
Imagine que você está construindo uma casa. Existem muitas partes diferentes (fundação, paredes, telhado, etc.) e você nem sempre precisa de todas as partes em uma casa básica. O Builder permite que você construa a casa passo a passo e decida o que incluir.

---

### 🚗 Analogia: um pedido de pizza 🍕
Pedir pizza é um ótimo exemplo de como o Builder funciona. Ao pedir pizza, você não simplesmente diz "quero uma pizza". Você escolhe a massa, o molho, os _toppings_, e outros detalhes, passo a passo.

Com o padrão Builder, você criaria um "builder" para pizzas, onde cada etapa é responsável por adicionar um componente da pizza.

---

### 🧱 Estrutura do Builder
Aqui está como o padrão Builder pode ser representado:

- **Builder**: Define os passos necessários para construir o produto. Esses passos são geralmente métodos separados.

- **ConcreteBuilder**: Implementa as etapas necessárias da construção.

- **Produto**: O objeto final que está sendo construído.

- **Director (opcional)**: Controla o processo de construção, chamando os métodos do Builder em uma ordem específica.

### 👨‍💻 Como implementar o Builder em Go?
Vamos ver como o padrão Builder pode ser implementado em Go. Vamos criar um exemplo de construção de um carro.

#### Exemplo em Go:
```go
package main

import "fmt"

// Produto final: Carro
type Car struct {
	Engine   string
	Wheels   int
	Color    string
}

// Builder para Carro
type CarBuilder struct {
	car Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (b *CarBuilder) SetEngine(engine string) *CarBuilder {
	b.car.Engine = engine
	return b
}

func (b *CarBuilder) SetWheels(wheels int) *CarBuilder {
	b.car.Wheels = wheels
	return b
}

func (b *CarBuilder) SetColor(color string) *CarBuilder {
	b.car.Color = color
	return b
}

func (b *CarBuilder) Build() Car {
	return b.car
}

func main() {
	builder := NewCarBuilder()

	car := builder.SetEngine("V8").
		SetWheels(4).
		SetColor("Vermelho").
		Build()

	fmt.Printf("Carro criado: %+v\n", car)
}
```

### 🧠 O que acontece aqui?
- Criamos um `CarBuilder` que possui métodos para definir diferentes partes do carro (motor, rodas, cor).

- Cada método retorna o próprio builder, o que permite o encadeamento de chamadas (método conhecido como _method chaining_).

- O método `Build()` é chamado no final para construir e retornar o carro completo.

Esse exemplo é útil para criar objetos que possuem muitos detalhes e propriedades opcionais

---

### 💡 Quando usar o padrão Builder?
- 🏗️ Use o Builder quando precisar construir objetos complexos que possuem muitas configurações diferentes.

- 🌟 Ele é especialmente útil quando você tem muitos parâmetros opcionais ou quando a ordem de construção é importante.

---

### ⚖️ Prós e Contras

|👍 Vantagens                                        |👎 Desvantagens                           |
|----------------------------------------------------|------------------------------------------|
|✅ Facilita a criação de objetos complexos          |❌ Adiciona mais classes ao código        |
|✅ Ajuda a evitar construtores com muitos parâmetros|❌ Pode ser excessivo para objetos simples|
|✅ Permite a criação de produtos em etapas          |                                          |

---

### 🚶 Passos do Builder (com pizza!)
Vamos visualizar os passos de um Builder com um exemplo divertido: a construção de uma pizza!

```go
package main

import "fmt"

// Produto final: Pizza
type Pizza struct {
	Dough  string
	Sauce  string
	Topping string
}

// Builder para Pizza
type PizzaBuilder struct {
	pizza Pizza
}

func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{}
}

func (b *PizzaBuilder) SetDough(dough string) *PizzaBuilder {
	b.pizza.Dough = dough
	return b
}

func (b *PizzaBuilder) SetSauce(sauce string) *PizzaBuilder {
	b.pizza.Sauce = sauce
	return b
}

func (b *PizzaBuilder) SetTopping(topping string) *PizzaBuilder {
	b.pizza.Topping = topping
	return b
}

func (b *PizzaBuilder) Build() Pizza {
	return b.pizza
}

func main() {
	builder := NewPizzaBuilder()

	pizza := builder.SetDough("massa fina").
		SetSauce("molho de tomate").
		SetTopping("queijo").
		Build()

	fmt.Printf("Pizza criada: %+v\n", pizza)
}
```

Agora, você pode montar a pizza passo a passo, escolhendo cada ingrediente. Ao final, o Builder monta a pizza para você!

---

### 🎯 Conclusão
O padrão Builder é ideal para construir objetos complexos com muitos detalhes opcionais, ajudando a manter o código limpo e fácil de entender. Ele é especialmente útil quando há várias combinações de valores possíveis ou quando a construção do objeto precisa seguir uma ordem específica.