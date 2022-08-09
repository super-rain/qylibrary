package gcodex

// func main() {
// 	fs := flag.NewFlagSet("gcodex", flag.ExitOnError)
// 	var (
// 		kind = fs.String("kind", "", "generate project Layout,eg:backend,frontend")
// 		name = fs.String("name", "appname", "generate backend/frontend services/app project layout with ent/hugo")
// 		src  = fs.String("src", "", "metadata toml file path")
// 		// dist=fs.String("dist","","target file path")
// 	)
// 	fs.Usage = usageFor(fs, os.Args[0]+" [flags] <a> <b>")
// 	fs.Parse(os.Args[1:])

// 	if *src == "" {
// 		fmt.Fprintf(os.Stderr, "error:%v\n", ErrIDLFileNotExists)
// 	}
// 	err := LoadIDLConfig(*src, *name)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error:%v\n", err)
// 		os.Exit(1)
// 	}
// 	IDL.DB[0].TmplName = "ent.tmpl"
// 	IDL.DB[0].FilePath = tmplFilePath
// 	IDL.DB[0].TarFilePath = "examples/backend/" + *name + "/"
// 	switch *kind {
// 	case "backend":
// 		err := GenBackendLayout()
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "error:%v\n", err)
// 			os.Exit(1)
// 		}
// 	case "frontend":
// 		// TODO: frontend app layout
// 	default:
// 		fmt.Fprintf(os.Stderr, "error: invalid method %q\n", *kind)
// 		os.Exit(1)
// 	}

// }

func Run(kind, name, srcPath, output string) error {
	if kind == "" || name == "" || srcPath == "" {
		return ErrInvalidParams
	}
	err := LoadIDLConfig(srcPath, name)
	if err != nil {
		return err
	}
	IDL.DB[0].TmplName = "ent.tmpl"
	IDL.DB[0].FilePath = tmplFilePath
	IDL.DB[0].TarFilePath = output + name + "/"
	switch kind {
	case "backend":
		err := GenBackendLayout()
		if err != nil {
			return err
		}
	case "frontend":
		// TODO: frontend app layout
	default:
		return ErrInvalidParams
	}
	return nil
}
