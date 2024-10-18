# 📖 O que é o padrão Adapter?

O **Adapter** é um padrão de projeto **estrutural** que permite que classes com interfaces incompatíveis trabalhem juntas. Ele faz isso convertendo a interface de uma classe em outra interface esperada pelos clientes. Em resumo, o Adapter atua como uma "ponte" entre duas interfaces que, de outra forma, seriam incompatíveis.

### 🤔 Qual problema o Adapter resolve?

Imagine que você tem um sistema onde novos componentes precisam ser integrados, mas eles não seguem o mesmo padrão de interface usado pelo sistema atual. O Adapter resolve o problema de incompatibilidade entre diferentes interfaces, sem alterar o código original das classes existentes.

---

### 📦 Exemplo prático

Suponha que você está desenvolvendo um sistema de pagamento. Seu sistema atualmente utiliza uma interface de pagamentos antiga, mas você precisa integrar uma nova API de terceiros. As interfaces não são compatíveis, e refatorar todo o sistema seria muito trabalhoso. O Adapter permite fazer essa integração sem a necessidade de modificar todo o código existente.

---

### 🚗 Analogia: Tomadas Elétricas ⚡

Pense em um adaptador de tomadas: você tem um aparelho com um plugue de dois pinos, mas a tomada no local tem três pinos. O adaptador converte o plugue de dois pinos para três, permitindo que seu aparelho funcione corretamente. O padrão Adapter faz o mesmo com interfaces no código.

---

### 🧱 Estrutura do Adapter

Aqui está a estrutura do padrão Adapter:

- **Target**: Define a interface que o cliente espera usar.

- **Adaptee**: Classe existente que precisa ser adaptada.

- **Adapter**: Implementa a interface Target e traduz as chamadas para o Adaptee.

- **Client**: Interage com a interface Target e não sabe que está lidando com um Adapter.

---

### 👨‍💻 Como implementar o Adapter em Go?

Vamos implementar o padrão Adapter em Go com um exemplo de integração de sistemas de pagamento.

#### Exemplo em Go:
```go
package main

import "fmt"

// Target define a interface esperada pelo cliente.
type PaymentProcessor interface {
    Pay(amount float64)
}

// Adaptee é a nova API que precisa ser adaptada.
type NewPaymentAPI struct{}

func (api *NewPaymentAPI) MakePayment(amountInCents int) {
    fmt.Printf("Payment of %d cents processed by new API\n", amountInCents)
}

// Adapter adapta a interface da nova API para a interface esperada.
type PaymentAdapter struct {
    api *NewPaymentAPI
}

// Pay é a implementação do método esperado, adaptando para a nova API.
func (adapter *PaymentAdapter) Pay(amount float64) {
    amountInCents := int(amount * 100) // Converter para centavos
    adapter.api.MakePayment(amountInCents)
}

func main() {
    // Cliente usa a interface antiga, mas queremos usar a nova API.
    newAPI := &NewPaymentAPI{}
    adapter := &PaymentAdapter{api: newAPI}

    // Cliente chama o método na interface esperada, que é adaptada.
    adapter.Pay(23.75) // Pagamento de 23.75 reais (ou 2375 centavos)
}
```

### 🧠 O que acontece aqui?
- **Target**: A interface `PaymentProcessor` define o método `Pay()` que o cliente espera usar.

- **Adaptee**: A nova API `NewPaymentAPI` usa um método diferente (`MakePayment()`), que espera o valor em centavos.

- **Adapter**: `PaymentAdapter` converte o valor de reais para centavos e chama o método da `NewPaymentAPI`.

- **Client**: O cliente faz a chamada através da interface `PaymentProcessor`, mas internamente o Adapter ajusta para usar a nova API.

---

### 💡 Quando usar o padrão Adapter?
- 🔌 Use o Adapter quando você quiser integrar uma classe existente com uma nova interface sem modificar o código original.

- 🏗️ Quando precisar reutilizar classes ou APIs externas que possuem interfaces incompatíveis com o sistema atual.

- 🔄 Útil quando você deseja aplicar um novo comportamento ou funcionalidade a uma interface antiga sem quebrar o código.

---

### ⚖️ Prós e Contras
| 👍 Vantagens | 👎 Desvantagens |
| ---- | ---- |
| ✅ Facilita a integração de novas APIs sem modificar o código existente | ❌ Pode adicionar complexidade extra ao código |
| ✅ Promove a reutilização de código e flexibilidade | ❌ Pode resultar em múltiplos Adapters para diferentes APIs |
| ✅ Permite manter o código legado intacto | |

---

### 🎯 Conclusão
O **padrão Adapter** é ideal para situações em que você precisa integrar novos componentes, APIs ou bibliotecas ao seu sistema sem modificar as classes existentes. Ele converte a interface de um sistema em uma interface que o cliente espera, garantindo compatibilidade entre interfaces diferentes. Em Go, isso pode ser implementado de maneira simples e eficiente usando structs e interfaces.