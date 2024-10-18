# Abstract Factory em Go

### 📖 O que é o padrão Abstract Factory?
O **Abstract Factory** é um padrão de projeto **criacional** que fornece uma interface para criar famílias de objetos relacionados ou dependentes, sem especificar suas classes concretas. Ou seja, ele permite que você crie grupos de objetos que compartilham algum relacionamento (por exemplo, UI para diferentes sistemas operacionais) sem expor a lógica de criação concreta.

### 🤔 Qual problema o Abstract Factory resolve?
Quando você está trabalhando com **famílias de objetos** (como botões, janelas, e menus de diferentes sistemas operacionais), o Abstract Factory oferece uma maneira de criar objetos sem ter que se preocupar com as classes exatas que serão instanciadas. Ele agrupa as fábricas relacionadas para garantir que os objetos criados sejam compatíveis uns com os outros.

---

### 🚗 Analogia: Móveis para Estilos Diferentes 🛋️
Imagine que você está mobiliando uma casa. Se escolher um estilo **moderno**, todos os móveis (cadeira, sofá, mesa) devem seguir esse estilo. Se optar por um estilo vintage, os móveis também devem estar em harmonia. A **Abstract Factory** é como uma loja de móveis que oferece várias linhas de produtos, cada uma com estilos compatíveis.

---

### 🧱 Estrutura do Abstract Factory
Aqui está a estrutura básica de um Abstract Factory:

- **AbstractFactory**: Interface que declara métodos para criar objetos abstratos (familiares entre si).

- **ConcreteFactory**: Implementação concreta da fábrica que cria objetos específicos.

- **AbstractProduct**: Interface para produtos individuais (como cadeira, mesa).

- **ConcreteProduct**: Implementação concreta dos produtos.

- **Client**: Interage apenas com as interfaces da fábrica e dos produtos, sem conhecer os detalhes da implementação concreta.

---

### 👨‍💻 Como implementar o Abstract Factory em Go?
Vamos criar um exemplo em que temos uma fábrica que pode criar móveis modernos ou móveis vitorianos. Esses móveis incluem cadeiras e sofás.

#### Exemplo em Go:
```go
package main

import "fmt"

// AbstractProduct: Interface de Cadeira
type Chair interface {
	SitOn() string
}

// AbstractProduct: Interface de Sofá
type Sofa interface {
	LayOn() string
}

// ConcreteProduct: Cadeira Moderna
type ModernChair struct{}

func (mc *ModernChair) SitOn() string {
	return "Sentando em uma cadeira moderna!"
}

// ConcreteProduct: Sofá Moderno
type ModernSofa struct{}

func (ms *ModernSofa) LayOn() string {
	return "Deitando em um sofá moderno!"
}

// ConcreteProduct: Cadeira Vitoriana
type VictorianChair struct{}

func (vc *VictorianChair) SitOn() string {
	return "Sentando em uma cadeira vitoriana!"
}

// ConcreteProduct: Sofá Vitoriano
type VictorianSofa struct{}

func (vs *VictorianSofa) LayOn() string {
	return "Deitando em um sofá vitoriano!"
}

// AbstractFactory: Fábrica de Móveis
type FurnitureFactory interface {
	CreateChair() Chair
	CreateSofa() Sofa
}

// ConcreteFactory: Fábrica de Móveis Modernos
type ModernFurnitureFactory struct{}

func (mf *ModernFurnitureFactory) CreateChair() Chair {
	return &ModernChair{}
}

func (mf *ModernFurnitureFactory) CreateSofa() Sofa {
	return &ModernSofa{}
}

// ConcreteFactory: Fábrica de Móveis Vitorianos
type VictorianFurnitureFactory struct{}

func (vf *VictorianFurnitureFactory) CreateChair() Chair {
	return &VictorianChair{}
}

func (vf *VictorianFurnitureFactory) CreateSofa() Sofa {
	return &VictorianSofa{}
}

// Client: Usa a fábrica abstrata para criar produtos
func main() {
	var furnitureFactory FurnitureFactory

	// Criando móveis modernos
	furnitureFactory = &ModernFurnitureFactory{}
	modernChair := furnitureFactory.CreateChair()
	modernSofa := furnitureFactory.CreateSofa()

	fmt.Println(modernChair.SitOn())
	fmt.Println(modernSofa.LayOn())

	// Criando móveis vitorianos
	furnitureFactory = &VictorianFurnitureFactory{}
	victorianChair := furnitureFactory.CreateChair()
	victorianSofa := furnitureFactory.CreateSofa()

	fmt.Println(victorianChair.SitOn())
	fmt.Println(victorianSofa.LayOn())
}
```
### 🧠 O que acontece aqui?
- Temos duas famílias de produtos: **Modernos** e **Vitorianos**.

- O cliente pode solicitar móveis modernos ou vitorianos sem saber os detalhes de como eles são criados. Cada **fábrica concreta** (moderna ou vitoriana) cria um conjunto de objetos compatíveis entre si.

---

### 💡 Quando usar o Abstract Factory?
- Quando seu sistema precisa criar múltiplos objetos relacionados ou dependentes, e você quer garantir que eles sejam compatíveis entre si.

- Quando você deseja isolar a lógica de criação de famílias de objetos concretos.

- Quando deseja fornecer ao cliente uma interface para criar famílias de objetos sem expor a implementação concreta.

---

### ⚖️ Prós e Contras
|👍 Vantagens| 👎 Desvantagens |
|----|----|
|✅ Garante a compatibilidade entre objetos de uma família |❌ Pode aumentar a complexidade do código|
|✅ Desacopla a criação dos objetos concretos do código cliente | ❌ Exige mais código e classes adicionais|
|✅ Facilita a troca de famílias de produtos em tempo de execução | ❌ Adicionar novas famílias pode ser difícil |

---

### 🎨 Diferença entre Factory e Abstract Factory
Apesar de ambos os padrões estarem relacionados à criação de objetos, há uma diferença crucial:

| Padrão | Descrição | Quando Usar |
|----|----|----|
| Factory | Cria objetos de uma única categoria, geralmente relacionados a um único produto (ex: transportes, como carros ou barcos). | Quando você precisa de uma única etapa de criação de objetos simples. |
| Abstract Factory | Cria objetos de **múltiplas categorias** relacionados entre si, que podem formar uma família de produtos (ex: móveis modernos ou vitorianos). | Quando você precisa garantir a compatibilidade entre famílias de objetos e deseja criar múltiplos tipos de produtos. |

#### 🔑 Diferenças principais:
1. **Escopo da criação**:
    - O **Factory** cria objetos de uma única família ou categoria.

    - O **Abstract Factory** cria várias famílias de objetos relacionados entre si.

2. **Complexidade**:
    - O **Factory** é mais simples e geralmente lida com a criação de um tipo de objeto.

    - O **Abstract Factory** lida com a criação de grupos de objetos inter-relacionados e, portanto, é mais complexo.

---

### 🛠️ Exemplo mais complexo: Fábrica de interfaces gráficas
Vamos imaginar um exemplo mais real: uma fábrica de **interfaces gráficas** (GUI), onde temos fábricas que podem criar botões e _checkboxes_ para diferentes sistemas operacionais.

```go
package main

import "fmt"

// AbstractProduct: Botão
type Button interface {
	Render() string
}

// AbstractProduct: Checkbox
type Checkbox interface {
	Toggle() string
}

// ConcreteProduct: Botão para Windows
type WindowsButton struct{}

func (wb *WindowsButton) Render() string {
	return "Renderizando botão no estilo Windows"
}

// ConcreteProduct: Checkbox para Windows
type WindowsCheckbox struct{}

func (wc *WindowsCheckbox) Toggle() string {
	return "Alternando checkbox no estilo Windows"
}

// ConcreteProduct: Botão para MacOS
type MacOSButton struct{}

func (mb *MacOSButton) Render() string {
	return "Renderizando botão no estilo MacOS"
}

// ConcreteProduct: Checkbox para MacOS
type MacOSCheckbox struct{}

func (mc *MacOSCheckbox) Toggle() string {
	return "Alternando checkbox no estilo MacOS"
}

// AbstractFactory: Fábrica de GUI
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// ConcreteFactory: Fábrica de GUI para Windows
type WindowsGUIFactory struct{}

func (wg *WindowsGUIFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (wg *WindowsGUIFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

// ConcreteFactory: Fábrica de GUI para MacOS
type MacOSGUIFactory struct{}

func (mg *MacOSGUIFactory) CreateButton() Button {
	return &MacOSButton{}
}

func (mg *MacOSGUIFactory) CreateCheckbox() Checkbox {
	return &MacOSCheckbox{}
}

// Client: Usa a fábrica abstrata para criar produtos
func main() {
	var guiFactory GUIFactory

	// Criando GUI para Windows
	guiFactory = &WindowsGUIFactory{}
	winButton := guiFactory.CreateButton()
	winCheckbox := guiFactory.CreateCheckbox()

	fmt.Println(winButton.Render())
	fmt.Println(winCheckbox.Toggle())

	// Criando GUI para MacOS
	guiFactory = &MacOSGUIFactory{}
	macButton := guiFactory.CreateButton()
	macCheckbox := guiFactory.CreateCheckbox()

	fmt.Println(macButton.Render())
	fmt.Println(macCheckbox.Toggle())
}
```
Nesse exemplo, temos uma fábrica que pode criar botões e _checkboxes_, mas o estilo depende do sistema operacional (Windows ou MacOS). Usamos o **Abstract Factory** para garantir que o botão e o _checkbox_ do mesmo sistema operacional sejam compatíveis.

### 🎯 Conclusão
O padrão **Abstract Factory** é uma solução poderosa para criar **famílias de objetos inter-relacionados**, promovendo a flexibilidade e a extensibilidade em sistemas complexos. Ele é útil quando você precisa garantir que vários objetos de diferentes categorias funcionem bem em conjunto e quer desacoplar a lógica de criação do cliente.