package main

type plugin struct{}

var JavaScriptPlugin plugin

func (*plugin) Name() string {
	return "Leftpad"
}

func (*plugin) Setup() error {
	return nil
}

func (*plugin) Teardown() error {
	return nil
}

func (*plugin) GetModules() map[string]interface{} {
	mods := map[string]interface{}{
		"leftpad": New(),
	}
	return mods
}

func init() {
	JavaScriptPlugin = plugin{}
}
