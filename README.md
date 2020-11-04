# Pretty Table
**Pretty Table** is a simple Go table view app to show tables prettier in terminal .

Usage
---
You can create tables by create an instance of "Table" struct and configure it.
```go
//All the information used in this example is untrue

type Table struct {
	values       [][]string
	headers      []string
	lockedHeader bool
	bodyTheme    string
	color        string
	headerSplit  bool
	paddingRight int
	paddingLeft  int
}

//Create an instance
clients := Table{
	values: [][]string{
		{"112.83.68.215", "be:ec:ad:ea:b4:d6", "USA"},
		{"168.31.13.241", "fc:ee:bd:21:eb:e2", "Canada"},
		{"137.114.50.162", "3a:55:06:c8:e3:4b", "England"},
	},
	headers:      []string{"IP", "Mac Address", "Country"},
	lockedHeader: false, //No support currently
	bodyTheme:    "t1", //"t1" & "t2" only currently
	color:        "#FFFFFF", //No support currently
	headerSplit:  false, //No support currently
	paddingLeft:  2,
	paddingRight: 3,
}
```

Also you can use **pushValue** method to add data to table :

```go

clients.pushValue([]string{"32.255.101.12", "93:31:80:fd:42:b7", "USA"})

```

Print
--
To print the table in terminal use **print** method . 
```go
clients.print()
```
The Output will be like this :
```text
┌──────────────────┬─────────────────────┬───────────┐
│  IP              │  Mac Address        │  Country  │
├──────────────────┼─────────────────────┼───────────┤
│  112.83.68.215   │  be:ec:ad:ea:b4:d6  │  USA      │
├──────────────────┼─────────────────────┼───────────┤
│  168.31.13.241   │  fc:ee:bd:21:eb:e2  │  Canada   │
├──────────────────┼─────────────────────┼───────────┤
│  137.114.50.162  │  3a:55:06:c8:e3:4b  │  England  │
├──────────────────┼─────────────────────┼───────────┤
│  32.255.101.12   │  93:31:80:fd:42:b7  │  USA      │
└──────────────────┴─────────────────────┴───────────┘
```

If you choose "t2" theme , the output will be like this :

```text
╔══════════════════╤═════════════════════╤═══════════╗
║  IP              │  Mac Address        │  Country  ║
╟──────────────────┼─────────────────────┼───────────╢
║  112.83.68.215   │  be:ec:ad:ea:b4:d6  │  USA      ║
╟──────────────────┼─────────────────────┼───────────╢
║  168.31.13.241   │  fc:ee:bd:21:eb:e2  │  Canada   ║
╟──────────────────┼─────────────────────┼───────────╢
║  137.114.50.162  │  3a:55:06:c8:e3:4b  │  England  ║
╟──────────────────┼─────────────────────┼───────────╢
║  32.255.101.12   │  93:31:80:fd:42:b7  │  USA      ║
╚══════════════════╧═════════════════════╧═══════════╝
```