/*
 * Find words in long unbroken English language strings
 * using the viterbi algorithm.
 *
 * Author: Kris H <github.com/sourcekris>
 */

package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "regexp"
  "strings"
)

type corpus struct {
  words map[string]int
  maxlen int
  total float64
}

type prob struct {
  prob_k float64
  k int
}

// newCorpus returns a new corpus struct initialized with nil values.
func newCorpus() *corpus {
  return &corpus{}
}

// newProb takes a prob_k and a k and returns a pointer to a prob struct.
func newProb(pk float64, k int) *prob {
  return &prob{
    prob_k: pk,
    k: k,
  }
}

// max finds the maximum of two integers.
func max(x, y int) int {
  if x > y {
    return x
  }

  return y
}

// wordFreq populates the words map in a corpus struct with word frequencies.
func (c *corpus) wordFreq(w []string) {
  out := make(map[string]int)

  var total int
  for _, x := range w {
    if _, ok := out[x]; ok {
      out[x]++
      total++
    } else {
      out[x] = 1
      total++
    }
  }

  c.words = out
  c.total = float64(total)
}

// wordProb calculates how probable word is in context of corpus c.
func (c *corpus) wordProb(word string) float64 {
  return float64(c.words[word]) / c.total
}

// maxProb finds the largest prob_k value from a slice of prob structs and 
// returns the prob_k, k.
func maxProb(ps []*prob) (float64, int) {
  var (
    prob_k float64
    k int
  )

  for _, z := range ps {
    if z.prob_k > prob_k {
      prob_k = z.prob_k
      k = z.k
    }
  }

  return prob_k, k
}

// reverse returns slice s in reverse.
func reverse(s []string) []string {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
    return s
}

// String specifies how a prob struct is to be represented as a string.
func (p prob) String() string {
  return fmt.Sprintf("(%v), %d", p.prob_k, p.k)
}

// words reads words from a file with filename f and packs a corpus c.
func (c *corpus) loadWords(f string) {

  re := regexp.MustCompile("[A-Za-z]+")
  b, err := ioutil.ReadFile(f)
  if err != nil {
    log.Fatal(err)
  }

  wb := re.FindAll(b, -1)
  if wb == nil {
    log.Fatal("Failed to extract words from corpus")
  }

  var ( 
    words []string
    maxlen int
  )
  for _, w := range wb {
    if len(w) > maxlen {
      maxlen = len(w)
    }
    words = append(words, strings.ToLower(string(w)))
  }

  c.wordFreq(words)
  c.maxlen = maxlen
}

// viterbi 
func (c *corpus) viterbi(text string) []string {
  var (
    probs = []float64{1.0}
    lasts = []int{0}
  )

  for i := 1; i < len(text) + 1; i++ {
    var y []*prob
    for j := max(0, i - c.maxlen); j < i; j++ {
      y = append(y, newProb(probs[j] * c.wordProb(text[j:i]), j))
    }

    prob_k, k := maxProb(y)

    probs = append(probs, prob_k)
    lasts = append(lasts, k)
  }

  var (
    words []string
    i = len(text)
  )

  for 0 < i {
    words = append(words, text[lasts[i]:i])
    i = lasts[i]
  }

  return reverse(words)
}

func main() {
  c := newCorpus()
  c.loadWords("eng_news_2015_1M-sentences.txt")
  fmt.Println(c.viterbi("elevenpickleswentintotownandfoundapubopen"))
}