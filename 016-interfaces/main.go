package main

import "fmt"

type shadowMonarch interface {
	ShadowExtract() []string
}

type jinwoo struct {
	shadows []string
}

type suhoo struct {
	shadows []string
}

func (j jinwoo) ShadowExtract() []string {
	j.shadows = append(j.shadows, "Igris")
	return j.shadows
}

func (s suhoo) ShadowExtract() []string {
	s.shadows = append(s.shadows, "Beru")
	return s.shadows
}

func ExtractedShadows(sm []shadowMonarch) []string {
	es := []string{}
	for _, v := range sm {
		es = append(es, v.ShadowExtract()...)
	}

	return es
}

func main() {
	sj := jinwoo{}
	ss := suhoo{}

	sj.ShadowExtract()
	ss.ShadowExtract()

	sm := []shadowMonarch{sj, ss}

	es := ExtractedShadows(sm)

	fmt.Println(es)

}
