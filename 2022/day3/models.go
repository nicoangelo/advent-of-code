package day3

import "log"

type CompartmenContents struct {
	Compartment1 *map[rune]int
	Compartment2 *map[rune]int
}

func (cc *CompartmenContents) init() {
	cc.Compartment1 = &map[rune]int{}
	cc.Compartment2 = &map[rune]int{}
}

type ElfGroup struct {
	Elfs []*CompartmenContents
}

func (eg *ElfGroup) init() {
	eg.Elfs = make([]*CompartmenContents, ELF_GROUP_SIZE)
	for i := range eg.Elfs {
		eg.Elfs[i] = &CompartmenContents{}
		eg.Elfs[i].init()
	}
}

func (eg *ElfGroup) findSharedItem() (res rune) {
	runeCounter := func(k rune) (currentItemCount int) {
		currentItemCount = 1 // it at least exists in the first elf's rucksack
		for i := 1; i < len(eg.Elfs); i++ {
			if _, ok := (*eg.Elfs[i].Compartment1)[k]; ok {
				currentItemCount += 1
			}
			if _, ok := (*eg.Elfs[i].Compartment2)[k]; ok {
				currentItemCount += 1
			}
		}
		return currentItemCount
	}
	for k := range *eg.Elfs[0].Compartment1 {
		if runeCounter(k) == ELF_GROUP_SIZE {
			return k
		}
	}
	for k := range *eg.Elfs[0].Compartment2 {
		if runeCounter(k) == ELF_GROUP_SIZE {
			return k
		}
	}
	log.Fatalf("No shared item found in %d elves.", ELF_GROUP_SIZE)
	return 0
}
