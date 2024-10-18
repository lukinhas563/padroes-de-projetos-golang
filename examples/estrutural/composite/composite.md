# 📖 O que é o padrão Composite?

O **Composite** é um padrão de projeto **estrutural** que permite compor objetos em estruturas de árvore para representar hierarquias. Ele trata objetos individuais e composições de objetos de maneira uniforme, permitindo que clientes manipulem objetos simples e compostos sem precisar diferenciá-los.

### 🤔 Qual problema o Composite resolve?
O Composite resolve o problema de manipular coleções de objetos, como se fossem objetos únicos, sem precisar lidar com a complexidade de distinguir entre objetos individuais e agregados de objetos. Ele permite estruturar os objetos em forma de árvore, criando uma hierarquia onde cada "nó" pode ser um objeto simples ou composto.

### 📦 Exemplo prático
Imagine um sistema de gerenciamento de arquivos, onde uma pasta pode conter tanto arquivos quanto outras pastas. O padrão Composite permite que você trate arquivos e pastas de forma uniforme, realizando operações como listar o conteúdo de uma pasta ou subpasta sem precisar se preocupar se é um arquivo ou outra pasta.

---

### 🚗 Analogia: Estrutura de Pastas 📂
Imagine uma estrutura de pastas e arquivos em seu computador. Uma pasta pode conter arquivos simples ou outras pastas, e você quer realizar operações sobre elas de forma consistente. O padrão Composite trata arquivos e pastas da mesma maneira, permitindo que você os manipule sem se preocupar com o que são individualmente.

---

### 🧱 Estrutura do Composite
Aqui está como o padrão Composite pode ser representado:

- **Component**: Define a interface que será implementada por todos os elementos, sejam eles simples ou compostos.

- **Leaf**: Representa um objeto simples que não contém outros objetos.

- **Composite**: Representa um objeto composto que pode conter outros componentes.

- **Client**: Manipula objetos através da interface comum `Component`, sem precisar saber se é um objeto simples ou composto.

---

### 👨‍💻 Como implementar o Composite em Go?
Vamos ver como o padrão Composite pode ser implementado em Go com um exemplo de gerenciamento de arquivos.

#### Exemplo em Go:
```go
package main

import "fmt"

// Component define a interface comum para folhas e nós compostos.
type Component interface {
    Display(indent string)
}

// File é uma "folha" da árvore de composição, ou seja, um objeto simples.
type File struct {
    name string
}

// Display exibe o nome do arquivo.
func (f *File) Display(indent string) {
    fmt.Println(indent + f.name)
}

// Directory é um "composto" que pode conter outros componentes (arquivos ou diretórios).
type Directory struct {
    name       string
    components []Component
}

// Add adiciona um componente (arquivo ou diretório) ao diretório.
func (d *Directory) Add(component Component) {
    d.components = append(d.components, component)
}

// Display exibe o diretório e todos os seus componentes, aplicando recursão para exibir subdiretórios e arquivos.
func (d *Directory) Display(indent string) {
    fmt.Println(indent + d.name + "/")
    for _, component := range d.components {
        component.Display(indent + "  ")
    }
}

func main() {
    // Criando arquivos individuais (folhas)
    file1 := &File{name: "file1.txt"}
    file2 := &File{name: "file2.txt"}
    file3 := &File{name: "file3.txt"}

    // Criando diretórios compostos
    folder1 := &Directory{name: "folder1"}
    folder2 := &Directory{name: "folder2"}
    folder3 := &Directory{name: "folder3"}

    // Adicionando arquivos e pastas ao folder1
    folder1.Add(file1)
    folder1.Add(folder2)

    // Adicionando arquivos ao folder2
    folder2.Add(file2)
    folder2.Add(folder3)

    // Adicionando um arquivo ao folder3
    folder3.Add(file3)

    // Exibindo a estrutura de arquivos e pastas
    folder1.Display("")
}
```

### 🧠 O que acontece aqui?
- Temos uma interface `Component` que define o método `Display()`, usado para exibir o nome dos arquivos ou diretórios.

- `File` é a folha da estrutura, representando arquivos individuais.

- `Directory` é o objeto composto, que pode conter outros componentes (sejam eles arquivos ou outros diretórios).

- Usamos o método `Add()` no `Directory` para adicionar componentes (arquivos ou diretórios) ao diretório.

- O método `Display()` é chamado recursivamente para exibir toda a estrutura de pastas e arquivos de forma organizada.

---

### 💡 Quando usar o padrão Composite?
- 🏢 Use o Composite quando precisar representar estruturas hierárquicas de objetos.

- 📂 Quando quiser tratar objetos simples e compostos de maneira uniforme.

- 🌲 Ideal para sistemas que precisam manipular árvores de objetos, como estruturas de pastas, menus ou gráficos.

---

### ⚖️ Prós e Contras

| 👍 Vantagens | 👎 Desvantagens |
| ---- | ---- |
| ✅ Facilita a manipulação de hierarquias complexas | ❌ Pode aumentar a complexidade do código |
| ✅ Trate objetos simples e compostos da mesma forma | ❌ Pode ser ineficiente para certas operações |
| ✅ Extensível: fácil de adicionar novos tipos de componentes | |

---

### 🎯 Conclusão
O **padrão Composite** é ideal para trabalhar com **hierarquias** de objetos, permitindo tratar objetos individuais e composições de forma uniforme. Sua implementação em Go é bastante flexível, e o padrão é particularmente útil em sistemas que manipulam árvores de objetos, como sistemas de arquivos, menus ou estruturas organizacionais.