package gcodex

// import(
// 	"golang.org/x/tools/cmd/stringer"
// )

// func genEnum(){
// 	// stringer
// }


// var (
// 	typeNames   = flag.String("type", "", "comma-separated list of type names; must be set")
// 	output      = flag.String("output", "", "output file name; default srcdir/<type>_string.go")
// 	trimprefix  = flag.String("trimprefix", "", "trim the `prefix` from the generated constant names")
// 	linecomment = flag.Bool("linecomment", false, "use line comment text as printed text when present")
// 	buildTags   = flag.String("tags", "", "comma-separated list of build tags to apply")
// )

// type Generator struct {
// 	buf bytes.Buffer // Accumulated output.
// 	pkg *Package     // Package we are scanning.

// 	trimPrefix  string
// 	lineComment bool
// }

// func (g *Generator) Printf(format string, args ...interface{}) {
// 	fmt.Fprintf(&g.buf, format, args...)
// }