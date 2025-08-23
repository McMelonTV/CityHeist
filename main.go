package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Node"
	"graphics.gd/startup"
)

type Main struct {
	Node.Extension[Main] `gd:"GameMain"`
}

func main() {
	classdb.Register[Player]()
	classdb.Register[Main]()

	startup.LoadingScene()

	//hello := Label.New()
	//hello.AsControl().SetAnchorsPreset(Control.PresetFullRect)
	//hello.SetHorizontalAlignment(GUI.HorizontalAlignmentCenter)
	//hello.SetVerticalAlignment(GUI.VerticalAlignmentCenter)
	//hello.SetText("Hello, World!")
	//
	//SceneTree.Add(hello)

	startup.Scene()
}
