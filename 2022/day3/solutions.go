package day3

var itemPriorityMap = calculateItemPriorityMap()

const ELF_GROUP_SIZE = 3

func part1(lines []string) (sharedItemPrioritySum int) {
	for _, v := range lines {
		splitIndex := len(v) / 2
		compartmentItemMap := &CompartmenContents{}
		compartmentItemMap.init()
		for i, r := range v {
			if i < splitIndex {
				(*compartmentItemMap.Compartment1)[r] += 1
			} else {
				(*compartmentItemMap.Compartment2)[r] += 1
			}
		}
		sharedItemPrioritySum += getSharedItemPriority(compartmentItemMap)
	}
	return sharedItemPrioritySum
}

func part2(lines []string) (sharedItemPrioritySum int) {
	elfGroupCompartments := getNewElfGroup()

	for i, v := range lines {
		currentElfIndex := i % ELF_GROUP_SIZE
		splitIndex := len(v) / 2
		currentElf := elfGroupCompartments.Elfs[currentElfIndex]
		for j, r := range v {
			if j < splitIndex {
				(*currentElf.Compartment1)[r] += 1
			} else {
				(*currentElf.Compartment2)[r] += 1
			}
		}
		if currentElfIndex == ELF_GROUP_SIZE-1 {
			sharedItem := elfGroupCompartments.findSharedItem()
			sharedItemPrioritySum += itemPriorityMap[sharedItem]
			elfGroupCompartments = getNewElfGroup()
		}
	}
	return sharedItemPrioritySum
}

func getNewElfGroup() (eg *ElfGroup) {
	eg = &ElfGroup{}
	eg.init()
	return eg
}

func getSharedItemPriority(itemMap *CompartmenContents) int {
	for r := range *itemMap.Compartment1 {
		if (*itemMap.Compartment2)[r] > 0 {
			return itemPriorityMap[r]
		}
	}
	return 0
}

func calculateItemPriorityMap() (res map[rune]int) {
	res = make(map[rune]int, 0)
	const letterCount = 26
	for i, startLetter := range []rune{'a', 'A'} {
		for j := 0; j < letterCount; j++ {
			res[rune(startLetter+rune(j))] = j + (i * letterCount) + 1
		}
	}
	return res
}
