package main

import (
  //"fmt"
  "testing"
  "reflect"
  "regexp"
  "strings"
)

const english = `508618  Nicole Kidman was banned from daughter's wedding because of Scientology -- Report Wonderwall 17 hrs ago By Mark Gray "The wedding was a Scientology ceremony," a source told the New York Post's Page Six.
508626  Nicole Thebeau is the co-founder of the Kent County Animal Rescue, a volunteer group that rescues at-risk animals in eastern New Brunswick.
508638  Nico Muhly at His Best The Philadelphia Orchestra performs Nico Muhly’s “Mixed Messages,” at Carnegie Hall.
508648  Nicosia (AFP) - The European Central Bank will start buying government debt in its new quantitative easing programme on March 9, ECB chief Mario Draghi said on Thursday.`

func TestMax(t *testing.T) {
  a := 0
  b := 1

  if max(a,b) != b {
    t.Errorf("expected b > a (%v > %v)", b, a)
  }
}

func TestMaxProb(t *testing.T) {
  a := newProb(0.999, 1)
  b := newProb(0.111, 2)

  var c = []*prob{a,b}

  maxpk, k := maxProb(c)

  if maxpk != a.prob_k || k != a.k {
    t.Errorf("Expected maxprob() to be %v", a)
  }
}

func TestReverse(t *testing.T) {
  var a = []string{"aa", "bb", "cc", "dd"}
  var b = []string{"dd", "cc", "bb", "aa"}

  c := reverse(a)

  if !reflect.DeepEqual(c, b) {
    t.Errorf("Expected %v == %v", c, b)
  }
}

func TestWordFreq(t *testing.T) {
  c := newCorpus()
  var words = []string{"aa", "bb", "ccc", "aa", "aa"}

  c.wordFreq(words)

  if c.words["aa"] != 3 {
    t.Errorf("expected c.words[aa] == 3, got %v", c.words["aa"])
  }
}

func TestWordProb(t *testing.T) {
  c := newCorpus()
  var words = []string{"aa", "bb", "ccc", "aa", "aa"}

  c.wordFreq(words)

  if c.wordProb("aa") != 0.6 {
    t.Errorf("expected c.wordprob(aa) == 0.6, got %v", c.wordProb("aa"))
  } 
}

func TestViterbi(t *testing.T) {
  var ( 
    words []string
    maxlen int
    re = regexp.MustCompile("[A-Za-z]+")
    c = newCorpus()
    expected = []string{"a","central","european","debt",
                        "programme", "is", "because", "of",
                        "the", "bank"}
  )

  wb := re.FindAllString(english, -1)
  if wb == nil {
    t.Error("Failed to extract words from corpus")
  }

  for _, w := range wb {
    if len(w) > maxlen {
      maxlen = len(w)
    }
    words = append(words, strings.ToLower(w))
  }

  c.wordFreq(words)
  c.maxlen = maxlen

  actual := c.viterbi("acentraleuropeandebtprogrammeisbecauseofthebank")

  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("Expected %v got %v", expected, actual)
  }
}