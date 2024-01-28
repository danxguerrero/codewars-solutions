// 7 kyu Complementary DNA

// Deoxyribonucleic acid (DNA) is a chemical found in the nucleus of cells and carries the "instructions" for the development and functioning of living organisms.

// If you want to know more: http://en.wikipedia.org/wiki/DNA

// In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G". Your function receives one side of the DNA (string, except for Haskell); you need to return the other complementary side. DNA strand is never empty or there is no DNA at all (again, except for Haskell).

// More similar exercise are found here: http://rosalind.info/problems/list-view/ (source)

// Example: (input --> output)

// "ATTGC" --> "TAACG"
// "GTAT" --> "CATA"

package kata

func DNAStrand(dna string) string {
  var strand string
  options := []rune{'A', 'T', 'C', 'G'}
  for _, s := range dna {
    switch s {
      case options[0]:
        strand = strand + "T"
      case options[1]:
        strand = strand + "A"
      case options[2]:
        strand = strand + "G"
      case options[3]:
        strand = strand + "C"
    }
  }
  return strand
}