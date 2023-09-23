package csv_to_slice

import (
	"encoding/csv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"strings"
)

func CsvToSlice(csvBody string, skipHeader bool) (header []string, result [][]string, finalErr error) {
	f := strings.NewReader(csvBody)
	csvReader := csv.NewReader(f)
	lineNumber := 0
	for {
		lineNumber++
		line, errR := csvReader.Read()
		if errR == io.EOF {
			break
		}
		if errR != nil {
			return header, nil, errors.WithStack(errR)
		}
		// do something with read line
		logrus.Tracef("%+v line: %+v", lineNumber, strings.Join(line, ", "))
		if skipHeader && lineNumber == 1 {
			header = line
			logrus.Tracef("skipping zero line %+v", line)
			continue
		}
		//var date := line[0]
		result = append(result, line)
	}
	//res.Body
	//	-- #2020_income_statement_quarter.csv
	//	-- #"date","symbol","reportedCurrency","cik","fillingDate","acceptedDate","calendarYear","period","revenue","costOfRevenue","grossProfit","grossProfitRatio","ResearchAndDevelopmentExpenses","GeneralAndAdministrativeExpenses","SellingAndMarketingExpenses","SellingGeneralAndAdministrativeExpenses","otherExpenses","operatingExpenses","costAndExpenses","interestExpense","depreciationAndAmortization","EBITDA","EBITDARatio","operatingIncome","operatingIncomeRatio","totalOtherIncomeExpensesNet","incomeBeforeTax","incomeBeforeTaxRatio","incomeTaxExpense","netIncome","netIncomeRatio","EPS","EPSDiluted","weightedAverageShsOut","weightedAverageShsOutDil","link","finalLink","interestIncome"
	//	-- #"2020-09-30","CKX","USD","0000352955","2020-11-06","2020-11-06 14:52:17",2020,"Q3",160842,7416,153426,0.9538926399820943,0,121426,0,121426,0,121426,128842,6914,1326,88485,0.5501361584660723,73331,0.4559194737692891,6914,80245,0.498905758446177,30191,50054,0.3111998109946407,0.03,0.03,1942495,1942495,"https://www.sec.gov/Archives/edgar/data/352955/000143774920022977/0001437749-20-022977-index.htm","https://www.sec.gov/Archives/edgar/data/352955/000143774920022977/ckx20200930_10q.htm",6914
	return header, result, nil
}
