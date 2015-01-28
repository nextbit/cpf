# cpf
[![Build status](https://secure.travis-ci.org/ustrajunior/cpf.png)](https://secure.travis-ci.org/ustrajunior/cpf)

Brazilian CPF validator

## Instalation

go get github.com/ustrajunior/cpf

## Usage

Call the valid func passing the cpf string
Returns bool

	cpf.Valid("886.100.254-46") => true
	cpf.Valid("88610025446") => true

	cpf.Valid("123.123.123.12") => false
	cpf.Valid("12312312") => false

