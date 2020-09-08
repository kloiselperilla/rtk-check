package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

type RtkListing struct {
	Kanji string
	Number int
}

func (r RtkListing) print() {
	fmt.Printf("%d: %s\n", r.Number, r.Kanji)
}

func pos(slice []string, value string) int {
    for p, v := range slice {
        if (v == value) {
            return p
        }
    }
    return -1
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func readKanjiList() []string {
   dat, err := ioutil.ReadFile("./rtk6-list.txt")
   check(err)
   return strings.Split(string(dat), " ")
}

func main() {
    rtkProgress, err := strconv.Atoi(os.Args[1])
    check(err)
    kanjiToCheck := os.Args[2]
	kanjiList := readKanjiList()
	var knownJoyoKanji []RtkListing
	var unknownJoyoKanji []RtkListing
	var nonJoyoKanji []string
    for _, kanji := range kanjiToCheck {
		if (kanji >= 0x4E00 && kanji <= 0x9FBF) {
	    	rtkNumber := pos(kanjiList, string(kanji))
	    	if rtkNumber == -1 {
		    	nonJoyoKanji = append(nonJoyoKanji, string(kanji))
	    	} else if rtkNumber > rtkProgress {
				unknownJoyoKanji = append(unknownJoyoKanji, RtkListing{string(kanji), rtkNumber})
	    	} else {
				knownJoyoKanji = append(knownJoyoKanji, RtkListing{string(kanji), rtkNumber})
			}
		}
    }
    if len(nonJoyoKanji) != 0 {
		fmt.Println("Contained these non-Joyo kanji:")
		fmt.Println(nonJoyoKanji)
    } else if len(unknownJoyoKanji) == 0 {
	    fmt.Println("You know these ones, my guy")
		for _, listing := range(knownJoyoKanji) {
			listing.print()
		}
    } else {
		fmt.Println("Not quite, you have yet to learn these kanji:")
		for _, listing := range(unknownJoyoKanji) {
			listing.print()
		}
    }
}