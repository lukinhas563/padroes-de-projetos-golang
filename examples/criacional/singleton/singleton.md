# Singleton em Go
### 📖 O que é o padrão Singleton?
O **Singleton** é um padrão de projeto **criacional** que garante que uma classe tenha apenas uma instância durante a execução do programa, oferecendo um ponto de acesso global a essa instância.

### 🤔 Por que usar o Singleton?
O Singleton resolve dois problemas principais:

1. **Instância única**: Em alguns cenários, é essencial garantir que apenas uma instância de uma classe seja criada, como ao acessar um recurso compartilhado (banco de dados, por exemplo).

2. **Acesso global**: Em vez de usar variáveis globais (que podem ser sobrescritas facilmente), o Singleton oferece um ponto de acesso seguro.

#### Um exemplo prático 📂
Imagine que sua aplicação precisa de um logger que registra eventos. Você não quer várias instâncias desse logger, já que isso poderia criar confusão nos logs. O Singleton garante que só exista um logger funcionando!

---

### 🚗 Analogia: o governo como Singleton
Pense no governo de um país. Não importa quem está no poder, sempre há um governo oficial. Ele é um ponto de acesso global e único para o controle do país. No código, o Singleton funciona da mesma maneira: uma única instância global!

---

### 👨‍💻 Como implementar o Singleton em Go?
Em Go, podemos implementar o padrão Singleton com o pacote `sync.Once` para garantir que nossa instância seja criada apenas uma vez, mesmo em um ambiente concorrente.

Aqui está o exemplo:

```golang
package singleton

import "sync"

type singleton struct{}

var singleInstance *singleton
var once sync.Once

func GetInstance() *singleton {
    once.Do(func() {
        singleInstance = &singleton{}
    })
    
    return singleInstance
}
```

### 🧠 Por que usar `sync.Once`?
Go permite que o código rode de forma **concorrente** através de **goroutines**. Se várias goroutines tentarem acessar o Singleton simultaneamente, sem controle de concorrência, poderiam ser criadas múltiplas instâncias da `struct`. O `sync.Once` resolve isso garantindo que o Singleton seja instanciado apenas uma vez, mesmo com múltiplas goroutines.

---

### 💡 Quando usar o Singleton?
- 🐞 Quando uma classe precisa ter apenas uma única instância durante todo o ciclo de vida da aplicação.

- 💫 Quando você percebe que está usando variáveis globais para guardar estados importantes (como variáveis de configuração) e quer uma solução mais segura.

---

### ⚖️ Vantagens e desvantagens

|👍 Vantagens                               |👎 Desvantagens                                                |
| ------------------------------------------| ---------------------------------------------------------------|
|✅ Instância única garantida               |❌ Viola o princípio de responsabilidade única                 |
|✅ Acesso global à instância               |❌ Pode ser difícil de testar                                  |
|✅ Inicialização sob demanda (lazy loading)|❌ Necessita de tratamento especial em ambiente multithread    |

---

### 📺 Quer saber mais?
Veja o vídeo do **HunCoding** sobre como usar o `sync.Once` em Go:

[📺 Assista no YouTube](https://www.youtube.com/watch?v=qtXD_R4d1rY&ab_channel=HunCoding)

Ou leia mais sobre o pacote `sync` na [documentação oficial do Go](https://pkg.go.dev/sync).

### 🎯 Conclusão
O Singleton é um padrão simples, mas muito útil para casos em que você precisa garantir uma única instância de um recurso em toda a aplicação. Embora tenha suas limitações, como dificuldade de teste e violação do princípio de responsabilidade única, ele resolve problemas comuns de maneira eficiente.