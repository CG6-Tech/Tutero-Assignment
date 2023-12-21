package main

type Skill struct {
	Name          string
	Progress      float64
	Prerequisites []*Skill
}
