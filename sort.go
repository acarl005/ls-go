package main

type BySize []*DisplayItem

func (s BySize) Less(i, j int) bool {
	return s[i].info.Size() < s[j].info.Size()
}
func (s BySize) Len() int {
	return len(s)
}
func (s BySize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByTime []*DisplayItem

func (s ByTime) Less(i, j int) bool {
	return s[i].info.ModTime().Unix() < s[j].info.ModTime().Unix()
}
func (s ByTime) Len() int {
	return len(s)
}
func (s ByTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByKind []*DisplayItem

func (s ByKind) Less(i, j int) bool {
	var kindi, kindj string
	if s[i].basename == "" {
		kindi = "."
	} else {
		kindi = s[i].ext
	}
	if s[j].basename == "" {
		kindj = "."
	} else {
		kindj = s[j].ext
	}
	return kindi < kindj
}
func (s ByKind) Len() int {
	return len(s)
}
func (s ByKind) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
