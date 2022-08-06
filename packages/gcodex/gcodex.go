package gcodex

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	fs := flag.NewFlagSet("gcodex", flag.ExitOnError)
	var (
		kind = fs.String("kind", "", "generate project Layout,eg:backend,frontend")
		name = fs.String("name", "appname", "generate backend/frontend services/app project layout with ent/hugo")
		src  = fs.String("src", "", "metadata toml file path")
		// dist=fs.String("dist","","target file path")
	)
	fs.Usage = usageFor(fs, os.Args[0]+" [flags] <a> <b>")
	fs.Parse(os.Args[1:])

	if *src == "" {
		fmt.Fprintf(os.Stderr, "error:%v\n", ErrIDLFileNotExists)
	}
	err := LoadIDLConfig(*src, *name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:%v\n", err)
		os.Exit(1)
	}
	IDL.DB[0].TmplName = "ent.tmpl"
	IDL.DB[0].FilePath = tmplFilePath
	IDL.DB[0].TarFilePath = "examples/backend/" + *name + "/"
	switch *kind {
	case "backend":
		err := GenBackendLayout()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error:%v\n", err)
			os.Exit(1)
		}
	case "frontend":
		// TODO: frontend app layout
	default:
		fmt.Fprintf(os.Stderr, "error: invalid method %q\n", *kind)
		os.Exit(1)
	}

}

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}
