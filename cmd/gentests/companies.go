package main

import (
	"bufio"
	"embed"
)

//go:embed _companies/*
var companiesFS embed.FS

type Companies struct {
	PlatinumGolds []string
	Golds         []string
}

func readCompaniesAll() (*Companies, error) {
	pgs, err := readCompanies("_companies/platinum_gold.txt")
	if err != nil {
		return nil, err
	}

	gs, err := readCompanies("_companies/gold.txt")
	if err != nil {
		return nil, err
	}

	return &Companies{
		PlatinumGolds: pgs,
		Golds:         gs,
	}, nil
}

func readCompanies(name string) ([]string, error) {
	txt, err := companiesFS.Open(name)
	if err != nil {
		return nil, err
	}
	defer txt.Close()

	var companies []string
	s := bufio.NewScanner(txt)
	for s.Scan() {
		companies = append(companies, s.Text())
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
