package main

import "github.com/faiface/pixel/pixelgl"

type Group struct {
	entities []Entity
	enabled  bool
}

func MakeGroup(entities ...Entity) Group {
	return Group{entities: entities, enabled: true}
}

func (g *Group) GetEntities() []Entity {
	return g.entities
}

func (g *Group) GetEnabled() bool {
	return g.enabled
}

func (g *Group) SetEnabled(enabled bool) {
	g.enabled = enabled
}

func (g *Group) AddEntity(entity Entity) {
	g.entities = append(g.entities, entity)
}

func (g *Group) RemoveEntity(entity Entity) {
	for i, e := range g.entities {
		if e == entity {
			g.entities = append(g.entities[:i], g.entities[i+1:]...)
			break
		}
	}
}

func (g *Group) Update(dt float64, win *pixelgl.Window) {
	if g.enabled {
		for _, e := range g.entities {
			e.Update(dt, win)
		}
	}
}

func (g *Group) Draw(win *pixelgl.Window) {
	if g.enabled {
		for _, e := range g.entities {
			e.Draw(win)
		}
	}
}
