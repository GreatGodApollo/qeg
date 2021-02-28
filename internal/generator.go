/*
 * MIT License
 *
 * Copyright (c) 2021 Brett Bender
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package internal

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"strings"
)

type EstimateGenerator struct {
	Title          string
	Customer       string
	CurrencyFormat string
	Items          []Item
	Sizes          []int
}

func NewGenerator(title, customer string) *EstimateGenerator {
	return &EstimateGenerator{
		Title:    title,
		Customer: customer,
		CurrencyFormat: "$%.2f",
		Items:    make([]Item, 0),
		Sizes:    make([]int, 3),
	}
}

func (eg *EstimateGenerator) AddItem(description string, price float64) *EstimateGenerator {
	eg.Items = append(eg.Items, Item{Description: description, Price: price})
	eg.CalculateSizes()

	return eg
}

func (eg *EstimateGenerator) AddItems(items []Item) *EstimateGenerator {
	eg.Items = append(eg.Items, items...)
	eg.CalculateSizes()
	
	return eg
}

func (eg *EstimateGenerator) CalculateSizes() {
	itemSize := 0
	priceSize := 0

	tLen := runewidth.StringWidth("Invoice: " + eg.Title)
	cLen := runewidth.StringWidth("Customer: " + eg.Customer)
	if cLen > tLen {
		tLen = cLen
	}

	for _, k := range eg.Items {
		iLen := runewidth.StringWidth(k.Description)
		pLen := runewidth.StringWidth(fmt.Sprintf(eg.CurrencyFormat, k.Price))

		if iLen > itemSize {
			itemSize = iLen
		}

		if pLen > priceSize {
			priceSize = pLen
		}
	}

	if tLen < (priceSize + itemSize + 5) {
		tLen = priceSize + itemSize + 5
	}

	eg.Sizes[0] = itemSize
	eg.Sizes[1] = priceSize
	eg.Sizes[2] = tLen
}

func (eg *EstimateGenerator) dash() string {
	return strings.Repeat("=", eg.Sizes[2]+2)
}

func (eg *EstimateGenerator) StringEstimate() string {
	builder := ""

	builder += "Invoice: " + eg.Title + "\n"
	builder += "Customer: " + eg.Customer + "\n"

	builder += eg.dash() + "\n\n"

	total := 0.0
	for _, item := range eg.Items {
		spacesLeft := (eg.Sizes[2] - eg.Sizes[1] - 4) - runewidth.StringWidth(item.Description)
		builder += strings.Repeat(" ", spacesLeft)
		builder += item.Description
		builder += "  | "

		price := fmt.Sprintf(eg.CurrencyFormat, item.Price)
		spacesLeft = eg.Sizes[1] - runewidth.StringWidth(price)
		builder += strings.Repeat(" ", spacesLeft)
		builder += price + "\n"

		total += item.Price
	}

	builder += eg.dash() + "\n"

	totalLine := fmt.Sprintf("Total: " + eg.CurrencyFormat, total)
	spacesLeft := eg.Sizes[2] - runewidth.StringWidth(totalLine)
	builder += strings.Repeat(" ", spacesLeft) + totalLine

	return builder
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
