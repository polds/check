# check

Validation Library for Go. [API Documentation](http://godoc.org/github.com/polds/check) on go.pkgdoc.org. Inspiration and partial port from [is.js](https://github.com/rthor/isjs)

# [Installation](https://github.com/polds/check#installation)

```
go get github.com/polds/check
```

## [Usage](https://github.com/polds/check#usage)

### [Credit Card](https://github.com/polds/check#credit-card)

To check if a credit card is a valid card

```golang
check.CreditCard("4111111111111111").CardIs("Any") // Returns true
```

To validate if a card is a specific type of card (Supported cards: American Express, Visa, Discover, Mastercard)

```golang
check.CreditCard("4111111111111111").CardIs("MasterCard") // Returns false
```

To determine the merchant of credit card

```golang
check.CreditCard("4111111111111111").Merchant() // Returns Visa
```

### [Email](https://github.com/polds/check#email)

To validate an email address

```golang
check.Email("test@test.com").IsValid() // Returns true
```


### [Issues](https://github.com/polds/check#issues)

- Email regex fails: `user@[IPv6:2001:db8:1ff::a0b:dbd0]`, `postbox@com`, `admin@mailserver1`. Besides those examples, regex is RFC compliant.