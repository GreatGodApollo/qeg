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

package cmd

import (
	"fmt"
	"os"

	"github.com/GreatGodApollo/qeg/internal"
	"github.com/atotto/clipboard"
	"github.com/jawher/mow.cli"
	"io/ioutil"
)

func Exec() {
	app := cli.App("qeg", "Quick Estimate Generator")

	app.Spec = "[-c] [-d] FILE"
	app.Version("v version", "qeg v0.2.0")

	var (
		copy     = app.BoolOpt("c copy", false, "Copy to clipboard?")
		discord  = app.BoolOpt("d discord", false, "Surround in codeblock?")
		fileName = app.StringArg("FILE", "", "What file should it read from?")
	)

	app.Action = func() {
		data, err := ioutil.ReadFile(*fileName)
		if err != nil {
			fmt.Println("Failed to read input file: \n", err.Error())
			return
		}

		estimateJson, err := internal.UnmarshalEstimateJSON(data)
		if err != nil {
			fmt.Println("Failed to decode JSON: \n", err.Error())
			return
		}
		
		estimate := internal.NewGenerator(estimateJson.Title, estimateJson.Customer)
		if estimateJson.CurrencyFormat != nil {
			estimate.CurrencyFormat = *estimateJson.CurrencyFormat
		}

		estimate.AddItems(estimateJson.Items)

		builder := ""
		if *discord {
			builder += "```\n"
		}

		builder += estimate.StringEstimate()

		if *discord {
			builder += "\n```"
		}

		fmt.Println(builder)

		if *copy {
			if clipboard.Unsupported {
				fmt.Println("\nYour platform is unsupported for clipboard copying!")
				return
			}
			_ = clipboard.WriteAll(builder)
			fmt.Println("\nEstimate copied to clipboard!")
		}
	}

	app.Run(os.Args)
}
