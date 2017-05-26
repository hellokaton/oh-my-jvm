package cmd

/**
	定义结构体，用于存储命令参数。
	java [-options] class [args...]
 */
type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	Class       string
	Args        []string
}
