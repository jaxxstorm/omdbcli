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
	"strings"

	"github.com/jaxxstorm/gomdb"
	"github.com/spf13/cobra"

	//"github.com/acidlemon/go-dumper"
	"github.com/olekukonko/tablewriter"
)

var ratings bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the details of a specific movie",
	Run: func(cmd *cobra.Command, args []string) {

		// get the movie
		query := &gomdb.QueryData{Title: strings.Join(args, " "), Year: year}

		// we're using the title
		res, err := gomdb.LookupByTitle(query)

		// sad face
		if err != nil {
			fmt.Println(err)
			return
		}

		// more table fun
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Title", "Year", "Director", "Writer", "Actors", "Plot", "Type"})
		// add the entry to the table
		table.Append([]string{res.Title, res.Year, res.Director, res.Writer, res.Actors, res.Plot, res.Type})

		// this is nasty, but time is a factor..
		// loop through the ratings, and then match on the source
		// if it's RT, add it to the footer
		if ratings {
			for _, v := range res.Ratings {
				if v.Source == "Rotten Tomatoes" {
					table.SetFooter([]string{"", "", "", "", "", "Rotten Tomatoes Rating", v.Value})
				}
			}
		}
		table.Render()

	},
}

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&year, "year", "y", "", "Specify a year to search")
	getCmd.Flags().BoolVarP(&ratings, "rating", "r", false, "Show the rotten tomatos rating")
}
