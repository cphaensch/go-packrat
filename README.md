go-packrat
============================

This library allows to construct backtracking top down packrat parsers in Go using parser combination. Packrat parsing enables the parsing of PEG Grammars in linear time. Parsers are combinated using the following basic parsers:

- `AtomParser`: Matches only a specified UTF8 string
- `RegexParser`: Matches a regular expression
- `AndParser`: Matches a given list of parsers sequentially
- `OrParser`: Matches if any of a given list of parsers matches
- `KleeneParser`: Matches a parser 0 to `n` times, optionally separated by another parser
- `ManyParser`: Matches a parser 1 to `n` times, optionally separated by another parser
- `MaybeParser`: Matches a parser 0 or 1 times

By default, `Atom` and `Regex` parsers skip (but do not match on) leading whitespace. This can be configured per parser.

If a parser matches, it returns an AST `*Node`. Every node points to the parser that produced it, the matched text, and a list of child nodes.

This library is currently used in production, but some rarely used features may be broken. Additional documentation is ToDo.

Example
-----------

A full example in form of a working json parser is provided in `json_test.go`.

```go
import (
    packrat "github.com/launix-de/go-packrat"
)

func main(){
    input := "Hello World"
    scanner := NewScanner(input, true)

    helloParser := NewAtomParser("Hello", true)
    worldParser := NewAtomParser("World", true)
    helloAndWorldParser := NewAndParser(helloParser, worldParser)

    n, err := Parse(helloAndWorldParser, scanner)
    // n is the AST root node
    // n.Children is a slice containing a node for both the Hello and World parser
}
```

Use case
-----------
Using this library, you can dynamically define and parse PEG grammars at runtime. Parsing time is proportional to the input length and grammar complexity. Note that if you do not need to build grammars at runtime, a parser generator like [gocc](https://github.com/goccmack/gocc) will produce a static LR parser that is both faster and uses less memory than `go-packrat`. If you want to use a parser combinator, worst-case exponential runtime is not a problem or low memory consumption is required, consider using [goparsec](https://github.com/prataprc/goparsec). `go-packrat` was written as a replacement for `goparsec` that solves the time complexity issue using a simple memoizatino cache. 

License
------------
Dual licensed with custom aggreements or GPLv3. See `LICENSE`.