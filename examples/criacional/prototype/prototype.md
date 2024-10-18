# Prototype em Go
### 📖 O que é o padrão Prototype?
O **Prototype** é um padrão de projeto **criacional** que permite copiar objetos existentes sem depender de suas classes exatas. Em vez de criar um objeto do zero, você faz cópias (ou clones) de um objeto já existente.

### 🤔 Qual problema o Prototype resolve?
O Prototype é útil quando a criação de um objeto pode ser cara em termos de tempo ou recursos. Ao invés de instanciar um novo objeto repetidamente, o Prototype permite clonar objetos pré-existentes, facilitando a criação rápida de cópias.

### 📦 Exemplo prático
Imagine que você tem um objeto que requer várias operações dispendiosas para ser criado, como carregar dados de um banco de dados ou acessar uma API externa. O Prototype permite criar cópias desse objeto sem repetir todo o processo.

---

### 🚗 Analogia: criar personagens de videogame 🎮
Quando você cria um personagem em um jogo, ao invés de começar do zero toda vez, você pode usar um personagem modelo como base (com cabelo, roupas, habilidades) e só fazer pequenas mudanças. Isso é o Prototype em ação: ao invés de criar tudo de novo, você começa com um "protótipo" e o adapta conforme necessário.

---

### 🧱 Estrutura do Prototype
Aqui está como o padrão Prototype pode ser representado:

- **Prototype**: Interface que define o método de clonagem.
- **ConcretePrototype**: Implementa a clonagem do objeto.
- **Client**: Utiliza os objetos Prototype para criar cópias.

---

### 👨‍💻 Como implementar o Prototype em Go?
Vamos implementar o Prototype criando um exemplo de **formas geométricas**.

#### Exemplo em Go:
```go
package main

import "fmt"

// Prototype: Interface que define o método Clone
type Shape interface {
	Clone() Shape
	GetInfo() string
}

// ConcretePrototype: Círculo
type Circle struct {
	Radius int
	Color  string
}

func (c *Circle) Clone() Shape {
	return &Circle{
		Radius: c.Radius,
		Color:  c.Color,
	}
}

func (c *Circle) GetInfo() string {
	return fmt.Sprintf("Círculo com raio %d e cor %s", c.Radius, c.Color)
}

// ConcretePrototype: Quadrado
type Square struct {
	Side  int
	Color string
}

func (s *Square) Clone() Shape {
	return &Square{
		Side:  s.Side,
		Color: s.Color,
	}
}

func (s *Square) GetInfo() string {
	return fmt.Sprintf("Quadrado com lado %d e cor %s", s.Side, s.Color)
}

func main() {
	// Criando um Círculo e clonando ele
	originalCircle := &Circle{Radius: 5, Color: "vermelho"}
	cloneCircle := originalCircle.Clone()

	// Criando um Quadrado e clonando ele
	originalSquare := &Square{Side: 10, Color: "azul"}
	cloneSquare := originalSquare.Clone()

	fmt.Println("Original:", originalCircle.GetInfo())
	fmt.Println("Clone:", cloneCircle.GetInfo())

	fmt.Println("Original:", originalSquare.GetInfo())
	fmt.Println("Clone:", cloneSquare.GetInfo())
}
```

### 🧠 O que acontece aqui?
- Definimos a interface `Shape` com o método `Clone()` para permitir a clonagem de formas.

- As classes concretas `Circle` e `Square` implementam o método `Clone()` e o método `GetInfo()` para exibir informações sobre a forma.

- Quando clonamos os objetos, uma nova instância é criada com as mesmas propriedades.

---

### 💡 Quando usar o padrão Prototype?
- Quando a criação de objetos é **custosa** ou envolve processos **complexos**.
- Quando você deseja criar novas instâncias que sejam **cópias de objetos já existentes**.
- Quando você precisa garantir que um novo objeto seja uma duplicata exata, mas independente do original.

---

### ⚖️ Prós e Contras

|👍 Vantagens|👎 Desvantagens|
|----|----|
|✅ Reduz o custo de criação de novos objetos |❌ Pode ser difícil implementar cópias profundas|
|✅ Flexível para criar objetos similares rapidamente| ❌ Depende de uma interface adicional para clonagem|

---

### 🚶 Passos para criar Prototype (com personagens de videogame!)

Vamos aplicar o Prototype a um exemplo de criação de personagens de videogame:

```go
package main

import "fmt"

// Prototype: Interface que define o método Clone
type Character interface {
	Clone() Character
	GetStats() string
}

// ConcretePrototype: Guerreiro
type Warrior struct {
	Health int
	Armor  int
}

func (w *Warrior) Clone() Character {
	return &Warrior{
		Health: w.Health,
		Armor:  w.Armor,
	}
}

func (w *Warrior) GetStats() string {
	return fmt.Sprintf("Guerreiro com %d de saúde e %d de armadura", w.Health, w.Armor)
}

// ConcretePrototype: Mago
type Mage struct {
	Health int
	Mana   int
}

func (m *Mage) Clone() Character {
	return &Mage{
		Health: m.Health,
		Mana:   m.Mana,
	}
}

func (m *Mage) GetStats() string {
	return fmt.Sprintf("Mago com %d de saúde e %d de mana", m.Health, m.Mana)
}

func main() {
	// Criando um Guerreiro e clonando ele
	originalWarrior := &Warrior{Health: 100, Armor: 50}
	cloneWarrior := originalWarrior.Clone()

	// Criando um Mago e clonando ele
	originalMage := &Mage{Health: 80, Mana: 200}
	cloneMage := originalMage.Clone()

	fmt.Println("Original:", originalWarrior.GetStats())
	fmt.Println("Clone:", cloneWarrior.GetStats())

	fmt.Println("Original:", originalMage.GetStats())
	fmt.Println("Clone:", cloneMage.GetStats())
}
```
Neste exemplo:

- Criamos personagens (`Warrior` e `Mage`) com atributos distintos, e utilizamos o Prototype para clonar suas instâncias.
- O método `Clone()` faz uma cópia exata dos atributos.

---

### 🎯 Conclusão
O padrão **Prototype** é ideal para clonar objetos quando a criação do zero é cara ou complexa. Ele facilita a criação de cópias, mantendo flexibilidade e evitando a necessidade de recriar a lógica de inicialização dos objetos.