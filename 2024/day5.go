package aoc2024

import (
	"bufio"
	"strconv"
	"strings"
)

type PrintingRule struct {
	first int
	second int
}

type PrintingUpdate struct {
	inOrder bool
	updates []int
}

type PrintingManual struct {
	rules []PrintingRule
	updates []PrintingUpdate
}

func GetSumOfMiddlePageNumbers(scanner *bufio.Scanner) int {
	
	printingManual := generatePrintingManual(scanner)

	for i := 0; i < len(printingManual.updates); i++ {
		printingManual.updates[i].inOrder = isUpdateInOrder(printingManual.updates[i], printingManual.rules)
	}

	sumOfMiddlePageNumber := 0
	for _, update := range printingManual.updates {
		if update.inOrder {
			sumOfMiddlePageNumber += getMiddlePageNumber(update)
		}
	}

	return sumOfMiddlePageNumber
}

func GetSumOfIncorrectlyOrderedUpdateMiddlePages(scanner *bufio.Scanner) int {

	printingManual := generatePrintingManual(scanner)

	for i := 0; i < len(printingManual.updates); i++ {
		printingManual.updates[i].inOrder = isUpdateInOrder(printingManual.updates[i], printingManual.rules)
	}

	
	sumOfMiddlePageNumber := 0
	for _, update := range printingManual.updates {
		if !update.inOrder {
			reorderedUpdate := reorderUpdate(update, printingManual.rules)
			reorderedUpdate.inOrder = isUpdateInOrder(reorderedUpdate, printingManual.rules)
			if !reorderedUpdate.inOrder {
				for i := 0; i < len(printingManual.rules) * len(reorderedUpdate.updates); i++ {
					if !isUpdateInOrder(reorderedUpdate, printingManual.rules) {
						reorderedUpdate = reorderUpdate(reorderedUpdate, printingManual.rules)
					} else {
						reorderedUpdate.inOrder = isUpdateInOrder(reorderedUpdate, printingManual.rules)
						sumOfMiddlePageNumber += getMiddlePageNumber(reorderedUpdate)
						break
					}
				}
			} else {
				sumOfMiddlePageNumber += getMiddlePageNumber(reorderedUpdate)
			}
		}
	}

	return sumOfMiddlePageNumber
}

func reorderUpdate(update PrintingUpdate, rules []PrintingRule) PrintingUpdate {
	for _, rule := range rules {
		firstIndex := -1
		secondIndex := -1


		for index, pageNumber := range update.updates {
			if rule.first == pageNumber {
				firstIndex = index
			}

			if rule.second == pageNumber {
				secondIndex = index
			}
		}

		if firstIndex == -1 || secondIndex == -1 {
			continue
		}

		if firstIndex > secondIndex {
			update.updates[firstIndex], update.updates[secondIndex] = update.updates[secondIndex], update.updates[firstIndex]
		}
	}

	return update
}

func generatePrintingManual(scanner *bufio.Scanner) PrintingManual {
	var printingManual PrintingManual
	for scanner.Scan() {
		line := scanner.Text()
		rule := strings.Split(line, "|")
		var printingRule PrintingRule
		if len(rule) > 1 {
			page1, err := strconv.Atoi(rule[0])

			if err == nil {
				printingRule.first = page1
			}

			page2, err := strconv.Atoi(rule[1])

			if err == nil {
				printingRule.second = page2
			}
			printingManual.rules = append(printingManual.rules, printingRule)
		}

		update := strings.Split(line, ",")

		if len(update) > 1 {
			var printingUpdate PrintingUpdate
			for i := 0; i < len(update); i++ {
				pageNumber, err := strconv.Atoi(update[i])


				if err == nil {
					printingUpdate.inOrder = false
					printingUpdate.updates = append(printingUpdate.updates, pageNumber)

				}
			}
			printingManual.updates = append(printingManual.updates, printingUpdate)
		}
	}

	return printingManual
}

func isUpdateInOrder(update PrintingUpdate, rules []PrintingRule) bool {
	var areRulesMet []bool
	for _, rule := range rules {
		firstIndex := -1
		secondIndex := -1
		for index, pageNumber := range update.updates {
			if rule.first == pageNumber {
				firstIndex = index
			}

			if rule.second == pageNumber {
				secondIndex = index
			}
		}

		if firstIndex == -1 || secondIndex == -1 {
			continue
		}

		if firstIndex < secondIndex {
			areRulesMet = append(areRulesMet, true)
		} else {
			areRulesMet = append(areRulesMet, false)
		}
	}
	
	isMet := true
	for _, isRuleMet := range areRulesMet {
		if !isRuleMet {
			isMet = false
		}
	}
	return isMet
}

func getMiddlePageNumber(update PrintingUpdate) int {
	length := len(update.updates)
	index := length / 2
	return update.updates[index]
}