# ViterbiWords

An implementation of the Viterbi algorithm in Go to find words in unbroken strings using a corpus.

### Corpus
Choosing an appropriate corpus is the key to this algorithm's success. I used the English News sentences corpus from [Leipzig University's Corpora collection](http://wortschatz.uni-leipzig.de/en/download/).

### Example
```
  c := newcorpus()
  c.loadwords("eng_news_2015_1M-sentences.txt")
  fmt.Println(c.viterbi("ihaveadogandthedogiscool"))
 ```
 
 **Output**
 ```
 [i have a dog and the dog is cool]
 ```
 
 ### Author
  Kris Hunt
 
 ### License
  Apache 2.0 License
