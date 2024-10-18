# Factory em Go
### 📖 O que é o padrão Factory?
O **Factory** é um padrão de projeto **criacional** que fornece uma interface para criar objetos em uma superclasse, permitindo que as subclasses alterem o tipo de objeto que será instanciado. Em vez de instanciar objetos diretamente, o Factory delega essa responsabilidade para classes específicas, escondendo os detalhes de criação dos objetos.

### 🤔 Qual problema o Factory resolve?
Imagine que você está desenvolvendo uma aplicação que precisa criar instâncias de várias classes relacionadas, mas não quer lidar diretamente com os detalhes de construção de cada uma delas. O Factory permite que você delegue essa tarefa a uma classe especializada, simplificando a criação de objetos.

---

### 🚗 Analogia: Pizzaria 🍕
Pense em uma pizzaria. Quando você pede uma pizza, você não precisa saber como cada ingrediente é preparado, nem como a pizza é montada e assada. Você só faz o pedido, e a pizzaria entrega a pizza pronta. O **Factory** é como a pizzaria, ele gerencia os detalhes da criação para você.

---

### 🧱 Estrutura do Factory
Aqui está a estrutura básica de um Factory:

- **Product**: Interface que define o tipo de objeto que será criado.
- **ConcreteProduct**: Classes concretas que implementam o Product.
- **Creator**: Classe Factory que declara o método CreateProduct.
- **ConcreteCreator**: Implementação concreta do Creator que instanciará os ConcreteProduct.

---

### 👨‍💻 Como implementar o Factory em Go?
Vamos usar o exemplo de criação de **transportes** como carros e barcos.

#### Exemplo em Go:
```go
package main

import "fmt"

// Product: Interface de Transporte
type Transport interface {
	Deliver() string
}

// ConcreteProduct: Carro
type Car struct{}

func (c *Car) Deliver() string {
	return "Entrega por estrada com um carro!"
}

// ConcreteProduct: Barco
type Ship struct{}

func (s *Ship) Deliver() string {
	return "Entrega por mar com um barco!"
}

// Creator: Interface da Factory
type TransportFactory interface {
	CreateTransport() Transport
}

// ConcreteCreator: Factory para criar um carro
type CarFactory struct{}

func (cf *CarFactory) CreateTransport() Transport {
	return &Car{}
}

// ConcreteCreator: Factory para criar um barco
type ShipFactory struct{}

func (sf *ShipFactory) CreateTransport() Transport {
	return &Ship{}
}

func main() {
	// Criando uma fábrica de carros
	var transportFactory TransportFactory
	transportFactory = &CarFactory{}
	car := transportFactory.CreateTransport()
	fmt.Println(car.Deliver())

	// Criando uma fábrica de barcos
	transportFactory = &ShipFactory{}
	ship := transportFactory.CreateTransport()
	fmt.Println(ship.Deliver())
}
```

### 🧠 O que acontece aqui?
- Definimos a interface `Transport` que tem o método `Deliver()`.
- As classes concretas `Car` e `Ship` implementam essa interface.
- Criamos as fábricas `CarFactory` e `ShipFactory`, que sabem como instanciar `Car` e `Ship`, respectivamente.
- O cliente pode pedir para a fábrica criar um transporte, sem se preocupar com os detalhes de construção.

---

### 💡 Quando usar o padrão Factory?
- Quando você precisa criar objetos sem especificar a classe exata deles no código do cliente.
- Quando o processo de criação envolve lógica complexa, que seria difícil de gerenciar diretamente em vários lugares.
- Quando você precisa de maior flexibilidade na criação de objetos, permitindo que sejam trocados em tempo de execução.

---

### ⚖️ Prós e Contras
|👍 Vantagens|👎 Desvantagens|
|----|----|
|✅ Permite maior flexibilidade na criação de objetos| ❌ Pode resultar em muitas classes e complexidade|
|✅ Desacopla a lógica de criação do cliente| ❌ Pode esconder dependências excessivas no sistema|

---

### 🎨 Diferença entre Factory e Builder
Os padrões **Factory** e **Builder** têm objetivos semelhantes — criar objetos —, mas eles diferem em como fazem isso e quando devem ser usados.

|Padrão|Descrição|Quando Usar|
|----|----|----|
|Factory|Cria objetos de uma família de classes. O cliente apenas solicita um objeto, e a factory sabe como criá-lo.|Cria objetos de uma família de classes. O cliente apenas solicita um objeto, e a factory sabe como criá-lo.|
|Builder|Constrói objetos complexos passo a passo. O cliente controla o processo de construção, escolhendo componentes.|Constrói objetos complexos passo a passo. O cliente controla o processo de construção, escolhendo componentes.|

#### 🔑 Diferenças principais:
1. Complexidade do objeto:
    - O **Factory** é útil quando a criação é simples e direta (ex: criar um carro ou um barco).

    - O **Builder** é útil quando o objeto tem muitas partes ou precisa ser configurado de forma personalizada (ex: construir uma casa com várias opções).

2. Processo de criação:
    - O **Factory** cria o objeto de uma só vez, sem precisar de intervenções.

    - O **Builder** permite criar o objeto por etapas, passo a passo, permitindo customizações durante a construção.

---

### 🚶 Passos para criar uma Factory (exemplo com computadores)
Agora, vamos criar uma fábrica de computadores onde o cliente pode pedir um computador **PC** ou **Server**:

```go
package main

import "fmt"

// Product: Interface de Computador
type Computer interface {
	GetDetails() string
}

// ConcreteProduct: PC
type PC struct{}

func (p *PC) GetDetails() string {
	return "PC: 16GB RAM, 512GB SSD, CPU i7"
}

// ConcreteProduct: Servidor
type Server struct{}

func (s *Server) GetDetails() string {
	return "Server: 64GB RAM, 2TB SSD, CPU Xeon"
}

// Creator: Interface de Factory de Computadores
type ComputerFactory interface {
	CreateComputer() Computer
}

// ConcreteCreator: Factory para criar PCs
type PCFactory struct{}

func (pcf *PCFactory) CreateComputer() Computer {
	return &PC{}
}

// ConcreteCreator: Factory para criar Servidores
type ServerFactory struct{}

func (sf *ServerFactory) CreateComputer() Computer {
	return &Server{}
}

func main() {
	// Criando uma fábrica de PCs
	var factory ComputerFactory
	factory = &PCFactory{}
	pc := factory.CreateComputer()
	fmt.Println(pc.GetDetails())

	// Criando uma fábrica de Servidores
	factory = &ServerFactory{}
	server := factory.CreateComputer()
	fmt.Println(server.GetDetails())
}
```
Aqui, o cliente pode criar um **PC** ou um **Servidor** sem saber como eles são construídos internamente.

---

### 🎯 Conclusão
O padrão **Factory** é ideal para criar objetos sem expor ao cliente os detalhes de construção, garantindo maior flexibilidade e facilitando a manutenção do código. Ele é especialmente útil quando o sistema precisa gerar objetos de várias classes relacionadas.

**Diferença principal**: o **Factory** cria objetos simples de uma única vez, enquanto o Builder constrói objetos complexos por etapas.