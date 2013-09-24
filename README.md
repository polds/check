# check

Validation Library for Go. [API Documentation](http://godoc.org/github.com/polds/check) on go.pkgdoc.org.

# Installation

```
go get github.com/polds/check
```

# Usage

# Credit Card

To check if a credit card is a valid card

```golang
check.CreditCard("4111111111111111").CardIs("Any") // Returns true
```

To validate if a card is a specific type of card (Supported cards: American Express, Visa, Discover, Mastercard)

```golang
check.CreditCard("4111111111111111").CardIs("MasterCard") // Returns false
```

To determine the type of credit card

```golang
check.CreditCard("4111111111111111").Card() // Returns Visa
```

# Email

To validate an email address

```golang
check.Email("test@test.com").IsValid() // Returns true
```