package fasta_test

import (
	. "github.com/MG-RAST/Shock/shock-server/store/type/sequence/fasta"
	//"regexp"
	"testing"
)

func TestRegex(t *testing.T) {
	println("valid:")
	for _, s := range valid {
		println(Regex.MatchString(s))
	}
	println("invalid:")
	for _, s := range invalid {
		println(Regex.MatchString(s))
	}

}

var valid = []string{`>S1 rank=0000056 x=2202.0 y=484.0 length=288
TGAATGTATTCCAGTAAACCGCCCGCGCAAGTAGGCTTCAAATGCCTGCACATTGTCCGTGCCGCGTTTTCAAAGTTTCTGTTCTTCGCCCGAAAGAATAGGAAGAATTGATTTGGCGACCTGTTCGGAAACAATATCTTCAAGCGTCAAAACGTCGCTGAATTTTCATCAAAGGTCTTCGCCCAACACGTCGAATTATCCTTGACGCTCAAAAGCTGCGCTGAAATGCGAATTCTATCACCGACGCGGCGGAGATTACCGTCAAGAATAAAATCAACGCCGAGTTCG`,
	`>S1 rank=0000056 x=2202.0 y=484.0 length=288
TGAATGTATTCCAGTAAACCGCCCGCGCAAGTAGGCTTCAAATGCCTGCACATTGTCCGT
GCCGCGTTTTCAAAGTTTCTGTTCTTCGCCCGAAAGAATAGGAAGAATTGATTTGGCGAC
CTGTTCGGAAACAATATCTTCAAGCGTCAAAACGTCGCTGAATTTTCATCAAAGGTCTTC
GCCCAACACGTCGAATTATCCTTGACGCTCAAAAGCTGCGCTGAAATGCGAATTCTATCA
CCGACGCGGCGGAGATTACCGTCAAGAATAAAATCAACGCCGAGTTCG`, `> S1 rank=0000056 x=2202.0 y=484.0 length=288
TGAATGTATTCCAGTAAACCGCCCGCGCAAGTAGGCTTCAAATGCCTGCACATTGTCCGT
GCCGCGTTTTCAAAGTTTCTGTTCTTCGCCCGAAAGAATAGGAAGAATTGATTTGGCGAC
CTGTTCGGAAACAATATCTTCAAGCGTCAAAACGTCGCTGAATTTTCATCAAAGGTCTTC
GCCCAACACGTCGAATTATCCTTGACGCTCAAAAGCTGCGCTGAAATGCGAATTCTATCA
CCGACGCGGCGGAGATTACCGTCAAGAATAAAATCAACGCCGAGTTCG
> S2 rank=0000091 x=3118.0 y=1043.0 length=340
TTTCGAGCTCCCGAACATTGCCCGGCCAATGGTAGCGGCGCAAGAAGTCGAGGGCTTCCC
TGCTGAGCCGCGGTGGTTGCAGACGCGTGGGCCGCGCGCGGCTTTGGCGATGAAGTGGGC
GGCGAGCGCTGCGATATCCGCTACGCGTTCGCGCAGTGGTGGCACCTGGATGTGAAGTAC
GGCGAGCCGATAGAACAGGTCCTCTCTGAAGGTTCCCTGTCGCACCATCGATAGTAAGTC
TCGATTGGTGGCCGCGATCACTCGAGTGTCGACCGAAAACGAGGTCTGACTTCCGACGCG
CTCGATGCGGTTATCCTGCAATGTGCGCAGCAACGCCACT
> S3 rank=0000153 x=2605.5 y=1640.0 length=430
TTTTATAACCGCGAGGACCTCTGGCAGTTCCCGCGCCAGCCGGGCGGCGATGGGATTGCA
ACGATGGCTCCCTATTACATCATCATGCGGCTGCCCGGTGAGCCGCAAGCCGAGTTCTTT
CTCATGGTCCCGATGGCGCCGAGTCGCCGCGACAACATGATCGCCTGGCTCGCGGCCCGT
TGCGACGCTCCCGACTACGGCAGACTGGTCGTCTACGAATTCCCCAAGGAGAAGCTCGTC
TACGGGCCGTTCCAGATCGAGGCGCGGATCAATCAGAGTACCGAGATATCGCAACAAATC
ACGCTGTGGAACCAGATGGGCTCGCGCGTGATACGCGGCGCAAAACCTGCTTGTGATCCC
GATCGAGAATTCGATACTTTATGTCACACCGCTCTATTTGCGAGCGGAACACGGACACTT
GCCGGAGCTG`, `
>S1 rank=0000056 x=2202.0 y=484.0 length=288
TGAATGTATTCCAGTAAACCGCCCGCGCAAGTAGGCTTCAAATGCCTGCACATTGTCCGT
GCCGCGTTTTCAAAGTTTCTGTTCTTCGCCCGAAAGAATAGGAAGAATTGATTTGGCGAC
CTGTTCGGAAACAATATCTTCAAGCGTCAAAACGTCGCTGAATTTTCATCAAAGGTCTTC
GCCCAACACGTCGAATTATCCTTGACGCTCAAAAGCTGCGCTGAAATGCGAATTCTATCA
CCGACGCGGCGGAGATTACCGTCAAGAATAAAATCAACGCCGAGTTCG

>S2 rank=0000091 x=3118.0 y=1043.0 length=340
TTTCGAGCTCCCGAACATTGCCCGGCCAATGGTAGCGGCGCAAGAAGTCGAGGGCTTCCC
TGCTGAGCCGCGGTGGTTGCAGACGCGTGGGCCGCGCGCGGCTTTGGCGATGAAGTGGGC
GGCGAGCGCTGCGATATCCGCTACGCGTTCGCGCAGTGGTGGCACCTGGATGTGAAGTAC
GGCGAGCCGATAGAACAGGTCCTCTCTGAAGGTTCCCTGTCGCACCATCGATAGTAAGTC
TCGATTGGTGGCCGCGATCACTCGAGTGTCGACCGAAAACGAGGTCTGACTTCCGACGCG
CTCGATGCGGTTATCCTGCAATGTGCGCAGCAACGCCACT

>S3 rank=0000153 x=2605.5 y=1640.0 length=430
TTTTATAACCGCGAGGACCTCTGGCAGTTCCCGCGCCAGCCGGGCGGCGATGGGATTGCA
ACGATGGCTCCCTATTACATCATCATGCGGCTGCCCGGTGAGCCGCAAGCCGAGTTCTTT
CTCATGGTCCCGATGGCGCCGAGTCGCCGCGACAACATGATCGCCTGGCTCGCGGCCCGT
TGCGACGCTCCCGACTACGGCAGACTGGTCGTCTACGAATTCCCCAAGGAGAAGCTCGTC
TACGGGCCGTTCCAGATCGAGGCGCGGATCAATCAGAGTACCGAGATATCGCAACAAATC
ACGCTGTGGAACCAGATGGGCTCGCGCGTGATACGCGGCGCAAAACCTGCTTGTGATCCC
GATCGAGAATTCGATACTTTATGTCACACCGCTCTATTTGCGAGCGGAACACGGACACTT
GCCGGAGCTG`, `>S1
TGAATGTATTCCAGTAAACCGCCCGCGCAAGTAGGCTTCAAATGCCTGCACATTGTCCGT
GCCGCGTTTTCAAAGTTTCTGTTCTTCGCCCGAAAGAATAGGAAGAATTGATTTGGCGAC
CTGTTCGGAAACAATATCTTCAAGCGTCAAAACGTCGCTGAATTTTCATCAAAGGTCTTC
GCCCAACACGTCGAATTATCCTTGACGCTCAAAAGCTGCGCTGAAATGCGAATTCTATCA
CCGACGCGGCGGAGATTACCGTCAAGAATAAAATCAACGCCGAGTTCG
>S2
TTTCGAGCTCCCGAACATTGCCCGGCCAATGGTAGCGGCGCAAGAAGTCGAGGGCTTCCC
TGCTGAGCCGCGGTGGTTGCAGACGCGTGGGCCGCGCGCGGCTTTGGCGATGAAGTGGGC
GGCGAGCGCTGCGATATCCGCTACGCGTTCGCGCAGTGGTGGCACCTGGATGTGAAGTAC
GGCGAGCCGATAGAACAGGTCCTCTCTGAAGGTTCCCTGTCGCACCATCGATAGTAAGTC
TCGATTGGTGGCCGCGATCACTCGAGTGTCGACCGAAAACGAGGTCTGACTTCCGACGCG
CTCGATGCGGTTATCCTGCAATGTGCGCAGCAACGCCACT
>S3
TTTTATAACCGCGAGGACCTCTGGCAGTTCCCGCGCCAGCCGGGCGGCGATGGGATTGCA
ACGATGGCTCCCTATTACATCATCATGCGGCTGCCCGGTGAGCCGCAAGCCGAGTTCTTT
CTCATGGTCCCGATGGCGCCGAGTCGCCGCGACAACATGATCGCCTGGCTCGCGGCCCGT
TGCGACGCTCCCGACTACGGCAGACTGGTCGTCTACGAATTCCCCAAGGAGAAGCTCGTC
TACGGGCCGTTCCAGATCGAGGCGCGGATCAATCAGAGTACCGAGATATCGCAACAAATC
ACGCTGTGGAACCAGATGGGCTCGCGCGTGATACGCGGCGCAAAACCTGCTTGTGATCCC
GATCGAGAATTCGATACTTTATGTCACACCGCTCTATTTGCGAGCGGAACACGGACACTT
GCCGGAGCTG`, `>s1 rank=0000056 x=2202.0 y=484.0 length=288
tgaatgtattccagtaaaccgcccgcgcaagtaggcttcaaatgcctgcacattgtccgt
gccgcgttttcaaagtttctgttcttcgcccgaaagaataggaagaattgatttggcgac?
ctgttcggaaacaatatcttcaagcgtcaaaacgtcgc?tgaattttcatcaaaggtcttc
gcccaacacgtcgaattatccttgacgctcaaaagctgcgctgaaatgcgaattctatca
ccgacgcggcggagattaccgtcaagaataaaatcaacgccgagttcg
>s2 rank=0000091 x=3118.0 y=1043.0 length=340
tttcgagctcccgaacattgcccggccaatggtagcggcgcaagaagtcgagggcttccc
tgctgagccgcggtggttgcagacgcgtgggccgcgcgcggctttggcgatgaagtgggc
ggcgagcgctgcgatatccgctacgcgttcgcgcagtggtggcacctggatgtgaagtac
ggcgagccgatagaacaggtcctctctgaaggttccctgtcgcaccatcgatagtaagtc
tcgattggtggccgcgatcactcgagtgtcgaccgaaaacgaggtctgacttccgacgcg
ctcgatgcggttatcctgcaatgtgcgcagcaacgccact
>s3 rank=0000153 x=2605.5 y=1640.0 length=430
ttttataaccgcgaggacctctggcagttcccgcgccagccgggcggcgatgggattgca
acgatggctccctattacatcatcatgcggctgcccggtgagccgcaagccgagttcttt
ctcatggtcccgatggcgccgagtcgccgcgacaacatgatcgcctggctcgcggcccgt
tgcgacgctcccgactacggcagactggtcgtctacgaattccccaaggagaagctcgtc
tacgggccgttccagatcgaggcgcggatcaatcagagtaccgagatatcgcaacaaatc
acgctgtggaaccagatgggctcgcgcgtgatacgcggcgcaaaacctgcttgtgatccc
gatcgagaattcgatactttatgtcacaccgctctatttgcgagcggaacacggacactt
gccggagctg`}

var invalid = []string{`@SEQ_ID
GATTTGGGGTTCAAAGCAGTATCGATCAAATAGTAAATCCATTTGTTCAACTCACAGTTT
+
!''*((((***+))%%%++)(%%%%).1***-+*''))**55CCF>>>>>>CCCCCCC65
@SEQ_ID
GATTTGGGGTTCAAAGCAGTATCGATCAAATAGTAAATCCATTTGTTCAACTCACAGTTT
+
!''*((((***+))%%%++)(%%%%).1***-+*''))**55CCF>>>>>>CCCCCCC65`, `>
TGAATGTATTCCAGTAAACCGCCCGCGCAAGTAGGCTTCAAATGCCTGCACATTGTCCGT
GCCGCGTTTTCAAAGTTTCTGTTCTTCGCCCGAAAGAATAGGAAGAATTGATTTGGCGAC
CTGTTCGGAAACAATATCTTCAAGCGTCAAAACGTCGCTGAATTTTCATCAAAGGTCTTC
GCCCAACACGTCGAATTATCCTTGACGCTCAAAAGCTGCGCTGAAATGCGAATTCTATCA
CCGACGCGGCGGAGATTACCGTCAAGAATAAAATCAACGCCGAGTTCG`, `>`, `S1
TGAATGTATTCCAGTAAACCGCCCGCGCAAGTAGGCTTCAAATGCCTGCACATTGTCCGT
GCCGCGTTTTCAAAGTTTCTGTTCTTCGCCCGAAAGAATAGGAAGAATTGATTTGGCGAC
CTGTTCGGAAACAATATCTTCAAGCGTCAAAACGTCGCTGAATTTTCATCAAAGGTCTTC
GCCCAACACGTCGAATTATCCTTGACGCTCAAAAGCTGCGCTGAAATGCGAATTCTATCA
CCGACGCGGCGGAGATTACCGTCAAGAATAAAATCAACGCCGAGTTCG
S2
TTTCGAGCTCCCGAACATTGCCCGGCCAATGGTAGCGGCGCAAGAAGTCGAGGGCTTCCC
TGCTGAGCCGCGGTGGTTGCAGACGCGTGGGCCGCGCGCGGCTTTGGCGATGAAGTGGGC
GGCGAGCGCTGCGATATCCGCTACGCGTTCGCGCAGTGGTGGCACCTGGATGTGAAGTAC
GGCGAGCCGATAGAACAGGTCCTCTCTGAAGGTTCCCTGTCGCACCATCGATAGTAAGTC
TCGATTGGTGGCCGCGATCACTCGAGTGTCGACCGAAAACGAGGTCTGACTTCCGACGCG
CTCGATGCGGTTATCCTGCAATGTGCGCAGCAACGCCACT
S3
TTTTATAACCGCGAGGACCTCTGGCAGTTCCCGCGCCAGCCGGGCGGCGATGGGATTGCA
ACGATGGCTCCCTATTACATCATCATGCGGCTGCCCGGTGAGCCGCAAGCCGAGTTCTTT
CTCATGGTCCCGATGGCGCCGAGTCGCCGCGACAACATGATCGCCTGGCTCGCGGCCCGT
TGCGACGCTCCCGACTACGGCAGACTGGTCGTCTACGAATTCCCCAAGGAGAAGCTCGTC
TACGGGCCGTTCCAGATCGAGGCGCGGATCAATCAGAGTACCGAGATATCGCAACAAATC
ACGCTGTGGAACCAGATGGGCTCGCGCGTGATACGCGGCGCAAAACCTGCTTGTGATCCC
GATCGAGAATTCGATACTTTATGTCACACCGCTCTATTTGCGAGCGGAACACGGACACTT
GCCGGAGCTG`, `>s1 rank=0000056 x=2202.0 y=484.0 length=288
tgaatgtattccagtaaaccgcccgcgcaagtaggcttcaaatgcctgcacattgtccgt
gccgcgttttcaaagt$%#%ttctgttcttcgcccgaaagaataggaagaattgatttggcgac?
ctgttcggaaacaatatcttcaagcgtcaaaacgtcgctgaattttcatcaaaggtcttc
gcccaacacgtcgaattatccttgacgctcaaaagctgcgctgaaatgcgaattctatca
ccgacgcggcggagattaccgtcaagaataaaatcaacgccgagttcg
>s2 rank=0000091 x=3118.0 y=1043.0 length=340
tttcgagctcccgaacattgcccggccaatggtagcggcgcaagaagtcgagggcttccc
tgctgagccgcggtggttgcagacgcgtgggccgcgcgcggctttggcgatgaagtgggc
ggcgagcgctgcgatatccgctacgcgttcgcgcagtggtggcacctggatgtgaagtac
ggcgagccgatagaacaggtcctctctgaaggttg%%ccctgtcgcaccatcgatagtaagtc
tcgattggtggccgcgatcactcgagtgtcgaccgaaaacgaggtctgacttccgacgcg
ctcgatgcggttatcctgcaatgtgcgcagcaacgccact
>s3 rank=0000153 x=2605.5 y=1640.0 length=430
ttttataaccgcgaggacctctggcagttcccgcgccagccgggcggcgatgggattgca
acgatggctccctattacatcatcatgcggctgcccggtgagccgcaagccgagttcttt
ctcatggtcccgatggcgccgagtcgccgcgacaacatgatcgcctggctcgcggcccgt
tgcgacgctcccgactacggcagactgg%%tcgtctacgaattccccaaggagaagctcgtc
tacgggccgttccagatcgaggcgcggatcaatcagagtaccgagatatcgcaacaaatc
acgctgtggaaccagatgggctcgcgcgtgatacgcggcgcaaaacctgcttgtgatccc
gatcgagaattcgatactttatgtcacaccgctctatttgcgagcggaacacggacactt
gccggagctg`}
