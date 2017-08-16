xmlBeautyfier
=============


xmlBeautyfier takes a xml file and prints it with indent (and prefix, if set)

```
go get github.com/scusi/xmlBeautyfier
cd $GOPATH/src/github.com/scusi/xmlBeautyfier
go install
```

From an arbitrary directory you can now:

```
xmlBeautyfier some.xml
```

Flags:
------

-i 	indent (string); sets the indet used, default: 2 whitespaces, "  "
-p	prefix (string); sets the prefix used, default: empty string, ""

	NOTE: if prefix is set to ```FILENAME``` the name of the current file is used as prefix.
	This is useful when parsing a lot of xml files through xmlBeautyfier and pipe the output to grep.
