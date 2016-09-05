# cpf
[![Build status](https://secure.travis-ci.org/nextbit/cpf.png)](https://secure.travis-ci.org/nextbit/cpf)

Brazilian CPF validator

## Instalation

```
go get github.com/nextbit/cpf
```

## Usage

Call the valid func passing the cpf string

	cpf.Valid("886.100.254-46") => true
	cpf.Valid("88610025446") => true

	cpf.Valid("123.123.123.12") => false
	cpf.Valid("12312312") => false

