package ling

import (
	"fmt"
)

// A Pipeline contains configured annotators and taggers for nl processing
type Pipeline struct {
	Annotators []string

	taggers []Processor
}

// AddTagger adds a new processor t to Pipeline p
func (p *Pipeline) AddTagger(t Processor) error {
	if t == nil {
		return fmt.Errorf("cannot add nil tagger")
	}
	p.taggers = append(p.taggers, t)
	return nil
}

// Annotate tags the Document by each configured processors and taggers
func (p *Pipeline) Annotate(d *Document) error {
	err := Processors["_"].Process(d)
	if err != nil {
		return err
	}
	for _, anno := range p.Annotators {
		if err = Processors[anno].Process(d); err != nil {
			return err
		}
	}

	for _, tagger := range p.taggers {
		if 