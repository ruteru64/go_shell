package shell

func setScript() []string {
	//　ここに使用できる変数を追加してください
	s := []string{
		"a",
		"pwd",
		"cat",
		"ls",
		"mkdir",
		"touch",
		"echo",
	}
	return s
}

func setflag() []string {
	//　ここに使用できる変数を追加してください
	s := []string{
		"help",
	}
	return s
}
