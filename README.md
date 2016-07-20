# Poupaniquel
> A financial software in just one binary for local or cloud usage

## What is it?

Poupaniquel is a Brazilian Portuguese name. In English means something like "save coins".

## Why?

I wish a simple software to help me to manage my finances but I don't want to pay for it. Also, I'd like an API to use with my shell scripts and I need to run the software on my Macbook or in my private cloud with no complicated configurations.
Poupaniquel should answer the questions:

1. How much money I spent?
2. How much money I have today?
3. How much money I should have in the future?

## Building

```
curl https://glide.sh/get | sh
go get github.com/gustavohenrique/poupaniquel
cd $GOPATH/src/github.com/gustavohenrique/poupaniquel
glide install
go build -o poupaniquel main.go
```

## How to contribute?

First, you need to open an issue to talk about your proposed. After I agree that your idea is great for the project, you should fork this repository and send me a Pull Request.

## License

MIT
