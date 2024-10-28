# 📖 O que é o padrão Façade?

O **Façade** é um padrão de projeto **estrutural** que fornece uma interface simplificada para um conjunto de classes, bibliotecas ou funcionalidades complexas. Ele age como uma "fachada" que oculta a complexidade do sistema, oferecendo uma interface unificada que facilita o uso para o cliente.

### 🤔 Qual problema o Façade resolve?

O Façade resolve o problema de simplificar o uso de sistemas complexos, onde várias classes e interações são necessárias para realizar uma tarefa. Ao fornecer uma interface mais simples, o padrão Façade esconde a complexidade do cliente e evita o acoplamento direto do cliente a várias classes internas do sistema.

---

### 📦 Exemplo prático

Imagine um sistema de home theater onde você precisa ligar o projetor, ajustar o som, escurecer as luzes, e preparar o filme para começar. Sem o Façade, o usuário teria que executar cada passo manualmente. Com o Façade, você cria uma interface simplificada, como `HomeTheaterFacade`, que organiza e executa todos esses passos para o usuário com apenas uma chamada de método.

---

### 🚗 Analogia: Controle Remoto da TV 📺

O Façade é como o controle remoto de uma TV. Internamente, a TV possui vários componentes complexos (circuitos, tela, som, entradas, etc.), mas o controle remoto oferece uma interface simplificada para o usuário, que pode ligar e desligar, ajustar o volume ou mudar o canal sem interagir com cada componente individualmente.

---

### 🧱 Estrutura do Façade

A estrutura do padrão Façade inclui:

- **Subsystem Classes**: Classes internas que possuem funcionalidades específicas do sistema e executam tarefas complexas.

- **Façade**: A classe que fornece uma interface simplificada para o cliente, invocando as operações necessárias nas classes internas.

- **Client**: A classe que interage com a fachada em vez de interagir diretamente com as classes internas do sistema

---

### 👨‍💻 Como implementar o Façade em Go?

Vamos ver como o padrão Façade pode ser implementado em Go com um exemplo de um sistema de Home Theater.

#### Exemplo em Go:

```go
package main

import "fmt"

// Subsistema 1: Sistema de som
type SoundSystem struct{}

func (s *SoundSystem) On() {
    fmt.Println("Sistema de som ligado.")
}

func (s *SoundSystem) SetVolume(level int) {
    fmt.Printf("Volume ajustado para %d.\n", level)
}

// Subsistema 2: Projetor
type Projector struct{}

func (p *Projector) On() {
    fmt.Println("Projetor ligado.")
}

func (p *Projector) WideScreenMode() {
    fmt.Println("Modo widescreen ativado no projetor.")
}

// Subsistema 3: Luzes
type Lights struct{}

func (l *Lights) Dim(level int) {
    fmt.Printf("Luzes escurecidas para %d%%.\n", level)
}

// Façade: HomeTheaterFacade
type HomeTheaterFacade struct {
    soundSystem *SoundSystem
    projector   *Projector
    lights      *Lights
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
    return &HomeTheaterFacade{
        soundSystem: &SoundSystem{},
        projector:   &Projector{},
        lights:      &Lights{},
    }
}

func (h *HomeTheaterFacade) WatchMovie() {
    fmt.Println("Preparando para assistir ao filme...")
    h.lights.Dim(10)
    h.projector.On()
    h.projector.WideScreenMode()
    h.soundSystem.On()
    h.soundSystem.SetVolume(5)
    fmt.Println("Tudo pronto! Aproveite o filme.")
}

func main() {
    // Usando o Façade para iniciar o sistema de home theater
    homeTheater := NewHomeTheaterFacade()
    homeTheater.WatchMovie()
}
```

### 🧠 O que acontece aqui?

- **Subsystem** Classes: `SoundSystem`, `Projector` e `Lights` são os subsistemas com operações específicas, como ligar o som, ajustar o volume, ligar o projetor, etc.

- **Façade**: `HomeTheaterFacade` encapsula a interação com os subsistemas e fornece o método `WatchMovie()` para o cliente, simplificando o processo.

- **Client**: No `main()`, o cliente usa a fachada para iniciar o sistema de home theater de forma simplificada, sem precisar conhecer os detalhes de cada subsistema.

Neste exemplo, o Façade permite que o cliente apenas chame `WatchMovie()` para preparar o sistema de home theater de forma eficiente.

---

### 💡 Quando usar o padrão Façade?

- 📲 Use o Façade quando precisar simplificar uma biblioteca ou um sistema complexo para o cliente.

- 🏛️ Ideal para sistemas onde múltiplas interações com várias classes são comuns, mas o cliente precisa de uma interface única e fácil de usar.

- 🚀 Quando quiser desacoplar o cliente de subsistemas específicos, permitindo mudanças internas sem impacto direto no cliente.

---

### ⚖️ Prós e Contras

| 👍 Vantagens | 👎 Desvantagens |
| ---- | ---- |
| ✅ Simplifica a interação do cliente com o sistema | ❌ Pode introduzir mais uma camada de abstração |
| ✅ Reduz o acoplamento entre cliente e subsistemas | ❌ Nem sempre esconde toda a complexidade interna |
| ✅ Flexível para mudanças nos subsistemas sem impactar o cliente | |

---

### 🎯 Conclusão

O **padrão Façade** oferece uma interface simplificada para sistemas complexos, melhorando a usabilidade e reduzindo o acoplamento do cliente com os subsistemas. Ele é ideal para sistemas onde uma interface unificada pode representar múltiplas operações internas. Em Go, o Façade permite organizar as interações entre subsistemas, tornando o código mais fácil de manter e expandir.