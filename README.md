# modhello

Um exemplo simples da utilização de modules em golang na versão 1.11beta2

#[Estrutura]

```go
   modhello
      modlib/lib
              go.mod
              modlib.go
      modapp
          go.mod
          main.go
```
#[go run]
```go

   $ cd modhello/modapp
   $ go run main.go
   $ 2018/07/27 07:50:35 start
   $ 2018/07/27 07:50:35 Sum(1,2) = 3

```

Mais Detalhes pode ser visualizado aqui *[golang modules](https://github.com/golang/go/wiki/Modules)
