package main

import "sort"

type ByProgress []*Skill

func (a ByProgress) Len() int           { return len(a) }
func (a ByProgress) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByProgress) Less(i, j int) bool { return a[i].Progress > a[j].Progress }

func LearningRoadmap(graph map[string]*Skill) []*Skill {
	var result []*Skill
	visited := make(map[*Skill]bool)

	var visit func(skill *Skill)
	visit = func(skill *Skill) {
		visited[skill] = true
		for _, prerequisite := range skill.Prerequisites {
			if !visited[prerequisite] {
				visit(prerequisite)
			}
		}
		result = append(result, skill)
	}

	var skills ByProgress
	for _, skill := range graph {
		skills = append(skills, skill)
	}
	sort.Sort(skills)

	for _, skill := range skills {
		if !visited[skill] {
			visit(skill)
		}
	}

	return result
}
