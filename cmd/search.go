// Copyright © 2017 Lee Briggs <lee@leebriggs.co.uk>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	//"github.com/acidlemon/go-dumper"
	"github.com/jaxxstorm/gomdb"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"strings"
)

var searchType string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for omdb entires",
	Run: func(cmd *cobra.Command, args []string) {

		// join all the args we pass, and join them into a search string
		query := &gomdb.QueryData{Title: strings.Join(args, " "), Year: year, SearchType: searchType}

		// do the search
		res, err := gomdb.Search(query)
		if err != nil {
			fmt.Println(err)
			return
		}

		// create a slice for the tablewrite
		// it only takes a slice of strings
		var entry []string

		// build the table
		table := tablewriter.NewWriter(os.Stdout)
		// set some headers
		table.SetHeader([]string{"Title", "Year", "Imdb ID", "Type"})

		// I don't love this, but it'll do
		// loop through the results, and append each []SearchResult
		// to the above slice
		for _, v := range res.Search {
			entry = append(entry, v.Title)
			entry = append(entry, v.Year)
			entry = append(entry, v.ImdbID)
			entry = append(entry, v.Type)
			// add the entry to the table
			table.Append(entry)
			// now clear the slice, ready to go again
			entry = entry[:0]
		}
		table.Render()
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringVarP(&year, "year", "y", "", "Specify a year to search")
	searchCmd.Flags().StringVarP(&searchType, "type", "t", "", "Specify a search type")

}
