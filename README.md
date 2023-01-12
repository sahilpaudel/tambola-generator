## Tambola Generator

This is Go Library that can used to generate upto 100 tambola 
tickets at a time.

### How it works

First add the dependency in your code
```
go get github.com/sahilpaudel/tambola-generator
```

Import as the below

```
tambolagenerator "github.com/sahilpaudel/tambola-generator"
```

Call the generate function passing number of tickets desired

```
tickets, err := tambolagenerator.GenerateTickets(10)
```

It will return tickets which is an array (size 10) of 3X9 tambola ticket.

```
[
	[
	    [0 11 25 35 42 0 63 0 0]
	    [3 0 26 36 0 51 0 72 0]
	    [9 18 0 0 43 57 0 0 86]
	]
]
```