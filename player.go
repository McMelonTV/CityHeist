package main

import (
	"graphics.gd/classdb/AnimatedSprite2D"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/CollisionShape2D"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/SpriteFrames"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

type Player struct {
	CharacterBody2D.Extension[Player] `gd:"GamePlayer"`

	Speed Vector2.XY

	AnimatedSprite2D AnimatedSprite2D.Instance
	CollisionShape2D CollisionShape2D.Instance
}

var PlayerControls = struct{ MoveUp, MoveLeft, MoveDown, MoveRight string }{
	MoveUp:    "ui_up",
	MoveLeft:  "ui_left",
	MoveDown:  "ui_down",
	MoveRight: "ui_right",
}

var PlayerAnimationDirection = struct {
	Up    string
	Left  string
	Down  string
	Right string
}{
	Up:    "up",
	Left:  "left",
	Down:  "down",
	Right: "right",
}

var PlayerAnimationType = struct {
	Default SpriteFrames.Instance
	Idle    SpriteFrames.Instance
	Walk    SpriteFrames.Instance
}{
	Default: Resource.Load[SpriteFrames.Instance]("res://animations/Player/default.tres"),
	Idle:    Resource.Load[SpriteFrames.Instance]("res://animations/Player/idle.tres"),
	Walk:    Resource.Load[SpriteFrames.Instance]("res://animations/Player/walk.tres"),
}

func (p *Player) Ready() {
	p.Speed = Vector2.New(100, 100)
}

func (p *Player) Process(delta Float.X) {
	var velocity Vector2.XY
	if Input.IsActionPressed(PlayerControls.MoveRight, false) {
		velocity.X += 1
	}
	if Input.IsActionPressed(PlayerControls.MoveLeft, false) {
		velocity.X -= 1
	}
	if Input.IsActionPressed(PlayerControls.MoveDown, false) {
		velocity.Y += 1
	}
	if Input.IsActionPressed(PlayerControls.MoveUp, false) {
		velocity.Y -= 1
	}

	if Vector2.Length(velocity) > 0 {
		velocity = Vector2.Mul(Vector2.Normalized(velocity), p.Speed)
		p.AnimatedSprite2D.SetSpriteFrames(PlayerAnimationType.Walk)
	} else {
		p.AnimatedSprite2D.SetSpriteFrames(PlayerAnimationType.Idle)
	}
	position := p.AsNode2D().Position()
	position = Vector2.Add(position, Vector2.MulX(velocity, delta))

	p.AsNode2D().SetPosition(Vector2.Clamp(position, Vector2.Zero, p.AsCanvasItem().GetViewportRect().Size))

	if velocity.X != 0 {
		if velocity.X < 0 {
			p.AnimatedSprite2D.SetAnimation(PlayerAnimationDirection.Left)
		} else if velocity.X > 0 {
			p.AnimatedSprite2D.SetAnimation(PlayerAnimationDirection.Right)
		}
	} else if velocity.Y != 0 {
		if velocity.Y < 0 {
			p.AnimatedSprite2D.SetAnimation(PlayerAnimationDirection.Up)
		} else if velocity.Y > 0 {
			p.AnimatedSprite2D.SetAnimation(PlayerAnimationDirection.Down)
		}
	}

	p.AnimatedSprite2D.Play()
}

func (p *Player) Start(pos Vector2.XY) {
	p.AsNode2D().SetPosition(pos)
	p.AsCanvasItem().Show()

	Engine.Println("start")
}
