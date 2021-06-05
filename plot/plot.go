package plot

import(
	"yagan.bot/mysql"
	"yagan.bot/telegram"
	"container/list"
	"time"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"strconv"
)

func LastHourPlot(){
	var plotPath = "/mnt/md0/shared/points.png"

	time.Sleep(5 * time.Second)

	for {
		tickers := mysql.RetrieveTickers()

		p := plot.New()

		p.Title.Text = "Valores del Mercado"
		plotutil.AddLinePoints(p,
			"Compra", createSellPoints(tickers),
			"Venta", createBuyPoints(tickers))

		p.NominalX(extractNames(tickers)...)

		if err := p.Save(6*vg.Inch, 3*vg.Inch, plotPath); err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)

		telegram.SendImage(plotPath)

		time.Sleep(1 * time.Hour)
	}
}

func createBuyPoints(values *list.List) plotter.XYs {
	pts := make(plotter.XYs, values.Len())

	var i int = 0
	for e := values.Front(); e != nil; e=e.Next() {
		pts[i].X = float64(i)
		
		s, _ := strconv.ParseFloat(e.Value.(mysql.Ticker).Min_ask, 64)
		pts[i].Y = s

		i++
	}

	return pts
} 

func createSellPoints(values *list.List) plotter.XYs {
	pts := make(plotter.XYs, values.Len())

	var i int = 0
	for e := values.Front(); e != nil; e=e.Next() {
		pts[i].X = float64(i)

		s, _ := strconv.ParseFloat(e.Value.(mysql.Ticker).Max_bid, 64)
		pts[i].Y = s

		i++
	}

	return pts
}

func extractNames(values *list.List) []string {
	names := make([]string, values.Len())

	var i int = 0
	for e := values.Front(); e != nil; e=e.Next() {
		hour, min, _ := e.Value.(mysql.Ticker).Date_time.Clock()
		names[i] = fmt.Sprintf("%02d:%02d",hour,min)
		i++
	}

	return names
}