package datastructure_test

import (
	"math/rand"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/alanzeng6181/autocomplete/datastructure"
)

var trie *datastructure.Trie[datastructure.SearchString]
var arr []datastructure.SearchString

func TestMain(m *testing.M) {
	trie = datastructure.NewTrie[datastructure.SearchString]()
	arr = make([]datastructure.SearchString, 0)
	os.Exit(m.Run())
}
func FuzzTrieCreate(f *testing.F) {
	for _, s := range []string{"a", "c", "dez", "dhze", "dhasdfjk", "adfaw", "asdfasdfi", "dhsk", "zzzz", "hzc"} {
		f.Add(s)
	}

	rand.Seed(time.Now().Unix())
	f.Fuzz(func(t *testing.T, a string) {
		searchString := datastructure.NewSearchString(a, int(rand.Int31n(2000)))
		topNRecords := trie.GetTopN(*searchString)
		for _, r := range topNRecords {
			if len(r.Text) < len(a) || r.Text[:len(a)] != a {
				t.Errorf("%s is not a superstring of %s", r.Text, a)
			}
		}

		superStrings := make([]*datastructure.SearchString, 0)
		for _, c := range arr {
			if len(c.Text) >= len(a) && c.Text[:len(a)] == a {
				superStrings = append(superStrings, &c)
			}
		}

		superStrings = superStrings[:datastructure.MinInt(len(superStrings), datastructure.DEFAULT_TOPN)]
		sort.Sort(datastructure.Data[datastructure.SearchString](topNRecords))
		sort.Sort(datastructure.Data[datastructure.SearchString](superStrings))
		if len(topNRecords) != len(superStrings) {
			t.Errorf("%s returned %d records, when there should be %d records. Returned:%v, Expected:%v",
				a, len(topNRecords), len(superStrings), topNRecords, superStrings)
		}
		for i := 0; i < len(topNRecords); i++ {
			if topNRecords[i].Text != topNRecords[i].Text {
				t.Errorf("%dth element of topN records not matched, expected %v, but got %v", i, superStrings[i], topNRecords[i])
			}
		}
		arr = append(arr, *searchString)
		trie.Add(*searchString)
	})
}

func TestDataT(t *testing.T) {
	arr := []*datastructure.SearchString{
		datastructure.NewSearchString("aabc", 1),
		datastructure.NewSearchString("zabc", 21),
		datastructure.NewSearchString("babc", 11),
		datastructure.NewSearchString("cabc", 51),
		datastructure.NewSearchString("eabc", 71),
	}

	sort.Sort(datastructure.Data[datastructure.SearchString](arr))

	for i := 0; i < len(arr)-1; i++ {
		if arr[i].Count < arr[i+1].Count {
			t.Errorf("array is not sorted")
		}
	}
}

func TestTopNT(t *testing.T) {
	arr1 := []*datastructure.SearchString{
		datastructure.NewSearchString("aabc", 1),
		datastructure.NewSearchString("zabc", 21),
		datastructure.NewSearchString("babc", 11),
		datastructure.NewSearchString("cabc", 151),
		datastructure.NewSearchString("eabc", 171),
		datastructure.NewSearchString("cabc", 251),
		datastructure.NewSearchString("eabc", 271),
		datastructure.NewSearchString("cabc", 1251),
		datastructure.NewSearchString("eabc", 371),
		datastructure.NewSearchString("cabc", 2351),
		datastructure.NewSearchString("eabc", 471),
		datastructure.NewSearchString("cabc", 551),
		datastructure.NewSearchString("eabc", 1571),
	}

	tn1 := datastructure.MakeTopN2(arr1)

	arr2 := []*datastructure.SearchString{
		datastructure.NewSearchString("aabc", 91),
		datastructure.NewSearchString("zabc", 901),
		datastructure.NewSearchString("babc", 111),
		datastructure.NewSearchString("zcabc", 511),
		datastructure.NewSearchString("deabc", 971),
		datastructure.NewSearchString("ababc", 211),
		datastructure.NewSearchString("cabc", 519),
		datastructure.NewSearchString("dfedabc", 731),
		datastructure.NewSearchString("dfbabc", 311),
		datastructure.NewSearchString("dfcabc", 951),
		datastructure.NewSearchString("adfeabc", 871),
	}

	tn2 := datastructure.MakeTopN2(arr2)

	tn1.Merge(tn2)

	for i := 0; i < tn1.GetData().Len()-1; i++ {
		if tn1.GetData()[i].Count < tn1.GetData()[i+1].Count {
			t.Error("TopN must be descending")
		}
	}

	greaterThanEqualsCount := 0
	lastElement := tn1.GetData()[tn1.GetData().Len()-1]
	expected := make([]*datastructure.SearchString, 0)

	for _, s := range append(arr1, arr2...) {
		if s.Count > lastElement.Count || (s.Text == lastElement.Text && s.Count == lastElement.Count) {
			expected = append(expected, s)
			greaterThanEqualsCount += 1
		}
	}

	sort.Sort(datastructure.Data[datastructure.SearchString](expected))
	if greaterThanEqualsCount != tn1.GetData().Len() {
		t.Error("element that should've been in TopN is not in TopN")
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != tn1.GetData()[i] {
			t.Errorf("expected %v at %dth element, but got %v", expected[i], i, tn1.GetData()[i])
		}
	}
}

func TestTrieAdd(t *testing.T) {
	arr := []*datastructure.SearchString{
		datastructure.NewSearchString("aabc", 1),
		datastructure.NewSearchString("zabc", 21),
		datastructure.NewSearchString("babc", 11),
		datastructure.NewSearchString("xcabc", 151),
		datastructure.NewSearchString("eabc", 171),
		datastructure.NewSearchString("cyabc", 251),
		datastructure.NewSearchString("eabc", 2471),
		datastructure.NewSearchString("czabc", 1251),
		datastructure.NewSearchString("eabc", 371),
		datastructure.NewSearchString("cawbc", 2351),
		datastructure.NewSearchString("eabc", 4721),
		datastructure.NewSearchString("caobc", 3551),
		datastructure.NewSearchString("eabck", 13571),
		datastructure.NewSearchString("aabcww", 9851),
		datastructure.NewSearchString("zabc", 901),
		datastructure.NewSearchString("babc", 1110),
		datastructure.NewSearchString("zcabc", 511),
		datastructure.NewSearchString("deabc", 9171),
		datastructure.NewSearchString("ababc", 2011),
		datastructure.NewSearchString("cabc", 519),
		datastructure.NewSearchString("dfedabc", 731),
		datastructure.NewSearchString("dfbabc", 311),
		datastructure.NewSearchString("dfcabc", 951),
		datastructure.NewSearchString("adfeabc", 871),
		datastructure.NewSearchString("kkaabc", 91),
		datastructure.NewSearchString("zwabc", 9201),
		datastructure.NewSearchString("badfbc", 1611),
		datastructure.NewSearchString("zcaasdbc", 5511),
		datastructure.NewSearchString("deasdabc", 971),
		datastructure.NewSearchString("abzwabc", 2141),
		datastructure.NewSearchString("caqjdhobc", 519),
		datastructure.NewSearchString("dfedazbc", 7231),
		datastructure.NewSearchString("dfbkdklabc", 311),
		datastructure.NewSearchString("dwfcabc", 951),
		datastructure.NewSearchString("wwnadfeabc", 8731),
		datastructure.NewSearchString("wwn", 333),
	}
	trie := datastructure.NewTrie[datastructure.SearchString]()
	for _, a := range arr {
		trie.Add(*a)
	}
	sort.Sort(datastructure.Data[datastructure.SearchString](arr))

	keyWords := []string{"aaab", "df", "z", "ca", "dfb", "b"}

	for _, k := range keyWords {
		expected := make([]*datastructure.SearchString, 0)
		topN := trie.GetTopN(datastructure.MakeSearchString1(k))
		addedToExpected := make(map[string]bool)
		for i := 0; i < len(arr); i++ {
			s := arr[i]
			if _, ok := addedToExpected[s.Text]; ok {
				continue
			}

			if len(k) <= len(s.Text) &&
				k == s.Text[:len(k)] {
				expected = append(expected, s)
				addedToExpected[s.Text] = true
			}
		}

		if len(topN) > len(expected) || len(topN) < len(expected) && len(topN) < datastructure.DEFAULT_TOPN {
			t.Errorf("expecte length of %d items in topN search of keyword %s, but got %d. Expected:%v, Got:%v", len(expected), k, len(topN), expected, topN)
		}
	}
}

func TestTrieAdd2(t *testing.T) {
	arr := []*datastructure.SearchString{
		datastructure.NewSearchString("as", 1),
		datastructure.NewSearchString("as is", 1),
		datastructure.NewSearchString("at this", 1),
		datastructure.NewSearchString("tomorrow", 1),
		datastructure.NewSearchString("tomorrow", 1),
		datastructure.NewSearchString("tomorrow", 1),
		datastructure.NewSearchString("as", 1),
		datastructure.NewSearchString("as", 1),
		datastructure.NewSearchString("as", 1),
		datastructure.NewSearchString("as", 1),
	}
	trie := datastructure.NewTrie[datastructure.SearchString]()
	for _, a := range arr {
		trie.Add(*a)
	}
	sort.Sort(datastructure.Data[datastructure.SearchString](arr))

	keyWords := []string{""}

	for _, k := range keyWords {
		expected := make([]datastructure.SearchString, 0)
		topN := trie.GetTopN(datastructure.MakeSearchString1(k))
		addedToExpected := make(map[string]bool)
		for i := 0; i < len(arr); i++ {
			s := arr[i]
			if _, ok := addedToExpected[s.Text]; ok {
				continue
			}

			if len(k) <= len(s.Text) &&
				k == s.Text[:len(k)] {
				expected = append(expected, *s)
				addedToExpected[s.Text] = true
			}
		}

		if len(topN) > len(expected) || len(topN) < len(expected) && len(topN) < datastructure.DEFAULT_TOPN {
			t.Errorf("expecte length of %d items in topN search of keyword %s, but got %d. Expected:%v, Got:%v", len(expected), k, len(topN), expected, topN)
		}
	}
}

func TestTrieAddIncrementing(t *testing.T) {
	arr := []*datastructure.SearchString{
		datastructure.NewSearchString("as", 1),
		datastructure.NewSearchString("as is", 1),
		datastructure.NewSearchString("at this", 1),
		datastructure.NewSearchString("tomorrow", 1),
		datastructure.NewSearchString("as", 1),
	}
	trie := datastructure.NewTrie[datastructure.SearchString]()
	for _, a := range arr {
		trie.Add(*a)
	}
	sort.Sort(datastructure.Data[datastructure.SearchString](arr))

	keyWords := []string{""}

	for _, k := range keyWords {
		expected := make([]datastructure.SearchString, 0)
		topN := trie.GetTopN(datastructure.MakeSearchString1(k))
		addedToExpected := make(map[string]bool)
		for i := 0; i < len(arr); i++ {
			s := arr[i]
			if _, ok := addedToExpected[s.Text]; ok {
				continue
			}

			if len(k) <= len(s.Text) &&
				k == s.Text[:len(k)] {
				expected = append(expected, *s)
				addedToExpected[s.Text] = true
			}
		}

		if len(topN) > len(expected) || len(topN) < len(expected) && len(topN) < datastructure.DEFAULT_TOPN {
			t.Errorf("expecte length of %d items in topN search of keyword %s, but got %d. Expected:%v, Got:%v", len(expected), k, len(topN), expected, topN)
		}
	}
}

func TestIncrementCount(t *testing.T) {
	tri := datastructure.NewTrie[datastructure.SearchString]()
	tri.Add(*datastructure.NewSearchString("abc", 10))
	tri.Add(*datastructure.NewSearchString("abc", 11))
	topN := tri.GetTopN(datastructure.MakeSearchString1("a"))
	if len(topN) != 1 {
		t.Errorf("expected searchString of 'a' to return 1 element but got %d", len(topN))
	}

	if topN[0].Text != "abc" {
		t.Errorf("expected the result element to be 'abc', but got %s", topN[0].Text)
	}

	if topN[0].Count != 21 {
		t.Errorf("expected string 'abc' to have count of 21, but was %d", topN[0].Count)
	}
}
