package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Skill struct {
	Name          string
	Progress      float64
	Prerequisites []*Skill
}

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

func parseInput(inputFilePath string) (map[string]*Skill, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	graph := make(map[string]*Skill)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "->")
		if len(parts) == 2 {
			childName := strings.TrimSpace(parts[1])
			parentName := strings.TrimSpace(parts[0])
			childSkill, exists := graph[childName]
			if !exists {
				childSkill = &Skill{Name: childName}
				graph[childName] = childSkill
			}
			parentSkill, exists := graph[parentName]
			if !exists {
				parentSkill = &Skill{Name: parentName}
				graph[parentName] = parentSkill
			}
			childSkill.Prerequisites = append(childSkill.Prerequisites, parentSkill)
		} else if len(parts) == 1 && strings.Contains(parts[0], "=") {
			progressParts := strings.Split(parts[0], "=")
			name := strings.TrimSpace(progressParts[0])
			progress := strings.TrimSpace(progressParts[1])
			progressValue := 0.0
			fmt.Sscanf(progress, "%f", &progressValue)
			skill, exists := graph[name]
			if exists {
				skill.Progress = progressValue
			} else {
				graph[name] = &Skill{Name: name, Progress: progressValue}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}

func InitFlags() {
	flag.Parse()
}

var (
	DockerContext = flag.String("dockerContext", ".", "Relative path to the context of the docker build. This directory should contain your Dockerfile")
	Dockerfile    = flag.String("dockerfile", "Dockerfile", "Relative path to the directory containing your Dockerfile")
)

func main() {
	InitFlags()
	inputFilePath := "input.txt"
	if len(os.Args) > 1 {
		inputFilePath = os.Args[1]
	}

	graph, err := parseInput(inputFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %s\n", err)
		os.Exit(1)
	}

	order := LearningRoadmap(graph)

	for _, skill := range order {
		fmt.Printf("%s\n", skill.Name)
	}

}
