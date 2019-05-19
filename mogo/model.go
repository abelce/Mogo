/*
 * File: model.go
 * Project: Mogo
 * File Created: Thursday, 16th May 2019 2:49:38 pm
 * Author: zxtang (1061225829@qq.com)
 * -----
 * Last Modified: Thursday, 16th May 2019 2:49:38 pm
 * Modified By: zxtang (1061225829@qq.com>)
 * -----
 * Copyright 2017 - 2019 Your Company, Your Company
 */
package mogo

import "time"

type ArticleConfig struct {
	Title     string
	Date      string
	ShortDate string
	Category  string
	Tags      []TagConfig
	Abstract  string
	Author    string
	Time      time.Time
	Link      string
	Content   string
	Nav       []NavConfig
}

type NavConfig struct {
	Name   string
	Href   string
	Target string
}

type ArticleBase struct {
	Link  string
	Title string
}

type TagConfig struct {
	Name         string
	ArticleTitle string
	ArticleLink  string
}

type Tag struct {
	Name     string
	Articles []ArticleBase
	Length   int
}

type Category struct {
	Name     string
	Articles []ArticleBase
	Length   int
}

type CustomPage struct {
	Id      string
	Title   string
	Content string
}

type YearArchive struct {
	Year   string
	Months []*MonthArchive
	months map[string]*MonthArchive
}

type MonthArchive struct {
	Month    string
	month    time.Month
	Articles []*ArticleBase
}

type Artic []*ArticleConfig
type ByDate struct{ Artic }

func (a Artic) Len() int            { return len(a) }
func (a Artic) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a.Artic[i].Time.After(a.Artic[j].Time) }

type YearArchives []*YearArchive

func (y YearArchives) Len() int {
	return len(y)
}

func (y YearArchives) Swap(i, j int) {
	y[i], y[j] = y[j], y[i]
}

func (y YearArchives) Less(i, j int) bool {
	return y[i].Year > y[j].Year
}

type MonthArchives []*MonthArchive

func (m MonthArchives) Len() int {
	return len(m)
}

func (m MonthArchives) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MonthArchives) Less(i, j int) bool {
	return m[i].month > m[j].month
}
