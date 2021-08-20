package models

import (
	"testing"

	"github.com/shopspring/decimal"
)

var (
	orderBook = OrderBook{}
)

func parse(p string, v string) []decimal.Decimal {
	dp, _ := decimal.NewFromString(p)
	dv, _ := decimal.NewFromString(v)
	return []decimal.Decimal{dp, dv}
}

func TestOrderBook_Update(t *testing.T) {
	tests := []struct {
		name string
		b    *OrderBook
		want int64
	}{
		{
			"snapshot",
			&OrderBook{
				Asks: [][]decimal.Decimal{parse("44663", "0.0108"), parse("44664", "0.1919"), parse("44665", "0.4226"), parse("44666", "0.2307"), parse("44677", "7.8185"),
					parse("44678", "0.4"), parse("44679", "0.137"), parse("44681", "1.0702"), parse("44682", "0.0281"), parse("44685", "0.2225"),
					parse("44686", "0.9423"), parse("44687", "1.5"), parse("44690", "0.0112"), parse("44695", "0.0234"), parse("44696", "0.6066"),
					parse("44698", "0.0662"), parse("44699", "0.044"), parse("44700", "2.4098"), parse("44703", "0.1051"), parse("44706", "0.14"),
					parse("44707", "0.0156"), parse("44712", "0.358"), parse("44713", "0.0009"), parse("44714", "0.0467"), parse("44716", "0.034"),
					parse("44718", "0.0234"), parse("44720", "0.0001"), parse("44723", "0.0467"), parse("44729", "0.039"), parse("44732", "0.5052"),
					parse("44735", "0.0051"), parse("44736", "20.4916"), parse("44737", "0.235"), parse("44739", "0.361"), parse("44744", "0.0002"),
					parse("44749", "0.0018"), parse("44750", "0.0632"), parse("44752", "0.1294"), parse("44753", "0.0045"), parse("44754", "0.0003"),
					parse("44761", "0.0006"), parse("44763", "0.0009"), parse("44764", "12.9036"), parse("44771", "0.0055"), parse("44777", "0.0001"),
					parse("44779", "0.0001"), parse("44780", "0.0761"), parse("44782", "0.195"), parse("44785", "0.0001"), parse("44787", "0.0017"),
					parse("44791", "0.0005"), parse("44792", "0.0005"), parse("44793", "0.0019"), parse("44796", "0.001"), parse("44798", "23.3279"),
					parse("44800", "0.092"), parse("44802", "0.079"), parse("44809", "0.0121"), parse("44813", "0.0009"), parse("44815", "0.0001"),
					parse("44816", "0.0004"), parse("44825", "0.0001"), parse("44828", "0.0237"), parse("44830", "0.0001"), parse("44832", "0.0003"),
					parse("44839", "1.389"), parse("44840", "0.0001"), parse("44841", "0.0082"), parse("44843", "0.0106"), parse("44847", "0.0006"),
					parse("44848", "0.02"), parse("44850", "0.0589"), parse("44859", "0.0059"), parse("44860", "0.554"), parse("44861", "18.3891"),
					parse("44862", "0.0003"), parse("44863", "0.0009"), parse("44864", "0.0001"), parse("44868", "0.0003"), parse("44873", "0.0001"),
					parse("44875", "0.0005"), parse("44878", "0.0016"), parse("44882", "0.0001"), parse("44884", "0.002"), parse("44885", "0.0056"),
					parse("44887", "0.0012"), parse("44889", "32.2195"), parse("44891", "27.6158"), parse("44893", "0.002"), parse("44895", "4.3145"),
					parse("44900", "0.0624"), parse("44907", "0.0001"), parse("44909", "0.1114"), parse("44911", "0.0005"), parse("44912", "0.001"),
					parse("44913", "8.5105"), parse("44916", "0.0004"), parse("44922", "28.7159"), parse("44928", "0.0022"), parse("44931", "0.003")},
				Bids: [][]decimal.Decimal{parse("44655", "0.31"), parse("44654", "0.6022"), parse("44650", "6.5661"), parse("44648", "0.008"), parse("44646", "0.008"),
					parse("44644", "0.028"), parse("44640", "1.5"), parse("44639", "0.9329"), parse("44638", "0.4177"), parse("44630", "0.0453"),
					parse("44628", "8.4713"), parse("44625", "0.0112"), parse("44623", "0.409"), parse("44620", "0.5599"), parse("44615", "0.0468"),
					parse("44613", "0.0009"), parse("44611", "0.0599"), parse("44607", "0.3"), parse("44606", "2.05"), parse("44605", "0.0234"),
					parse("44604", "0.039"), parse("44603", "0.0468"), parse("44602", "13.5048"), parse("44600", "0.123"), parse("44599", "0.0327"),
					parse("44598", "0.0467"), parse("44595", "0.346"), parse("44594", "0.0001"), parse("44592", "0.1051"), parse("44588", "0.235"),
					parse("44586", "0.0054"), parse("44583", "0.0001"), parse("44581", "0.1052"), parse("44567", "0.0005"), parse("44565", "0.079"),
					parse("44564", "0.001"), parse("44563", "0.0009"), parse("44559", "0.0001"), parse("44553", "0.002"), parse("44552", "0.0001"),
					parse("44550", "21.5564"), parse("44545", "0.0001"), parse("44540", "0.0001"), parse("44519", "0.1627"), parse("44518", "0.542"),
					parse("44514", "0.0004"), parse("44513", "0.0009"), parse("44501", "0.001"), parse("44500", "0.0618"), parse("44496", "23.5431"),
					parse("44488", "0.0001"), parse("44486", "0.0003"), parse("44485", "0.0107"), parse("44483", "0.0008"), parse("44480", "0.0001"),
					parse("44470", "0.0025"), parse("44464", "0.0005"), parse("44463", "0.0009"), parse("44459", "0.0001"), parse("44450", "0.0516"),
					parse("44448", "0.001"), parse("44447", "0.0001"), parse("44444", "0.02"), parse("44443", "0.0001"), parse("44434", "23.9983"),
					parse("44430", "0.0035"), parse("44424", "26.4719"), parse("44421", "0.0001"), parse("44420", "0.0001"), parse("44419", "0.0001"),
					parse("44413", "0.0031"), parse("44411", "32.7509"), parse("44403", "24.701"), parse("44400", "0.0883"), parse("44396", "0.0003"),
					parse("44387", "8.7491"), parse("44382", "0.002"), parse("44380", "0.003"), parse("44375", "0.0004"), parse("44374", "0.0001"),
					parse("44365", "0.0121"), parse("44363", "0.001"), parse("44362", "0.0005"), parse("44360", "0.0019"), parse("44350", "0.0595"),
					parse("44345", "0.235"), parse("44333", "0.0018"), parse("44332", "0.0011"), parse("44327", "0.0001"), parse("44324", "0.0001"),
					parse("44323", "0.0004"), parse("44320", "43.5662"), parse("44318", "41.4537"), parse("44315", "0.0107"), parse("44313", "0.0009"),
					parse("44309", "0.001"), parse("44305", "0.112"), parse("44301", "0.0006"), parse("44300", "0.0576"), parse("44296", "0.0057")}},
			int64(775538454),
		}, {
			"update 1",
			&OrderBook{
				Asks: [][]decimal.Decimal{parse("44706", "0.1237")},
				Bids: [][]decimal.Decimal{parse("44650", "0"), parse("44292", "0.0006")}},
			int64(62906094),
		}, {
			"update 2",
			&OrderBook{
				Asks: [][]decimal.Decimal{parse("44664", "0"), parse("44941", "0.0001")},
				Bids: [][]decimal.Decimal{}},
			int64(2563962883),
		}, {
			"update 3",
			&OrderBook{
				Asks: [][]decimal.Decimal{parse("44712", "0.4635"), parse("44728", "0.4"), parse("44941", "0")},
				Bids: [][]decimal.Decimal{}},
			int64(2045348734),
		}, {
			"update 4",
			&OrderBook{
				Asks: [][]decimal.Decimal{parse("44682", "6.1216")},
				Bids: [][]decimal.Decimal{}},
			int64(1834229990),
		}, {
			"update 5",
			&OrderBook{
				Asks: [][]decimal.Decimal{parse("44676", "0.008"), parse("44931", "0")},
				Bids: [][]decimal.Decimal{}},
			int64(144826912),
		}, {
			"update 6",
			&OrderBook{
				Asks: [][]decimal.Decimal{parse("44706", "0.0182"), parse("44712", "0.358"), parse("44717", "0.0447"), parse("44928", "0")},
				Bids: [][]decimal.Decimal{}},
			int64(729799368),
		}, {
			"update 7",
			&OrderBook{
				Asks: [][]decimal.Decimal{parse("44715", "0.0163"), parse("44717", "0")},
				Bids: [][]decimal.Decimal{}},
			int64(171811415),
		}, {
			"update 8",
			&OrderBook{
				Asks: [][]decimal.Decimal{},
				Bids: [][]decimal.Decimal{parse("44625", "0"), parse("44289", "0.0001")}},
			int64(3679650162),
		}, {
			"update 9",
			&OrderBook{
				Asks: [][]decimal.Decimal{},
				Bids: [][]decimal.Decimal{parse("44648", "4.7843")}},
			int64(3966390855),
		},
	}
	a := OrderBook{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := a.Update(tt.b); got != tt.want {
				t.Errorf("OrderBook checksum = %v, want %v", got, tt.want)
			}
		})
	}
}
