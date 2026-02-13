package txn

import "slices"

const SelectDebts = `
SELECT issuer, aquirer, SUM(amt)
FROM txns
WHERE state = 'AUTH_RESPONDED'
GROUP BY issuer, aquirer
`

type Debt struct {
	issuer  string
	aquirer string
	amt     float32
}

type Net struct {
	issuer      string
	aquirer     string
	issuer_amt  float32
	aquirer_amt float32
	net         float32
}

func FetchNets(debts []Debt) (nets []Net) {
	for _, debt := range debts {
		idx := slices.IndexFunc(nets, func(x Net) bool {
			return (x.issuer == debt.issuer && x.aquirer == debt.aquirer) || (x.aquirer == debt.issuer && x.issuer == debt.aquirer)
		})
		if idx == -1 {
			nets = append(nets, Net{
				issuer:     debt.issuer,
				aquirer:    debt.aquirer,
				issuer_amt: debt.amt,
			})
			continue
		}
		if debt.issuer == nets[idx].issuer {
			nets[idx].issuer_amt = debt.amt
		} else {
			nets[idx].aquirer_amt = debt.amt
		}
	}
	for i := range nets {
		nets[i].net = nets[i].issuer_amt - nets[i].aquirer_amt
		if nets[i].net > 0 {
			continue
		}
		nets[i].aquirer, nets[i].issuer = nets[i].issuer, nets[i].aquirer
		nets[i].aquirer_amt, nets[i].issuer_amt = nets[i].issuer_amt, nets[i].aquirer_amt
		nets[i].net = -nets[i].net
	}
	return
}
