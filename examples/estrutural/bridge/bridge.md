# 📖 O que é o padrão Bridge?

O **Bridge** é um padrão de projeto **estrutural** que separa a abstração da implementação, permitindo que ambos variem independentemente. Ele permite que você altere tanto a interface de alto nível (a abstração) quanto as implementações (a parte concreta) de forma isolada, sem que uma dependa diretamente da outra.

### 🤔 Qual problema o Bridge resolve?

O padrão Bridge resolve o problema de criar classes que podem ter várias implementações e abstrações, sem que a alteração de uma afete a outra. Sem ele, a combinação de múltiplas abstrações e implementações poderia resultar em uma explosão de subclasses para cobrir todas as combinações possíveis.

---

### 📦 Exemplo prático

Imagine que você está desenvolvendo um sistema de envio de mensagens. Existem diferentes plataformas (e-mail, SMS, etc.) e diferentes tipos de mensagens (aviso, promoção). Se tentássemos criar uma classe para cada combinação de plataforma e tipo de mensagem, acabaríamos com muitas classes. O Bridge permite separar a abstração (o tipo de mensagem) da implementação (a plataforma de envio), criando um design flexível

---

### 🚗 Analogia: Controle Remoto 🖥️

Pense em um controle remoto universal que pode funcionar com qualquer marca de TV. O controle remoto é a abstração, e a TV é a implementação. Ao usar o padrão Bridge, você pode criar diferentes controles remotos e diferentes TVs, permitindo que o controle opere em qualquer TV, independentemente de sua marca.

---

### 🧱 Estrutura do Bridge

Aqui está a estrutura do padrão Bridge:

- **Abstraction**: Define a interface de alto nível que mantém uma referência para o implementador (Implementor).

- **RefinedAbstraction**: Extensão da Abstraction, com mais funcionalidades.

- **Implementor**: Interface para a implementação de baixo nível.

- **ConcreteImplementor**: Implementação concreta do Implementor.

---

### 👨‍💻 Como implementar o Bridge em Go?

Vamos ver um exemplo de como o padrão Bridge pode ser implementado em Go com um sistema de envio de mensagens.

#### Exemplo em Go:

```go
package main

import "fmt"

// Implementor define a interface para as plataformas de envio de mensagens.
type MessageSender interface {
    SendMessage(message string)
}

// ConcreteImplementor: Implementação concreta para envio por e-mail.
type EmailSender struct{}

func (e *EmailSender) SendMessage(message string) {
    fmt.Println("Enviando e-mail com a mensagem:", message)
}

// ConcreteImplementor: Implementação concreta para envio por SMS.
type SMSSender struct{}

func (s *SMSSender) SendMessage(message string) {
    fmt.Println("Enviando SMS com a mensagem:", message)
}

// Abstraction: Define a interface de mensagem, que usa o implementador.
type Message struct {
    sender MessageSender
}

// RefinedAbstraction: Mensagem de aviso.
type AlertMessage struct {
    *Message
}

func (a *AlertMessage) Send(text string) {
    fmt.Println("Mensagem de alerta:")
    a.sender.SendMessage(text)
}

// RefinedAbstraction: Mensagem promocional.
type PromotionMessage struct {
    *Message
}

func (p *PromotionMessage) Send(text string) {
    fmt.Println("Mensagem promocional:")
    p.sender.SendMessage(text)
}

func main() {
    // Criando os implementadores
    emailSender := &EmailSender{}
    smsSender := &SMSSender{}

    // Enviando uma mensagem de alerta por e-mail
    alertEmail := &AlertMessage{&Message{sender: emailSender}}
    alertEmail.Send("Aviso de manutenção programada!")

    // Enviando uma mensagem promocional por SMS
    promoSMS := &PromotionMessage{&Message{sender: smsSender}}
    promoSMS.Send("Desconto exclusivo de 50%!")
}
```

### 🧠 O que acontece aqui?

- **Abstraction**: `Message` mantém uma referência para `MessageSender`, a interface de envio.

- **Implementor**: `MessageSender` define a interface para as diferentes plataformas de envio.

- **ConcreteImplementors**: `EmailSender` e `SMSSender` são as implementações concretas para envio de mensagens.

- **RefinedAbstraction**: `AlertMessage` e `PromotionMessage` são especializações da abstração que adicionam funcionalidades específicas de mensagens.

Com essa implementação, você pode facilmente adicionar novas plataformas de envio (por exemplo, WhatsApp) ou novos tipos de mensagens (por exemplo, mensagens transacionais) sem modificar as classes existentes.

---

### 💡 Quando usar o padrão Bridge?
- 🏗️ Use o Bridge quando precisar separar a abstração da implementação para permitir que ambas evoluam de forma independente.

- 🚀 Quando tiver muitas variações de uma abstração e implementação que precisam ser combinadas de forma flexível.

- 🔄 Ideal quando você deseja reduzir a complexidade de herança, evitando a criação de muitas subclasses para cobrir todas as combinações possíveis.

---

### ⚖️ Prós e Contras

| 👍 Vantagens | 👎 Desvantagens |
|----|----|
| ✅ Separa a abstração da implementação, aumentando a flexibilidade | ❌ Pode aumentar a complexidade inicial do design |
| ✅ Facilita a extensão de ambas as partes sem modificar o código existente | ❌ Pode ser exagerado em sistemas simples |
|  ✅ Reduz a quantidade de subclasses necessárias | |

---

### 🎯 Conclusão

O **padrão Bridge** é uma ótima escolha para situações em que você tem múltiplas abstrações e implementações que precisam ser combinadas de maneira flexível. Ele desacopla essas duas partes, permitindo que ambas evoluam de forma independente e reduzindo a complexidade causada pela multiplicação de subclasses. Em Go, o Bridge é implementado de maneira eficiente usando interfaces e composição.