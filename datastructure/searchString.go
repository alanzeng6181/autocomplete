package datastructure

import (
	"errors"
	"fmt"
)

type SearchString struct {
	Text  string
	Count int
}

func (searchString SearchString) String() string {
	return fmt.Sprintf("{text:%s, count:%d}", searchString.Text, searchString.Count)
}

func MakeSearchString(text string, count int) SearchString {
	return SearchString{text, count}
}

func MakeSearchString1(text string) SearchString {
	return SearchString{Text: text}
}

func (searchString SearchString) CanBeRoot() bool {
	return len(searchString.Text) == 1
}

func (searchString SearchString) Add(another Trier) (Trier, error) {
	anotherSearchString, ok := another.(SearchString)
	if !ok {
		return nil, errors.New("parameter is not of type SearchString")
	}
	return SearchString{Text: searchString.Text, Count: searchString.Count + anotherSearchString.Count}, nil
}

func (searchString SearchString) Compare(anotherTier Trier) CompareResult {
	anotherSearchString, ok := anotherTier.(SearchString)
	if !ok {
		return Unrelated
	}

	if anotherSearchString.Text == searchString.Text {
		return Equal
	}

	if len(anotherSearchString.Text) == len(searchString.Text) {
		return Unrelated
	}

	if len(anotherSearchString.Text) > len(searchString.Text) && anotherSearchString.Text[:len(searchString.Text)] == searchString.Text {
		if len(anotherSearchString.Text) == len(searchString.Text)+1 {
			return IsParent
		}
		return IsAncestor
	}

	if len(searchString.Text) > len(anotherSearchString.Text) && searchString.Text[:len(anotherSearchString.Text)] == anotherSearchString.Text {
		if len(searchString.Text) == len(anotherSearchString.Text)+1 {
			return IsChild
		}
		return IsDescendent
	}
	return Unrelated
}

func (searchString SearchString) GetChildOf(ancestor Trier) (Trier, error) {

	anotherSearchString, ok := ancestor.(SearchString)
	if !ok {
		return nil, errors.New("parameter is not of type SearchString")
	}

	compareResult := anotherSearchString.Compare(searchString)
	if compareResult != IsParent && compareResult != IsAncestor {
		return nil, fmt.Errorf("parameter text:%s must be a substring of searchString.Text %s", anotherSearchString.Text, searchString.Text)
	}

	return SearchString{
		Text: searchString.Text[:len(anotherSearchString.Text)+1],
	}, nil
}

func (searchString SearchString) Less(t Trier) bool {
	anotherSearchString, _ := t.(SearchString)
	//descending
	return searchString.Count > anotherSearchString.Count
}
