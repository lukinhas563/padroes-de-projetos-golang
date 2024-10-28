# 📖 O que é o padrão Proxy?

O **Proxy** é um padrão de projeto **estrutural** que fornece um substituto ou um representante para outro objeto a fim de controlar o acesso a ele. Ele cria uma "cópia controlada" do objeto real, oferecendo uma camada intermediária entre o cliente e o objeto final. O Proxy pode ser usado para adicionar lógica extra antes ou depois de uma operação, como verificação de acesso, controle de cache, ou até mesmo gerenciamento de recursos de forma mais eficiente.

### 🤔 Qual problema o Proxy resolve?

O Proxy resolve o problema de controle de acesso e otimização de recursos em objetos que podem ser custosos de inicializar, acessar ou manipular diretamente. Com o Proxy, você pode gerenciar como e quando um objeto é criado, fornecendo mecanismos para evitar operações desnecessárias ou repetidas.

---

### 📦 Exemplo prático

Imagine um sistema de download de vídeos, onde a visualização de uma miniatura do vídeo só será carregada se o usuário clicar para assisti-la. Nesse caso, o Proxy pode ser usado para carregar as miniaturas de forma leve, inicializando o vídeo completo apenas quando realmente necessário.

---

### 🧱 Estrutura do Proxy

A estrutura do padrão Proxy inclui:

- **Subject**: Interface comum para o objeto real e o proxy.

- **RealSubject**: O objeto real que realiza as operações.

- **Proxy**: Classe intermediária que controla o acesso ao `RealSubject`, adicionando comportamentos adicionais, como controle de acesso ou cache.

- **Client**: Classe que usa o `Proxy` para acessar o `RealSubject`.

---

### 👨‍💻 Como implementar o Proxy em Go?

Vamos ver como o padrão Proxy pode ser implementado em Go com um exemplo de carregamento de imagens.

#### Exemplo em Go:

```go
package main

import "fmt"

// Subject: Interface comum para o RealImage e o ProxyImage
type Image interface {
    Display()
}

// RealSubject: RealImage, o objeto real que carrega e exibe a imagem
type RealImage struct {
    fileName string
}

// NewRealImage é o construtor do RealImage que carrega a imagem
func NewRealImage(fileName string) *RealImage {
    fmt.Printf("Carregando imagem %s...\n", fileName)
    return &RealImage{fileName: fileName}
}

// Display exibe a imagem
func (img *RealImage) Display() {
    fmt.Printf("Exibindo imagem %s\n", img.fileName)
}

// Proxy: ProxyImage, que controla o acesso ao RealImage
type ProxyImage struct {
    fileName   string
    realImage  *RealImage
}

// Display exibe a imagem usando o proxy, carregando-a apenas quando necessário
func (proxy *ProxyImage) Display() {
    if proxy.realImage == nil {
        proxy.realImage = NewRealImage(proxy.fileName)  // Carrega a imagem apenas se ainda não foi carregada
    }
    proxy.realImage.Display()
}

func main() {
    // Cliente usa o proxy para acessar a imagem
    image := &ProxyImage{fileName: "imagem_proxied.png"}
    
    // A imagem será carregada apenas na primeira exibição
    fmt.Println("Primeira chamada para Display():")
    image.Display()  // Carrega e exibe a imagem

    fmt.Println("Segunda chamada para Display():")
    image.Display()  // Exibe a imagem sem recarregar
}
```

### 🧠 O que acontece aqui?

- **Subject**: `Image` é a interface comum para `RealImage` e `ProxyImage`.

- **RealSubject**: `RealImage` é o objeto real que realiza o carregamento da imagem e exibe seu conteúdo.

- **Proxy**: `ProxyImage` controla o acesso ao `RealImage`, carregando a imagem apenas na primeira vez que `Display()` é chamado. Nas chamadas seguintes, `Display()` exibe a imagem sem precisar recarregá-la.

- **Client**: O cliente acessa o `ProxyImage` para exibir a imagem, o que evita o carregamento desnecessário da imagem nas exibições subsequentes.

Neste exemplo, o Proxy atua como um **carregamento atrasado** (ou _lazy loading_), carregando a imagem apenas quando necessário.

---

### 💡 Quando usar o padrão Proxy?

- 🚦 Use o Proxy quando precisar de controle de acesso a um objeto complexo.

- 📲 Ideal para operações de cache, carregamento atrasado, controle de recursos ou verificações de permissões.

- 💾 Útil quando o custo de criação ou execução de um objeto é alto e você quer otimizar o desempenho.

---

### ⚖️ Prós e Contras

| 👍 Vantagens | 👎 Desvantagens |
| ---- | ---- |
| ✅ Adiciona um controle extra sobre o acesso ao objeto | ❌ Pode aumentar a complexidade do sistema |
| ✅ Otimiza o uso de recursos | ❌ O desempenho pode ser impactado se houver muitas chamadas intermediárias |
| ✅ Permite adição de lógica adicional (ex.: cache ou verificação de permissões) | |

---

### 🎯 Conclusão

O **padrão Proxy** é uma solução eficaz para situações em que o controle de acesso ou a otimização de recursos é necessária antes de permitir o acesso a um objeto real. Em Go, o Proxy permite que o cliente interaja com o objeto sem saber se ele está acessando o objeto real ou um substituto controlado, tornando o sistema mais flexível e eficiente.