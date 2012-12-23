// Copyright ©2012 The bíogo.talks Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"code.google.com/p/biogo.external/muscle"
	"code.google.com/p/biogo/exp/alphabet"
	"code.google.com/p/biogo/exp/seq"
	"code.google.com/p/biogo/exp/seq/linear"
	"code.google.com/p/biogo/exp/seq/multi"
	"code.google.com/p/biogo/exp/seqio/fasta"
	"fmt"
	"io"
	"strings"
)

var s = `>71.2259 lcl|scaffold_41:8288143+
CCCCAAATTCTCATAAAAAGACCAGACTTAATGGTCTGACTGAGACTAGAGGAATCCCGG
TGGTCATGGTCCCCAAACCTTCTGTTGGCCCAGGACAGGAACCATTCCCGAAGACAACTC
ATCAGACACGGAAGGGACTGGACAATGGGTAGGAGAGAGATGCTGACGAAGAGTGAGCTA
CTTGTATCAGGTGGACACTTGAGACTGTGTTGGCATCTCCTGTCTGGAGGGGAGATAGGA
GGGTAGAGAGGGTTAGAAACTGGCAAAATCGTCATGAAAGGAGGGACTGGAAGGAGGGAG
CGGGCTGACTCAGTAGGGGGAGAGTAAGTGGGAGTATGGAGTAAGGTGTATATAAGCTTA
TATGTGACAGATTGACTTGATTTGTAAACTTTCACTTAAAGCACAATAAAAATTATTTTT
TAAAAAATTGTTT
>71.2259 lcl|scaffold_41:11597466-
ATTATTATTTTTTTAAATAATTTTTATTGTGTTTTAAGGGAAAGTTTGCAAATCAAGTCA
GTCTCTCACATATAACCTTATATACACCTTACTCCATACTCCCATTTACTCTCCCCCTAA
TGAGTCAGCCCGCTCCCTCCTTCCGGTCTCTCCTTTCTTGACGATTTTGTCAGTTTCTAA
CCCTCTCTACCCTTCTATCTCTCCTCCAGACAGGAGATGCCAACACTGTCTCAAGTGTCC
ACTTGATACAAGTAGCTCACTCTTCGTCAGCATCTCTCTCCAACCCATTGTCCAGTCCCT
GCCATGTCTGATGAGTTGTCTTTGGGAATGGTTCCTGTCCTGGGCCAACAGAAGGTTTGG
GGACCATGACCGCTGGGATTCCTCTAGTCTCAGTCAGACCATTAAGTCTGGTCTTTTTAT
GAGA
>71.2259 lcl|scaffold_45:2724255+
ATAAAAAGACCAGACTTAATGGTCTGACTGAGACTAGAAGAATCCCGGTGGCCATGGTCC
CCAAACCTTCTGTTGGCCCAGGACAGGAACCATTCCCGAAGACAATTCATCAGACATGGA
AGGGACTGGACAATGGGTTGGAGAGAGATGCTGATAAAGAGTGAGCTACTTGTATCAGGT
GGACGTTTGAGACTGTATTGGCATCTCCTGTCTGGAGGGGAGATAGGGTAGAGAGGGTTA
GAAACTGGCAAAACGGTCACGAAAGGAGAGACTGGAAGAAGGGAGCAGGCTGACTCATTA
GGGGGAGAGTAAATGGGAGTATGTAGTAAGGTGTATATAAGCTTACATGTGACAGACTGA
CTTGATTTGTAAACTTTCACTTAAAGCACAATAAAAATTATTTTTTAAAAATTTGCC
`

func main() {
	m, err := muscle.Muscle{Quiet: true}.BuildCommand()
	if err != nil {
		panic(err)
	}
	m.Stdin = strings.NewReader(s)
	m.Stdout = &bytes.Buffer{}
	m.Run()
	// {CONS OMIT
	var (
		r = fasta.NewReader(m.Stdout.(io.Reader), &linear.Seq{
			Annotation: seq.Annotation{Alpha: alphabet.DNA},
		})
		ms = &multi.Multi{
			ColumnConsense: seq.DefaultQConsensus,
		}
	)
	for {
		s, err := r.Read()
		if err != nil {
			break
		}
		ms.SetName(s.Name())
		ms.Add(s)
	}
	c := ms.Consensus(false)
	c.Threshold = 42
	c.QFilter = seq.CaseFilter
	fmt.Printf("%60a\n", c)
	// CONS} OMIT
}
